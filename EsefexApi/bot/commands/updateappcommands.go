package commands

import (
	"log"

	"github.com/pkg/errors"
)

// get hash of current application commands and compare to the stored hash
// if the hashes are the same, return
// if the hashes are different, delete all application commands and register new ones
// update the stored hash
// This is done to avoid hitting the rate limit for registering application commands (200 per day)
// The rate limit is not usually a problem in production, but it can be if you are developing and testing a lot
func (c *CommandHandlers) UpdateApplicationCommands() error {
	log.Println("Updating application commands...")
	oldHash, err := c.dbs.CmdHashStore.GetCommandHash()
	if err != nil {
		return errors.Wrap(err, "Cannot get command hash")
	}

	newHash, err := c.ApplicationCommandsHash()
	if err != nil {
		return errors.Wrap(err, "Cannot get application command hash")
	}

	if oldHash == newHash {
		log.Println("Command hashes are the same, skipping update")
		return nil
	}
	log.Println("Command hashes are different, updating commands...")

	err = c.DeleteAllApplicationCommands()
	if err != nil {
		return errors.Wrap(err, "Cannot delete all commands")
	}

	err = c.RegisterApplicationCommands()
	if err != nil {
		return errors.Wrap(err, "Cannot register commands")
	}

	err = c.dbs.CmdHashStore.SetCommandHash(newHash)
	if err != nil {
		return errors.Wrap(err, "Cannot set command hash")
	}

	log.Println("Finished updating application commands")

	return nil
}
