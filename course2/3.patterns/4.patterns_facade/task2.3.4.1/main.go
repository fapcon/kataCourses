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

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --all
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

type Indicatorer interface {
	SMA(pair string, limit, period int, from, to time.Time) ([]float64, error)
	EMA(pair string, limit, period int, from, to time.Time) ([]float64, error)
}

type Indicator struct {
	exchange     Exchanger
	calculateSMA func(data []float64, period int) []float64
	calculateEMA func(data []float64, period int) []float64
}

type IndicatorOption func(*Indicator)

func NewIndicator(exchange Exchanger, opts ...IndicatorOption) *Indicator {
	i := &Indicator{
		exchange:     exchange,
		calculateEMA: calculateEMA,
		calculateSMA: calculateSMA,
	}
	for _, opt := range opts {
		opt(i)
	}
	return i
}

func WithSMA(calculateSMA func(data []float64, period int) []float64) func(*Indicator) {
	return func(indicator *Indicator) {
		indicator.calculateSMA = calculateSMA
	}
}

func WithEMA(calculateEMA func(data []float64, period int) []float64) func(*Indicator) {
	return func(indicator *Indicator) {
		indicator.calculateEMA = calculateEMA
	}
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

func (i *Indicator) SMA(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	var sl []float64
	res, err := i.exchange.GetCandlesHistory(pair, limit, from, to)
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
	return sum, err
}

func (i *Indicator) EMA(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	res, err := i.exchange.GetCandlesHistory(pair, limit, from, to)
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

func calculateSMA(data []float64, period int) []float64 {
	return nil
}

func calculateEMA(data []float64, period int) []float64 {
	return nil
}

// Dashboarder должен возвращать историю свечей и индикаторы с несколькими периодами, заданными через opts
type Dashboarder interface {
	GetDashboard(pair string, opts ...func(*Dashboard)) (DashboardData, error)
}

type DashboardData struct {
	Name           string
	CandlesHistory CandlesHistory
	Indicators     map[string][]IndicatorData
	limit          int
	from           time.Time
	to             time.Time
}

type IndicatorData struct {
	Name     string
	Period   int
	Indicate []float64
}

type IndicatorOpt struct {
	Name      string
	Periods   []int
	Indicator Indicatorer
}

type Dashboard struct {
	exchange           Exchanger
	withCandlesHistory bool
	IndicatorOpts      []IndicatorOpt
	limit              int
	from               time.Time
	to                 time.Time
}

func (d *Dashboard) GetDashboard(pair string, opts ...func(*Dashboard)) (DashboardData, error) {
	for _, opt := range opts {
		opt(d)
	}
	fmt.Println(d)
	indicators := make(map[string][]IndicatorData)
	var iData []IndicatorData
	for i := 0; i < len(d.IndicatorOpts); i++ {
		for j := 0; j < len(d.IndicatorOpts[i].Periods); j++ {
			q, err := d.IndicatorOpts[i].Indicator.SMA(pair, d.limit, d.IndicatorOpts[i].Periods[j], d.from, d.to)
			if err != nil {
				return DashboardData{}, err
			}
			iData = append(iData, IndicatorData{Indicate: q, Period: d.IndicatorOpts[i].Periods[j], Name: "SMA"})
		}
		for j := 0; j < len(d.IndicatorOpts[i].Periods); j++ {
			q, err := d.IndicatorOpts[i].Indicator.EMA(pair, d.limit, d.IndicatorOpts[i].Periods[j], d.from, d.to)
			if err != nil {
				return DashboardData{}, err
			}
			iData[j] = IndicatorData{Indicate: q, Period: d.IndicatorOpts[i].Periods[j], Name: "EMA"}
		}
	}
	indicators[pair] = iData
	cHis, err := d.exchange.GetCandlesHistory(pair, d.limit, d.from, d.to)
	if err != nil {
		return DashboardData{}, nil
	}
	data := DashboardData{Name: pair, limit: d.limit, from: d.from, to: d.to, Indicators: indicators, CandlesHistory: cHis}
	return data, nil
}

func WithCandlesHistory(limit int, from, to time.Time) func(*Dashboard) {
	return func(d *Dashboard) {
		d.to = to
		d.from = from
		d.limit = limit
		d.withCandlesHistory = true
	}
}

func WithIndicatorOpts(opts ...IndicatorOpt) func(*Dashboard) {
	return func(d *Dashboard) {
		for _, val := range opts {
			d.IndicatorOpts = append(d.IndicatorOpts, val)
		}
	}
}

func NewDashboard(exchange Exchanger) Dashboarder {
	return &Dashboard{exchange: exchange}
}

func main() {
	exchange := NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{Timeout: time.Minute}))
	dashboard := NewDashboard(exchange)
	data, err := dashboard.GetDashboard("BTC_USD", WithCandlesHistory(30, time.Now().Add(-time.Hour*24), time.Now()), WithIndicatorOpts(
		IndicatorOpt{
			Name:      "SMA",
			Periods:   []int{5, 10, 20},
			Indicator: NewIndicator(exchange),
		},
		IndicatorOpt{
			Name:      "EMA",
			Periods:   []int{5, 10, 20},
			Indicator: NewIndicator(exchange),
		},
	))
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
