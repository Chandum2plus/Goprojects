/*
What is Logging to programs-:

the standard library package log provides a basic infrastructure for log
management in Go language that can be used for logging our Go programs.
the main purpose of logging is to get trace of what's happening in the program,
where it's happening, and when it's happening. Logs can be providing code
tracing,profiling and analytics. Logging (eye and ear of a programmer) is a way to find
those bugs and learn more about how the program is functioning.

*/

// To work with package log, we must add it to the list of imports

package main

import "log"

func init() {
	log.SetPrefix("LOG:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	// Println write to the standard logger
	log.Println("Init Started")
}
func main() {
	log.Println("main Started")
	// fatalln is println() followed by a call to os.Exit(1)
	log.Fatalln("Fatal message")

	// Panicln is Println() followed by a call to panic()
	log.Panicln("Panic message")
}
