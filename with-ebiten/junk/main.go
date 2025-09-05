package main

import "some/package/junk"

func main() {
	j := junk.NewJunk()
	junk.PublicFunc(*j)
}
