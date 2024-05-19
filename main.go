package main

import (
	"context"
	"fmt"
	"os"

	"github.com/foureyez/linkbook/internal"
)

func main() {
	ctx := context.Background()
	app := internal.NewApp()
	if err := app.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
