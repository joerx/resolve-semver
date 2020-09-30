package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"

	"github.com/Masterminds/semver"
)

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}

	constraint := os.Args[1]
	versions := parseVersions(os.Args[2])

	result, err := FindLatestMatching(constraint, versions)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}

	fmt.Println(result)
}

func parseVersions(raw string) []string {
	return regexp.MustCompile(`\s+`).Split(raw, -1)
}

//
func usage() {
	fmt.Printf("%s <constraint> <versions>\n", os.Args[0])
	fmt.Printf(`E.g. %s "~> 0.12.0" "0.12.1 0.12.2 0.12.3"`+"\n", os.Args[0])
}

// FindLatestMatching given a list of availableVersions returns the latest version matching
// the given constraint
func FindLatestMatching(constraint string, availableVersions []string) (string, error) {
	// Filter list of available versions against constraints
	c, err := semver.NewConstraint(constraint)
	if err != nil {
		return "", err
	}

	compat := make(semver.Collection, 0)

	for _, s := range availableVersions {
		v, err := semver.NewVersion(s)
		if err != nil {
			return "", err
		}
		if c.Check(v) {
			compat = append(compat, v)
		}
	}

	if len(compat) == 0 {
		return "", fmt.Errorf("could not find a version matching constraint %s", constraint)
	}

	// Sort compat versions
	sort.Sort(compat)

	// Return the latest compat version
	return compat[len(compat)-1].String(), nil
}
