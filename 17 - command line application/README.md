# CLI Application - IP and Server Finder

This is a command line application developed in Go that allows you to search for information about IP addresses and domain name servers on the internet.

## Features

The application offers two main commands:

### 1. Search IPs (`ip`)
Searches and displays all IP addresses associated with a specific domain.

### 2. Search Servers (`servers`)
Searches and displays the names of name servers (Name Servers) of a specific domain.

## How to Use

### Prerequisites
- Go 1.24.3 or higher
- Dependency: `github.com/urfave/cli`

### Compilation
```bash
go build -o command-line main.go
```

### Command Syntax

#### Search IPs of a domain:
```bash
./command-line ip --host <domain>
```

**Examples:**
```bash
./command-line ip --host amazon.com.br
./command-line ip --host google.com
./command-line ip --host github.com
```

#### Search name servers of a domain:
```bash
./command-line servers --host <domain>
```

**Examples:**
```bash
./command-line servers --host amazon.com.br
./command-line servers --host google.com
./command-line servers --host github.com
```

### Parameters

- `--host`: Specifies the domain to be queried (default: "devbook.com.br")

### Output Examples

#### Searching IPs:
```bash
$ ./command-line ip --host amazon.com.br
205.251.242.103
52.84.0.0
52.84.0.1
...
```

#### Searching servers:
```bash
$ ./command-line servers --host amazon.com.br
ns-1234.awsdns-12.org
ns-567.awsdns-34.com
ns-890.awsdns-56.net
ns-123.awsdns-78.co.uk
```

## Project Structure

```
17 - command line application/
├── app/
│   └── app.go          # Main CLI application logic
├── main.go             # Application entry point
├── go.mod              # Project dependencies
└── README.md           # This documentation
```

## Technologies Used

- **Go**: Main programming language
- **github.com/urfave/cli**: Library for creating CLI applications
- **net**: Go standard package for network operations (LookupIP, LookupNS)

## Credits

This project was developed during the course **"Learn Go (Golang) from scratch, 100% hands-on and with a real application to add to your developer portfolio!"** taught by **Otávio Augusto Gallego** on the Udemy platform.

- **Instructor**: Otávio Augusto Gallego
- **Platform**: Udemy
- **Course**: Learn Go (Golang) from scratch, 100% hands-on and with a real application to add to your developer portfolio! 