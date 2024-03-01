package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
)

// swagger:route POST /api/address/search  addr RequestAddressSearch
// getting address
// responses:
// 200:

func main() {
	router := chi.NewRouter()

	proxy := NewReverseProxy("hugo", "1313")

	//go TimeUpdate()
	//go BinTreeBuilt()
	//go graphRandomBuilt()

	os.Setenv("HOST", proxy.host)

	router.Use(proxy.ReverseProxy)

	router.Mount("/", newApiRouter())

	http.ListenAndServe(":8080", router)

}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func newApiRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/api/address/search*", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		addrReq := r.FormValue("query") //Header.Get("query")

		fmt.Println("query parsed:", addrReq)

		addrReq = `{"query":"нижний советс 3"}`

		resp, err := addressGetting([]byte(addrReq))

		status := 200

		if err != nil {
			resp = append(resp, []byte(fmt.Sprintf("%v", err))...)
			if strings.Contains(err.Error(), "format") {
				status = 400
			} else {
				if strings.Contains(err.Error(), "reach") {
					status = 500
				} else {
					status = 405
				}
			}

		}

		w.WriteHeader(status)
		w.Write(resp)

	})

	r.Post("/api/address/geocode*", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		geoSuggReq := r.FormValue("query")

		fmt.Println("query parsed:", geoSuggReq)

		//fmt.Println("query:", geoSuggReq)

		geoSuggReq = `{ "lat": "55.878", "lng": "37.653" }`

		resp, err := geoSuggGetting([]byte(geoSuggReq))

		status := 200

		if err != nil {
			resp = append(resp, []byte(fmt.Sprintf("%v", err))...)
			if strings.Contains(err.Error(), "format") {
				status = 400
			} else {
				if strings.Contains(err.Error(), "reach") {
					status = 500
				} else {
					status = 405
				}
			}

		}

		w.WriteHeader(status)
		w.Write(resp)

	})

	r.Get("/swagger*", func(w http.ResponseWriter, r *http.Request) {

		//http.ServeFile(w, r, "./swagger.json")

		w.Header().Set("Content-Type", "text/html	; charset=utf-8")
		tmpl, err := template.New("swagger").Parse(swaggerTemplate)

		//fmt.Println("template prepared")
		if err != nil {
			fmt.Println("template error:", err)
			return
		}
		err = tmpl.Execute(w, struct {
			Time int64
		}{
			Time: time.Now().Unix(),
		})
		if err != nil {
			fmt.Println("template execution error:", err)
			return
		}
	})

	return r

}

func (rp *ReverseProxy) ReverseProxy(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var proxyUrl *url.URL
		var err error

		if !strings.HasPrefix(r.URL.Path, "/api") && !strings.HasPrefix(r.URL.Path, "/swagger") {

			proxyUrl, err = url.Parse(fmt.Sprintf("http://%s:%s", rp.host, rp.port))
			if err != nil {
				fmt.Println("url parsing error:", err)
			}

			proxy := httputil.NewSingleHostReverseProxy(proxyUrl)

			proxy.ServeHTTP(w, r)

			return

		}

		handler.ServeHTTP(w, r)

	})
}

const content1 = `---
menu:
    before:
        name: tasks
        weight: 5
title: Обновление данных в реальном времени
---

# Задача: Обновление данных в реальном времени

Напишите воркер, который будет обновлять данные в реальном времени, на текущей странице.
Текст данной задачи менять нельзя, только время и счетчик.

Файл данной страницы: ` + "`" + `/app/static/tasks/_index.md` + "`" + `

Должен меняться счетчик и время:

Текущее время:` // 2021-10-13 15:00:00

const content2 = `
Счетчик:` // 0

const content3 = `
## Критерии приемки:
- [ ] Воркер должен обновлять данные каждые 5 секунд
- [ ] Счетчик должен увеличиваться на 1 каждые 5 секунд
- [ ] Время должно обновляться каждые 5 секунд`

