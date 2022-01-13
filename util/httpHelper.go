package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送GET请求
// url       请求地址
// return    请求返回的内容
func Get(url string) (string, error) {
	// 超时时间：10秒
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
	}

	return result.String(), nil
}

func GetObj(url string, obj interface{}) error {
	strObj, err := Get(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(strObj), obj)
	if err != nil {
		return err
	}
	return nil
}

// 发送POST请求
// url         	请求地址
// data			POST请求提交的数据
// contentType 	请求体格式，如：application/json
// return   	请求返回的内容
func Post(url string, data interface{}, contentType string) (string, error) {
	// 超时时间：10秒
	client := &http.Client{Timeout: 10 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}

func PostObj(url string, data interface{}, contentType string, obj interface{}) error {
	strObj, err := Post(url, data, contentType)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(strObj), obj)
	if err != nil {
		return err
	}
	return nil
}
