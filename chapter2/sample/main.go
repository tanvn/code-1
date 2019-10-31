package main

import (
	"log"
	"os"

	_ "github.com/tanvn/goinaction-code/chapter2/sample/matchers"
	"github.com/tanvn/goinaction-code/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func init(){
	log.Println("from main package")

}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("Trump")
}
