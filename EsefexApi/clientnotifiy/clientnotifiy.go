package clientnotifiy

import (
	"esefexapi/types"
	"esefexapi/util/dcgoutil"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type IClientNotifier interface {
	// UpdateNotificationUsers notifies the clients that some data has been updated
	// This should cause the client to refetch the data
	// if users is empty, then all clients should be notified
	// this function will handle the case where a user does not have any connections
	UpdateNotificationUsers(users ...types.UserID) error
	UpdateNotificationGuilds(guilds ...types.GuildID) error
	UpdateNotificationChannels(channels ...types.ChannelID) error
}

var _ IClientNotifier = &WsClientNotifier{}

// implements ClientNotifier
type WsClientNotifier struct {
	userConnections map[types.UserID][]*websocket.Conn
	ds              *discordgo.Session
	stop            chan struct{}
	ready           chan struct{}
}

func NewWsClientNotifier(ds *discordgo.Session) *WsClientNotifier {
	return &WsClientNotifier{
		userConnections: make(map[types.UserID][]*websocket.Conn),
		ds:              ds,
		stop:            make(chan struct{}),
		ready:           make(chan struct{}),
	}
}

// UpdateNotificationChannels implements IClientNotifier.
func (w *WsClientNotifier) UpdateNotificationChannels(channels ...types.ChannelID) error {
	for _, channel := range channels {
		users, err := dcgoutil.ChannelUserIDs(w.ds, channel)
		if err != nil {
			return errors.Wrap(err, "error getting channel user ids")
		}

		err = w.UpdateNotificationUsers(users...)
		if err != nil {
			return errors.Wrap(err, "error updating notification")
		}
	}

	return nil
}

// UpdateNotificationGuilds implements IClientNotifier.
func (w *WsClientNotifier) UpdateNotificationGuilds(guilds ...types.GuildID) error {
	for _, guild := range guilds {
		users, err := dcgoutil.GuildUserIDs(w.ds, guild)
		if err != nil {
			return errors.Wrap(err, "error getting channel user ids")
		}

		err = w.UpdateNotificationUsers(users...)
		if err != nil {
			return errors.Wrap(err, "error updating notification")
		}
	}

	return nil
}

func (w *WsClientNotifier) UpdateNotificationUsers(users ...types.UserID) error {
	if len(users) == 0 {
		for k := range w.userConnections {
			err := w.writeUpdate(k)
			if err != nil {
				return errors.Wrap(err, "error writing update")
			}
		}
		return nil
	}

	for _, user := range users {
		if _, ok := w.userConnections[user]; !ok {
			continue
		}

		err := w.writeUpdate(user)
		if err != nil {
			return errors.Wrap(err, "error writing update")
		}
	}
	return nil
}

func (w *WsClientNotifier) writeUpdate(user types.UserID) error {
	var causedError error = nil

	for _, conn := range w.userConnections[user] {
		err := conn.WriteMessage(websocket.TextMessage, []byte("update"))
		if err != nil {
			conn.Close()
			w.RemoveConnection(user, conn)
			causedError = errors.Wrap(err, "error writing message to websocket, removing connection")
		}
	}
	return causedError
}

func (w *WsClientNotifier) AddConnection(user types.UserID, conn *websocket.Conn) {
	w.userConnections[user] = append(w.userConnections[user], conn)
}

func (w *WsClientNotifier) RemoveConnection(user types.UserID, conn *websocket.Conn) {
	for i, c := range w.userConnections[user] {
		if c == conn {
			w.userConnections[user] = append(w.userConnections[user][:i], w.userConnections[user][i+1:]...)
			break
		}
	}
}
