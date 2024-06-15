/*
Copper.co - Interview

The task typically involves identifying arbitrage opportunities in currency exchange rates. Arbitrage involves exploiting the price differences of currencies in different markets to make a profit. This would involve:

Finding a Currency Conversion Cycle: That results in more of the original currency than you started with after completing a series of trades starting and ending with the same currency.

Calculation: Multiply the exchange rates for a cycle of currency conversions. If the product of these rates exceeds 1 when you return to the original currency, an arbitrage opportunity exists.

example:

GBP -> USD -> EUR -> GBP  

[2 0 1 2] 1.012 --- (Round) ---> 1.01

 USD: {USD: 1.00, EUR: 0.92, GBP: 0.79, CHF: 0.89},
 EUR: {USD: 1.09, EUR: 1.00, GBP: 0.86, CHF: 0.90},
 GBP: {USD: 1.28, EUR: 1.19, GBP: 1.00, CHF: 1.13},
 CHF: {USD: 1.13, EUR: 1.04, GBP: 0.88, CHF: 1.00},
*/

/*
Basic Concept:
- Select a Currency: Start from one currency.
- Find Cycles: Explore all cycles starting and ending at this currency.
- Calculate Product: Multiply the exchange rates for each cycle.
- Check for Arbitrage: If the product of the rates in any cycle is greater than 1, it indicates an arbitrage opportunity.

Explanation:
- checkArbitrage Function: This function recursively checks for cycles starting from a specific currency. It multiplies the exchange rates as it traverses through the cycle and checks if the product exceeds 1 when it completes the cycle.
- findArbitrage Function: This iterates over each currency and uses checkArbitrage to detect any arbitrage starting from each.
- Depth: A depth check is included to avoid infinite recursion and only allow revisiting the start currency when completing a cycle.
*/

package main

import (
	"fmt"
)

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	CHF Currency = "CHF"
)

var exchange = map[Currency]map[Currency]float64{
	USD: {EUR: 0.92, GBP: 0.79, CHF: 0.89},
	EUR: {USD: 1.09, GBP: 0.86, CHF: 0.90},
	GBP: {USD: 1.28, EUR: 1.19, CHF: 1.13},
	CHF: {USD: 1.13, EUR: 1.04, GBP: 0.88},
}

// Check for arbitrage starting from a specific currency
func checkArbitrage(start Currency, current Currency, product float64, visited map[Currency]bool, depth int) bool {
	if depth > 0 && current == start {
		return product > 1.0 // Arbitrage condition
	}
	visited[current] = true
	for next, rate := range exchange[current] {
		if !visited[next] || (depth > 0 && next == start) {
			if checkArbitrage(start, next, product*rate, visited, depth+1) {
				return true
			}
		}
	}
	visited[current] = false
	return false
}

func findArbitrage() bool {
	for currency := range exchange {
		visited := make(map[Currency]bool)
		if checkArbitrage(currency, currency, 1.0, visited, 0) {
			return true
		}
	}
	return false
}

func main() {
	if findArbitrage() {
		fmt.Println("Arbitrage opportunity detected!")
	} else {
		fmt.Println("No arbitrage opportunity found.")
	}
}
