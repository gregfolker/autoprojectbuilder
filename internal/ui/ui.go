package ui

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/gregfolker/auto-project-builder/pkg/project"
)

const (
	PromptPrefix = ">> "
)

func GetUserInput(p *project.Project) {
	fmt.Printf("What is the name of this project?\n")
	p.Name = getProjectName()

	fmt.Printf("Who is the author of this project?\n")
	p.Author = getProjectAuthor()

	fmt.Printf("What language will this project be written in?\n")
	p.Language = getProjectLanguage()
}

func getProjectAuthor() string {
	return prompt.Input(PromptPrefix, authorAutoCompleter)
}

func getProjectName() string {
	return prompt.Input(PromptPrefix, emptyAutoCompleter)
}

func getProjectLanguage() string {
	return prompt.Input(PromptPrefix, languageAutoCompleter)
}

func authorAutoCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest {
		{Text: "Greg Folker"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func languageAutoCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest {
		{Text: "Go"},
		{Text: "C"},
		{Text: "Python"},
		{Text: "Rust"},
		{Text: "Bash"},
		{Text: "Perl"},
		{Text: "Java"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func emptyAutoCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest {
		{Text: ""},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
