package equal

import "strings"

func Equal(s, t []string) bool {
	return strings.Join(s, "") == strings.Join(t, "")
}