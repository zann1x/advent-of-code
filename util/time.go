package util

import (
	"log"
	"time"
)

// Track tracks the time a code block took for execution.
// what is an arbitrary string telling what exactly is measured now.
func Track(what string) func() {
	start := time.Now()
	return func() {
		log.Println(what, "took", time.Since(start))
	}
}
