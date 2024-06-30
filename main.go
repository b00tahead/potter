package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/b00tahead/potter/internal"
)

func main() {
	// Command-line flags
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	questionCmd := flag.NewFlagSet("question", flag.ExitOnError)

	// Subcommands
	if len(os.Args) < 2 {
		fmt.Println("Potter AI CLI")
		fmt.Println("Usage:")
		fmt.Println("	potter init <your_api_key> - Initialized Potter with your API key")
		fmt.Println("	potter question <your question> - Ask a question to the Potter AI")
		return
	}

	switch os.Args[1] {
	case "init":
		initCmd.Parse(os.Args[2:])
		initializePotter(initCmd.Args())
	case "question":
		questionCmd.Parse(os.Args[2:])
		askQuestion(questionCmd.Args())
	default:
		fmt.Println("Unknown command:", os.Args[1])
		return
	}
}

func initializePotter(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: potter init <your_api_key>")
		return
	}

	apiKey := args[0]

	err := internal.SaveAPIKey(apiKey)
	if err != nil {
		fmt.Println("Error initializing Potter AI")
		return
	}

	fmt.Println("Potter AI initialized successfully.")
}

func askQuestion(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: potter question <your_question>")
		return
	}

	question := args[0]

	response, err := internal.GetChatGPTResponse(question)
	if err != nil {
		fmt.Println("Error getting response from ChatGPT:", err)
		return
	}

	fmt.Println(response)
}
