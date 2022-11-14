package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("Введите свой ник")
	nickname := ""
	fmt.Scanln(&nickname)

	for {
		buf := make([]byte, 256) // создаем буфер
		_, err = conn.Read(buf)
		if err == io.EOF {
			break
		}

		_, err = io.WriteString(os.Stdout, fmt.Sprintf("%s : %s", nickname, string(buf)))
		// выводим измененное сообщение сервера в консоль
		if err != nil {
			break
		}
	}
}
