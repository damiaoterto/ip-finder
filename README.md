# IP Finder

A command-line tool written in Go that helps you find IP addresses and nameservers for a given domain.

## Features

- Lookup IP addresses for a domain
- Find nameservers associated with a domain
- Simple command-line interface
- Zero external dependencies for core functionality

## Installation

### Prerequisites

- Go 1.x or higher
- Make (optional, for using Makefile commands)

### Building from source

1. Clone the repository:
```bash
git clone github.com/damiaoterto/ip-finder
cd ip-finder
```

2. Build the binary:
```bash
make build
```

This will create an executable named `ip-finder` in your current directory.

## Usage

The application provides two main commands:

### Finding IP addresses

```bash
./ip-finder ip --host example.com
```

This will display all IP addresses associated with the domain.

### Finding nameservers

```bash
./ip-finder servers --host example.com
```

This will list all nameservers for the specified domain.

## Makefile Commands

- `make build`: Compiles the application
- `make clear`: Cleans build artifacts and removes the binary

## Project Structure

```
.
├── main.go         # Entry point
├── app.go         # Application logic and CLI setup
└── Makefile       # Build automation
```

## Dependencies

- [github.com/urfave/cli](https://github.com/urfave/cli) - CLI application framework

## License

This project is open source and available under the MIT License.
