package main

import (
	"flag"
	"fmt"
	"os"
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
	s := "### Technologies Used\n"
	for _, tech := range techStack {
		s += fmt.Sprintf("- %s\n", tech)
	}
	return s
}

//Set add a new technology to the techList
func (t *techList) Set(val string) error {
	*t = append(*t, val)
	return nil
}

func setFlags() {
	flag.BoolVar(&help, "h", false, "print this help message (shorthand)")
	flag.BoolVar(&help, "help", false, "print this help message")
	flag.StringVar(&title, "t", "# Title", "the title of the project (shorthand)")
	flag.StringVar(&title, "title", "# Title", "the title of the project")
	flag.StringVar(&intro, "intro", "### Intro to be added", "the intro/explanation of the project")
	flag.Var(&techStack, "s", "provide a list of technologies used on the project (shorthand)")
	flag.Var(&techStack, "stack", "provide a list of technologies used on the project")
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
	_, err = f.WriteString(fmt.Sprintf("%s\n\n%s\n\n%s", title, intro, techStack.String()))
	check(err)
	f.Sync()
}

func main() {
	setFlags()
	flag.Parse()
	if help {
		printHelp()
	}
	filePath := "./README.md"
	if len(flag.Args()) != 0 {
		filePath = flag.Args()[0]
	}
	if title[0] != '#' {
		title = "# " + title
	}
	writeFile(filePath)
}
