package secret

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type InkeV1Signature struct {
	authFormat string
}

func init() {
	s := InkeV1Signature{}
	s.authFormat = "InkeV1 %s:%s"
	RegisterSign(s.Protocol(), &s)

}

func (s *InkeV1Signature) Protocol() string {
	return "InkeV1"

}

// 签名接口
func (s *InkeV1Signature) Signature(request *http.Request, accessKeyId, accessSecret string, bundleID string) (authorization string, bodyMD5 string, err error) {

	bodyMD5 = ""

	// 获取body
	var body []byte
	body, err = ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}
	request.Body = ioutil.NopCloser(bytes.NewReader(body))
	// 计算body的md5
	bodyMD5 = Md5base64(body)

	rawQuery := request.URL.RawQuery
	// 去掉nginx上附加的xrealip
	realipIdx := strings.LastIndex(rawQuery, "&xrealip")
	if realipIdx > 0 {
		rawQuery = rawQuery[:realipIdx]
	}

	//uri := fmt.Sprintf("%s?%s", request.URL.Path, rawQuery)
	// 需要签名的字段
	stringToSign := bundleID + "\n" +
		bodyMD5 + "\n" +
		strings.ToUpper(request.Method) + "\n" +
		request.Header.Get("Client-TimeStamp") + "\n" +
		rawQuery

	// 计算签名
	signature := Sha1base64(stringToSign, accessSecret)
	authorization = fmt.Sprintf(s.authFormat, accessKeyId, signature)
	return authorization,bodyMD5,nil
}


func Sha1base64(strToSign string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(strToSign))
	return string(base64.StdEncoding.EncodeToString(h.Sum(nil)))
}

func Md5base64(body []byte) string {
	m := md5.New()
	m.Write(body)
	return string(base64.StdEncoding.EncodeToString(m.Sum(nil)))
}