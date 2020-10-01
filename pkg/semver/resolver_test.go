package semver

import "testing"

var versions = []string{"0.11.0", "0.11.1", "0.12.1", "0.12.2", "0.12.3", "0.12.4-alpha1", "0.13.1", "0.13.2"}
var constraints = map[string]string{
	">= 0.12, < 0.13": "0.12.3",
	">= 0.12":         "0.13.2",
	"~> 0.12":         "0.12.3",
	"~> 0.12-0":       "0.12.4-alpha1",
	"< 0.12":          "0.11.1",
	"0.12.2":          "0.12.2",
}

func TestFindLatestMatching(t *testing.T) {
	for constraint, expected := range constraints {
		result, err := FindLatestMatching(constraint, versions)
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	}
}

func TestCannotFindMatchingVersion(t *testing.T) {
	result, err := FindLatestMatching(">= 1.0", versions)
	if err == nil {
		t.Errorf("Expected error but got result '%s'", result)
	}
}
