package geo

import (
	"bufio"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"

	geo "github.com/kellydunn/golang-geo"
)

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type PolygonChecker interface {
	Contains(point Point) bool // проверить, находится ли точка внутри полигона
	Allowed() bool             // разрешено ли входить в полигон
	RandomPoint() Point        // сгенерировать случайную точку внутри полигона
}

type Polygon struct {
	polygon *geo.Polygon
	allowed bool
}

func (p Polygon) Contains(point Point) bool {

	return p.polygon.Contains(geo.NewPoint(point.Lat, point.Lng))

}

func (p Polygon) Allowed() bool {

	return p.allowed
}

func (p Polygon) RandomPoint() Point {
	var outp Point

	points := p.polygon.Points()
	if len(points) == 0 {
		log.Fatal("empty polygon error")
		return outp
	}

	i1 := rand.Intn(len(points))
	i2 := (len(points)*3/2 - i1) % len(points)

	outp.Lat = rand.Float64()*math.Abs(points[i1].Lat()-points[i2].Lat()) + math.Min(points[i1].Lat(), points[i2].Lat())
	outp.Lng = rand.Float64()*math.Abs(points[i1].Lng()-points[i2].Lng()) + math.Min(points[i1].Lng(), points[i2].Lng())

	//fmt.Println("random point generated:", outp)

	return outp
}

func NewPolygon(points []Point, allowed bool) *Polygon {
	// используем библиотеку golang-geogrpc для создания полигона

	geoPoints := make([]*geo.Point, len(points))

	for i, point := range points {
		geoPoints[i] = geo.NewPoint(point.Lat, point.Lng)
	}

	return &Polygon{
		polygon: geo.NewPolygon(geoPoints),
		allowed: allowed,
	}
}

func CheckPointIsAllowed(point Point, allowedZone PolygonChecker, disabledZones []PolygonChecker) bool {
	// проверить, находится ли точка в разрешенной зоне

	for _, dZone := range disabledZones {

		if dZone.Contains(point) {
			return false
		}
	}

	return allowedZone.Contains(point)
}

func GetRandomAllowedLocation(allowedZone PolygonChecker, disabledZones []PolygonChecker) Point {
	var point Point
	// получение случайной точки в разрешенной зоне

	for {
		point = allowedZone.RandomPoint()

		if !allowedZone.Contains(point) {
			continue
		}

		for _, dZone := range disabledZones {

			if dZone.Contains(point) {
				continue
			}
		}

		break

	}

	return point
}

func ParseJS(path, name string) *[]Point {
	points := []Point{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("file opening error:" + err.Error())
		return &points
	}

	scnr := bufio.NewScanner(file)

	for scnr.Scan() {
		textLine := scnr.Text()

		if strings.Contains(textLine, name) {
			for scnr.Scan() {
				textLine := scnr.Text()
				if strings.Contains(textLine, "];") {
					return &points
				}
				textLine = strings.Replace(textLine, "[", "", 1)
				textLine = strings.Replace(textLine, "],", "", 1)
				textLine = strings.ReplaceAll(textLine, " ", "")
				pts := strings.Split(textLine, ",")
				p1, err := strconv.ParseFloat(pts[0], 64)
				if err != nil {
					log.Fatal("coordinates file parsing error:" + err.Error())
				}
				p2, err := strconv.ParseFloat(pts[1], 64)
				if err != nil {
					log.Fatal("coordinates file parsing error:" + err.Error())
				}
				points = append(points, Point{p1, p2})

			}
		}

	}

	return &points

}

func NewDisAllowedZone1() *Polygon {
	// добавить полигон с разрешенной зоной
	// полигоны лежат в /public/js/polygons.js

	points := ParseJS("/app/public/js/polygon.js", "noOrdersPolygon1")

	return NewPolygon(*points, false)
}

func NewDisAllowedZone2() *Polygon {
	// добавить полигон с разрешенной зоной
	// полигоны лежат в /public/js/polygons.js

	points := ParseJS("/app/public/js/polygon.js", "noOrdersPolygon2")

	return NewPolygon(*points, false)
}

func NewAllowedZone() *Polygon {
	// добавить полигон с разрешенной зоной
	// полигоны лежат в /public/js/polygons.js

	points := ParseJS("/app/public/js/polygon.js", "mainPolygon")

	return NewPolygon(*points, true)
}
