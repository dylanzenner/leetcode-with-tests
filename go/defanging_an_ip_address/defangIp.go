package defangIp

import "strings"

func DefangIp(s string) string{
	return strings.ReplaceAll(s, ".", "[.]")
}