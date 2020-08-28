package main

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
)

const version string = "0.0.1"

type userInfo struct {
	// get uinfo from user.User
	uinfo *user.User
	ginfo string
}

func main() {
	a, err := checkArgs(os.Args)
	errCheck(err, 2)

	u, err := grabUserInfo(checkUserArg(a))
	errCheck(err, 2)

	var ui userInfo
	ui.uinfo = u
	display(grabUserGroups(ui))
}

// checkArgs check if it is not empty
func checkArgs(a []string) (string, error) {

	switch {
	case len(a) <= 1:
		u, err := user.Current()
		errCheck(err, 2)
		return u.Username, nil
	case len(a) > 2:
		usage()
		return "", errors.New("Too many args")
	case a[1] == "-h" || a[1] == "--help":
		usage()
		os.Exit(0)
		return "", nil
	case a[1] == "-v" || a[1] == "--version":
		fmt.Println("version:", version)
		os.Exit(0)
		return version, nil
	default:
		return a[1], nil
	}

}

// checkUserArg checks if id passed is username or uid
func checkUserArg(u string) (string, string) {
	_, err := strconv.ParseInt(u, 10, 64)

	if err != nil {
		return u, "username"
	}
	return u, "uid"
}

// grabUserInfo get user information based on checkUserArg returns
func grabUserInfo(u string, uType string) (*user.User, error) {
	if uType == "uid" {
		return user.LookupId(u)
	}
	return user.Lookup(u)

}

// errCheck in case of error it exist with provide error code ex
func errCheck(e error, ex int) {
	if e != nil {
		fmt.Println(e)
		os.Exit(ex)
	}
}

// grabUserGroups grabe the user groups for given ID
func grabUserGroups(ui userInfo) userInfo {

	g, err := ui.uinfo.GroupIds()
	errCheck(err, 4)

	for _, v := range g {
		ug, err := user.LookupGroupId(v)
		errCheck(err, 4)
		ui.ginfo += ug.Name + ":" + ug.Gid + " "
	}

	return ui
}

// display shows formated output
func display(u userInfo) {

	fmt.Println("username: ", u.uinfo.Username)
	fmt.Println("Name:     ", u.uinfo.Name)
	fmt.Println("homeDir:  ", u.uinfo.HomeDir)
	fmt.Println("uid:      ", u.uinfo.Uid)
	fmt.Println("gid:      ", u.uinfo.Gid)

	fmt.Println(strings.Repeat("-", 30))
	fmt.Print("Groups:    ")
	fmt.Println(u.ginfo)

}

// usage Well, usage only
func usage() {

	ut := `Print user and group information for the specified USER
or (when USER omitted) for the current user.

  -v, --version     output version information and exit
  -h, --help        display this help and exit
	`

	fmt.Printf("Usage: %s [OPTION]... [USER]\n", os.Args[0])
	fmt.Printf("%s \n", ut)

}
