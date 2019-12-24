package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	// request line
	m := strings.Fields(ln)[0] // method
	u := strings.Fields(ln)[1] // uri
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", u)

	// multiplexer
	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/user" {
		user(conn)
	}
	if m == "GET" && u == "/users" {
		users(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>INDEX</strong><br>
	<a href="/">index</a><br>
	<a href="/user">user</a><br>
	<a href="/users">users</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func user(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<h1>USER</h1><br>
	<p>Paulo</p>
	<p>Keller</p>
	<a href="/">index</a><br>
	<a href="/user">user</a><br>
	<a href="/users">users</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func users(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<h1>USERS</h1><br>
	<ul>
		<li>Paulo Keller</li>
		<li>User 2</li>
		<li>User 3</li>
	</ul>
	
	<a href="/">index</a><br>
	<a href="/user">user</a><br>
	<a href="/users">users</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
