package libs

import "strings"

func SearchPerson(users []string, name string) []string {
	for x := range users {
		if y := strings.ToLower(name); y == strings.ToLower(users[x]) {
			return []string{users[x]}
		}
	}
	return []string{""}
}
