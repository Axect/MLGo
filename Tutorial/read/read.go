package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// =============================================================================
	// Open CSV
	// =============================================================================

	f, err := os.Open("simple.csv")
	if err != nil {
		log.Fatal(err)
	}

	// =============================================================================
	// Read in the CSV records
	// =============================================================================

	r := csv.NewReader(f)
	records, err2 := r.ReadAll()
	if err2 != nil {
		log.Fatal(err)
	}

	// ============================================================================
	// Get the Maximum value in the integer column
	// ============================================================================

	var intMax int
	for _, record := range records {
		intVal, err3 := strconv.Atoi(record[0])
		if err3 != nil {
			log.Fatal(err3)
		}

		if intVal > intMax {
			intMax = intVal
		}
	}

	fmt.Println(intMax)
}
