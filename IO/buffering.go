package io

import (
	"bufio"
	"os"
)

const MAX_WRITE = 1_000_000
func Bad() {
	f , err := os.OpenFile("file.bin", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for range MAX_WRITE {
		f.WriteString("sfdlkjhdsfkjhshdfhkjasd")
	}
}
func Good(){
	f , err := os.OpenFile("file.bin", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	for range MAX_WRITE {
		writer.WriteString("sfdlkjhdsfkjhshdfhkjasd")
	}
}
