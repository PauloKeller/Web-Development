package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func handleIndex() string {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<meta charset="UTF-8">
	<title>GET INDEX</title>
	<body>
		<h1>"GET INDEX"</h1>
		<a href="/">index</a><br>
		<a href="/apply">apply</a><br>
	</body>
	</html>
`
	return body
}

func handleApply() string {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<meta charset="UTF-8">
	<title>POST APPLY</title>
	<body>
			<h1>"GET APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			<form action="/apply" method="POST">
			<input type="hidden" value="In my good death">
			<input type="submit" value="submit">
			</form>
		</body>
	</html>
`
	return body
}

func handleApplyPost() string {
	body := `
	<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>POST APPLY</title>
		</head>
		<body>
			<h1>"POST APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
	</html>
`
	return body
}

func handleDefault() string {
	body := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>default</title>
</head>
<body>
	<h1>"default"</h1>
</body>
</html>
`
	return body
}

func router(c net.Conn, method string, uri string) {
	var body string

	switch {
	case method == "GET" && uri == "/":
		body = handleIndex()
	case method == "GET" && uri == "/apply":
		body = handleApply()
	case method == "POST" && uri == "/apply":
		body = handleApplyPost()
	default:
		body = handleDefault()
	}

	io.WriteString(c, "HTTP/1.1 200 OK\n")
	fmt.Fprintf(c, "Content-Length: %d\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\n")
	io.WriteString(c, body)
	fmt.Println("Code got here.")
	io.WriteString(c, "I see you connected.")
}

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	var i int
	var rMethod string
	var rURI string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URI:", rURI)
		}

		if ln == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
	}
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
