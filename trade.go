package huobi

import "strconv"

type Session struct {
	Addr             string
	AccessKeyId      string
	SecretKey        string
	SignatureVersion int
}

func (this *Session) GetAccountInfo() (string, error) {
	req, err := httpRequest(this, "get_account_info")
	if err != nil {
		return "", err
	}
	signRequest(req)
	bodyStr, err := sendRequest(req)
	return bodyStr, err
}

func (this *Session) GetOrders(coinType CoinType) (string, error) {
	req, err := httpRequest(this, "get_orders")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("coin_type", string(coinType))
	req.URL.RawQuery = q.Encode()
	signRequest(req)
	bodyStr, err := sendRequest(req)
	return bodyStr, err
}

func (this *Session) OrderInfo(id string, coinType CoinType) (string, error) {
	req, err := httpRequest(this, "order_info")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("coin_type", string(coinType))
	q.Add("id", id)
	req.URL.RawQuery = q.Encode()
	signRequest(req)
	bodyStr, err := sendRequest(req)
	return bodyStr, err
}

func (this *Session) Buy(price float64, amount float64, coinType CoinType, tradePassword string) (string, error) {
	req, err := httpRequest(this, "buy")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("coin_type", string(coinType))
	q.Add("price", strconv.FormatFloat(price, 'G', -1, 64))
	q.Add("amount", strconv.FormatFloat(amount, 'G', -1, 64))

	req.URL.RawQuery = q.Encode()
	signRequest(req)

	q = req.URL.Query()
	q.Add("trade_password", tradePassword)
	req.URL.RawQuery = q.Encode()
	bodyStr, err := sendRequest(req)

	return bodyStr, err
}

func (this *Session) Sell(price float64, amount float64, coinType CoinType, tradePassword string) (string, error) {
	req, err := httpRequest(this, "sell")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("coin_type", string(coinType))
	q.Add("price", strconv.FormatFloat(price, 'G', -1, 64))
	q.Add("amount", strconv.FormatFloat(amount, 'G', -1, 64))

	req.URL.RawQuery = q.Encode()
	signRequest(req)

	q = req.URL.Query()
	q.Add("trade_password", tradePassword)
	req.URL.RawQuery = q.Encode()
	bodyStr, err := sendRequest(req)

	return bodyStr, err
}

func (this *Session) BuyMarket(amount float64, coinType CoinType, tradePassword string) (string, error) {
	req, err := httpRequest(this, "buy_market")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("coin_type", string(coinType))
	q.Add("amount", strconv.FormatFloat(amount, 'G', -1, 64))

	req.URL.RawQuery = q.Encode()
	signRequest(req)

	q = req.URL.Query()
	q.Add("trade_password", tradePassword)
	req.URL.RawQuery = q.Encode()
	bodyStr, err := sendRequest(req)

	return bodyStr, err
}

func (this *Session) SellMarket(amount float64, coinType CoinType, tradePassword string) (string, error) {
	req, err := httpRequest(this, "sell_market")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("coin_type", string(coinType))
	q.Add("amount", strconv.FormatFloat(amount, 'G', -1, 64))

	req.URL.RawQuery = q.Encode()
	signRequest(req)

	q = req.URL.Query()
	q.Add("trade_password", tradePassword)
	req.URL.RawQuery = q.Encode()
	bodyStr, err := sendRequest(req)

	return bodyStr, err
}

func (this *Session) CancelOrder(id string, coinType CoinType) (string, error) {
	req, err := httpRequest(this, "cancel_order")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("coin_type", string(coinType))
	q.Add("id", id)
	req.URL.RawQuery = q.Encode()
	signRequest(req)
	bodyStr, err := sendRequest(req)
	return bodyStr, err
}

func (this *Session) GetNewDealOrders(coinType CoinType) (string, error) {
	req, err := httpRequest(this, "get_new_deal_orders")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("coin_type", string(coinType))
	req.URL.RawQuery = q.Encode()
	signRequest(req)
	bodyStr, err := sendRequest(req)
	return bodyStr, err
}

func (this *Session) GetOrderIdByTradeId(trade_id string, coinType CoinType) (string, error) {
	req, err := httpRequest(this, "get_order_id_by_trade_id")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("coin_type", string(coinType))
	q.Add("trade_id", trade_id)
	req.URL.RawQuery = q.Encode()
	signRequest(req)
	bodyStr, err := sendRequest(req)
	return bodyStr, err
}

func (this *Session) WithdrawCoin(address string, amount float64, coinType CoinType, tradePassword string) (string, error) {
	req, err := httpRequest(this, "withdraw_coin")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("coin_type", string(coinType))
	q.Add("amount", strconv.FormatFloat(amount, 'G', -1, 64))
	q.Add("address", address)

	req.URL.RawQuery = q.Encode()
	signRequest(req)

	q = req.URL.Query()
	q.Add("trade_password", tradePassword)
	req.URL.RawQuery = q.Encode()

	bodyStr, err := sendRequest(req)
	return bodyStr, err
}

func (this *Session) CancelWithdrawCoin(withdraw_coin_id string) (string, error) {
	req, err := httpRequest(this, "cancel_withdraw_coin")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("withdraw_coin_id", withdraw_coin_id)
	req.URL.RawQuery = q.Encode()
	signRequest(req)
	bodyStr, err := sendRequest(req)
	return bodyStr, err
}

func (this *Session) GetWithdrawCoinResult(withdraw_coin_id string) (string, error) {
	req, err := httpRequest(this, "get_withdraw_coin_result")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("withdraw_coin_id", withdraw_coin_id)
	req.URL.RawQuery = q.Encode()
	signRequest(req)
	bodyStr, err := sendRequest(req)
	return bodyStr, err
}

func (this *Session) Transfer(account_form AccountType, account_to AccountType, amount float64, coinType CoinType) (string, error) {
	req, err := httpRequest(this, "transfer")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("account_form", string(account_form))
	q.Add("account_to", string(account_to))
	q.Add("coin_type", string(coinType))
	q.Add("amount", strconv.FormatFloat(amount, 'G', -1, 64))

	req.URL.RawQuery = q.Encode()
	signRequest(req)

	bodyStr, err := sendRequest(req)
	return bodyStr, err
}

func (this *Session) Loan(amount float64, loan_type LoanType) (string, error) {

	req, err := httpRequest(this, "loan")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("loan_type", string(loan_type))
	q.Add("amount", strconv.FormatFloat(amount, 'G', -1, 64))

	req.URL.RawQuery = q.Encode()
	signRequest(req)

	bodyStr, err := sendRequest(req)
	return bodyStr, err
}

func (this *Session) Repayment(loan_id string, amount float64) (string, error) {
	req, err := httpRequest(this, "repayment")
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("amount", strconv.FormatFloat(amount, 'G', -1, 64))
	q.Add("loan_id", loan_id)
	req.URL.RawQuery = q.Encode()
	signRequest(req)

	bodyStr, err := sendRequest(req)
	return bodyStr, err
}

func (this *Session) GetLoanAvailable() (string, error) {
	req, err := httpRequest(this, "get_loan_available")
	if err != nil {
		return "", err
	}
	signRequest(req)
	bodyStr, err := sendRequest(req)
	return bodyStr, err
}

func (this *Session) GetLoans() (string, error) {
	req, err := httpRequest(this, "get_loans")
	if err != nil {
		return "", err
	}
	signRequest(req)
	bodyStr, err := sendRequest(req)
	return bodyStr, err
}
