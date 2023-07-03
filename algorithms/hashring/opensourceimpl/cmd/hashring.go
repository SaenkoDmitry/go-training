package main

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/serialx/hashring"
)

var memcachedServers = map[string]string{
	"192.168.0.246:11212": "A",
	"192.168.0.247:11212": "B",
	"192.168.0.249:11212": "C",
}

func main() {
	memcachedServersList := lo.MapToSlice(memcachedServers, func(key string, value string) string {
		return key
	})

	// print source data
	fmt.Println("\nserver ring:")
	for k, v := range memcachedServers {
		fmt.Println(v, "=", k)
	}

	// initialize hash ring struct
	ring := hashring.New(memcachedServersList)

	// find server for key
	key := "my_key"
	server, _ := ring.GetNode(key)
	fmt.Println("\nkey:", key, "=>", "server:", memcachedServers[server])
}
