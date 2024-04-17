package main

import (
	"fmt"

	"github.com/ZafirProjects/QuodOrbisChallenge/server"
)

func main() {
	server := server.NewServer()

	fmt.Println("Server listening on port :3000")

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