func TimeUpdate() {
	t := time.NewTicker(5 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:

			//fmt.Println("writing to file...")
			err := os.WriteFile("/app/static/tasks/_index.md", []byte(fmt.Sprint(content1, time.Now().Format(time.RFC1123), content2, b, content3)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}

func TreeUpdate() {
	t := time.NewTicker(5 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:

			//fmt.Println("writing to file...")
			err := os.WriteFile("/app/static/tasks/_index.md", []byte(fmt.Sprint(content1, time.Now().Format(time.RFC1123), content2, b, content3)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}

const content4 = `---
menu:
    after:
        name: binary_tree
        weight: 2
title: Построение сбалансированного бинарного дерева
---

# Задача построить сбалансированное бинарное дерево
Используя AVL дерево, постройте сбалансированное бинарное дерево, на текущей странице.

Нужно написать воркер, который стартует дерево с 5 элементов, и каждые 5 секунд добавляет новый элемент в дерево.

Каждые 5 секунд на странице появляется актуальная версия, сбалансированного дерева.

При вставке нового элемента, в дерево, нужно перестраивать дерево, чтобы оно оставалось сбалансированным.

Как только дерево достигнет 100 элементов, генерируется новое дерево с 5 элементами.
` +
	"```" + `go
package binary

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Key    int
	Height int
	Left   *Node
	Right  *Node
}

type AVLTree struct {
	Root *Node
}

func NewNode(key int) *Node {
	return &Node{Key: key, Height: 1}
}

func (t *AVLTree) Insert(key int) {
	t.Root = insert(t.Root, key)
}

func (t *AVLTree) ToMermaid() string {

}

func height(node *Node) int {

}

func max(a, b int) int {

}

func updateHeight(node *Node) {

}

func getBalance(node *Node) int {

}

func leftRotate(x *Node) *Node {

}

func rightRotate(y *Node) *Node {

}

func insert(node *Node, key int) *Node {

}

func GenerateTree(count int) *AVLTree {

}
` +
	"```" + `

Не обязательно использовать выше описанный код, можно использовать любую реализацию, выдающую сбалансированное бинарное дерево.

## Mermaid Chart

[MermaidJS](https://mermaid-js.github.io/) is library for generating svg charts and diagrams from text.

## Вывод:

{{< columns >}}
` +
	"```" + `tpl
{{</*/* mermaid [class="text-center"]*/*/>}}
graph TD
`

const content5 = `
{{/*/* /mermaid */*/}}
` + "```" + `
{{< /columns >}}

{{< mermaid >}}
graph TD
`

type Node struct {
	value int
	left  *Node
	right *Node
}

func BinTreeBuilt() {
	output := content4

	arr := make([]int, 0)

	arr = sortIntArray(arr)

	t := time.NewTicker(5 * time.Second)

	for i := len(arr); i < 100; i++ {

	link1:
		newElement := rand.Intn(999)

		for i := 0; i < len(arr); i++ {
			if newElement == arr[i] {
				goto link1
			}
		}
		arr = append(arr, newElement)
		arr = sortIntArray(arr)

		balancedTree := recurseSplit(arr)

		output = binTreePrintRecurse(balancedTree)

		output = content4 + output + content5 + output + "\n{{< /mermaid >}}"

		err := os.WriteFile("/app/static/tasks/binary.md", []byte(fmt.Sprint(output)), 0644)
		if err != nil {
			log.Println(err)
		}

		<-t.C

	}

}

func recurseSplit(arr []int) *Node {

	if len(arr) == 0 {
		return nil
	}

	newNode := &Node{}

	if len(arr) == 1 {
		newNode.value = arr[0]
		//fmt.Println(newNode.value)
		return newNode
	}

	var max, ind int

	for i, element := range arr {
		if element > max {
			max = element
			ind = i
		}

		//fmt.Println("max element, index, arr:", max, ind, arr)

	}

	ind /= 2

	newNode.value = arr[ind]
	//fmt.Println(arr[ind])

	newNode.left = recurseSplit(arr[:ind])

	newNode.right = recurseSplit(arr[ind+1:])

	//fmt.Println(newNode.value, newNode.left, newNode.right)
	return newNode

}

func binTreePrintRecurse(bt *Node) string {
	var output string

	if bt.left != nil {
		output += strconv.Itoa(bt.value) + "-->" + strconv.Itoa(bt.left.value) + "\n"
		output += binTreePrintRecurse(bt.left)
	}

	if bt.right != nil {
		output += strconv.Itoa(bt.value) + "-->" + strconv.Itoa(bt.right.value) + "\n"
		output += binTreePrintRecurse(bt.right)
	}

	return output

}

func sortIntArray(arr []int) []int {
	var tmp int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				tmp = arr[i]
				arr[i] = arr[j]
				arr[j] = tmp

			}

		}
	}
	return arr
}

const content6 = `
---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

# Построение графа

Нужно написать воркер, который будет строить граф на текущей странице, каждые 5 секунд
От 5 до 30 элементов, случайным образом. Все ноды графа должны быть связаны.
` +
	"```" + `go
type Node struct {
    ID int
    Name string
	Form string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
    Links []*Node
}
` + "```" + `

## Mermaid Chart

[MermaidJS](https://mermaid-js.github.io/) is library for generating svg charts and diagrams from text.

## Граф

{{< columns >}}
` + "```" + `tpl
{{</*/* mermaid [class="text-center"]*/*/>}}
graph LR
`

const content7 = `
{{</*/* /mermaid */*/>}}
` + "```" + `

<--->

{{< mermaid >}}
graph LR
`
const content8 = `
{{< /mermaid >}}

{{< /columns >}}
`

type GraphElement struct {
	ID    int
	Name  string
	Form  string
	Links []*GraphElement
}

func graphRandomBuilt() {
	var output string

	form := []string{"circle", "rect", "square", "ellipse", "round-rect", "rhombus"}

	t := time.NewTicker(5 * time.Second)

	for {
		graphArr := make([]*GraphElement, 0)
		for i := 0; i < 15; i++ {
			newElement := &GraphElement{}
			newElement.ID = i
			newElement.Name = string(byte(i + 65))
			newElement.Form = form[rand.Intn(len(form)-1)]
			graphArr = append(graphArr, newElement)

			if i > 3 {

				for j := i - 3; j < i; j++ {

					if graphArr[j].Links != nil && len(graphArr[j].Links) > 1 {
						continue
					}

					rnd := rand.Intn(j)
					rnd1 := rand.Intn(i-j) + j + 1

					if graphArr[rnd].Links == nil {

						graphArr[rnd].Links = append(graphArr[rnd].Links, graphArr[j])

					} else {
						if len(graphArr[rnd].Links) < 2 && graphArr[rnd].Links[0] != graphArr[j] {
							graphArr[rnd].Links = append(graphArr[rnd].Links, graphArr[j])
						}
					}

					if graphArr[j].Links == nil {

						graphArr[j].Links = append(graphArr[j].Links, graphArr[rnd1])

					} else {

						if len(graphArr[j].Links) < 2 && graphArr[j].Links[0] != graphArr[rnd1] {

							graphArr[j].Links = append(graphArr[j].Links, graphArr[rnd1])

						}

					}

				}

			}

			if len(graphArr) > 4 {

				output = graphPrint(&graphArr)

				output = content6 + output + content7 + output + content8

				err := os.WriteFile("/app/static/tasks/graph.md", []byte(fmt.Sprint(output)), 0644)
				if err != nil {
					log.Println(err)
				}

				<-t.C

				//fmt.Println(output)

			}

		}

	}

}

func graphPrint(grArr *[]*GraphElement) string {
	var output, br1, br2 string

	mapForm := make(map[string]int)

	for i, gr := range *grArr {
		if i == len(*grArr)-1 {
			break
		}
		//fmt.Println("i", i)
		for j := 0; j < len(gr.Links); j++ {
			br1, br2 = "", ""

			if _, ok := mapForm[gr.Name]; !ok {
				output += gr.Name

				mapForm[gr.Name] = 1

				switch gr.Form {
				case "circle":
					br1 = "(("
					br2 = "))"
				case "rect":
					br1 = "["
					br2 = "]"
				case "square":
					br1 = "["
					br2 = "]"

				case "ellipse":
					br1 = "(["
					br2 = "])"

				case "round-rect":
					br1 = "("
					br2 = ")"

				case "rhombus":

					br1 = "{"
					br2 = "}"

				}

			}

			output += br1 + gr.Name + br2 + " --> "

			br1, br2 = "", ""

			if _, ok := mapForm[gr.Links[j].Name]; !ok {

				output += gr.Links[j].Name

				mapForm[gr.Links[j].Name] = 1

				switch gr.Form {
				case "circle":
					br1 = "(("
					br2 = "))"
				case "rect":
					br1 = "["
					br2 = "]"
				case "square":
					br1 = "["
					br2 = "]"

				case "ellipse":
					br1 = "(["
					br2 = "])"

				case "round-rect":
					br1 = "("
					br2 = ")"

				case "rhombus":

					br1 = "{"
					br2 = "}"

				}

			}

			output += br1 + gr.Links[j].Name + br2 + "\n"

		}

	}

	return output
}

// swagger:parameters RequestAddressSearch
type RequestAddressSearch struct {
	// in:query
	Query string `json:"query"`
}

func addressRequestUnmarshal(req []byte) (*RequestAddressSearch, error) {
	reqUnmrsh := RequestAddressSearch{}

	err := json.Unmarshal(req, &reqUnmrsh)

	return &reqUnmrsh, err
}

// swagger:parameters ResponseAddress
type ResponseAddress struct {
	// in:addresses
	Addresses []*Address `json:"addresses"`
}

type Address string

type GeocodeRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

func GeocodeRequestUnmarshal(req []byte) (*GeocodeRequest, error) {
	reqUnmrsh := GeocodeRequest{}

	err := json.Unmarshal(req, &reqUnmrsh)

	return &reqUnmrsh, err
}

type GeocodeResponse struct {
	Addresses []*Address `json:"addresses"`
}

type SuggestionsByGeo struct {
	Suggestions []Suggestion `json:"suggestions"`
}

type Suggestion struct {
	Address     string             `json:"value"`
	FullAddress string             `json:"unrestricted_value"`
	Data        map[string]*string `json:"data"`
}

type Location struct {
	Location LocationClass `json:"location"`
}
type LocationClass struct {
	CityName   string             `json:"value"`
	CityPostal string             `json:"unrestricted_value"`
	Data       map[string]*string `json:"data"`
}

type RequestAddressGeocode struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Geocode []GeocodeElement

type GeocodeElement struct {
	Source       string `json:"source"`
	Result       string `json:"result"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	Region       string `json:"region"`
	CityArea     string `json:"city_area"`
	CityDistrict string `json:"city_district"`
	Street       string `json:"street"`
	House        string `json:"house"`
	GeoLat       string `json:"geo_lat"`
	GeoLon       string `json:"geo_lon"`
	QcGeo        int64  `json:"qc_geo"`
}

func UnmarshalGeocode(data []byte) (Geocode, error) {
	var r Geocode
	err := json.Unmarshal(data, &r)
	return r, err
}

func getAddress(request *RequestAddressSearch) []byte {

	client := &http.Client{}

	var data = strings.NewReader(fmt.Sprintf("[ \"%s\" ]", request.Query)) // `[ "мск сухонска 11/-89" ]`
	req, err := http.NewRequest("POST", "https://cleaner.dadata.ru/api/v1/clean/address", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token 1283c56b743aef1844dfa719ad707dda5e3d1dfc")
	req.Header.Set("X-Secret", "1d344f41d0eb75f7160466819d19da2be383d1cb")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bodyText

}

func UnmarshalAddress(data []byte) (*AddressSearch, error) {
	var adr AddressSearch
	err := json.Unmarshal(data, &adr)
	return &adr, err
}

func getResponseAddress(addr *AddressSearch) ([]byte, error) {

	addrResp := make([]Address, 0)

	for _, el := range *addr {

		addrResp = append(addrResp, Address(el.Result))

	}

	return json.Marshal(addrResp)

}

func getSuggByGeoCode(inp *GeocodeRequest) ([]byte, error) {
	client := &http.Client{}
	lat := inp.Lat
	long := inp.Lng
	var data = strings.NewReader(fmt.Sprintf("{ \"lat\":%s, \"lon\":%s }", lat, long)) //(`{ "lat": 55.878, "lon": 37.653 }`)
	req, err := http.NewRequest("POST", "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token 1283c56b743aef1844dfa719ad707dda5e3d1dfc")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bodyText, err
}

func UnmarshalSuggestionsByGeo(data []byte) (*SuggestionsByGeo, error) {
	var suggs SuggestionsByGeo
	err := json.Unmarshal(data, &suggs)
	return &suggs, err
}

func getResponseAddressByGeo(suggs *SuggestionsByGeo) ([]byte, error) {

	addrResp := make([]Address, 0)

	for _, sugg := range *&suggs.Suggestions {

		addrResp = append(addrResp, Address(sugg.Address))

	}

	return json.Marshal(addrResp)

}

func addressGetting(inp []byte) ([]byte, error) {

	addrReqUnm, err := addressRequestUnmarshal(inp)
	if err != nil {
		fmt.Println("request unmarshaling error occured", err)

	}

	addr := getAddress(addrReqUnm)

	addrUnmr, err := UnmarshalAddress(addr)

	if err != nil {
		fmt.Println("respond unmarshalling error occured", err)

	}

	outp, err := getResponseAddress(addrUnmr)
	if err != nil {
		fmt.Println("respond marshalling error occured", err)

	}
	return outp, err

}
func geoSuggGetting(inp []byte) ([]byte, error) {

	georeq, err := GeocodeRequestUnmarshal(inp)
	if err != nil {
		fmt.Println("request unmarshaling error occured", err)
	}

	resp, err := getSuggByGeoCode(georeq)
	if err != nil {
		fmt.Println("respond error occured", err)
	}

	respUnmr, err := UnmarshalSuggestionsByGeo(resp)
	if err != nil {
		fmt.Println("respond unmarshalling error occured", err)
	}

	outp, err := getResponseAddressByGeo(respUnmr)
	if err != nil {
		fmt.Println("respond marshalling error occured", err)
	}

	return outp, err

}

type AddressSearch []AdressSearchElement

type AdressSearchElement struct {
	Source               string      `json:"source"`
	Result               string      `json:"result"`
	PostalCode           string      `json:"postal_code"`
	Country              string      `json:"country"`
	CountryISOCode       string      `json:"country_iso_code"`
	FederalDistrict      string      `json:"federal_district"`
	RegionFiasID         string      `json:"region_fias_id"`
	RegionKladrID        string      `json:"region_kladr_id"`
	RegionISOCode        string      `json:"region_iso_code"`
	RegionWithType       string      `json:"region_with_type"`
	RegionType           string      `json:"region_type"`
	RegionTypeFull       string      `json:"region_type_full"`
	Region               string      `json:"region"`
	AreaFiasID           interface{} `json:"area_fias_id"`
	AreaKladrID          interface{} `json:"area_kladr_id"`
	AreaWithType         interface{} `json:"area_with_type"`
	AreaType             interface{} `json:"area_type"`
	AreaTypeFull         interface{} `json:"area_type_full"`
	Area                 interface{} `json:"area"`
	CityFiasID           interface{} `json:"city_fias_id"`
	CityKladrID          interface{} `json:"city_kladr_id"`
	CityWithType         interface{} `json:"city_with_type"`
	CityType             interface{} `json:"city_type"`
	CityTypeFull         interface{} `json:"city_type_full"`
	City                 interface{} `json:"city"`
	CityArea             string      `json:"city_area"`
	CityDistrictFiasID   interface{} `json:"city_district_fias_id"`
	CityDistrictKladrID  interface{} `json:"city_district_kladr_id"`
	CityDistrictWithType string      `json:"city_district_with_type"`
	CityDistrictType     string      `json:"city_district_type"`
	CityDistrictTypeFull string      `json:"city_district_type_full"`
	CityDistrict         string      `json:"city_district"`
	SettlementFiasID     interface{} `json:"settlement_fias_id"`
	SettlementKladrID    interface{} `json:"settlement_kladr_id"`
	SettlementWithType   interface{} `json:"settlement_with_type"`
	SettlementType       interface{} `json:"settlement_type"`
	SettlementTypeFull   interface{} `json:"settlement_type_full"`
	Settlement           interface{} `json:"settlement"`
	StreetFiasID         string      `json:"street_fias_id"`
	StreetKladrID        string      `json:"street_kladr_id"`
	StreetWithType       string      `json:"street_with_type"`
	StreetType           string      `json:"street_type"`
	StreetTypeFull       string      `json:"street_type_full"`
	Street               string      `json:"street"`
	SteadFiasID          interface{} `json:"stead_fias_id"`
	SteadKladrID         interface{} `json:"stead_kladr_id"`
	SteadCadnum          interface{} `json:"stead_cadnum"`
	SteadType            interface{} `json:"stead_type"`
	SteadTypeFull        interface{} `json:"stead_type_full"`
	Stead                interface{} `json:"stead"`
	HouseFiasID          string      `json:"house_fias_id"`
	HouseKladrID         string      `json:"house_kladr_id"`
	HouseCadnum          string      `json:"house_cadnum"`
	HouseType            string      `json:"house_type"`
	HouseTypeFull        string      `json:"house_type_full"`
	House                string      `json:"house"`
	BlockType            interface{} `json:"block_type"`
	BlockTypeFull        interface{} `json:"block_type_full"`
	Block                interface{} `json:"block"`
	Entrance             interface{} `json:"entrance"`
	Floor                interface{} `json:"floor"`
	FlatFiasID           string      `json:"flat_fias_id"`
	FlatCadnum           string      `json:"flat_cadnum"`
	FlatType             string      `json:"flat_type"`
	FlatTypeFull         string      `json:"flat_type_full"`
	Flat                 string      `json:"flat"`
	FlatArea             string      `json:"flat_area"`
	SquareMeterPrice     string      `json:"square_meter_price"`
	FlatPrice            string      `json:"flat_price"`
	PostalBox            interface{} `json:"postal_box"`
	FiasID               string      `json:"fias_id"`
	FiasCode             string      `json:"fias_code"`
	FiasLevel            string      `json:"fias_level"`
	FiasActualityState   string      `json:"fias_actuality_state"`
	KladrID              string      `json:"kladr_id"`
	CapitalMarker        string      `json:"capital_marker"`
	Okato                string      `json:"okato"`
	Oktmo                string      `json:"oktmo"`
	TaxOffice            string      `json:"tax_office"`
	TaxOfficeLegal       string      `json:"tax_office_legal"`
	Timezone             string      `json:"timezone"`
	GeoLat               string      `json:"geo_lat"`
	GeoLon               string      `json:"geo_lon"`
	BeltwayHit           string      `json:"beltway_hit"`
	BeltwayDistance      interface{} `json:"beltway_distance"`
	QcGeo                int64       `json:"qc_geo"`
	QcComplete           int64       `json:"qc_complete"`
	QcHouse              int64       `json:"qc_house"`
	Qc                   int64       `json:"qc"`
	UnparsedParts        interface{} `json:"unparsed_parts"`
	Metro                []Metro     `json:"metro"`
}

type Metro struct {
	Distance float64 `json:"distance"`
	Line     string  `json:"line"`
	Name     string  `json:"name"`
}

const (
	swaggerTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <script src="//unpkg.com/swagger-ui-dist@3/swagger-ui-standalone-preset.js"></script>
    <!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui-standalone-preset.js"></script> -->
    <script src="//unpkg.com/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
    <!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui-bundle.js"></script> -->
    <link rel="stylesheet" href="//unpkg.com/swagger-ui-dist@3/swagger-ui.css" />
    <!-- <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui.css" /> -->
	<style>
		body {
			margin: 0;
		}
	</style>
    <title>Swagger</title>
</head>
<body>
    <div id="swagger-ui"></div>
    <script>
        window.onload = function() {
          SwaggerUIBundle({
            url: "./swagger.json?{{.Time}}",
            dom_id: '#swagger-ui',
            presets: [
              SwaggerUIBundle.presets.apis,
              SwaggerUIStandalonePreset
            ],
            layout: "StandaloneLayout"
          })
        }
    </script>
</body>
</html>
`
)
