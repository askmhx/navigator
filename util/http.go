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
	datas := MapToUrlValues(dataMap)
	datas.Add("Sign", GetSign(dataMap, appKey))
	return http.PostForm(notifyUrl, datas)
}

func GetSign(dataMap map[string]string, appKey string) string {
	keys := make([]string, 0, len(dataMap))
	values := make([]string, 0, len(dataMap))
	for k := range dataMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		value := dataMap[k]
		values = append(values, value)
	}
	plainText := fmt.Sprintf("%s|%s", strings.Join(values, "|"), appKey)
	hasher := sha512.New()
	hasher.Write([]byte(plainText))
	sign := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sign
}

func MapToUrlValues(dataMap map[string]string) url.Values {
	datas := url.Values{}
	for k := range dataMap {
		datas.Add(k, dataMap[k])
	}
	return datas
}
