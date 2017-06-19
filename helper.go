package huobi

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func httpRequest(session *Session, method string) (*http.Request, error) {
	req, err := http.NewRequest("GET", session.Addr, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("method", method)
	q.Add("access_key", session.AccessKeyId)
	now := strconv.FormatInt(time.Now().Unix(), 10)
	q.Add("secret_key", session.SecretKey)
	q.Add("created", now)
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func signRequest(req *http.Request) {
	q := req.URL.Query()
	b := []byte(q.Encode())
	hash := md5.Sum(b)
	q.Add("sign", hex.EncodeToString(hash[:]))
	req.URL.RawQuery = q.Encode()
}

func sendRequest(req *http.Request) (string, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 { // OK
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(bodyBytes), nil
	}
	return "", errors.New("server error")
}
