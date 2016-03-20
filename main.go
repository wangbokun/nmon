package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var str string

func md5sum(AppSecret string) string {

	srcData := []byte(AppSecret)
	hash := md5.New()
	hash.Write(srcData)
	cipherText2 := hash.Sum(nil)
	hexText := make([]byte, 32)
	hex.Encode(hexText, cipherText2)

	return string(hexText)
}

func reg_wecenter(reg_url string, user_name string, password string, email string) {
	v := url.Values{"user_name": {user_name}, "password": {password}, "email": {email}}

	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}
	req, _ := http.NewRequest("POST", reg_url, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func login_wecenter(login_url string, user_name string, password string) string {

	v := url.Values{"user_name": {user_name}, "password": {password}}

	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}
	req, _ := http.NewRequest("POST", login_url, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	resp, _ := client.Do(req)

	cookies := resp.Cookies()

	for _, v := range cookies {
		str = v.Name + "=" + v.Value + ";" + str
	}

	fmt.Println("cookies:", cookies)

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("this is login :", string(data))
	return str

}

func head_upload_wecenter(head_url string, str string, sex string) {
	fmt.Println(head_url)
	v := url.Values{"sex": {sex}}

	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}
	req, _ := http.NewRequest("POST", head_url, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Set("Cookie", str)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("this is head :", string(data))
}

func main() {

	AppSecret := "12884789df747d7affbcd6a7cadd9359"
	md5String := md5sum("account" + AppSecret)

	user_name := "wang23523"
	password := "123456"
	email := "411241127@qq.com"

	url := "http://wecenter.dev.hihwei.com/"

	register := "api/account/register_process/?"
	login := "api/account/login_process/?"
	head := "api/people/profile_setting/?"

	reg_wecenter(url+register+"mobile_sign="+md5String, user_name, password, email)

	str := login_wecenter(url+login+"mobile_sign="+md5String, user_name, password)

	sex := "3"
	md5String = md5sum("people" + AppSecret)

	head_upload_wecenter(url+head+"mobile_sign="+md5String, str, sex)
}
