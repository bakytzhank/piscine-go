# Math Skills Program

This Go program calculates statistical measures such as average, median, variance, and standard deviation from a set of numerical data in a text file.

## Usage

1. **Execute**:
   - Run the program with a data file path as an argument:
    ```
    go run mathskills.go "path/filename.txt"
    ```
    Ensure data file contains numerical data, each value on a separate line.

2. **Results**:
   - The program prints rounded integers:
     - Average
     - Median
     - Variance
     - Standard Deviation

## Structure

- `readData`: Reads data from a text file.
- `calculateAvg`: Calculates the average.
- `calculateMdn`: Calculates the median.
- `calculateVar`: Calculates the variance.
- `calculateStdDev`: Calculates the standard deviation.
- `main`: Entry point, reads data, calculates measures, and prints results.

## Error Handling

- Manages file reading and data parsing errors.


