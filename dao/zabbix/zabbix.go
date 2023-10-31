package zabbix

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"multiovirt-admin/settings"
	"net/http"
	"time"
)

type ZabbixClient struct {
	url      string
	username string
	password string
	client   *http.Client
}

type ZabbixAPI struct {
	zabbixclient *ZabbixClient
	auth         string
}

func InitZabbix(cfg *settings.ZabbixConfig) (*ZabbixAPI, error) {
	zc, err := NewZabbixClient(cfg.Url, cfg.User, cfg.Password)
	if err != nil {
		fmt.Println("NewZabbixClient error!")
	}
	auth, _ := zc.authenticate()
	if auth == "" {
		return nil, nil
	}
	return &ZabbixAPI{
		zabbixclient: zc.zabbixclient,
		auth:         auth,
	}, nil
}

func NewZabbixClient(url string, username string, password string) (*ZabbixAPI, error) {
	tr := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        10,               // 控制最大空闲连接数
			MaxIdleConnsPerHost: 10,               // 控制每个目标主机的最大空闲连接数
			IdleConnTimeout:     30 * time.Second, // 控制空闲连接的超时时间
		},
	}
	if url != "" && username != "" && password != "" {
		return &ZabbixAPI{
			zabbixclient: &ZabbixClient{
				url:      url,
				username: username,
				password: password,
				client:   tr,
			},
		}, nil
	}
	zap.L().Error("Input parameter error! Please config url,username,password!")
	return nil, errors.New("Input parameter error!")
}

func (z *ZabbixAPI) authenticate() (string, error) {
	authRequest := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "user.login",
		"params": map[string]string{
			"user":     z.zabbixclient.username,
			"password": z.zabbixclient.password,
		},
		"id": 1,
	}
	jsonData, err := json.Marshal(authRequest)
	if err != nil {
		zap.L().Error("zabbix authenticate jsonData error!", zap.Error(err))
		return "", err
	}
	resp, err := z.zabbixclient.client.Post(z.zabbixclient.url, "application/json-rpc", bytes.NewBuffer(jsonData))
	if err != nil {
		zap.L().Error("zabbixclient Post error!", zap.Error(err))
		return "", err
	}
	defer resp.Body.Close()
	var authResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		zap.L().Error("zabbixclient authenticate json.NewDecoder error!", zap.Error(err))
		return "", err
	}
	return authResponse["result"].(string), nil
}
