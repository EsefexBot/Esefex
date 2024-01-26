package subcommandbuilder

import (
	"esefexapi/bot/commands"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Build builds the subcommand
// Any errors encountered during any of the steps will be returned here to improve usability
func (b *SubCommandBuilder) Build() (commands.Command, error) {
	cmd := commands.Command{
		ApplicationCommand: b.applicationCommand(),
		Handler:            b.handlerProxy(),
	}

	if len(b.errors) > 0 {
		return cmd, b.errors[0]
	}

	return cmd, nil
}

func (b *SubCommandBuilder) applicationCommand() discordgo.ApplicationCommand {
	cmd := discordgo.ApplicationCommand{
		Name:        b.name,
		Description: b.description,
		Options:     []*discordgo.ApplicationCommandOption{},
	}

	// for _, sc := range b.subcommands {
	// 	aco := discordgo.ApplicationCommandOption{
	// 		Type: discordgo.ApplicationCommandOptionSubCommand,
	// 		Name: sc.ApplicationCommand.Name,
	// 		Description: ,
	// 	}

	// 	cmd.Options = append(cmd.Options, &sc.ApplicationCommand)
	// }

	return cmd
}

func (b *SubCommandBuilder) handlerProxy() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		opt0 := i.ApplicationCommandData().Options[0]

		// check if subcommand exists
		for _, c := range b.subcommands {
			if c.ApplicationCommand.Name == opt0.Name {
				// remove the subcommand from the options (so it doesn't get passed to the handler)
				// and replace it with the subcommand's options

				newI := *i
				newData := newI.ApplicationCommandData()
				newData.Options = opt0.Options[1:]
				newI.Data = newData

				c.Handler(s, &newI)
				return
			}
		}

		// otherwise check if subcommand exists in group
		for _, g := range b.groups {
			if g.Name == opt0.Name {
				opt1 := opt0.Options[0]

				for _, c := range g.Commands {
					if c.ApplicationCommand.Name == opt1.Name {
						c.Handler(s, i)
						return
					}
				}
			}
		}

		// if we get here, the subcommand doesn't exist
		// that means we have a bug somewhere

		log.Println("SubCommandBuilder: subcommand not found")
		log.Println("SubCommandBuilder: subcommand name:", opt0.Name)

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Subcommand not found, this is a bug, please report it to the bot developer",
			},
		})

		if err != nil {
			log.Println("SubCommandBuilder: error sending response:", err)
		}
	}
}
