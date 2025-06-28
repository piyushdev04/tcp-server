package server

import (
    "fmt"
    "log"
    "os"
)

var logger *log.Logger

func InitLogger() {
    os.MkdirAll("logs", os.ModePerm)
    file, err := os.OpenFile("logs/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    logger = log.New(file, "", log.LstdFlags)
}

func logInfo(format string, v ...any) {
    fmt.Printf("[INFO] "+format+"\n", v...)
    logger.Printf("[INFO] "+format, v...)
}

func logError(format string, v ...any) {
    fmt.Printf("[ERROR] "+format+"\n", v...)
    logger.Printf("[ERROR] "+format, v...)
}

func logFatal(format string, v ...any) {
    fmt.Printf("[FATAL] "+format+"\n", v...)
    logger.Fatalf("[FATAL] "+format, v...)
}

func logCommand(addr, cmd string) {
    logger.Printf("[COMMAND] %s: %s", addr, cmd)
}