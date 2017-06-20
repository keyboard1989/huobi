package lab

import "huobi"

// btc ltc
const addr = "wss://api.huobi.com/ws"

// eth
const labaddr = "wss://be.huobi.com/ws"

func NewConnection(marketType huobi.MarketType) *Connection {
	c := Connection{}
	if marketType == huobi.CNYBTC || marketType == huobi.CNYLTC {
		c.addr = addr
	} else if marketType == huobi.CNYETH {
		c.addr = labaddr
	} else {
		panic("coin type error")
	}
	c.init()
	return &c
}
