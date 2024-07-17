package main

// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// 	"os/signal"
// 	"strings"
// 	"syscall"

// 	"github.com/bwmarrin/discordgo"
// )

// var conn net.Conn

// func main() {
// 	// Connect to the TCP server
// 	var err error
// 	conn, err = connectToServer("localhost:8080")
// 	if err != nil {
// 		fmt.Println("Error connecting:", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close()

// 	// Initialize Discord session
// 	sess, err := discordgo.New(os.Getenv("DISCORD_BOT_TOKEN"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Add message handler
// 	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
// 		// Ignore messages from the bot itself
// 		if m.Author.ID == s.State.User.ID {
// 			return
// 		}

// 		// Check for the !send command prefix
// 		if strings.HasPrefix(m.Content, "!send ") {
// 			msg := strings.TrimPrefix(m.Content, "!send ")
// 			sendMessageToTCPServer(msg)
// 			s.ChannelMessageSend(m.ChannelID, "Message sent to TCP server!")
// 		}
// 	})

// 	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

// 	err = sess.Open()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer sess.Close()

// 	fmt.Println("The bot is online!")

// 	sc := make(chan os.Signal, 1)
// 	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
// 	<-sc
// }

// func connectToServer(address string) (net.Conn, error) {
// 	return net.Dial("tcp", address)
// }

// func sendMessageToTCPServer(message string) {
// 	_, err := fmt.Fprintf(conn, message+"\n")
// 	if err != nil {
// 		fmt.Println("Error sending message:", err)
// 	}
// }
