# Quadchecker

This program identifies the quad functions based on the input string and its structure. It categorizes the input string into five different quad functions (A, B, C, D, and E) depending on their layout patterns. The program allows users to input a string and then identifies which quad the string corresponds to.

## Usage

To use the program, follow these steps:

1. Clone the repository to your local machine.
2. Ensure you have Go installed on your machine.
3. Run the program with the command: go run main.go
4. Input the string you want to categorize into quad types. You can also pipe the outputs of quad functions or other commands that print into command line console.
5. The program will identify the quad(s) that the string belongs to and display the result.

## Program Structure

The program consists of the following components:

1. **Main Package:** The main package contains the core logic for reading the input string from command line console, categorizing it into quads, and displaying the results.

2. **GetXY Function:** This function calculates the width and height of the input string based on the number of characters and newline characters.

3. **Itoa Function:** This function converts an integer to its corresponding string representation. Only positive integers handled.

4. **Quad Functions (A, B, C, D, E):** These functions generate specific pattern layouts for each quad based on the width and height of the input string.

## Functionality

The program works as follows:

1. It reads the input string from the user or from the command line console.
2. It calculates the width and height of the input string.
3. It compares the input string with predefined patterns for each quad (A, B, C, D, and E) to identify the corresponding quad(s).
4. It displays the identified quad(s) along with their respective x and y coordinates.

If the input string does not match any of the predefined quadrant patterns, the program displays "Not a quad function."

Test folder contains programs that print quad layout patterns with provided x and y arguments and can be used for testing purposes.

### Development Team:
- Yerdaulet Tasbolat
- Bakytzhan Kaziyev
- Suindik Taspolatov
- Zhanserik Rakhmetuly

