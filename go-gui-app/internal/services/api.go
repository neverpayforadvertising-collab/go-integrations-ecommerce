package services // Defines service layer package

import (
	"encoding/json" // Provides JSON encoding/decoding
	"net/http" // Enables HTTP client functionality
)

// FetchData handles external API call and returns formatted response
func FetchData() string {
	resp, err := http.Get("https://api.publicapis.org/entries") // Sends GET request to external API
	if err != nil { // Checks if request failed
		return "Error: " + err.Error() // Returns error message as string
	}
	defer resp.Body.Close() // Ensures response body is closed after function exits

	var data interface{} // Declares generic variable to hold JSON response
	err = json.NewDecoder(resp.Body).Decode(&data) // Decodes JSON response into variable
	if err != nil { // Checks if decoding failed
		return "Error parsing JSON: " + err.Error() // Returns parsing error
	}

	bytes, err := json.MarshalIndent(data, "", "  ") // Formats JSON into readable indented form
	if err != nil { // Checks if formatting failed
		return "Error formatting JSON: " + err.Error() // Returns formatting error
	}

	return string(bytes) // Converts byte slice to string and returns result
}