# Esefex Backend
This is the backend for the esefex bot. This is a go application that uses discordgo and gorilla/mux. It manages http/ws api and the discord bot.

## Api
To use most of the api endpoints you need to have the `User-Token` cookie set to your esefex token (generated with the `/user link` command)



### `/dump` (GET)
Dumps the Request, used for debugging purposes
```go
Status Codes: 200
Body: string
Cross Origin: *
Authorization Required: false
```
<hr>

### `/api/guild` (GET)
Returns the guilds the user is in (if any)
```go
Status Codes: 200, 401, 403, 500
Body: string | error
Cross Origin: *
Authorization Required: true
```
<hr>

### `/api/guilds` (GET)
Returns all guilds the user is in
```go
Status Codes: 200, 401, 403, 500
Body: []{
    guildID: string,
    guildName: string
} | error
Cross Origin: *
Authorization Required: true
```
<hr>

### `/` (GET)
Returns the index page
```go
Status Codes: 200
Body: string
Cross Origin: *
Authorization Required: false
```
<hr>

### `/link?<guild_id>` (GET)
Returns the link page, where you can link your discord account to your esefex account, this is meant to be used with the `/user link` command by the user. 

This site does not set the `User-Token` cookie, you will have to click the button which . This is to prevent the discord embed crawler from using the one time link token.
#### Params
`<guild_id>`: The guild id to link to

```go
Status Codes: 200 | 400 | 401 | 500
Body: html | error
Cross Origin: *
Authorization Required: false
```
<hr>

### `/api/link?<guild_id>` (GET)
The second part of the link process, this is meant to be used by the `/user link` command by the user.

#### Params
`<guild_id>`: The guild id to link to

```go
Status Codes: 200 | 400 | 401 | 500
Body: html | error
Cross Origin: *
Authorization Required: false
```
<hr>

### `/api/sounds/<guild_id>` (GET)
Returns all sounds for the guild

#### Params
`<guild_id>`: The guild id to get the sounds for

```go
Status Codes: 200 | 403 | 500
Body: []{
    id: string, // The id is the sha256 hash of the name, used to play the sound
    guildId: string,
    name: string, // The name of the sound, this is how users identify the sound
    icon: {
        regularEmoji: bool, // If the icon is a regular emoji or a custom one
        name: string, // The name of the emoji or the unicode character (ie. "🤡")
        id: string, // The id of the emoji, if it is a custom one
        url: string // The url of the emoji image
    },
    length: float32 // The length of the sound in seconds
    created: int64 // The unix timestamp of when the sound was created
} | error
Cross Origin: *
Authorization Required: true
```
<hr>

### `api/ws` (GET)
Connect to the websocket, this is used by the api to send events to the client (currently only "update" events without any data)
```go
Status Codes: 200 | 401
Message: "update"
Cross Origin: *
Authorization Required: true
```
<hr>

### `api/playsound/<sound_id>` (POST)
Plays a sound in the voice channel the user is in

#### Params
`<sound_id>`: The id of the sound to play

```go
Status Codes: 200 | 401 | 500
Body: string | error
Cross Origin: *
Authorization Required: true
```


## Bot Commands
```go
/help -> shows general Help

/sound -> shows sound help
/sound upload -> upload sound
/sound list -> list sounds
/sound delete -> delete sounds
/sound modify -> modify sounds

/permission -> shows permissions help
/permission set role <roleid> <permissionID> <value> -> sets permissions for user
/permission set user <userID> <permissionID> <value>
/permission set channel <channelID> <permissionID> <value>

/permission get role/user/channel -> gets permission value for role/user/channel
/permission list role/user/channel -> list all permissions for role/user/channel
/permission clear role/user/channel -> clears permissions for role/user/channel

/bot -> shows bot help
/bot join -> makes bot join
/bot leave -> makes bot leave
/bot stats -> displays bot stats

/config -> change bot config

/user link -> link user
/user unlink -> unlink user
/user stats -> show user stats
```
