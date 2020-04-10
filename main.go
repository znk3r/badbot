package main

import (
	"fmt"
)

// Version is the build version
var Version string

// GitTag is the git tag of the build
var GitTag string

// BuildDate is the date when the build was created
var BuildDate string

func main() {
	fmt.Printf("Badbot version %s, build %s %s\n", Version, GitTag, BuildDate)
}