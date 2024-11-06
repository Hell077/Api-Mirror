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
	path := flag.String("path", "", "path to file to mirror")
	port := flag.Int("port", 0, "Порт для запуска сервера (0 для случайного)")

	flag.Parse()

	config, err := parser.ParseYAML(*path)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if err := generator.Generator(config, "api_documentation.html"); err != nil {
		fmt.Println("Ошибка генерации HTML:", err)
		return
	}

	if *port == 0 {
		var err error
		*port, err = server.FindFreePort()
		if err != nil {
			fmt.Println("Ошибка нахождения свободного порта:", err)
			return
		}
	}

	htmlContent, err := os.ReadFile("api_documentation.html")
	if err != nil {
		fmt.Println("Ошибка чтения HTML файла:", err)
		return
	}

	srv := &server.Server{}
	if err := srv.StartServer(*port, string(htmlContent)); err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
		return
	}

	srv.Wait()
}
