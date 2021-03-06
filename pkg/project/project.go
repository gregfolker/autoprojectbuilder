package project

import (
	"os"
	"path"
	"strings"
	"fmt"
	"errors"
	"github.com/gregfolker/auto-project-builder/pkg/files"
	"github.com/gregfolker/auto-project-builder/pkg/languages"
	"github.com/gregfolker/auto-project-builder/pkg/colors"
)

type Project struct {
	Name string
	Author string
	Language string
	Path string
}

func NewProject() *Project {
	p := &Project{}

	return p
}

func (prog Project) PrintProjectInfo() {
	fmt.Printf("\n--- New Project Details ---\n")
	fmt.Printf("Project Name: %s\n", prog.Name)
	fmt.Printf("Author: %s\n", prog.Author)
	fmt.Printf("Language: %s\n", prog.Language)
	fmt.Printf("Path: %s\n", prog.Path)
}

func (prog *Project) CreateNewProjectDir() error {
	var dir string
	userHome, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	// Avoid creating directories with spaces
	if strings.Contains(prog.Name, " ") {
		dir = strings.ReplaceAll(strings.ToLower(prog.Name), " ", "-")
	} else {
		dir = strings.ToLower(prog.Name)
	}

	prog.Path = path.Join(userHome, strings.ToLower(prog.Language), "src", dir)

	fmt.Printf("Creating %s...\n", prog.Path)

	if err := os.MkdirAll(prog.Path, os.FileMode(0755)); err != nil {
		return err
	}

	fmt.Printf(colors.ColorText("Generated: ", colors.ANSI_GREEN) + "%s\n", prog.Path)

	return nil
}

func (prog *Project) CreateNewProjectSubDirs() error {
	var err error

	if prog.Path == "" {
		return errors.New("Cannot create project subdirectory structure; Project path does not exist\n")
	}

	fmt.Printf("Initializing project structure...\n")

	switch strings.ToLower(prog.Language) {
	case languages.GOLANG:
		err = languages.InitGolangDirs(prog.Path)
	case languages.C:
		err = languages.InitCDirs(prog.Path)
	case languages.JAVA:
		err = languages.InitJavaDirs(prog.Path)
	default:
		fmt.Printf("Nothing additional to create for %s project...\n", prog.Language)
		err = nil
	}

	return err
}

func (prog *Project) CreateProjectFile(filename string) error {
	if prog.Path == "" {
		return errors.New("Project directory does not exist\n")
	}

	f := path.Join(prog.Path, filename)

	switch filename {
	case "README":
		return files.GenerateReadMe(f + languages.MARKDOWN_EXT, prog.Name, prog.Author)
	case "TODO":
		return files.GenerateTODO(f + languages.MARKDOWN_EXT, prog.Name)
	case "main":
		return files.GenerateMain(f, prog.Name, prog.Author, prog.Language)
	default:
		return errors.New("Unknown file " + filename + ", unable to create\n")
	}

	return nil
}
