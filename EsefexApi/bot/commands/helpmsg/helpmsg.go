package helpmsg

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

var helpMessages = map[string]string{
	"":                    "General.txt",
	"UI":                  "UI.txt",
	"Commands:Bot":        "Commands:Bot.txt",
	"Commands:Sound":      "Commands:Sound.txt",
	"Commands:User":       "Commands:User.txt",
	"Commands:Permission": "Commands:Permission.txt",
}

// GetHelpMessage returns the help message for the given category.
// If the category is empty, it returns the general help message.
// it will open the file in the helpmsg folder with the name of the category.
func GetHelpMessage(category string) (string, error) {
	fname, ok := helpMessages[category]
	if !ok {
		fname = helpMessages[""]
	}

	p := fmt.Sprintf("bot/commands/helpmsg/%s", fname)
	buf, err := os.ReadFile(p)
	if err != nil {
		return "", errors.Wrap(err, "Error reading help message")
	}

	return string(buf), nil
}
