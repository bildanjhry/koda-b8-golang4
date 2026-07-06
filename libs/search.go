package libs

func SearchPerson(users []string, name string) []string {
	for x := range len(users) {
		if y := name; y == users[x] {
			return []string{y}
		}
	}
	return []string{""}
}
