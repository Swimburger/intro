# Swimburger - Python Package

Personal intro card for the terminal.

## Usage

With pipx (recommended):
```bash
pipx install swimburger
swimburger
```

Or run directly without installing:
```bash
pipx run swimburger
```

With pip:
```bash
pip install swimburger
swimburger
```

## Local Development

```bash
pip install -e .
swimburger
```

## Publishing

```bash
pip install build twine
python -m build
twine upload dist/*
```
