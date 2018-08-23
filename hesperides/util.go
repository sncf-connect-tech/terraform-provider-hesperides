package hesperides

import (
	"fmt"
	"strings"
)

// return the pieces of id `a-b` as a, b
func parseTwoPartID(id string) (string, string) {
	parts := strings.SplitN(id, "-", 2)
	return parts[0], parts[1]
}

// format the strings into an id `a-b`
func buildTwoPartID(a, b *string) string {
	return fmt.Sprintf("%s-%s", *a, *b)
}

// return the pieces of id `a-b-c` as a, b, c
func parseThreePartID(id string) (string, string, string) {
	parts := strings.SplitN(id, "-", 3)
	return parts[0], parts[1], parts[2]
}

// format the strings into an id `a-b-c`
func buildThreePartID(a, b, c *string) string {
	return fmt.Sprintf("%s-%s-%s", *a, *b, *c)
}

const (
	WorkingCopy = "workingcopy"
	Release     = "release"
)
