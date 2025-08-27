package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./messages.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	counterChannel := getLinesChannel(f)
	for i := range counterChannel {
		fmt.Printf("read: %s\n", i)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		defer f.Close()
		defer close(ch)
		buf := make([]byte, 8)
		strBuf := ""
		for {
			res, err := f.Read(buf)
			if err != nil {
				if err == io.EOF {
					return
				}
				fmt.Printf("errors: %s\n", err.Error())
			}
			strs := strings.Split(string(buf[:res]), "\n")
			if len(strs) == 1 {
				strBuf += strs[0]
			}
			if len(strs) == 2 {
				strBuf += strs[0]
				ch <- strBuf
				strBuf = ""
				strBuf += strs[1]
			}
		}
	}()
	return ch
}
