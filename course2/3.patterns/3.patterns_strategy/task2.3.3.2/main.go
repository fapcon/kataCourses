package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=Exchanger
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

type IndicatorSMA struct {
	exchange     Exchanger
	calculateSMA func(data []float64, period int) []float64
}

type IndicatorEMA struct {
	exchange     Exchanger
	calculateEMA func(data []float64, period int) []float64
}

type IndicatorOptionSMA func(*IndicatorSMA)
type IndicatorOptionEMA func(*IndicatorEMA)

func NewIndicatorSMA(exchange Exchanger, opts ...IndicatorOptionSMA) *IndicatorSMA {
	i := &IndicatorSMA{
		exchange: exchange,
	}
	for _, opt := range opts {
		opt(i)
	}
	return i
}
func NewIndicatorEMA(exchange Exchanger, opts ...IndicatorOptionEMA) *IndicatorEMA {
	i := &IndicatorEMA{
		exchange: exchange,
	}
	for _, opt := range opts {
		opt(i)
	}
	return i
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
	//https://api.exmo.com/v1.1/candles_history?symbol=BTC_USD&resolution=30&from=1585556979&to=1585557979
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

func calculateSMA(data []float64, period int) []float64 {
	sma := make([]float64, 0, len(data))
	var total float64

	for i := 0; i+period <= len(data); i += period {
		for j := 0; j < period; j++ {
			total += data[j+i]
		}
		sma = append(sma, math.Round(total/float64(period)))
		total = 0
	}
	return sma
}

func calculateEMA(data []float64, period int) []float64 {
	mult := 2.0 / float64(period+1)
	ema := make([]float64, len(data))
	ema[0] = data[0]
	for i := 1; i < len(data); i++ {
		ema[i] = math.Round((data[i]-ema[i-1])*mult + ema[i-1])
	}
	return ema
}

type GeneralIndicator struct {
	//indSMA    *IndicatorSMA
	//indEMA    *IndicatorEMA
	exchange  Exchanger
	indicator Indicatorer
}

func NewGeneralIndicator(ex Exchanger, strategy Indicatorer) *GeneralIndicator {
	return &GeneralIndicator{exchange: ex, indicator: strategy}
}

func (ie *IndicatorEMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	res, err := ie.exchange.GetCandlesHistory(pair, limit, from, to)
	if err != nil {
		return nil, err
	}
	var sl []float64
	for j := range res.Candles {
		sl = append(sl, res.Candles[j].C)
	}
	var e []float64
	for o := 0; o < len(res.Candles)-1; o++ {
		ema := res.Candles[o+1].C*float64((period+1)/2) + res.Candles[o].C*(1-float64((period+1)/2))
		e = append(e, ema)
	}
	return e, err
}

func (is *IndicatorSMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	var sl []float64
	res, err := is.exchange.GetCandlesHistory(pair, limit, from, to)
	if err != nil {
		return nil, err
	}
	for j := range res.Candles {
		sl = append(sl, res.Candles[j].C)
	}
	//fmt.Println(sl)
	var sum, mas []float64
	for _, data := range sl {
		mas = append(mas, data)
	}
	var r float64
	if period >= len(mas) {
		a := 0.0
		for i := 0; i < len(mas); i++ {
			a = a + mas[i]
		}
		r = a / float64(len(mas))
		sum = append(sum, r)
		return sum, err
	} else {
		for j := 0; j < len(mas)-(period-1); j += period {
			a := 0.0
			for i := 0; i+j < period+j; i++ {
				a = a + mas[i+j]
			}
			r = a / float64(period)
			sum = append(sum, r)
		}
		return sum, err
	}
}

type GeneralIndicatorer interface {
	GetData(pair string, limit, period int, from, to time.Time, indicator Indicatorer) ([]float64, error)
}

func (gi *GeneralIndicator) GetData(pair string, limit, period int, from, to time.Time, indicator Indicatorer) ([]float64, error) {
	return indicator.GetData(pair, limit, period, from, to)
}

type Indicatorer interface {
	GetData(pair string, limit, period int, from, to time.Time) ([]float64, error)
}

func WithSMA(calculateSMA func(data []float64, period int) []float64) func(*IndicatorSMA) {
	return func(i *IndicatorSMA) {
		i.calculateSMA = calculateSMA
	}
}

func WithEMA(calculateEMA func(data []float64, period int) []float64) func(*IndicatorEMA) {
	return func(i *IndicatorEMA) {
		i.calculateEMA = calculateEMA
	}
}

func main() {
	var exchange Exchanger
	exchange = NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{Timeout: time.Minute}))
	indicatorSMA := NewIndicatorSMA(exchange)
	generalIndicator := NewGeneralIndicator(exchange, indicatorSMA)
	sma, err := generalIndicator.GetData("BTC_USD", 30, 10, time.Now().Add(-time.Hour*24), time.Now(), indicatorSMA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sma)

	indicatorEMA := NewIndicatorEMA(exchange)
	generalIndicator = NewGeneralIndicator(exchange, indicatorEMA)
	ema, err := generalIndicator.GetData("BTC_USD", 30, 10, time.Now().Add(-time.Hour*24), time.Now(), indicatorEMA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ema)
}
