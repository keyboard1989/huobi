package huobi

const defaultAPIAddress = "https://api.huobi.com/apiv3"

type CoinType string

const BTC CoinType = "1"
const LTC CoinType = "2"

type AccountType string

const CNY AccountType = "1"
const USD AccountType = "2"

type LoanType string

const LOANBTC LoanType = "1"
const LOANLTC LoanType = "2"
const LOANCNY LoanType = "3"
const LOANUSD LoanType = "4"

type MarketType string

const CNYBTC MarketType = "1"
const CNYLTC MarketType = "2"
const USDBTC MarketType = "3"

func NewSession(accessKey string, secertKey string) *Session {
	return &Session{defaultAPIAddress, accessKey, secertKey, 2}
}
