package main

import (
    "net/http"
    "strings"
)

func SendToNtfy(message string) {
    http.Post("https://ntfy.sh/corposistemas", "text/plain",
        strings.NewReader(message))
}
