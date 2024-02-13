package main

import (
    "bufio"
    "fmt"
    "os"
    "math"
    "sort"
    "strconv"
)

// read data from a file and return a slice of float64 values
func readData(dataFile string) ([]float64, error) {
    file, err := os.Open(dataFile)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var data []float64
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // Convert text to float64
        val, err := strconv.ParseFloat(scanner.Text(), 64)
        if err != nil {
            return nil, err
        }
        data = append(data, val)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return data, nil
}

// calculate the average of the provided data
func calculateAvg(data []float64) float64 {
    sum := 0.0
    for _, val := range data {
        sum += val
    }
    return sum / float64(len(data))
}

// calculate the median of the provided data
func calculateMdn(data []float64) float64 {
    sort.Float64s(data)
    n := len(data)
    if n%2 == 0 {
        return (data[n/2-1] + data[n/2]) / 2
    }
    return data[n/2]
}

// calculate the variance of the provided data based on the average
func calculateVar(data []float64, avg float64) float64 {
    sum := 0.0
    for _, val := range data {
        sum += (val - avg) * (val - avg)
    }
    return sum / float64(len(data))
}

// calculate the standard deviation based on the variance
func calculateStdDev(variance float64) float64 {
    return math.Sqrt(variance)
}

func main() {
    // Ensure correct number of command-line arguments
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run mathskills.go data.txt")
        os.Exit(1)
    }

    dataFile := os.Args[1]
    // Read data from the provided file
    data, err := readData(dataFile)
    if len(data)==0 {
        fmt.Println("No data in file")
        os.Exit(1)
    }
    if err != nil {
        fmt.Println("Error reading data:", err)
        os.Exit(1)
    }

    // Calculate statistical measures
    avg := calculateAvg(data)
    mdn := calculateMdn(data)
    variance := calculateVar(data, avg)
    stdDev := calculateStdDev(variance)

    // Print results
    fmt.Printf("Average: %.f\n", math.Round(avg))
    fmt.Printf("Median: %.f\n", math.Round(mdn))
    fmt.Printf("Variance: %.f\n", math.Round(variance))
    fmt.Printf("Standard Deviation: %.f\n", math.Round(stdDev))
}
