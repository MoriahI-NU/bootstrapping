package helper

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/distuv"
)

type Mustang struct {
	Price float64
}

type R_results struct {
	Sample_Mean []float64 `json:"Sample_Mean"`
	Confidence  []struct {
		Lower    float64 `json:"lower"`
		Upper    float64 `json:"upper"`
		CentralP float64 `json:"central.p"`
		Row      string  `json:"_row"`
	} `json:"Confidence"`
	Tstar []float64 `json:"tstar"`
	Tmarg []float64 `json:"t_marg"`
}

///////////////Helper functions for Go////////////////////

func ReadCSV(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var mustangs []float64
	for _, record := range records[1:] {
		price, err := strconv.ParseFloat(record[2], 64)

		if err != nil {
			return nil, err
		}

		mustangs = append(mustangs, price)
	}

	return mustangs, nil
}

func CalculateMean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

func BootstrapResample(data []float64) []float64 {
	resample := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		index := rand.Intn(len(data))
		resample[i] = data[index]
	}
	return resample
}

func CalculateConfidenceInterval(data []float64, level float64) [2]float64 {
	alpha := (1 - level) / 2
	lowerIndex := int(alpha * float64(len(data)-1))
	upperIndex := int((1 - alpha) * float64(len(data)-1))

	return [2]float64{data[lowerIndex], data[upperIndex]}
}

func CalculateTStar(alpha float64, degreesOfFreedom int) float64 {
	tDist := distuv.StudentsT{
		Mu:    0,
		Sigma: 1,
		Nu:    float64(degreesOfFreedom),
	}
	criticalValue := tDist.Quantile(1 - alpha/2)

	roundedcriticalValue := math.Round(criticalValue*1000) / 1000
	return roundedcriticalValue
}

func CalculateStandardDeviation(data []float64) float64 {
	sum := 0.0
	mean := CalculateMean(data)

	for _, value := range data {
		sum += math.Pow(value-mean, 2)
	}

	variance := sum / float64(len(data)-1)

	stdev := math.Sqrt(variance)

	roundedstdev := math.Round(stdev*1000) / 1000

	return roundedstdev
}

///////////////Helper functions for R////////////////////

// For some reason, the warnings/notifications from importing the mosaic library are included in outR (found in the last function here: "RunR")
// To remedy this, I've created the following functions to parse the output of the R script and return only the JSON output

func findJSONStartIndex(output string) int {
	// Look for the first occurrence of "{"
	return strings.Index(output, "{")
}

func findJSONEndIndex(output string, startIndex int) int {
	// Start from the character after "{"
	braceCount := 1
	for i := startIndex + 1; i < len(output); i++ {
		switch output[i] {
		case '{':
			braceCount++
		case '}':
			braceCount--
			if braceCount == 0 {
				return i
			}
		}
	}
	return -1
}

func RunR(data string) R_results {
	cmdR := exec.Command("Rscript", "bootstrap.R", data)
	outR, err := cmdR.CombinedOutput()
	if err != nil {
		fmt.Println(err, string(outR))
		return R_results{}
	}

	//convert output to string
	outStr := string(outR)

	// Find the start and end indices of the JSON parts
	startIndex := findJSONStartIndex(outStr)
	if startIndex == -1 {
		fmt.Println("No JSON found in the output")
		return R_results{}
	}

	endIndex := findJSONEndIndex(outStr, startIndex)
	if endIndex == -1 {
		fmt.Println("Invalid JSON structure in the output")
		return R_results{}
	}

	// Extract ONLY the JSON part
	jsonStr := outStr[startIndex : endIndex+1]

	// Unmarshal the jsonStr
	var results R_results
	if err := json.Unmarshal([]byte(jsonStr), &results); err != nil {
		fmt.Println(err)
		return R_results{}
	}

	return results
}
