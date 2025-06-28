package client

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func StartREPL() {
    var conn net.Conn
    var err error

    fmt.Println(">> Welcome to TCP Client")
    reader := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print(">> ")
        if !reader.Scan() {
            break
        }

        input := reader.Text()
        parts := strings.SplitN(input, " ", 2)
        cmd := strings.ToLower(parts[0])

        switch cmd {
        case "connect":
            if len(parts) < 2 {
                fmt.Println("Usage: connect <host:port>")
                continue
            }
            conn, err = net.Dial("tcp", parts[1])
            if err != nil {
                fmt.Println("Connection failed:", err)
                continue
            }
            fmt.Println("Connected to", parts[1])

        case "send", "ping", "time", "exit":
            if conn == nil {
                fmt.Println("Not connected.")
                continue
            }

            var toSend string
            if cmd == "send" {
                if len(parts) < 2 {
                    fmt.Println("Usage: send <message>")
                    continue
                }
                toSend = "ECHO " + parts[1]
            } else {
                toSend = strings.ToUpper(cmd)
            }

            fmt.Fprintf(conn, toSend+"\n")
            reply := make([]byte, 4096)
            n, _ := conn.Read(reply)
            fmt.Println("<<", string(reply[:n]))

            if cmd == "exit" {
                conn.Close()
                return
            }

        default:
            fmt.Println("Commands: connect, send, ping, time, exit")
        }
    }
}
