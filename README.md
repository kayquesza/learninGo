# 🚀 LearninGo - My Go Learning Journey

<div align="center">

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GitHub](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white)
![Udemy](https://img.shields.io/badge/Udemy-EC5252?style=for-the-badge&logo=Udemy&logoColor=white)

</div>

## 📖 About the Project

This repository is dedicated to my learning of the **Go (Golang)** programming language. Here you will find all my progress, exercises, projects and documentation from my learning journey.

### 🎯 Objectives

- Document my progress in learning Go
- Share knowledge and experiences
- Create a portfolio of Go projects
- Establish a solid foundation in the language

## 📚 Study Resources

### Main Course
- **Course**: [Learn Golang from Zero - Develop a Complete Application](https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa)
- **Instructor**: [Otavio Gallego](https://github.com/OtavioGallego/curso-golang)
- **Platform**: Udemy

## 🌟 Featured Projects

### 🚀 **Complete Applications**
- **Project 17**: CLI Application for IP and Server lookup
- **Project 24**: Full CRUD REST API with MySQL database
- **Project 23**: Database connection and management

### 🧪 **Testing & Quality**
- **Project 19**: Comprehensive testing examples with coverage reports
- **Project 16**: Interface implementations and generic types

### ⚡ **Advanced Concepts**
- **Project 18**: Concurrency patterns, goroutines, and channels
- **Project 14**: Advanced function concepts (closures, defer, panic/recover)

## 📁 Repository Structure

```
learninGo/
├── 01-package/                    # Packages and modules
├── 02-variable/                   # Variables
├── 03-data-types/                 # Data types
├── 04-functions/                  # Functions
├── 05-operators/                  # Operators
├── 06-struct/                     # Structures
├── 07-inheritance/                # Inheritance
├── 08-arrays-slices/              # Arrays and Slices
├── 09-pointers/                   # Pointers
├── 10-maps/                       # Maps
├── 11-control-structure/          # Control structures
├── 12-switch/                     # Switch statements
├── 13-loops/                      # Loops
├── 14-advanced-functions/         # Advanced functions
│   ├── 14.1-named-return/         # Named return
│   ├── 14.2-variadic-functions/   # Variadic functions
│   ├── 14.3-anonymous-functions/  # Anonymous functions
│   ├── 14.4-recursive-functions/  # Recursive functions
│   ├── 14.5-defer/                # Defer
│   ├── 14.6-panic-recover/        # Panic and Recover
│   ├── 14.7-closure/              # Closure
│   ├── 14.8-pointers/             # Pointers
│   └── 14.9-init/                 # Init
├── 15-methods/                    # Methods
├── 16-interfaces/                 # Interfaces
│   ├── shapes/                    # Basic interfaces
│   └── generic-type/              # Generic type
├── 17-command-line-app/           # CLI Applications (IP & Server Finder)
├── 18-competition/                 # Concurrency & Competition
│   ├── 1-goroutines/              # Goroutines
│   ├── 2-waitgroup/               # WaitGroup
│   ├── 3-channels/                # Channels
│   ├── 4-buffered-channels/       # Buffered channels
│   ├── 5-select/                  # Select
│   └── 6-concurrency-patterns/    # Concurrency Patterns
│       ├── worker-pools/          # Worker Pools
│       ├── generators/            # Generators
│       └── multiplexer/           # Multiplexer
├── 19-automated-testing/          # Automated testing
│   ├── 1-introduction/            # Introduction
│   └── 2-advanced-testing/        # Advanced Testing
├── 20-json/                       # Working with JSON
│   ├── 1-marshal/                 # Marshal
│   └── 2-unmarshal/               # Unmarshal
├── 21-http/                       # HTTP servers
├── 22-html/                       # HTML templates
├── 23-data-base/                  # Database connections
└── 24-basic-crud/                 # Basic CRUD operations
    ├── database/                   # Database connection module
    └── server/                     # HTTP server with CRUD endpoints
```

## 🛠️ Technologies Used

- **Go (Golang)** - Main language
- **Git** - Version control
- **GitHub** - Repository hosting
- **VS Code or Cursor** - Code editor(s)

### Additional Dependencies
- **MySQL** - Database for projects 23 & 24
- **github.com/go-sql-driver/mysql** - MySQL driver for Go
- **github.com/gorilla/mux** - HTTP router and URL matcher
- **github.com/urfave/cli** - CLI application framework

## 🚀 How to Run the Projects

### Prerequisites
- Go installed (version 1.21 or higher)
- Git installed
- MySQL database (for projects 23 and 24)

### Database Setup (Projects 23 & 24)
```bash
# Create MySQL database
mysql -u root -p
CREATE DATABASE learningo;
CREATE USER 'golang'@'localhost' IDENTIFIED BY 'golang';
GRANT ALL PRIVILEGES ON learningo.* TO 'golang'@'localhost';
FLUSH PRIVILEGES;
EXIT;

# Create users table
mysql -u golang -p learningo
CREATE TABLE users (
    id_user INT AUTO_INCREMENT PRIMARY KEY,
    name_user VARCHAR(255) NOT NULL,
    email_user VARCHAR(255) NOT NULL UNIQUE
);
EXIT;
```

### Running a project
```bash
# Clone the repository
git clone https://github.com/kayquesza/learninGo.git

# Navigate to the project directory
cd learninGo/[project-name]

# Run the project
go run main.go
```

### Running the CRUD API (Project 24)
```bash
cd "24 - basic crud"
go run main.go

# The server will start on port 8080
# Test the endpoints:
# POST /users - Create user
# GET /users - List all users
# GET /users/{id} - Get specific user
# PUT /users/{id} - Update user
# DELETE /users/{id} - Delete user
```

### Running the Database Connection (Project 23)
```bash
cd "23 - data base"
go run data-base.go

# Make sure MySQL is running and the database is configured
```

### Running the CLI Application (Project 17)
```bash
cd "17 - command line application"
go run main.go

# Available commands:
# ./linha-de-comando ip --host google.com
# ./linha-de-comando servers --host github.com

# Or build and run:
# go build -o app main.go
# ./app ip --host amazon.com
```

## 📈 Progress

### ✅ **Completed Topics**
- [x] **Fundamentals**: Variables, data types, operators
- [x] **Control Structures**: If, switch, loops
- [x] **Functions**: Basic and advanced
- [x] **Structs and Methods**: Object-oriented programming
- [x] **Pointers**: Memory management
- [x] **Collections**: Arrays, slices, maps
- [x] **Interfaces**: Polymorphism
- [x] **Concurrency**: Goroutines and channels
- [x] **Testing**: Automated testing
- [x] **Web**: HTTP servers and JSON
- [x] **Database**: MySQL connections and basic operations
- [x] **CRUD API**: Complete REST API with Create, Read, Update, Delete operations
- [x] **Practical Projects**: Complete applications

### 🎯 **Key Skills Developed**
- **Backend Development**: HTTP servers, REST APIs, database operations
- **Concurrency**: Goroutines, channels, wait groups, select statements
- **Testing**: Unit tests, table-driven tests, coverage reports
- **CLI Development**: Command-line applications with flags and commands
- **Database**: MySQL connections, prepared statements, CRUD operations

### 🚀 **Next Steps & Future Learning**
- **Microservices**: Building distributed systems with Go
- **GraphQL**: Implementing GraphQL APIs
- **Docker**: Containerization of Go applications
- **Kubernetes**: Deployment and orchestration
- **Cloud Platforms**: AWS, GCP, or Azure integration
- **Performance**: Profiling and optimization techniques


## 🤝 Contributing

This is a learning repository, but suggestions and improvements are welcome! Feel free to:
- Report bugs or issues
- Suggest new features or improvements
- Share additional learning resources
- Contribute code examples

## 📞 Contact

- **GitHub**: [@kayquesza](https://github.com/kayquesza)
- **LinkedIn**: [Kayque Mendes de Souza](https://linkedin.com/in/kayquesza)

---

<div align="center">

**⭐ If this project helped you, consider giving it a star!**

</div>
