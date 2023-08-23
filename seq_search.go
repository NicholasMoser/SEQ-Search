package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	version := "1.0.0"
	fmt.Printf("SEQ Search %s\n", version)

	// Get user input
	fmt.Print("Bytes in hex: ")
	var hexBytes string
	in := bufio.NewReader(os.Stdin)
	hexBytes, err := in.ReadString('\n')
	hexBytes = strings.Replace(hexBytes, "0x", "", -1)
	hexBytes = strings.Replace(hexBytes, " ", "", -1)
	hexBytes = strings.TrimSpace(hexBytes)
	bytePattern, err := hex.DecodeString(hexBytes)
	if err != nil {
		fmt.Println("Invalid hex bytes")
		exit(1)
	}

	fmt.Print("SEQ files directory: ")
	var dir string
	in = bufio.NewReader(os.Stdin)
	dir, err = in.ReadString('\n')
	dir = strings.TrimSpace(dir)
	fileInfo, err := os.Stat(dir)
	if err != nil {
		fmt.Printf("%s is not a valid path\n", dir)
		exit(1)
	}
	if !fileInfo.IsDir() {
		fmt.Printf("%s is not a valid directory\n", dir)
		exit(1)
	}
	fmt.Println()

	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(path, ".seq") {
			searchFile(path, bytePattern)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Impossible to walk directories: %s", err)
		exit(1)
	}
	exit(0)
}

func searchFile(path string, bytePattern []byte) {
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Failed to read bytes from file %s\n", path)
	}
	index := bytes.Index(dat, bytePattern)
	if index != -1 {
		fmt.Printf("Offset 0x%X in %s\n", index, path)
	}
}

func exit(code int) {
	fmt.Println("\nPress enter to exit...")
	var output string
	fmt.Scanln(&output)
	os.Exit(code)
}
