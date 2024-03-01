package main

import (
	"fmt"
	"os"
	"time"
)
import "github.com/mattevans/dinero"

func main() {
	rate := currencyPairRate("USD", "EUR", 100.0)
	fmt.Println(rate) // 82.73

}
func currencyPairRate(val1 string, val2 string, q float64) float64 {
	client := dinero.NewClient(
		os.Getenv("OPEN_EXCHANGE_APP_ID"),
		val1,
		20*time.Minute,
	)
	rsp, err := client.Rates.Get(val2)
	if err != nil {
		return 0
	}
	return q**rsp
}