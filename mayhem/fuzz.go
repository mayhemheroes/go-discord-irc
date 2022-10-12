package fuzz

import "github.com/qaisjp/go-discord-irc/irc/format"

func mayhemit(bytes []byte) int {
    content := string(bytes)

    ircf.Parse(content)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}