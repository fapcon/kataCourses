package main

import "fmt"

const (
	ProductCocaCola = iota
	ProductPepsi
	ProductSprite
)

type Product struct {
	ProductID     int
	Sells         []float64
	Buys          []float64
	CurrentPrice  float64
	ProfitPercent float64
}

type Profitable interface {
	SetProduct(p *Product)
	GetAverageProfit() float64
	GetAverageProfitPercent() float64
	GetCurrentProfit() float64
	GetDifferenceProfit() float64
	GetAllData() []float64
	Average(prices []float64) float64
	Sum(prices []float64) float64
}

type StatisticProfit struct {
	product *Product
}

func (s *StatisticProfit) SetProduct(p *Product) {
	s.product = p
}
func (s *StatisticProfit) GetAverageProfit() float64 {
	sum := s.Sum(s.product.Sells) - s.Sum(s.product.Buys)
	return sum / float64(len(s.product.Sells))
}

func (s *StatisticProfit) GetAverageProfitPercent() float64 {
	averageProfit := s.GetAverageProfit()
	return (averageProfit / s.product.CurrentPrice) * 100
}

func (s *StatisticProfit) GetCurrentProfit() float64 {
	currentProfit := s.Sum(s.product.Sells) - s.Sum(s.product.Buys)
	return currentProfit
}

func (s *StatisticProfit) GetDifferenceProfit() float64 {
	differenceProfit := s.Sum(s.product.Sells) - s.Sum(s.product.Buys)
	return differenceProfit - (s.product.CurrentPrice * float64(len(s.product.Sells)))
}

func (s *StatisticProfit) GetAllData() []float64 {
	return []float64{s.GetAverageProfit(), s.GetCurrentProfit(), s.GetAverageProfitPercent(), s.GetDifferenceProfit()}
}

func (s *StatisticProfit) Average(prices []float64) float64 {
	sum := s.Sum(prices)
	avg := sum / float64(len(prices))
	return avg
}

func (s *StatisticProfit) Sum(prices []float64) float64 {
	var sum float64
	for _, val := range prices {
		sum += val
	}
	return sum
}

func main() {
	p := &Product{
		ProductID:    ProductCocaCola,
		Sells:        []float64{2.0, 4.0, 5.5},
		Buys:         []float64{1.5, 1.9, 4.6},
		CurrentPrice: 2.66,
	}

	statistic := &StatisticProfit{}
	statistic.SetProduct(p)

	fmt.Println("Average Profit:", statistic.GetAverageProfit())
	fmt.Println("Average Profit Percent:", statistic.GetAverageProfitPercent())
	fmt.Println("Current Profit:", statistic.GetCurrentProfit())
	fmt.Println("Difference Profit:", statistic.GetDifferenceProfit())
	data := statistic.GetAllData()
	fmt.Println("All Data:", data)
	fmt.Println("Average Sell Price:", statistic.Average(p.Sells))
	fmt.Println("Total Sell Amount:", statistic.Sum(p.Sells))
}
