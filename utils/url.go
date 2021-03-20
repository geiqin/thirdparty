package utils

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

// build url with param
type UrlBuilder struct {
	baseUrl string
	params  url.Values
}

func NewUrlBuilder(baseUrl string) *UrlBuilder {
	uv, err := url.ParseRequestURI(baseUrl)
	builder := &UrlBuilder{}
	if err != nil {
		log.Println(err)
		return builder
	}
	urls := strings.SplitN(uv.String(), "?", 2)
	builder.baseUrl = urls[0]
	builder.params = uv.Query()
	return builder
}

func (this *UrlBuilder) AddParam(key string, value interface{}) *UrlBuilder {
	if key == "" {
		return this
	}
	this.params.Add(key, fmt.Sprint(value))
	return this
}

func (this *UrlBuilder) Build() string {
	if this.baseUrl == "" {
		return ""
	}
	if len(this.params) == 0 {
		return this.baseUrl
	}
	return this.baseUrl + "?" + this.params.Encode()
}
