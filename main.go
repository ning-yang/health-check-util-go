package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error on getting current user:", err)
	}

	replyFilePath := path.Join(usr.HomeDir, "reply.data")

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for text != "q" { // break the loop if text == "q"
		fmt.Print("Enter reply: ")
		scanner.Scan()
		text = scanner.Text()
		if text != "q" {
			fmt.Println("Writing to", replyFilePath, ":", text)
			ioutil.WriteFile(replyFilePath, []byte(text), 0644)
		}
	}
}
