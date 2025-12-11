# Swimburger Intro

Multi-platform intro/business card that works across different package managers and platforms.

## Usage

### Node.js (npx)
```bash
npx swimburger
```

### .NET
```bash
dotnet tool install -g swimburger
swimburger
```

Or with the new .NET 10+ `dnx` tool execution script:
```bash
dnx swimburger
```

### Python
```bash
python -m pipx run swimburger
```

Or without installation:
```bash
```

### Go
```bash
go install github.com/swimburger/intro/packages/go/swimburger@latest
swimburger
```

## Development

This repository contains multiple packages, one for each platform:

- `packages/npx` - Node.js package for npx
- `packages/dotnet` - .NET global tool
- `packages/python` - Python package
- `packages/go` - Go module

Each package is independently versioned and published to its respective registry.

## Structure

Each package contains the same profile information but formatted for its platform's conventions and styling libraries.
