package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const HASH_TO_TEST = "Qmaisz6NMhDB51cCvNWa1GMS7LU1pAxdF4Ld6Ft9kZEP2a"
const HASH_STRING = "Hello from IPFS Gateway Checker"
const GATEWAY_LIST_URL = "https://raw.githubusercontent.com/ipfs/public-gateway-checker/master/gateways.txt"

var r = strings.NewReplacer(":hash", HASH_TO_TEST)

// Get get only one active gateway url
func Get() (r string, err error) {
	list := gatewayList()
	if len(list) < 1 {
		return "", fmt.Errorf("no gateway to test")
	}
	ch := make(chan string)
	for _, url := range list {
		go func(url string) {
			check(url, ch)
		}(url)
	}
	return <-ch, nil
}

// ActiveList get all active gateway list
func ActiveList() (result []string, err error) {
	list := gatewayList()
	if len(list) < 1 {
		return result, fmt.Errorf("no gateway to test")
	}

	ch := make(chan string)
	var wg sync.WaitGroup
	for _, url := range list {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			check(url, ch)
		}(url)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for validGw := range ch {
		result = append(result, validGw)
	}
	return result, nil
}

func check(url string, ch chan string) {
	realUrl := r.Replace(url)
	response, err := httpGet(realUrl)
	if err != nil {
		//fmt.Printf("query gateway err: %s\n", err)
		return
	}
	if strings.TrimSpace(string(response)) == HASH_STRING {
		ch <- url
	}
}

func gatewayList() (l []string) {
	response, err := httpGet(GATEWAY_LIST_URL)
	if err != nil {
		//fmt.Printf("get gateway list err: %s\n", err)
		os.Exit(1)
	}
	json.Unmarshal(response, &l)
	return
}

func httpGet(url string) (content []byte, err error) {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	response, err := netClient.Get(url)
	if err == nil {
		defer response.Body.Close()
		content, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("failed to parse response: %s\n", err)
			os.Exit(1)
		}
	}
	return
}
