Go-Reloaded

Overview

Go-Reloaded is a command-line text processing tool written in Go.

The program reads an input text file, applies a series of transformations, and writes the processed result to an output file.

This project demonstrates:
- CLI argument handling
- File reading and writing
- Text parsing and transformation pipeline
- String manipulation and formatting rules

---

Features

The program supports the following transformations:

1. Hexadecimal conversion
- `(hex)` converts the previous word from hexadecimal to decimal

2. Binary conversion
- `(bin)` converts the previous word from binary to decimal

3. Case transformations
- `(up)` → uppercase previous word  
- `(low)` → lowercase previous word  
- `(cap)` → capitalizes previous word  

4. Multi-word case
- `(up, N)` → uppercases previous N words  
- `(low, N)` → lowercases previous N words  
- `(cap, N)` → capitalizes previous N words  

---

5. Punctuation rules
- Punctuation marks `.,!?;:` are attached to the previous word
- A space is removed before punctuation
- A space is kept after punctuation (unless grouped)

---

6. Grouped punctuation
- `...` and `!?` are preserved as special groups

---

7. Quotes handling
- Removes unnecessary spaces inside single quotes `' '`
- Ensures quoted text is formatted correctly

---

8. Articles correction
- Converts `a` → `an` when the next word starts with:
  - a, e, i, o, u, or h

---

Usage : go run . input.txt output.txt
=======
# go-reloaded
Go-Reloaded is a command-line text processing tool written in Go that applies structured transformations such as hex/bin conversion, case formatting, punctuation fixing, quotes handling, and basic grammar corrections through a sequential processing pipeline.
