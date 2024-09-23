package utils

import (
	"fmt"
	"strings"
)

type RESPCommands struct {
	SET      string
	GET      string
	LPUSH    string
	RPUSH    string
	LPOP     string
	RPOP     string
	DEL      string
	INCR     string
	DECR     string
	PING     string
	LRANGE   string
	EXISTS   string
	FLUSHALL string
	EXPIRE   string
	TTL      string
	PERSIST  string
}

var respCmds = RESPCommands{
	SET:      "SET",
	GET:      "GET",
	LPUSH:    "LPUSH",
	RPUSH:    "RPUSH",
	LPOP:     "LPOP",
	RPOP:     "RPOP",
	DEL:      "DEL",
	INCR:     "INCR",
	DECR:     "DECR",
	PING:     "PING",
	LRANGE:   "LRANGE",
	EXISTS:   "EXISTS",
	FLUSHALL: "FLUSHALL",
	EXPIRE:   "EXPIRE",
	TTL:      "TTL",
	PERSIST:  "PERSIST",
}

type ReplyType struct {
	Int    rune
	Bulk   rune
	Error  rune
	Array  rune
	Status rune
}

var replyType ReplyType = ReplyType{
	Int:    ':',
	Bulk:   '$',
	Array:  '*',
	Error:  '-',
	Status: '+',
}

var clientMessage string = fmt.Sprintf(
	"Please use these following commands:\n%s\n%s\n%s\n",
	strings.Join([]string{
		respCmds.PING, respCmds.SET, respCmds.GET, respCmds.EXISTS,
	}, ", "),
	strings.Join([]string{
		respCmds.FLUSHALL, respCmds.DEL, respCmds.INCR, respCmds.DECR,
	}, ", "),
	strings.Join([]string{
		respCmds.LPUSH, respCmds.LPOP, respCmds.LRANGE, respCmds.RPUSH, respCmds.RPOP,
	}, ", "),
)
