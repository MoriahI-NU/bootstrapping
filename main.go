package main

import (
	"bootstrap/helper"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sort"
	"time"
)

func main() {

	fmt.Printf("---------R---------\n")

	// Set seed for consistentcy
	rand.Seed(time.Now().UnixNano())

	//start R timer
	start_time := time.Now()

	// Run R script
	R_results := helper.RunR("MustangPrice.csv")

	var memStats runtime.MemStats

	// Print results
	fmt.Println("Sample Mean:", R_results.Sample_Mean)
	fmt.Print("Confidence Intervals: ")
	for _, interval := range R_results.Confidence {
		fmt.Printf("[%f, %f] ", interval.Lower, interval.Upper)
	}
	fmt.Println()
	fmt.Println("tStar:", R_results.Tstar)
	fmt.Println("tMargin:", R_results.Tmarg)

	//end R timer
	r_elapsed_time := time.Since(start_time)
	fmt.Print("R elapsed time: ", r_elapsed_time, "\n")

	runtime.ReadMemStats(&memStats)
	R_bytes := memStats.TotalAlloc
	fmt.Print("R bytes: ", memStats.TotalAlloc, "\n")

	/////////////////////////////////////////

	fmt.Printf("---------Go---------\n")

	//start go timer
	start_time = time.Now()

	// Read data from CSV file
	mustangs, err := helper.ReadCSV("MustangPrice.csv")
	if err != nil {
		log.Fatal(err)
	}

	var prices []float64
	for _, mustang := range mustangs {
		prices = append(prices, mustang)
	}

	// Sample Mean
	sampleMean := helper.CalculateMean(prices)
	fmt.Printf("Sample Mean: %f\n", sampleMean)

	// Resample 2000 times
	numResamples := 2000
	resamplingDistribution := make([]float64, numResamples)
	for i := 0; i < numResamples; i++ {
		resample := helper.BootstrapResample(prices)
		resampleMean := helper.CalculateMean(resample)
		resamplingDistribution[i] = resampleMean
	}

	// Calculate confidence interval
	sort.Float64s(resamplingDistribution)
	confidenceInterval := helper.CalculateConfidenceInterval(resamplingDistribution, 0.95)

	fmt.Printf("Confidence Intervals (95%%): [%f, %f]\n", confidenceInterval[0], confidenceInterval[1])

	// Calculate critical values
	tStar := helper.CalculateTStar(0.95, 24)
	// tStar := helper.CalculateTStar(0.95, len(mustangs)-1)
	fmt.Printf("tStar: %f\n", tStar)

	// Calculate margin of error
	tMargin := tStar * helper.CalculateStandardDeviation(resamplingDistribution)
	fmt.Printf("tMargin: %f\n", tMargin)

	runtime.ReadMemStats(&memStats)

	//end go timer
	go_elapsed_time := time.Since(start_time)
	fmt.Print("Go elapsed time: ", go_elapsed_time, "\n")
	fmt.Print("Go bytes: ", memStats.TotalAlloc-R_bytes, "\n")
}
