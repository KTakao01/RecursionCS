// https://recursionist.io/dashboard/problems/432
package main

import (
	"fmt"
)

type Function func(int)

func main() {
	printTaxByState := func(state string) Function {
		switch state {
		case "Arizona":
			return func(income int) {
				fedTax := int(0.21 * float64(income))
				staTax := int(0.049 * float64(income))
				fmt.Printf("Federal Tax: %d\n", fedTax)
				fmt.Printf("%s State Tax: %d\n", state, staTax)
				fmt.Printf("Tax Amount: %d\n", fedTax+staTax)
			}

		case "California":
			return func(income int) {
				fedTax := int(0.21 * float64(income))
				staTax := int(0.088 * float64(income))
				fmt.Printf("Federal Tax: %d\n", fedTax)
				fmt.Printf("%s State Tax: %d\n", state, staTax)
				fmt.Printf("Tax Amount: %d\n", fedTax+staTax)
			}

		case "Northcarolina":
			return func(income int) {
				fedTax := int(0.21 * float64(income))
				staTax := int(0.025 * float64(income))
				fmt.Printf("Federal Tax: %d\n", fedTax)
				fmt.Printf("%s State Tax: %d\n", state, staTax)
				fmt.Printf("Tax Amount: %d\n", fedTax+staTax)
			}
		default:
			return func(income int) {
				fedTax := int(0.21 * float64(income))
				fmt.Printf("Federal Tax: %d\n", fedTax)
				fmt.Printf("Unknown State Tax\n")
				fmt.Printf("Unknown Tax Amount\n")
			}
		}
	}

	getTaxInAZ := printTaxByState("Arizona")
	getTaxInCA := printTaxByState("California")
	getTaxInNC := printTaxByState("Northcarolina")
	getTaxInAZ(400000)
	getTaxInCA(100000)
	getTaxInNC(500000)

}
