package server

import (
    "fmt"
    "sync/atomic"
    "time"
)

var totalConnections int32
var totalBytes int32
var totalCommands int32

func IncrementConnections() {
    atomic.AddInt32(&totalConnections, 1)
}

func UpdateBytes(n int) {
    atomic.AddInt32(&totalBytes, int32(n))
}

func IncrementCommands() {
    atomic.AddInt32(&totalCommands, 1)
}

func PrintMetrics() {
    for {
        time.Sleep(30 * time.Second)
        fmt.Println("\n--- Metrics ---")
        fmt.Printf("Connections: %d\n", totalConnections)
        fmt.Printf("Bytes Received: %d\n", totalBytes)
        fmt.Printf("Commands Handled: %d\n\n", totalCommands)
    }
}