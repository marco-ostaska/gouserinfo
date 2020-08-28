package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {

	a := os.Args
	fmt.Println(len(a))
	id, err := user.Lookup(a[1])

	if err != nil {
		fmt.Println(err, "FDP")
		os.Exit(2)
	}
	// id, _ := user.LookupId(a[1])

	fmt.Println("username: ", id.Username)
	fmt.Println("Name:     ", id.Name)
	fmt.Println("homeDir:  ", id.HomeDir)
	fmt.Println("uid:      ", id.Uid)
	fmt.Println("gid:      ", id.Gid)

	fmt.Println(strings.Repeat("-", 30))
	fmt.Print("Groups:    ")

	g, _ := id.GroupIds()
	var gs string

	for _, v := range g {
		a4, _ := user.LookupGroupId(v)
		gs += a4.Name + ":" + a4.Gid + " "

	}

	fmt.Println(gs)

}
