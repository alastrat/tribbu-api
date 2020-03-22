package timetracker

import (
	"fmt"
	"log"
	"time"
)

func timing(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
