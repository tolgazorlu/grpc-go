# Golang gRPC and Protobuf Project

### Layout

```tree
├── .gitignore
├── CHANGELOG.md
├── Makefile
├── README.md
├── bin
│   ├── calculator
│   │   ├── client
│   │   └── server
│   └── greet
│       ├── client
│       └── server
├── calculator
│   ├── client
│   │   ├── avg.go
│   │   ├── calculator.go
│   │   ├── main.go
│   │   ├── prime.go
│   │   └── sqrt.go
│   ├── proto
│   │   ├── avg.proto
│   │   ├── calculator.proto
│   │   ├── prime.proto
│   │   ├── sqrt.proto
│   │   └── sum.proto
│   ├── server
│   │   ├── avg.go
│   │   ├── calculator.go
│   │   ├── main.go
│   │   ├── prime.go
│   │   ├── sqrt.proto

```

A brief description of the layout:

- `.gitignore` varies per project, but all projects need to ignore `bin` directory.
- `Makefile` is used to build the project. **You need to tweak the variables based on your project**.
- `CHANGELOG.md` contains auto-generated changelog information.
- `README.md` is a detailed description of the project.
- `bin` is to hold build outputs.

## Notes

- Makefile **MUST NOT** change well-defined command semantics, see Makefile for details.
