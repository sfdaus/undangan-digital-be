package helpers

import (
	"agree-agreepedia/bin/config"
	httpError "agree-agreepedia/bin/pkg/http-error"
	"agree-agreepedia/bin/pkg/utils"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpRequest struct {
	Method        string
	Url           string
	RequestBody   io.Reader
	Authorization string
	Username      string
	Password      string
	Headers       map[string]string
}

type HttpResponse struct {
	Success bool        `json:"success" default:"false"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

func HTMLToString(filename string, data interface{}) (string, error) {
	t, err := template.ParseFiles(config.GlobalEnv.RootApp + filename)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func HTTPRequest(payload *HttpRequest) utils.Result {
	var result utils.Result

	request, err := http.NewRequest(payload.Method, payload.Url, payload.RequestBody)
	if err != nil {
		errObj := httpError.NewConflict()
		errObj.Message = "URL tidak ditemukan"
		result.Error = errObj
		return result
	}

	for k, e := range payload.Headers {
		request.Header.Set(k, e)
	}

	client := &http.Client{}
	if payload.Headers["Content-Type"] == "" {
		request.Header.Set("Content-Type", "application/json")
	}

	if payload.Authorization != "" {
		request.Header.Set("Authorization", payload.Authorization)
	} else {
		request.SetBasicAuth(payload.Username, payload.Password)
	}
	response, err := client.Do(request)
	if err != nil {
		errObj := httpError.NewConflict()
		errObj.Message = "Gagal fetch URL"
		result.Error = errObj
		return result
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var HttpResponse HttpResponse
	json.Unmarshal([]byte(body), &HttpResponse)

	result.Data = HttpResponse.Data

	return result
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(config.GlobalEnv.EncodeDecodeKey.EncodeKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, []byte(config.GlobalEnv.EncodeDecodeKey.IVKey))
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(config.GlobalEnv.EncodeDecodeKey.EncodeKey))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, []byte(config.GlobalEnv.EncodeDecodeKey.IVKey))
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
