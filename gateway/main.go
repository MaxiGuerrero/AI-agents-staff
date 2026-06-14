package main

import (
	"fmt"
	"os"
	"os/signal"

	config "github.com/MaxiGuerrero/AI-agents-staff/gateway/config"
	handlers "github.com/MaxiGuerrero/AI-agents-staff/gateway/handlers"
	messageplatform "github.com/MaxiGuerrero/AI-agents-staff/gateway/message-platform"
)

func main() {
	// Load configuration and print the Telegram token to verify that it is being read correctly
	var conf = config.LoadConfig()
	handlers := handlers.GetHandlers()
	bot := messageplatform.NewTelegramBot(conf.Telegram.Token, generateSecurityToken())
	bot.InitHandlers(handlers)
	bot.Start()
	// Initzialize shutdown handler to gracefully exit the application when an interrupt signal is received
	waitForShutdown(bot)
}

// waitForShutdown waits for an interrupt signal and performs cleanup before exiting
// Signal ctrl + c to trigger the shutdown process
func waitForShutdown(bot messageplatform.MessagePlatform) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	<-sig
	bot.Stop()
	fmt.Println("Shutdown...")
}
