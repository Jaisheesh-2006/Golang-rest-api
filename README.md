# Students API

Simple Go `students-api` project demonstrating a minimal command-line program.

## Prerequisites

- Go `1.25.0` or later installed and available in `PATH`

## Running the app

```powershell
go run ./cmd/students-api
```

That prints a greeting from the API entry point.

## Project structure

- `cmd/students-api/main.go`: entry point for the students API example.
- `.gitignore`: Go-specific ignore rules under the API submodule.
- `go.mod`: module configuration and Go version.

## Testing

This minimal project has no automated tests yet; you can add Go tests in `cmd/students-api` if desired.
