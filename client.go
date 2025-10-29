package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("❌ Error connecting to server:", err)
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🌟 Welcome to Amazing Chatroom 🌟")
	fmt.Println("Type your messages and press Enter to send.")
	fmt.Println("Type 'exit' anytime to leave the chat.\n")

	fmt.Print("🧑 Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// 🎉 Show welcome message after entering the name
	fmt.Printf("\n🎉 Welcome, %s! You’ve joined the chatroom. Start talking below ⬇️\n\n", name)

	for {
		fmt.Print("💬 You: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("👋 Goodbye! Thanks for chatting ❤️")
			break
		}

		fullMsg := fmt.Sprintf("%s: %s", name, text)

		var chatHistory []string
		err = client.Call("ChatServer.SendMessage", fullMsg, &chatHistory)
		if err != nil {
			log.Println("⚠️ Error calling RPC:", err)
			break
		}

		fmt.Println("\n===== 🗨️ Chat History =====")
		for _, msg := range chatHistory {
			fmt.Println(msg)
		}
		fmt.Println("===========================\n")
	}
}
