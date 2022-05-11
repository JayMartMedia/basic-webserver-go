# Basic Webserver (Go)

A basic go webserver which creates some endpoints which call functions to mutate an array in memory (in a production server, these functions would likely be reading/inserting into/from a database).

Expands upon this tutorial from the Go team: https://go.dev/doc/tutorial/web-service-gin

This repo adds delete and update endpoints.

# Setup

Clone this project: `git clone https://github.com/JayMartMedia/basic-webserver-go.git`

Install Go: https://go.dev/doc/tutorial/getting-started#install

Install dependancies: `go get .`

Start the server: `go run .`

# Testing the API endpoints

If you are using VSCode, you can add the 'REST Client' extension and use the included requests in the `./tests/albums.rest` file.

Otherwise, you can test the API endpoints by using curl, postman, etc.
