package util

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

func HttpPost(notifyUrl string, dataMap map[string]string, appKey string) (*http.Response, error) {
	if !IsValidUrl(notifyUrl) {
		return nil, errors.New("URL ERROR")
	}
	keys := make([]string, 0, len(dataMap))
	values := make([]string, 0, len(dataMap))
	for k := range dataMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	datas := url.Values{}
	for _, k := range keys {
		value := dataMap[k]
		datas.Add(k, value)
		values = append(values, value)
	}
	plainText := fmt.Sprintf("%s|%s", strings.Join(values, "|"), appKey)
	hasher := sha512.New()
	hasher.Write([]byte(plainText))
	sign := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	datas.Add("Sign", sign)
	return http.PostForm(notifyUrl, datas)
}
