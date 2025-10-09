// cmd/anonymitylayer/main.go
package main

import (
    "flag"
    "log"
    "os"

    "anonymitylayer/internal/anonymitylayer"
)

func main() {
    // Define flags with descriptions for better user experience
    verbose := flag.Bool("verbose", false, "Enable verbose logging for detailed output")
    input := flag.String("input", "", "Path to the input file")
    output := flag.String("output", "", "Path to the output file")
    flag.Parse()

    // Create a new app instance with the specified verbosity level
    app := anonymitylayer.NewApp(*verbose)
    
    // Run the app with the provided input and output file paths
    if err := app.Run(*input, *output); err != nil {
        log.Printf("Error: %v", err)
        os.Exit(1)
    }
}