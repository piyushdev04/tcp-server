package server

import (
    "strings"
    "time"
)

func ParseCommand(input string) string {
    parts := strings.SplitN(input, " ", 2)
    cmd := strings.ToUpper(parts[0])
    IncrementCommands()

    switch cmd {
    case "PING":
        return "PONG\n"
    case "ECHO":
        if len(parts) > 1 {
            return parts[1] + "\n"
        }
        return "Usage: ECHO <message>\n"
    case "TIME":
        return time.Now().Format(time.RFC1123) + "\n"
    case "EXIT":
        return "Goodbye!\n"
    default:
        return "Unknown command\n"
    }
}