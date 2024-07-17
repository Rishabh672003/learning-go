package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type getResult struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Todos []getResult

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("search query failed:")
	}

	var result Todos
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}

	const headerTemplate = `Todos:
	{{- range . }}
--------------------------------------------------------
    UserID: {{.UserID}}
    ID: {{.ID}}
    Title: {{.Title}}
    Completed: {{.Completed}}
	{{- end }}`

	// Parse the template
	tmpl, err := template.New("headers").Parse(headerTemplate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing template: %v\n", err)
		return
	}

	// Execute the template with the result data
	if err := tmpl.Execute(os.Stdout, result); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing template: %v\n", err)
		return
	}
}
