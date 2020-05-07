package main

import (
	"fmt"

	"github.com/nlopes/slack"
)


func main() {
    TOKEN := "xoxb-2154537752-944237489669-RZZ81t8Rt8JvC98aZ5A4tgsm"
    //TOKEN := "xoxp-2154537752-573993009606-928783508514-cabd0260729758fa9ae9f3de33d977e8"
    USERGROUP := "STW6TB5H6"
    api := slack.New(TOKEN)
    fmt.Println("ADDING USER hemanth TO SLACK CHANNEL: CTK8YGX4Y")
    // _, err := api.InviteUserToChannel("CTK8YGX4Y", "shemant")
    user, err := api.GetUserByEmail("arathore@vmware.com")
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
