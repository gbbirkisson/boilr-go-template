package main

import (
	"fmt"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Println(info())
	fmt.Println(message())
}

func info() string {
	return fmt.Sprintf("{{ProjectName}}\nversion\t%v\ncommit\t%v\ndate\t%v", version, commit, date)
}

func message() string {
	return "Hello, template"
}
