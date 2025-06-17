package main

import (
	"database/sql"
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
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read csv: %v", err)
	}

	var schools []School
	for _, record := range records {
		if len(record) < 5 {
			return nil, fmt.Errorf("incorrect number of fields in: %v", record)
		}
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
	// Simulated MySQL database connection
	db, err := sql.Open("mysql", "user:password@tcp(localhost)/database")
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}
	defer db.Close()

	schools, err := loadSchoolsFromCSV("data/schools_and_teachers.csv")
	if err != nil {
		log.Fatalf("failed to load schools from csv: %v", err)
	}

	http.HandleFunc("/schools", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		response := ApiResponse{Message: "List of Schools", Data: schools}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode schools to JSON", http.StatusInternalServerError)
		}
	})

	fmt.Println("Server is running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
