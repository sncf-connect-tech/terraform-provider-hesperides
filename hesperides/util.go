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

const (
	WorkingCopy = "workingcopy"
	Release     = "release"
)
