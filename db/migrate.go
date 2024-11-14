package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
)

type Config struct {
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
	DBHost     string `json:"db_host"`
	DBPort     int    `json:"db_port"`
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func connectDB(config *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	return sql.Open("postgres", connStr)
}

// Check if a table exists in PostgreSQL
func tableExists(db *sql.DB, tableName string) (bool, error) {
	var result *string
	query := `SELECT to_regclass($1)`
	err := db.QueryRow(query, tableName).Scan(&result)
	if err != nil {
		return false, err
	}
	return result != nil, nil
}

// Create a table dynamically based on the CSV header
func createTable(db *sql.DB, tableName string, columns []string) error {
	// Construct the CREATE TABLE SQL statement
	columnDefs := []string{}
	for _, col := range columns {
		columnDefs = append(columnDefs, fmt.Sprintf("%s TEXT", col))
	}
	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, strings.Join(columnDefs, ", "))

	_, err := db.Exec(stmt)
	fmt.Println(stmt)
	if err != nil {
		return fmt.Errorf("failed to create table %s: %v", tableName, err)
	}

	fmt.Printf("Table %s created successfully.\n", tableName)
	return nil
}

// Insert data into PostgreSQL with a dynamic table name
func insertData(db *sql.DB, tableName string, records [][]string) error {
	if len(records) < 1 {
		return fmt.Errorf("no records found in CSV file for table %s", tableName)
	}

	// Get columns from the first row (header)
	columns := records[0]
	placeholder := make([]string, len(columns))
	for i := range placeholder {
		placeholder[i] = fmt.Sprintf("$%d", i+1)
	}

	// Construct the INSERT SQL statement dynamically
	stmt := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`,
		tableName, strings.Join(columns, ", "), strings.Join(placeholder, ", "))

	for _, record := range records[1:] { // Skip the header row
		params := make([]interface{}, len(record))

		for i, v := range record {
			params[i] = v
		}

		_, err := db.Exec(stmt, params...)
		if err != nil {
			return fmt.Errorf("failed to insert record into %s: %v", tableName, err)
		}
	}
	return nil
}

// Read and process a CSV file
func processCSVFile(db *sql.DB, filedir string) error {
	// Extract table name from filename (without .csv extension)
	tableName := strings.TrimSuffix(filepath.Base(filedir), ".csv")

	file, err := os.Open(filedir)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Check if the table exists, create if not
	tableExistsFlag, err := tableExists(db, tableName)

	if err != nil {
		return err
	}

	if !tableExistsFlag {
		// Create the table using the first row (header) as the column names
		err = createTable(db, tableName, records[0])
		if err != nil {
			return err
		}
	}

	// Insert each record into the database
	err = insertData(db, tableName, records)
	if err != nil {
		return err
	}

	fmt.Printf("Processed file: %s\n", filedir)
	return nil
}

// Process all CSV files in a directory
func processCSVFilesInDir(db *sql.DB, dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".csv") {
			filepath := filepath.Join(dir, file.Name())
			err := processCSVFile(db, filepath)
			if err != nil {
				log.Printf("Failed to process file %s: %v", filepath, err)
			}
		}
	}
	return nil
}

func main() {
	// Load config
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to PostgreSQL
	db, err := connectDB(config)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Directory containing CSV files
	dir := "./csv"

	// Process all CSV files in the directory
	err = processCSVFilesInDir(db, dir)
	if err != nil {
		log.Fatalf("Failed to process CSV files: %v", err)
	}

	fmt.Println("All CSV files processed successfully!")

}
