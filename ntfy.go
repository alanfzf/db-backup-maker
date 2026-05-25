package main

import (
    "net/http"
    "strings"
    "log"
    "io"
)

func SendToNtfy(message string) {
    resp, err := http.Post("https://ntfy.sh/corposistemas", "text/plain",
        strings.NewReader(message))

    if err != nil {
        log.Printf("Error sending notification: %v", err)
    }

    defer resp.Body.Close()


    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        log.Printf("Failed to send notification. Status: %s, Response: %s", resp.Status, string(body))
    }
}
