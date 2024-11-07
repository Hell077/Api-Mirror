package main

import (
	"flag"
	"fmt"
	"github.com/Hell077/Api-Mirror/internal/generator"
	"github.com/Hell077/Api-Mirror/internal/parser"
	"github.com/Hell077/Api-Mirror/internal/server"
	"os"
)

func main() {
	path := flag.String("path", "", "path to the file to mirror")
	port := flag.Int("port", 0, "Port to run the server (0 for a random one)")

	flag.Parse()

	config, err := parser.ParseYAML(*path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := generator.Generator(config, "api_documentation.html"); err != nil {
		fmt.Println("Error generating HTML:", err)
		return
	}

	if *port == 0 {
		var err error
		*port, err = server.FindFreePort()
		if err != nil {
			fmt.Println("Error finding a free port:", err)
			return
		}
	}

	htmlContent, err := os.ReadFile("api_documentation.html")
	if err != nil {
		fmt.Println("Error reading HTML file:", err)
		return
	}

	srv := &server.Server{}
	if err := srv.StartServer(*port, string(htmlContent)); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	srv.Wait()
}
