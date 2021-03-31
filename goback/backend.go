package main

import ( 	
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"log"
	"net/http"
	"io/ioutil"
	"time"
	"encoding/json"
	"strconv"
)


// "Meta Data": {
// 	"1. Information": "Daily Time Series with Splits and Dividend Events",   
// 	"2. Symbol": "MSFT",
// 	"3. Last Refreshed": "2021-03-30",
// 	"4. Output Size": "Compact",
// 	"5. Time Zone": "US/Eastern"
// },

// "2020-11-04": {
// 	"1. open": "214.02",
// 	"2. high": "218.32",
// 	"3. low": "212.4185",
// 	"4. close": "216.39",
// 	"5. adjusted close": "215.323650798",
// 	"6. volume": "42311777",
// 	"7. dividend amount": "0.0000",
// 	"8. split coefficient": "1.0"
// }

type STOCK_DATA struct {
	MetaData METADATA_RESPONSE `json:"Meta Data"`
	TimeSeriesData map[string] TIME_SERIES `json:"Time Series (Daily)"`
}

type METADATA_RESPONSE struct {
	Information string `json:"1. Information"`
	Symbol string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize string `json:"4. Output Size"`
	TimeZone string `json:"5. Time Zone"`
}

type TIME_SERIES struct {
	Open string `json:"1. open"`
	High string `json:"2. high"`
	Low string `json:"3. low"`
	Close string `json:"4. close"`
	AdjustedClose string `json:"5. adjusted close"`
	Volume string `json:"6. volume"`
	DividendAmount string `json:"7. dividend amount"`
	SplitCoefficient string `json:"8. split coefficient"`
}


var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
        panic(err.Error())
    }

    return json.Unmarshal(body, target)
}

func main() {
	r := gin.Default()
	r.GET("/stonks", func(c *gin.Context) {

		// Grab ENV vars and make endpoint variable
		apikey := os.Getenv("API_KEY")
		symbol := os.Getenv("SYMBOL")
		endpoint := fmt.Sprintf("https://www.alphavantage.co/query?apikey=%s&function=TIME_SERIES_DAILY_ADJUSTED&symbol=%s", apikey, symbol)
		ndays, err := strconv.Atoi(os.Getenv("NDAYS"))
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("APIKey: %s, Symbol: %s, NDays: %d\n", apikey, symbol, ndays)

		// Call Stock API and Unmarshal results
		var stockdata STOCK_DATA
		getJson(endpoint, &stockdata)

		// Loop through stock data for NDAYS
		// TODO: Apparently they're not in order.
		i := 0
		var price_list []float64
		var total = 0.0
		for _, stockdata := range stockdata.TimeSeriesData {
			if i == ndays {
				break
			}
			// fmt.Printf("Date: %v\n", date)
			// fmt.Printf("Data: %v\n", stockdata)
			if closePrice, err := strconv.ParseFloat(stockdata.Close, 64); err == nil {
				price_list = append(price_list, closePrice)
				total += closePrice
			}
			i++
		}

		var average = total / float64(ndays)

		output := fmt.Sprintf("%v %v, average=%.2f", symbol, price_list, average)

		c.JSON(200, gin.H{
			"message": output,
		})
	})
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}