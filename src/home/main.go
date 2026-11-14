package main

import (
	"fmt"
	"disk_monitor/src/internal"
)


func main() {
	fmt.Println("Starting Disk Monitor...")
	file_handler.HandleFiles()
	fmt.Println("Done!")
}