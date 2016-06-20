// Package common contains subpackages that are shared amongst the mongo
// tools.
package common

import (
	"strings"
)

func SplitNamespace(ns string) (string, string) {
	i := strings.Index(ns, ".")
	if i != -1 {
		return ns[:i], ns[i+1:]
	}
	return "", ns
}
