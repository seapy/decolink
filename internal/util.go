package internal

import "fmt"

// ToLink ...
func ToLink(t string, s string, url string) string {
	if s != "" {
		t = fmt.Sprintf("%s - %s", t, s)
	}
	return fmt.Sprintf("[%s](%s)", t, url)
}
