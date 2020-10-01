package semver

import (
	"fmt"
	"sort"

	sv "github.com/Masterminds/semver"
)

// FindLatestMatching given a list of availableVersions returns the latest version matching
// the given constraint
func FindLatestMatching(constraint string, availableVersions []string) (string, error) {
	// Filter list of available versions against constraints
	c, err := sv.NewConstraint(constraint)
	if err != nil {
		return "", err
	}

	compat := make(sv.Collection, 0)

	for _, s := range availableVersions {
		v, err := sv.NewVersion(s)
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
