package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
	}
	body := ""
	io.WriteString(c, "HTTP/1.1 200 OK\n")
	fmt.Fprintf(c, "Content-Length: %d\n", len(body))
	fmt.Fprint(c, "Content-Type: text/plain\n")
	io.WriteString(c, body)
	fmt.Println("Code got here.")
	io.WriteString(c, "I see you connected.")
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(conn)
	}
}
