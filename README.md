# MNC Bank API

This is the documentation for the MNC Bank API, which allows you to interact with our banking services programmatically.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (Version X.X.X) - Download and install it from [golang.org](https://golang.org/).

### Installation

1. Clone the repository to your local machine with HTTP or SSH:
   ###### Using HTTP
   ```bash
   git clone https://github.com/EkoEdyPurwanto/mnc_bank.git
   ```
   ###### Using SSH
   ```bash
   git clone git@github.com:EkoEdyPurwanto/mnc_bank.git
   ```

2. change into the project directory:
   ```bash
   cd mnc_bank
   ```

3. Install the project dependencies:
   ```bash
   go mod download && go mod tidy
   ```

4. Change to the app directory and run the main.go file:
   ```bash
   go run main.go
   ```
   
5. Open the run.http file to try the API that has been created

### Acknowledgments
[echo labstack](), [validator](), [google uuid](), [jwt]()