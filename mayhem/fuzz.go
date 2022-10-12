package fuzz

import "github.com/qaisjp/go-discord-irc"

func mayhemit(bytes []byte) int {
    content := string(bytes)
    var test bridge.Bridge
    test.SetIRCListenerName(content)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}