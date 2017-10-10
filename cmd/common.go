package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/k82cn/activiti-client/api"
)

const (
	DefaultTimeFormat = "01/02/06 15:04"
)

func FormatTime(t *time.Time) string {
	if t == nil {
		return ""
	}

	return t.Format(DefaultTimeFormat)
}

type ActivitiClient struct {
	User     string
	Password string
	BaseURL  string
	client   *http.Client
}

var Client *ActivitiClient

func InitClient(user, password, url string) {
	Client = &ActivitiClient{
		User:     user,
		Password: password,
		BaseURL:  url,
		client:   &http.Client{},
	}
}

func (ac *ActivitiClient) Post(url string, data interface{}, obj interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", ac.BaseURL+"/"+url, bytes.NewReader(d))
	if err != nil {
		return err
	}
	req.SetBasicAuth(ac.User, ac.Password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := ac.client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if obj != nil {
		if err := json.Unmarshal(body, obj); err != nil {
			return err
		}
	}

	return nil
}

func (ac *ActivitiClient) Get(url string, obj interface{}) error {
	req, err := http.NewRequest("GET", ac.BaseURL+"/"+url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(ac.User, ac.Password)
	resp, err := ac.client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, obj); err != nil {
		return err
	}

	return nil
}

func CheckErr(err error, errmsg api.ErrMessageInterface) {
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	if errmsg != nil {
		if len(errmsg.GetException()) != 0 && len(errmsg.GetMessage()) != 0 {
			fmt.Printf("%s: %s\n", errmsg.GetMessage(), errmsg.GetException())
			os.Exit(1)
		}
	}
}
