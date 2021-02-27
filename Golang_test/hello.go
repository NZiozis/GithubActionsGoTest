package main

import "log"
import "os"
import "github.com/google/gopacket/pcap"
import "fmt"

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {

	InfoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stderr, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {
	InfoLogger.Println("Application Start")

	fmt.Println(pcap.FindAllDevs())

	InfoLogger.Println("Application Complete")
}

func add(x int, y int) int {
	InfoLogger.Printf("Running add command with, %v and %v", x, y)
	return x + y
}

func sub(x int, y int) int {
	InfoLogger.Printf("Running sub command with, %v and %v", x, y)
	return x - y
}
