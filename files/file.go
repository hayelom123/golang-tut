package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating a file ", err)
		return
	}

	defer file.Close()

	fmt.Println("File created:", file.Name())

	file.WriteString("Jesus christ is God!!!!\n")
	file.WriteString("In the beginning was the Word, and the Word was with God, and the Word was God.(john 1:1)\n")
	file.WriteString("He was with God in the beginning.Through him all things were made; without him nothing was made that has been made.\n")
	file.WriteString("In him was life, and that life was the light of all mankind.\nThe light shines in the darkness, and the darkness has not overcome[a] it\n")
	file.WriteString("14 The Word became flesh and made his dwelling among us. We have seen his glory, \nthe glory of the one and only Son, who came from the Father, full of grace and truth.\n")

	fmt.Println("The first file is done.")

	// 1. Open the file with specific flags:
	// os.O_CREATE: Create the file if it doesn't exist
	// os.O_WRONLY: Open the file for writing only
	// os.O_APPEND: Append data to the file instead of overwriting it (use os.O_TRUNC to overwrite)
	// 0644: Standard file permissions (readable by everyone, writable by owner)
	myfile, errr := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)

	if errr != nil {
		fmt.Println("Error creating/opening  a file ", errr)
		return
	}

	defer myfile.Close()

	if _, errrr := myfile.WriteString("Hello from Go!!!\n"); errrr != nil {
		fmt.Println("Failed to write:", errrr)
		return
	}

	// Read from the beginning of the file.
	if _, err := myfile.Seek(0, io.SeekStart); err != nil {
		fmt.Println("Failed to seek:", err)
		return
	}

	content, err := io.ReadAll(myfile)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	fmt.Println("Content:")
	fmt.Println(string(content))

}
