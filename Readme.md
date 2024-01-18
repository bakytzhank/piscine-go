# Go reloaded

This is a simple text completion/editing/auto-correction tool written in Go. It reads a text file, processes its content based on specified triggers, and writes the result to an output file.

## Features

- **Case Modification**: Converts text to uppercase, lowercase, or title case based on triggers like `(up)`, `(low)`, `(cap)`.
- **Numeric Conversion**: Converts numbers from hexadecimal or binary to decimal based on triggers like `(hex)`, `(bin)`.
- **A/An Conversion**: Converts "a" or "A" to "an" or "An" when followed by a word starting with a vowel.
- **Punctuation Handling**: Manages the placement of punctuation marks, ensuring proper spacing.
- **Single Quote Adjustment**: Handles the positioning of single quotes in the text.

## Usage

### Prerequisites

- Go installed on your machine.

### Installation

1. **Clone the repository:**

   ```
   git clone https://zero.academie.one/git/bkaziyev/go-reloaded
   ```

2. **Change into the project directory:**

    ```
    cd go-reloaded
    ```

### Run

Run the executable with the input and output file names as command-line arguments:
```
go run . input.txt output.txt
```
Replace input.txt with the name of your input file and output.txt with the desired name for the output file.

### Example

Assume you have an input file named sample.txt:
```
This is a (cap) sample text for 12 (hex) conversion. a apple, a cat.
```

Running the text processor:
```
go run . sample.txt processed_output.txt
```

The output will be written to processed_output.txt:
```
This is A Sample text for 18 conversion. an apple, a cat.
```