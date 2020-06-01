package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type techList []string

var (
	location  string
	help      bool
	title     string
	intro     string
	techStack techList
)

//String return a string representation of the techList
func (t *techList) String() string {
	s := "## Technologies Used\n\n"
	for _, tech := range techStack {
		s += fmt.Sprintf("- %s\n", tech)
	}
	return s
}

//Set add a new technology to the techList
func (t *techList) Set(val string) error {
	for _, s := range strings.Split(val, ", ") {
		*t = append(*t, s)
	}
	return nil
}

func setFlags() {
	flag.BoolVar(&help, "h", false, "print this help message (shorthand)")
	flag.BoolVar(&help, "help", false, "print this help message")
	flag.StringVar(&title, "t", "# Title", "the title of the project (shorthand)")
	flag.StringVar(&title, "title", "# Title", "the title of the project")
	flag.StringVar(&intro, "intro", "## Intro to be added", "the intro/explanation of the project")
	flag.Var(&techStack, "s", "a list of technologies used on the project (shorthand)\nSeparate multiple values with \", \"")
	flag.Var(&techStack, "stack", "a list of technologies used on the project\nSeparate multiple values with \", \"")
}

func printHelp() {
	fmt.Println("Usage of ./rewrite-me [location]:")
	flag.PrintDefaults()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeFile(path string) {
	f, err := os.Create(path)
	check(err)
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%s\n\n%s\n\n%s\n", title, intro, techStack.String()))
	check(err)
	_, err = f.WriteString("This README file was created by ReWrite-Me. Learn more [here](https://github.com/c4llmeco4ch/ReWrite-Me)")
	f.Sync()
}

func main() {
	setFlags()
	flag.Parse()
	if help {
		printHelp()
	}
	filePath := "./README.md"
	if len(flag.Args()) == 1 {
		filePath = flag.Args()[0]
	} else if len(flag.Args()) > 1 {
		fmt.Println("Invalid file path")
		printHelp()
		os.Exit(2)
	}
	if title[0] != '#' {
		title = "# " + title
	}
	writeFile(filePath)
}
