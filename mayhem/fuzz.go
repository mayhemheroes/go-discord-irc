package fuzz

import "github.com/qaisjp/go-discord-irc/dstate"
import "github.com/bwmarrin/discordgo"

func mayhemit(bytes []byte) int {
    content := string(bytes)
    var test discordgo.Session
    dstate.ChannelMessage(&test, content, content)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}