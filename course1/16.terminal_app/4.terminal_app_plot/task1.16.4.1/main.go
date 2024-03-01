package main

import (
	"encoding/json"
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/gosuri/uilive"
	"github.com/guptarohit/asciigraph"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int64  `json:"updated"`
}

var event keyboard.KeyEvent

/*func keys(wg *sync.WaitGroup, a <-chan keyboard.KeyEvent) {
	defer wg.Done()
	select {
	case event = <-a:
	default:
	}
	if event.Err != nil {
		panic(event.Err)
	}
}*/

func main() {
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	var graphBTC, graphLTC, graphETH string
	res := make(map[string]Response)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.exmo.com/v1/ticker?pair=BTC_USD", nil)
	if err != nil {
		fmt.Println("Ошибка при создании GET запроса:", err)
		return
	}

	_, _ = fmt.Fprintf(writer.Newline(), "1. BTC_USD\n")
	_, _ = fmt.Fprintf(writer.Newline(), "2. LTC_USD\n")
	_, _ = fmt.Fprintf(writer.Newline(), "3. ETH_USD\n")
	_, _ = fmt.Fprintf(writer.Newline(), "\n")
	_, _ = fmt.Fprintf(writer.Newline(), "Press 1-3 to change symbol, press q to exit\n")

	plotValuesBTC := make([]float64, 0)
	plotValuesLTC := make([]float64, 0)
	plotValuesETH := make([]float64, 0)

	//wg := new(sync.WaitGroup)
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Ошибка при выполнении GET запроса:", err)
			return
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			fmt.Println("Ошибка при распаковке JSON:", err)
			return
		}

		fBTC, err := strconv.ParseFloat(res["BTC_USD"].Avg, 64)
		if err != nil {
			err.Error()
		}

		//plotValuesLTC[0] = fBTC

		plotValuesBTC = append(plotValuesBTC, fBTC)

		graphBTC = asciigraph.Plot(plotValuesBTC, asciigraph.Width(100), asciigraph.Height(10), asciigraph.LowerBound(fBTC-fBTC/10000), asciigraph.UpperBound(fBTC+fBTC/10000), asciigraph.SeriesColors(asciigraph.Red), asciigraph.Precision(10000))

		fLTC, err := strconv.ParseFloat(res["LTC_USD"].Avg, 64)
		if err != nil {
			err.Error()
		}
		//plotValuesLTC[0] = fLTC

		plotValuesLTC = append(plotValuesLTC, fLTC)

		graphLTC = asciigraph.Plot(plotValuesLTC, asciigraph.Width(100), asciigraph.Height(10), asciigraph.LowerBound(fLTC-fLTC/10000), asciigraph.UpperBound(fLTC+fLTC/10000), asciigraph.SeriesColors(asciigraph.Red))

		fETH, err := strconv.ParseFloat(res["ETH_USD"].Avg, 64)
		if err != nil {
			err.Error()
		}
		//plotValuesETH[0] = fETH

		plotValuesETH = append(plotValuesETH, fETH)

		graphETH = asciigraph.Plot(plotValuesETH, asciigraph.Width(100), asciigraph.Height(10), asciigraph.LowerBound(fETH-fETH/10000), asciigraph.UpperBound(fETH+fETH/10000), asciigraph.SeriesColors(asciigraph.Red), asciigraph.Precision(10000))
		//wg.Add(1)
		//	go keys(wg, keysEvents)
		select {
		case event = <-keysEvents:
		default:
		}
		if event.Err != nil {
			panic(event.Err)
		}

		if event.Rune == 49 {
			_, _ = fmt.Fprintf(writer.Newline(), "BTC_USD: %v \n", res["BTC_USD"].Avg)
			_, _ = fmt.Fprintf(writer.Newline(), graphBTC)
			//_, _ = fmt.Fprintf(writer.Newline(), "%v \n", res["BTC_USD"].Avg)
		}
		if event.Rune == 50 {
			_, _ = fmt.Fprintf(writer.Newline(), "LTC_USD: %v \n", res["LTC_USD"].Avg)
			_, _ = fmt.Fprintf(writer.Newline(), graphLTC)
		}
		if event.Rune == 51 {
			_, _ = fmt.Fprintf(writer.Newline(), "ETH_USD: %v \n", res["ETH_USD"].Avg)
			_, _ = fmt.Fprintf(writer.Newline(), graphETH)
		}
		if event.Rune == '\x00' {
			_, _ = fmt.Fprintf(writer.Newline(), "1. BTC_USD\n")
			_, _ = fmt.Fprintf(writer.Newline(), "2. LTC_USD\n")
			_, _ = fmt.Fprintf(writer.Newline(), "3. ETH_USD\n")
			_, _ = fmt.Fprintf(writer.Newline(), "\n")
			_, _ = fmt.Fprintf(writer.Newline(), "Press 1-3 to change symbol, press q to exit\n")
		}
		if event.Rune == 113 {
			break
		}
		now := time.Now()
		_, _ = fmt.Fprintf(writer.Newline(), "\nТекущее время: %v\n", now.Format("15:04:05"))
		_, _ = fmt.Fprintf(writer.Newline(), "Текущая дата: %v\n", now.Format("2006-01-02"))
		time.Sleep(100 * time.Millisecond)
	}
	//	wg.Wait()
}

// https://api.exmo.com/v1/ticker?pair=BTC_USD
