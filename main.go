package main

import "fmt"

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
	builtBy = "unknown"
)

func main() {
	fmt.Printf("my app %s, Commit %s, built at %s by %s", Version, Commit, Date, builtBy)
}
