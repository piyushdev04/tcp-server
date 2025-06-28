package server

import (
    "bufio"
    "net"
    "strings"
)

func Start() {
    InitLogger()
    go PrintMetrics()

    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        logFatal("Failed to listen: %v", err)
    }
    logInfo("Server started on :8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            logError("Failed to accept: %v", err)
            continue
        }
        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    defer conn.Close()
    addr := conn.RemoteAddr().String()
    logInfo("Connected: %s", addr)
    IncrementConnections()

    reader := bufio.NewReader(conn)
    for {
        input, err := reader.ReadString('\n')
        if err != nil {
            logInfo("Disconnected: %s", addr)
            return
        }
        input = strings.TrimSpace(input)
        UpdateBytes(len(input))
        logCommand(addr, input)
        response := ParseCommand(input)
        conn.Write([]byte(response))
        if strings.ToUpper(input) == "EXIT" {
            return
        }
    }
}