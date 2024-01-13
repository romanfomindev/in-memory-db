package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("[spider]> ")
		read, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		if read == "exit\n" {
			return
		}

		fmt.Printf(read)
	}
}
