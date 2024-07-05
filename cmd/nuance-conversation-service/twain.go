package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"github.com/bwmarrin/discordgo"
// )

// func main() {
// 	sess, err := discordgo.New(os.Getenv("DISCORD_BOT_TOKEN"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
// 		if m.Author.ID == s.State.User.ID {
// 			return
// 		}

// 		conn, err := connectToServer("localhost:8080")
// 		if err != nil {
// 			fmt.Println("Error connecting:", err)
// 			os.Exit(1)
// 		}
// 		defer conn.Close()

// 		sendAndReceiveMessages(conn, m, s)

// 		// if m.Content == "hello" {
// 		// 	s.ChannelMessageSend(m.ChannelID, "world!")
// 		// }
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

// func sendAndReceiveMessages(conn net.Conn, m *discordgo.MessageCreate, s *discordgo.Session) {
// 	go readMessages(conn, m, s)

// 	// scanner := bufio.NewScanner(os.Stdin)
// 	// fmt.Println("Enter messages to send, type 'exit' to stop:")
// 	// for scanner.Scan() {
// 	// 	text := scanner.Text()
// 	// 	if strings.ToLower(text) == "exit" {
// 	// 		break
// 	// 	}
// 	// 	_, err := fmt.Fprintf(conn, text+"\n")
// 	// 	if err != nil {
// 	// 		fmt.Println("Error sending message:", err)
// 	// 		break
// 	// 	}
// 	// }

// 	fmt.Println("m.Content is: ", m.Content)
// 	_, err := fmt.Fprintf(conn, m.Content+"\n")
// 	if err != nil {
// 		fmt.Println("Error sending message:", err)
// 	}

// 	// if err := scanner.Err(); err != nil {
// 	// 	fmt.Println("Error reading from stdin:", err)
// 	// }
// }

// func readMessages(conn net.Conn, m *discordgo.MessageCreate, s *discordgo.Session) {
// 	reader := bufio.NewReader(conn)
// 	for {
// 		message, err := reader.ReadString('\n')
// 		fmt.Println("readMessages message is: ", message)
// 		if err != nil {
// 			fmt.Println("Error reading from connection:", err)
// 			return
// 		}
// 		fmt.Print("Received from server: ", message) // delete l8r
// 		s.ChannelMessageSend(m.ChannelID, message)
// 	}
// }
