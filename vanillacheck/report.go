package vanillacheck

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
)

func (tr *TestRunner) PrintSummary() {
	passed, failed := 0, 0
	for _, result := range tr.Results {
		if result.Passed {
			passed++
			fmt.Printf("\033[32m[PASS]\033[0m %s (%d ms)\n", result.Name, result.Runtime.Milliseconds())
		} else {
			failed++
			fmt.Printf("\033[31m[FAIL]\033[0m %s - %s\n", result.Name, result.ErrorMsg)
		}
	}
	fmt.Printf("\nTotal: %d | Passed: %d | Failed: %d\n", len(tr.Results), passed, failed)
}

func (tr *TestRunner) WriteJSONReport(filename string) {
	data, _ := json.MarshalIndent(tr.Results, "", "  ")
	os.WriteFile(filename, data, 0644)
	fmt.Println("JSON report written to", filename)
}

func (tr *TestRunner) WriteHTMLReport(filename string) {
	const tmpl = `
	<html>
	<head><title>Test Report</title></head>
	<body>
	<h1>Test Results</h1>
	<table border="1">
	<tr><th>Test Name</th><th>Status</th><th>Runtime (ms)</th><th>Error</th></tr>
	{{range .}}
	<tr>
	<td>{{.Name}}</td>
	<td>{{if .Passed}}PASS{{else}}FAIL{{end}}</td>
	<td>{{.Runtime.Milliseconds}}</td>
	<td>{{.ErrorMsg}}</td>
	</tr>
	{{end}}
	</table>
	</body>
	</html>`
	t, _ := template.New("report").Parse(tmpl)
	f, _ := os.Create(filename)
	defer f.Close()
	t.Execute(f, tr.Results)
	fmt.Println("HTML report written to", filename)
}
