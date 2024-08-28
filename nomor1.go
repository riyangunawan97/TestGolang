package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// Structure to match the JSON format
type Country struct {
	Name     string `json:"name"`
	DialCode string `json:"dialCode"`
	ISOCode  string `json:"isoCode"`
	Flag     string `json:"flag"`
}

// Handler function to render the HTML
func renderTable(w http.ResponseWriter, r *http.Request) {
	// Get the JSON data from the URL
	resp, err := http.Get("https://citcall.com/test/countries.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON data
	var countries []Country
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create an HTML template to display the data
	const tpl = `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Countries</title>
    </head>
    <body>
        <h1>List of Countries</h1>
        <table border="1">
            <tr>
                <th>Name</th>
                <th>Dial Code</th>
                <th>ISO Code</th>
                <th>Flag</th>
            </tr>
            {{range .}}
            <tr>
                <td>{{.Name}}</td>
                <td>{{.DialCode}}</td>
                <td>{{.ISOCode}}</td>
                <td><img src="{{.Flag}}" alt="Flag of {{.Name}}" width="50"/></td>
            </tr>
            {{end}}
        </table>
    </body>
    </html>
    `

	// Parse and execute the template
	tmpl, err := template.New("table").Parse(tpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template with the data
	err = tmpl.Execute(w, countries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Register the handler function
	http.HandleFunc("/", renderTable)

	// Start the web server on port 8080
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
