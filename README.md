# Grake

Manage and make groff manuscripts.

### Example: grake.toml

```toml
Title = "Lorem ipsum around the world"
Author = "TheDevtop"
Files = ["main.ms", "index.ms"]
Output = "manuscript.pdf"
```

### Command: Initialize

Create a new Groff project.

```
Usage of init:
  -a string
        Specify name of author (default "John Doe")
  -c uint
        Specify columns (default 2)
  -d string
        Specify working directory
  -f string
        Specify initial source file (default "manuscript.ms")
  -o string
        Specify output file (default "manuscript.pdf")
  -t string
        Specify title (default "Title")
```

### Command: Build

Compile a Groff project, output to pdf.

```
Usage of build:
  -d string
        Specify working directory
```

### Command: Clean

Remove output file.

```
Usage of clean:
  -d string
        Specify working directory
```
