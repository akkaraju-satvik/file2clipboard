package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"golang.design/x/clipboard"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("\nExiting the clipboard writer")
			os.Exit(0)
		}
	}()
	fileName := os.Args[1]
	if fileName == "" {
		log.Fatal("Please provide a file name")
	}
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	err = clipboard.Init()
	if err != nil {
		panic(err)
	}
	changed := clipboard.Write(clipboard.FmtText, file)
	fmt.Println("Copied file contents to clipboard! Press Ctrl+C after pasting the contents to exit the program.")
	x := <-changed
	fmt.Println(x)
}
