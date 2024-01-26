package commands

import (
	"log"

	"github.com/pkg/errors"
)

// Call after opening the session
func (c *CommandHandlers) RegisterApplicationCommands() error {
	log.Println("Registering application commands...")

	for _, v := range c.Commands {
		_, err := c.ds.ApplicationCommandCreate(c.ds.State.User.ID, "", v)
		if err != nil {
			return errors.Wrapf(err, "Cannot create '%v' command", v.Name)
		}

		log.Printf("Registered '%v' command", v.Name)
	}

	log.Println("Registered all application commands")

	return nil
}
