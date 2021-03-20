package oauth

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

func HttpGet(getUrl string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Transport: tr, Jar: cookieJar}

	resp, err := client.Get(getUrl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
func HttpPost(getUrl string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	reqest, err := http.NewRequest("POST", getUrl, nil)

	if err != nil {
		return "", err
	}

	response, _ := client.Do(reqest)

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
