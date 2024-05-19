package main

import (
	"fmt"
	"os"

	"github.com/foureyez/linkbook/internal"
)

func main() {
	app := internal.NewApp()
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
