package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Fprint(conn, " ")

		go func() {
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				ln := scanner.Text()
				// fmt.Fprintf(conn, "Someone says: %s\n", ln)
				// fmt.Printf("Someone says: %s\n", ln)
				io.WriteString(conn, "Someone says: "+ln+"\n")
			}

			// io.Copy(conn, conn)

			conn.Close()
		}()

	}

}
