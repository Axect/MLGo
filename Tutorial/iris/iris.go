package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	var rawCSVData [][]string

	for {
		record, err := reader.Read()
		// io.EOF is an example of error
		if err == io.EOF {
			break
		}

		rawCSVData = append(rawCSVData, record)
	}

	fmt.Println(rawCSVData)

}
