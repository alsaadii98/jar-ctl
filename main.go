package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: jar-ctl [options]")
		fmt.Println("Commands: add, list, remove")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		printBanner()
		fmt.Println("Adding a new jar entry...")
		// Logic to add a new jar entry
	case "list":
		printBanner()
		fmt.Println("Listing all jar entries...")
		// Logic to list all jar entries
	case "remove":
		printBanner()
		fmt.Println("Removing a jar entry...")
		// Logic to remove a jar entry
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		fmt.Println("Available commands: add, list, remove")
		os.Exit(1)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printBanner() {
	fmt.Println(`
      ██╗ █████╗ ██████╗      ██████╗████████╗██╗     
      ██║██╔══██╗██╔══██╗    ██╔════╝╚══██╔══╝██║     
      ██║███████║██████╔╝    ██║        ██║   ██║     
 ██   ██║██╔══██║██╔═══╝     ██║        ██║   ██║     
 ╚█████╔╝██║  ██║██║         ╚██████╗   ██║   ███████╗
  ╚════╝ ╚═╝  ╚═╝╚═╝          ╚═════╝   ╚═╝   ╚══════╝
                   CLI for Managing JARs

               Developed by @alsaadii98
	`)
}
