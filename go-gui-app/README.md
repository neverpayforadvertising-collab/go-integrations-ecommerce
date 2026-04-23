<!-- How everything connects (workflow)
main.go starts the app
ui.go builds the interface
Button click triggers FetchData()
api.go calls external API
Response is formatted → returned → displayed

Key engineering takeaways
Clear separation: UI vs Services
Non-blocking UI via goroutines
Clean API handling with error checks
Structured for real production scaling

Run the project
go mod tidy   // Install dependencies
go env -w CGO_ENABLED=1
go env CGO_ENABLED
go clean -cache
export CGO_ENABLED=1
export CC=gcc
go run ./cmd  // Start GUI app -->