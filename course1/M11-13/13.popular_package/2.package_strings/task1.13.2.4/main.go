package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateActivationKey() string {
	var res string
	mas := make([]string, 4)
	for j := 0; j < 4; j++ {
		sb := strings.Builder{}
		sb.Grow(4)
		for i := 0; i < 4; i++ {
			sb.WriteByte(charset[rand.Intn(len(charset))])
		}
		mas[j] = sb.String()
	}
	res = strings.Join(mas, "-")
	return res
}

func main() {
	activationKey := generateActivationKey()
	fmt.Println(activationKey) // UQNI-NYSI-ZVYB-ZEFQ
}
