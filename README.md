## Overview
**single-cms-boilerplate** is a RESTful API built with Golang

- **Language**: Go 1.24
- **Live Reload**: [Air](https://github.com/air-verse/air)
- **API Docs**: [Swag](https://github.com/swaggo/swag)
- **Testing**: [GoMock](https://github.com/uber-go/mock)
- **Makefile**: Build automation

## Installation

1. Install dependencies:
   ```sh
   go mod tidy
   ```
2. Set up environment variables (e.g., `.env` file):
   ```plaintext
   rename .env.example to .env and put your config
   ```
3. Run service watch mode:
   ```sh
   ## watch mode
   make dev
   
   ## prod mode
   make start
   ```
   
## Documentation

1. Generate Swagger Docs:
   ```sh
   make swag-generate
   ```
2. Format Swagger Docs:
   ```sh
   make swag-format
   ```

## Generate Mock File

1. Generate Mock Unit Test:
   ```sh
   make mock-create source=${your-file-path}
   ```
   
## Testing

1. Run Unit Test:
   ```sh
   make test
   ```

