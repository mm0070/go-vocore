package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mm0070/go-vocore/vocore"
)

func main() {
	// Set up display
	display, err := vocore.InitializeScreen()
	if err != nil {
		log.Fatal("Failed to initialize screen: %v", err)
	}
	defer display.Close()

	// Open the image file
	file, err := os.Open("example.png")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Read the entire file into a byte slice
	img, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}
	err = display.WriteToScreen(img)
	if err != nil {
		log.Fatal("Screen write failed: %s", err)
	}
}
