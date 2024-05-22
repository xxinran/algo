package main

import "fmt"

func parseFloatNum(s string) float64 {
	result := float64(0)
	decimal := false
	decimalFactor := 0.1
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			decimal = true
			continue
		}
		if decimal {
			result = result + float64(s[i]-'0')*decimalFactor
			decimalFactor *= 0.1
		} else {
			result = result*10 + float64(s[i]-'0')
		}

		//fmt.Println(result)

	}
	return result

}

func main() {
	fmt.Println(parseFloatNum("1.234"))
}
