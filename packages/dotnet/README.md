# Swimburger - .NET Global Tool

Personal intro card for the terminal.

## Usage

Install globally:
```bash
dotnet tool install -g swimburger
swimburger
```

Or with the new .NET 10+ `dnx` tool execution script:
```bash
dnx swimburger
```

## Local Development

```bash
dotnet run
```

## Publishing

```bash
dotnet pack
dotnet nuget push bin/Release/swimburger.1.0.0.nupkg --source https://api.nuget.org/v3/index.json --api-key YOUR_API_KEY
```
