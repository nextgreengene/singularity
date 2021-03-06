package shell

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/margostino/singularity/pkg/config"
	"github.com/margostino/singularity/pkg/context"
	"strconv"
	"strings"
)

type Shell struct {
	Suggestions []prompt.Suggest
}

var PowerShell *Shell

func Welcome() {
	fmt.Printf("%s (v.%s)\n", config.Welcome(), config.Version())
}

func (s *Shell) Prompt() string {
	username := context.GetUsername()
	state := context.GetState()
	clock := context.GetClock()
	prefix := fmt.Sprintf("%s@%s(%s)> ", state, username, clock)
	return prompt.Input(strings.ToLower(prefix), Completer(s.Suggestions))
}

func (s *Shell) Input() []string {
	commandLine := s.Prompt()
	return strings.Fields(commandLine)
}

func NewShell() *Shell {
	var suggestions = make([]prompt.Suggest, 0)
	commands := config.GetCommandsConfiguration().CommandList

	for _, command := range commands {
		var commandText string
		if command.Args > 0 {
			commandText = command.Id + " x" + strconv.Itoa(command.Args)
		} else {
			commandText = command.Id
		}
		suggestion := prompt.Suggest{
			Text:        commandText,
			Description: command.Description,
		}
		suggestions = append(suggestions, suggestion)
	}
	PowerShell = &Shell{Suggestions: suggestions}
	return PowerShell
}

func Completer(suggestions []prompt.Suggest) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
	}
}

func (s *Shell) GetOptions() []prompt.Suggest {
	return s.Suggestions
}
