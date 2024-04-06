package main

import (
	"fmt"
)

func main() {
	result := scrape("pertamina") // hanya contoh
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
