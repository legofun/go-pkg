package http

import (
	"bytes"
	"crypto/tls"
	"errors"
	"github.com/golang/glog"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type httpHandle struct {
	url         string
	contentType string
	timeout     time.Duration
	retryTimes  int8
	tlsVerify   bool
}

func NewHttpHandle(url string, opts ...httpOptions) *httpHandle {
	hd := &httpHandle{
		url: url,
	}

	opts = append([]httpOptions{defaultHttpOption()}, opts...)
	for _, v := range opts {
		v.apply(hd)
	}

	return hd
}

//发送Post请求
func (h *httpHandle) Post(data []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", h.url, bytes.NewReader(data))
	if err != nil {
		glog.Error("http new request error: ", err.Error(), ", ", string(data))
		return nil, err
	}

	if len(h.contentType) > 0 {
		request.Header.Set("Content-Type", h.contentType)
	}

	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: h.timeout,
		}).DialContext,
	}

	//跳过证书验证
	if !h.tlsVerify {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(request)
	if err != nil {
		glog.Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		glog.Error(resp.Status)
		return nil, errors.New(resp.Status)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Error(err.Error())
		return nil, err
	}
	return respBytes, nil
}

//发送Get请求
func (h *httpHandle) Get() ([]byte, error) {
	client := &http.Client{Timeout: h.timeout}
	resp, err := client.Get(h.url)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	var (
		buffer [1024]byte
		n      int
	)
	result := bytes.NewBuffer(nil)
	for {
		n, err = resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			glog.Error(err)
			return nil, err
		}
	}

	return result.Bytes(), nil
}
