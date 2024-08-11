package server

import (
	"fmt"
	"os"
)

func StartServer() {
	srv := New(":8080")

	if err := srv.Start(); err != nil {
		fmt.Println("Error starting server: ", err)
		os.Exit(1)
	}
}
