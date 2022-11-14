package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var (
a *bool
)

func init() {
a = flag.Bool("a", false, "the a value")
flag.Parse()
}

func main() {
	if *a  {
        fmt.Println(*a)
	}
	var str string
	ch := make(chan string)

	go run(ch)

	fmt.Print("можно начать общение в чате: ")
	for {
		fmt.Fscan(os.Stdin, &str)
		ch <- str
		// fmt.Print("Вы ввели: ", str)
	}

}

func handleConn(c net.Conn, ch <- chan string) {
	defer c.Close()

	for {
		message := <- ch
		message = message + "\n\r"
		_, err := io.WriteString(c, message, )
		if err != nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
}

func run(ch <- chan string) {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn,  ch)
	}
}
