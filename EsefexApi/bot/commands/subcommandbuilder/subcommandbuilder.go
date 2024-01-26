package subcommandbuilder

import (
	"esefexapi/bot/commands"
	"fmt"

	"github.com/pkg/errors"
)

// SubCommandBuilder is a builder for creating subcommands
type SubCommandBuilder struct {
	errors      []error
	name        string
	description string
	subcommands []commands.Command
	groups      []commands.SubcommandGroup
}

func New(name string, description string) *SubCommandBuilder {
	return &SubCommandBuilder{
		name:        name,
		description: description,
		subcommands: []commands.Command{},
		groups:      []commands.SubcommandGroup{},
	}
}

func (b *SubCommandBuilder) SubCommand(cmd commands.Command) *SubCommandBuilder {
	// check if subcommand already exists (by name)
	for _, c := range b.subcommands {
		if c.ApplicationCommand.Name == cmd.ApplicationCommand.Name {
			b.errors = append(b.errors, errors.WithStack(fmt.Errorf("SubCommand: subcommand %s already exists", cmd.ApplicationCommand.Name)))
			return b
		}
	}

	b.subcommands = append(b.subcommands, cmd)
	return b
}

func (b *SubCommandBuilder) SubCommandGroup(name string, description string) *SubCommandBuilder {
	// if group already exists, return error
	for _, g := range b.groups {
		if g.Name == name {
			b.errors = append(b.errors, errors.WithStack(fmt.Errorf("SubCommandGroup: group %s already exists", name)))
			return b
		}
	}

	b.groups = append(b.groups, commands.SubcommandGroup{
		Name:        name,
		Description: description,
		Commands:    []*commands.Command{},
	})
	return b
}

func (b *SubCommandBuilder) SubCommandGroupCommand(gname string, cmd commands.Command) *SubCommandBuilder {
	// check if subcommand already exists in group (by name)
	for _, c := range b.groups {
		if c.Name == gname {
			for _, sc := range c.Commands {
				if sc.ApplicationCommand.Name == cmd.ApplicationCommand.Name {
					b.errors = append(b.errors, errors.WithStack(fmt.Errorf("SubCommandGroupCommand: subcommand %s already exists in group %s", cmd.ApplicationCommand.Name, gname)))
					return b
				}
			}
		}
	}

	for i, g := range b.groups {
		if g.Name == gname {
			b.groups[i].Commands = append(b.groups[i].Commands, &cmd)
			return b
		}
	}

	b.errors = append(b.errors, errors.WithStack(fmt.Errorf("SubCommandGroupCommand: could not find group %s", gname)))
	return b
}
