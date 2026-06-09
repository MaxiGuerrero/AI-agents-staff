package main

import "fmt"

func main() {
	var config = LoadConfig()
	fmt.Printf("token: %s", config.Telegram.Token)
}
