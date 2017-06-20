package lab

import (
	"huobi"
	"log"

	"net/url"

	"bytes"
	"compress/gzip"
	"io/ioutil"

	"encoding/json"

	"time"

	"fmt"

	"github.com/gorilla/websocket"
)

type Connection struct {
	addr    string
	sub     string
	id      string
	market  huobi.MarketType
	c       *websocket.Conn
	handler func([]byte)
}

func (this *Connection) init() {

	u := url.URL{Scheme: "ws", Host: "api.huobi.com", Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	log.Printf("connecting to %s", this.addr)

	this.c = c

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			buf := bytes.NewReader(message)
			reader, err := gzip.NewReader(buf)

			if err != nil {
				log.Println(err)
				return
			}
			b, err := ioutil.ReadAll(reader)
			if err != nil {
				log.Println(err)
				return
			}
			var dat map[string]interface{}
			json.Unmarshal(b, &dat)
			_, ok := dat["ping"]
			if ok {
				dat2 := make(map[string]interface{})
				dat2["pong"] = dat["ping"]
				this.c.WriteJSON(dat2)
			} else {
				if _, ok := dat["subbed"]; ok {
					log.Println("订阅成功: ", dat["subbed"])
				} else if _, ok := dat["unsub"]; ok {
					log.Println("取消订阅: ", dat["unsub"])
				} else if _, ok := dat["err-msg"]; ok {
					log.Println(dat["err-msg"])
				} else if this.handler != nil {
					this.handler(b)
				}

			}
		}
	}()
}

func (this *Connection) ChangeSubDepth(step string, handler func([]byte)) {
	symbol := market2Symbol(this.market)
	subStr := fmt.Sprintf("market.%s.depth.%s", symbol, step)
	id, e := sub2Server(this, subStr)
	if e != nil {
		panic(e)
	}
	this.id = id
	this.handler = handler
	this.sub = subStr

}

func (this *Connection) ChangeSubDetail(handler func([]byte)) {
	symbol := market2Symbol(this.market)
	subStr := fmt.Sprintf("market.%s.detail", symbol)
	id, e := sub2Server(this, subStr)
	if e != nil {
		panic(e)
	}
	this.id = id
	this.handler = handler
	this.sub = subStr
}

func (this *Connection) ChangeSubKline(period string, handler func([]byte)) {
	symbol := market2Symbol(this.market)
	subStr := fmt.Sprintf("market.%s.kline.%s", symbol, period)
	id, e := sub2Server(this, subStr)
	if e != nil {
		panic(e)
	}
	this.id = id
	this.handler = handler
	this.sub = subStr
}

func (this *Connection) ChangeSubTrade(handler func([]byte)) {
	symbol := market2Symbol(this.market)
	subStr := fmt.Sprintf("market.%s.trade.detail", symbol)
	id, e := sub2Server(this, subStr)
	if e != nil {
		panic(e)
	}
	this.id = id
	this.handler = handler
	this.sub = subStr
}

func (this *Connection) Unsub() {
	topic := map[string]interface{}{}
	topic["id"] = this.id
	topic["unsub"] = this.sub
	e := this.c.WriteJSON(topic)
	if e != nil {
		panic(e)
	}
	this.handler = nil
	this.id = ""
	this.sub = ""
}

func (this *Connection) Close() {
	this.c.Close()
}

func sub2Server(c *Connection, subStr string) (string, error) {
	topic := map[string]interface{}{}
	id := "id" + time.Now().String()
	topic["id"] = id
	topic["sub"] = subStr
	e := c.c.WriteJSON(topic)
	if e != nil {
		return "", e
	}
	return id, nil
}

func market2Symbol(marketType huobi.MarketType) string {
	symbol := "btccny"
	switch marketType {
	case huobi.CNYBTC:
		symbol = "btccny"
	case huobi.CNYLTC:
		symbol = "ltccny"
	case huobi.CNYETH:
		symbol = "ethcny"
	}
	return symbol
}
