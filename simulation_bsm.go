package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Parameters
	S0 := 100.0  // initial value
	K := 105.0   // strike price
	T := 1.0     // maturity
	r := 0.05    //risk free short rate
	sigma := 0.2 //volatility
	numPoints := 250_000

	start := time.Now()
	optionPrice := bsmCallValue(S0, K, T, r, sigma, numPoints)
	duration := time.Since(start)

	fmt.Printf("European Option Value: %.3f\n", optionPrice)
	fmt.Println("Execution time: ", duration)
}

func bsmCallValue(S0, K, T, r, sigma float64, n int) float64 {
	d1 := math.Log(S0/K) + T*(r+0.5*sigma*sigma)/(sigma*math.Sqrt(T))
	d2 := math.Log(S0/K) + T*(r-0.5*sigma*sigma)/(sigma*math.Sqrt(T))

	mciD1 := monteCarloIntegrator(gaussian, -20.0, d1, n)
	mciD2 := monteCarloIntegrator(gaussian, -20.0, d2, n)

	return S0*mciD1 - K*math.Exp(-r*T)*mciD2
}

// MC integrator
func monteCarloIntegrator(f func(float64) float64, a float64, b float64, n int) float64 {
	var s float64
	for i := 0; i < n; i++ {
		ui := rand.Float64()
		xi := a + (b-a)*ui
		s += f(xi)
	}

	s = ((b - a) / float64(n)) * s
	return s
}

// function to be integrated
func gaussian(x float64) float64 {
	return (1 / math.Sqrt(2*math.Pi)) * math.Exp(-0.5*x*x)
}
