package main

import (
	"fmt"
	"github.com/bbucko/gobt/torrent"
	"github.com/marksamman/bencode"
)

func main() {
	fmt.Println("Hello, 世界 {}", torrent.New())

	dict := make(map[string]interface{})
	dict["string key"] = "hello world"
	dict["int key"] = 123456
	fmt.Printf("bencode encoded dict: %s\n", bencode.Encode(dict))

}
