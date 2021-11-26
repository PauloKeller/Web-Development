package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	var i int
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			xs := strings.Fields(ln)
			rMethod := xs[0]
			rURI := xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URI:", rURI)
		}

		if ln == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
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

		serve(conn)
	}
}
