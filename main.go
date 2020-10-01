package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joerx/resolve-semver/pkg/semver"
)

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}

	constraint := os.Args[1]
	versions := parseVersions(os.Args[2])

	result, err := semver.FindLatestMatching(constraint, versions)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}

	fmt.Println(result)
}

func parseVersions(raw string) []string {
	return regexp.MustCompile(`\s+`).Split(raw, -1)
}

func usage() {
	fmt.Printf("%s <constraint> <versions>\n", os.Args[0])
	fmt.Printf(`E.g. %s "~> 0.12.0" "0.12.1 0.12.2 0.12.3"`+"\n", os.Args[0])
}
