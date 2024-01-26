package commands

import (
	"log"

	"github.com/pkg/errors"
)

func (c *CommandHandlers) DeleteGuildApplicationCommands(guildID string) error {
	cmds, err := c.ds.ApplicationCommands(c.ds.State.User.ID, guildID)
	if err != nil {
		return errors.Wrapf(err, "Cannot get commands for guild '%v'", guildID)
	}

	for _, v := range cmds {
		err = c.ds.ApplicationCommandDelete(c.ds.State.User.ID, guildID, v.ID)
		if err != nil {
			return errors.Wrapf(err, "Cannot delete '%v' command", v.Name)
		}
		log.Printf("Deleted '%v' command", v.Name)
	}

	return nil
}

func (c *CommandHandlers) DeleteAllApplicationCommands() error {
	log.Println("Deleting all commands...")

	for _, g := range c.ds.State.Guilds {
		err := c.DeleteGuildApplicationCommands(g.ID)
		if err != nil {
			return errors.Wrapf(err, "Cannot delete commands for guild '%v'", g.ID)
		}
	}

	err := c.DeleteGuildApplicationCommands("")
	if err != nil {
		return errors.Wrap(err, "Cannot delete global commands")
	}

	log.Println("Deleted all commands")

	return nil
}
