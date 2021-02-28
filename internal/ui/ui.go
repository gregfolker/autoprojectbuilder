package ui

import (
	"github.com/c-bata/go-prompt"
)

const (
	PromptPrefix = ">> "
)

func GetProjectAuthor() string {
	return prompt.Input(PromptPrefix, authorAutoCompleter)
}

func GetProjectName() string {
	return prompt.Input(PromptPrefix, emptyAutoCompleter)
}

func authorAutoCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest {
		{Text: "Greg Folker"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func emptyAutoCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest {
		{Text: ""},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
