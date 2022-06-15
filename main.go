package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if nil != err {
		log.Println(err)
	}

	// go func(conn net.Conn) {
	// 	send := "hello"
	// 	for {
	// 		_, err = conn.Write([]byte(send))
	// 		if err != nil {
	// 			log.Println("Faild to write data : ", err)
	// 			break
	// 		}

	// 		time.Sleep(1 * time.Second)
	// 	}
	// }(conn)

	go func() {
		data := make([]byte, 4096)

		for {
			n, err := conn.Read(data)
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("Server send : " + string(data[:n]))
			time.Sleep(time.Duration(3) * time.Second)
		}
	}()

	for {
		var s string
		fmt.Scanln(&s)
		conn.Write([]byte(s))
		time.Sleep(time.Duration(3) * time.Second)
	}
}
