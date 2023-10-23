// https://recursionist.io/dashboard/problems/431
package main

import (
	"fmt"
)

type Function func(int)

func main() {
	weightFormulaByPlanet := func(planet string) Function {
		switch planet {
		case "Earth":
			return func(weight int) {
				mass := int(9.8 * float64(weight))
				fmt.Printf("The weight on Earth is %d\n", mass)
			}
		case "Mars":
			return func(weight int) {
				mass := int(3.7 * float64(weight))
				fmt.Printf("The weight on Mars is %d\n", mass)
			}
		case "Jupiter":
			return func(weight int) {
				mass := int(24.8 * float64(weight))
				fmt.Printf("The weight on Jupiter is %d\n", mass)
			}
		default:
			return func(weight int) {
				fmt.Println("Unknown planet")
			}
		}
	}

	getWeightOnEarth := weightFormulaByPlanet("Earth")
	getWeightOnMars := weightFormulaByPlanet("Mars")
	getWeightOnJupiter := weightFormulaByPlanet("Jupiter")
	getWeightOnEarth(50)
	getWeightOnMars(70)
	getWeightOnJupiter(90)

}
