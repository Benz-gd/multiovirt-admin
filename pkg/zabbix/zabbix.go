package zabbix

import (
	"bytes"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"multiovirt-admin/settings"
	"net/http"
)

type ZabbixClient struct {
	url      string
	username string
	password string
}
type JsonRPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Error   ZabbixError `json:"error"`
	Result  interface{} `json:"result"`
	Id      int         `json:"id"`
}

type JsonRPCRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Auth    string      `json:"auth,omitempty"`
	Id      int         `json:"id"`
}

type ZabbixError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (z *ZabbixError) Error() string {
	return z.Data
}

type API struct {
	url    string
	user   string
	passwd string
	id     int
	auth   string
}

func InitZabbix(cfg *settings.ZabbixConfig) (*ZabbixClient, error) {
	return newZabbixClient(cfg.Url, cfg.User, cfg.Password)
}

func newZabbixClient(url string, username string, password string) (*ZabbixClient, error) {
	if url != "" && username != "" && password != "" {
		return &ZabbixClient{
			url:      url,
			username: username,
			password: password,
		}, nil
	}
	zap.L().Error("Input parameter error! Please config url,username,password!")
	return nil, errors.New("Input parameter error!")
}

func (z *ZabbixClient) authenticate() (string, error) {
	authRequest := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "user.login",
		"params": map[string]string{
			"user":     z.username,
			"password": z.password,
		},
		"id": 1,
	}
	jsonData, err := json.Marshal(authRequest)
	if err != nil {
		zap.L().Error("zabbix authenticate jsonData error!", zap.Error(err))
		return "", err
	}
	resp, err := http.Post(z.url, "application/json-rpc", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
}
