package main

import (
    "os"
    "log"
    "strings"
    "github.com/brettbuddin/victor"
)

func main() {
    account := os.Getenv("CAMPFIRE_ACCOUNT")
    token   := os.Getenv("CAMPFIRE_TOKEN")
    rooms   := os.Getenv("CAMPFIRE_ROOMS")

    if account == "" || token == "" || rooms == "" {
        log.Panic("Please set CAMPFIRE_ACCOUNT, CAMPFIRE_TOKEN and CAMPFIRE_ROOMS")
    }

    roomIdStrings := strings.Split(rooms, ",")
    roomsArr      := make([]int, 0)

    for _, id := range roomIdStrings {
        j, _ := strconv.Atoi(id)
        roomsArr = append(roomsArr, j) 
    }

    r := victor.NewCampfire("victor", account, token, roomsArr)

    r.Hear("derp", func(msg *victor.TextMessage) {
        msg.Send("Derp!")
    })

    r.Run()
}