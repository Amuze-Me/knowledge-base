package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./kb.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for _, record := range records {
		if len(record) >= 3 {
			format := record[0]
			link := record[1]
			tags := record[2]

			fmt.Printf("Format: %s\nLink: %s\nTags: %s\n\n", format, link, tags)
		}
	}
}
