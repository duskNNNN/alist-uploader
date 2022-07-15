package Api

import (
	Response "github.com/duskNNNN/alist-uploader/response"
	// "github.com/duskNNNN/alist-uploader/token"
	"bytes"
	"encoding/json"

	// "fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

var (
	HttpClient *http.Client
)

func InitHttpClient() {
	HttpClient = &http.Client{
		Timeout: 60 * time.Second,
	}
}

// upload file
func ApiUploadFile(alist_url string, params map[string]string, file io.Reader, filename string) bool {
	alist_upload_url := alist_url + "/api/public/upload"
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	// create a file field in form
	formFile, err := writer.CreateFormFile("files", filename)
	if err != nil {
		log.Fatal(err)
	}
	// read file content to file field
	_, err = io.Copy(formFile, file)
	if err != nil {
		log.Fatal(err)
	}
	// write other params to form
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	if err = writer.Close(); err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", alist_upload_url, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	resp, err := HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	resp_body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr.Error())
		return false
	}
	var uploadMessage Response.Message
	json.Unmarshal(resp_body, &uploadMessage)
	if uploadMessage.Message != "success" {
		return false
	}
	return true
}
