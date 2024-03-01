package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}

type Currencies []string

type OrderBook map[string]OrderBookPair

type Ticker map[string]TickerValue

type Trades map[string][]Pair

type TickerValue struct {
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

type OrderBookPair struct {
	AskQuantity string     `json:"ask_quantity"`
	AskAmount   string     `json:"ask_amount"`
	AskTop      string     `json:"ask_top"`
	BidQuantity string     `json:"bid_quantity"`
	BidAmount   string     `json:"bid_amount"`
	BidTop      string     `json:"bid_top"`
	Ask         [][]string `json:"ask"`
	Bid         [][]string `json:"bid"`
}

type Pair struct {
	TradeID  int64  `json:"trade_id"`
	Type     string `json:"type"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
	Amount   string `json:"amount"`
	Date     int64  `json:"date"`
}

type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
}

type Exmo struct {
	client *http.Client
	url    string
}

//func (r *Trades) UnmarshalJSON(d []byte) error {
//	// first, decode just the object's keys and leave
//	// the values as raw, non-decoded JSON
//	var obj map[string]json.RawMessage
//	if err := json.Unmarshal(d, &obj); err != nil {
//		return err
//	}
//
//	// finally, decode the rest of the element values
//	// in the object and store them in the Pair field
//	res := make(map[string][]Pair, len(obj))
//	for key, val := range obj {
//		cc := []Pair{}
//		if err := json.Unmarshal(val, &cc); err != nil {
//			return err
//		}
//		res[key] = cc
//	}
//	return nil
//}

//func (c *Pair) UnmarshalJSON(d []byte) error {
//	tmp := []interface{}{&c.Quantity, &c.TradeID, &c.Amount, &c.Date, &c.Type, &c.Price}
//	length := len(tmp)
//	err := json.Unmarshal(d, &tmp)
//	if err != nil {
//		return err
//	}
//	g := len(tmp)
//	if g != length {
//		return fmt.Errorf("Lengths don't match: %d != %d", g, length)
//	}
//	return nil
//}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {
	data := url.Values{}
	for _, pair := range pairs {
		data.Set("pair", pair)
	}
	var tradess Trades
	//var tradesss map[string]json.RawMessage
	//res := make(map[string][]Pair)
	fmt.Println(e.url + trades)
	b, err := e.doPostRequest(e.url+trades, data)
	if err != nil {
		return Trades{}, err
	}
	err = json.Unmarshal(b, &tradess)
	if err != nil {
		return Trades{}, err
	}
	return tradess, nil
}

//err = json.Unmarshal(b, &tradess)
//if err != nil {
//	return Trades{}, err
//}
//var pairss []Pair
//err = json.Unmarshal(tradess[pairs[0]], &pairss)
//if err != nil {
//	log.Fatal(err)
//}
//err = json.Unmarshal(tradess["result"], &tradesss)
//if err != nil {
//	return Trades{}, err
//}
//for k, raw := range tradess {
//	var p []Pair
//	err = json.Unmarshal(raw, &p)
//	res[k] = p
//}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	data := url.Values{}
	for _, pair := range pairs {
		data.Set("limit", strconv.Itoa(limit))
		data.Set("pair", pair)
	}

	var ob OrderBook
	fmt.Println(e.url + orderBook)
	b, err := e.doPostRequest(e.url+orderBook, data)
	if err != nil {
		return OrderBook{}, err
	}
	err = json.Unmarshal(b, &ob)
	return ob, nil
}

func (e *Exmo) GetCurrencies() (Currencies, error) {
	b, err := e.doRequest(e.url + currency)
	var cur Currencies
	err = json.Unmarshal(b, &cur)
	if err != nil {
		return Currencies{}, err
	}
	return cur, nil
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	urll := e.url + candlesHistory + "?" + "symbol=" + pair + "&resolution=" + strconv.Itoa(limit) + "&from=" + strconv.Itoa(int(start.Unix())) + "&to=" + strconv.Itoa(int(end.Unix()))
	b, err := e.doRequest(urll)
	if err != nil {
		return CandlesHistory{}, err
	}
	var cHistory CandlesHistory
	err = json.Unmarshal(b, &cHistory)
	if err != nil {
		return CandlesHistory{}, err
	}
	return cHistory, nil
}

func (e *Exmo) GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error) {
	urll := e.url + candlesHistory + "?" + "symbol=" + pair + "&resolution=" + strconv.Itoa(limit) + "&from=" + strconv.Itoa(int(start.Unix())) + "&to=" + strconv.Itoa(int(end.Unix()))
	b, err := e.doRequest(urll)
	if err != nil {
		return nil, err
	}
	var cHis CandlesHistory
	err = json.Unmarshal(b, &cHis)
	if err != nil {
		return nil, err
	}
	var res []float64
	for i := range cHis.Candles {
		res = append(res, cHis.Candles[i].C)
	}
	return res, nil
}

func (e *Exmo) doRequest(url string) ([]byte, error) {
	resp, err := e.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (e *Exmo) doPostRequest(urll string, data url.Values) ([]byte, error) {
	r, _ := http.NewRequest(http.MethodPost, urll, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := e.client.Do(r)
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	return resBody, nil
}

func NewExmo(opts ...func(exmo *Exmo)) *Exmo {
	exmo := &Exmo{}
	for _, opt := range opts {
		opt(exmo)
	}
	return exmo
}

func WithClient(client *http.Client) func(exmo *Exmo) {
	return func(exmo *Exmo) {

		exmo.client = client
	}
}

func WithURL(url string) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.url = url
	}
}

func (e *Exmo) GetTicker() (Ticker, error) {
	// Реализация метода GetTicker
	b, err := e.doRequest(e.url + ticker)
	if err != nil {
		return Ticker{}, err
	}
	var ticker Ticker
	err = json.Unmarshal(b, &ticker)
	if err != nil {
		return Ticker{}, err
	}
	return ticker, nil
}

// Другие методы...

func main() {
	var exchange Exchanger
	exchange = NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{Timeout: time.Minute}))
	//a, _ := exchange.GetTicker()
	//fmt.Println(a)
	//b, _ := exchange.GetTrades("BTC_USD", "BTC_EUR")
	//fmt.Println(b)
	c, _ := exchange.GetOrderBook(100, "BTC_USD")
	fmt.Println(c)
	//d, _ := exchange.GetCurrencies()
	//fmt.Println(d)
	//t, err := exchange.GetClosePrice("BTC_USD", 30, time.Now().Add(-time.Hour*24), time.Now())
	//if err != nil {
	//	return
	//}
	//fmt.Println(t)
	//ticker, err := exchange.GetCandlesHistory("BTC_USD", 30, time.Now().Add(-time.Hour*24), time.Now())
	//if err != nil {
	//	return
	//}
	//fmt.Println(ticker)
}
