package main

import (
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/http"
)

func main() {
	http.StartFiberServer()
}
