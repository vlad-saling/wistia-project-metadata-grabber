package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	bearerToken := "XXXX" // Set your Bearer token here

	var records []struct {
		ID   float64
		Name string
	}

	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var inputs []string
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input.txt:", err)
		return
	}

	for _, input := range inputs {

		url := fmt.Sprintf("https://api.wistia.com/v1/projects/%s", input)

		fmt.Printf("Processing #%d: %s\n", len(records)+1, url)

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("accept", "application/json")
		req.Header.Add("authorization", "Bearer "+bearerToken)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error making HTTP request:", err)
			return
		}
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
		name, _ := result["name"].(string)
		id, _ := result["id"].(float64)

		records = append(records, struct {
			ID   float64
			Name string
		}{ID: id, Name: name})
	}

	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("id,name\n")
	if err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	for _, record := range records {
		_, err = file.WriteString(fmt.Sprintf("%.0f,%s\n", record.ID, record.Name))
		if err != nil {
			fmt.Println("Error writing data:", err)
			return
		}
	}

	fmt.Println("Output saved to output.csv")
}
