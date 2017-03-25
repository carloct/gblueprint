package main

import (
	"fmt"
	"os"
)

func main() {

	opts := ServerOptions{
		Port:    8080,
		Address: "localhost",
	}

	err := Server(opts)
	if err != nil {
		exitWithError("cannot start the server: %s", err)
	}

}

func exitWithError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args)
	os.Exit(1)
}
