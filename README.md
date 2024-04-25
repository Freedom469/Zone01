#  A Simple Text Completion/Editing/Auto-correction Tool

This Go program processes a text file, applies various transformations based on specified commands, and writes the result to an output file.

## Features

- **Punctuation Mark Correction:** Corrects spacing around punctuation marks.
- **Word Case Transformation:** Supports capitalizing, uppercasing, and lowercasing of words.
- **Numeric Conversion:** Converts hexadecimal and binary numbers to decimal.
- **Custom Commands:** Supports custom commands like `(cap, n)` to capitalize the last `n` words, `(up, n)` to uppercase the last `n` words, and `(low, n)` to lowercase the last `n` words.

## Usage

### Input File Format

The input file should contain the text to be processed.

### Output File Format

The processed text will be written to the output file.

### Command-Line Usage


Replace `<inputfile>` with the path to the input file and `<outputfile>` with the desired path for the output file.

## Custom Commands

The following custom commands are supported:

- **(cap):** Capitalizes the preceding word.
- **(up):** Converts the preceding word to uppercase.
- **(low):** Converts the preceding word to lowercase.
- **(hex):** Converts the preceding hexadecimal number to decimal.
- **(bin):** Converts the preceding binary number to decimal.
- **(cap, n):** Capitalizes the last `n` words.
- **(up, n):** Uppercases the last `n` words.
- **(low, n):** Lowercases the last `n` words.

## Example

Suppose we have an input file `input.txt` with the following content:


Running the program with the command:


will produce an output file `output.txt` with the following content:


## Dependencies

- Go 1.15 or higher

## Author

[Abraham Maiko King'oo]

## License

This project is licensed under the [MIT License](LICENSE).
