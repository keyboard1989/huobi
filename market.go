package huobi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const cnyMarket = "http://api.huobi.com/staticmarket/"
const usdMarket = "http://api.huobi.com/staticmarket/"

func GetKline(interval string, length uint, marketType MarketType) (string, error) {
	var addr string
	switch marketType {
	case CNYBTC:
		addr = fmt.Sprintf(cnyMarket+"btc_kline_%s_json.js?length=%d", interval, length)
	case CNYLTC:
		addr = fmt.Sprintf(cnyMarket+"ltc_kline_%s_json.js?length=%d", interval, length)
	case USDBTC:
		addr = fmt.Sprintf(usdMarket+"btc_kline_%s_json.js?length=%d", interval, length)

	}
	resp, err := http.Get(addr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", err2
	}
	return string(b), nil

}

func GetTickerRealTime(length uint, marketType MarketType) (string, error) {
	var addr string
	switch marketType {
	case CNYBTC:
		addr = fmt.Sprintf(cnyMarket+"ticker_btc_json.js?length=%d", length)
	case CNYLTC:
		addr = fmt.Sprintf(cnyMarket+"ticker_ltc_json.js?length=%d", length)
	case USDBTC:
		addr = fmt.Sprintf(usdMarket+"ticker_btc_json.js?length=%d", length)

	}
	resp, err := http.Get(addr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", err2
	}
	return string(b), nil
}

func GetDepth(depth uint, marketType MarketType) (string, error) {
	var addr string
	switch marketType {
	case CNYBTC:
		addr = fmt.Sprintf(cnyMarket+"depth_btc_%d.js", depth)
	case CNYLTC:
		addr = fmt.Sprintf(cnyMarket+"depth_ltc_%d.js", depth)
	case USDBTC:
		addr = fmt.Sprintf(usdMarket+"depth_btc_%d.jsd", depth)

	}
	resp, err := http.Get(addr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", err2
	}
	return string(b), nil
}

func GetOrderBookAndTAS(marketType MarketType) (string, error) {
	var addr string
	switch marketType {
	case CNYBTC:
		addr = fmt.Sprintf(cnyMarket + "detail_btc.js")
	case CNYLTC:
		addr = fmt.Sprintf(cnyMarket + "detail_ltc.js")
	case USDBTC:
		addr = fmt.Sprintf(usdMarket + "detail_btc.jsd")

	}
	resp, err := http.Get(addr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", err2
	}
	return string(b), nil
}
