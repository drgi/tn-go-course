package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:8000")
	if err != nil {
		fmt.Println("Connection failed. Error: ", err)
		return
	}
	defer conn.Close()

	c := make(chan bool, 1)
	querys := make(chan string)

	go copyOutput(c, conn)
	go readInput(c, querys)

	for {
		select {
		case <-c:
			fmt.Println("Connection closed")
			close(querys)
			return
		case q := <-querys:
			_, err := conn.Write([]byte(q))
			if err != nil {
				fmt.Println("Write failed. Error: ", err)
				return
			}

			_, err = conn.Write([]byte("\n"))
			if err != nil {
				fmt.Println("Write failed. Error: ", err)
				return
			}
		}
	}
}

func readInput(c chan bool, queries chan string) {
	for {
		stdInReader := bufio.NewReader(os.Stdin)
		query, _, err := stdInReader.ReadLine()
		if err != nil {
			fmt.Println("Read failed. Error: ", err)
			c <- true
			return
		}
		queries <- string(query)
	}
}

func copyOutput(close chan bool, conn net.Conn) {
	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Println("Read output failed. Error: ", err)
		close <- true
		return
	}
}
