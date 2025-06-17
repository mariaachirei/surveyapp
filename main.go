package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type School struct {
	Name        string
	Code        string
	Teacher     string
	Subject     string
	ClassGrades []string
}

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func loadSchoolsFromCSV(filePath string) ([]School, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var schools []School
	for _, record := range records {
		schools = append(schools, School{
			Name:        record[0],
			Code:        record[1],
			Teacher:     record[2],
			Subject:     record[3],
			ClassGrades: record[4:],
		})
	}
	return schools, nil
}

func main() {
	schools, err := loadSchoolsFromCSV("data/schools_and_teachers.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Loaded data from CSV")
	for _, school := range schools {
		fmt.Printf("School: %+v\n", school)
	}

	http.HandleFunc("/schools", func(w http.ResponseWriter, r *http.Request) {
		response := ApiResponse{Message: "List of Schools", Data: schools}
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server is running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
