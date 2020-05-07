package main

import (
	"fmt"

	"github.com/nlopes/slack"
)


func main() {
    api := slack.New(TOKEN)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(user.ID)
    }
    //UpdateUserGroupMembers
    _, err = api.UpdateUserGroupMembers(USERGROUP, user.ID)
    if err != nil {
        fmt.Println(err)
    }
}
