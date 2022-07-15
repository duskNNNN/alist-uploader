package Api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	File "github.com/duskNNNN/alist-uploader/file"
	Response "github.com/duskNNNN/alist-uploader/response"
)

// get all url for target path
func ApiPathGetUrl(alist_url string, params *map[string]string, file_names []string) bool {
	alist_pathGet_url := alist_url + "/api/public/path"
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	for key, val := range *params {
		writer.WriteField(key, val)
	}
	if err := writer.Close(); err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", alist_pathGet_url, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
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

	var pathGetMessage Response.Message
	json.Unmarshal(resp_body, &pathGetMessage)
	var urls []string
	if pathGetMessage.Message != "success" {
		return false
	} else {
		for _, v := range pathGetMessage.Data.Files {
			for _, vname := range file_names {
				if v.Name == vname {
					urls = append(urls, v.Thumbnail)
					break
				}
			}
		}
	}
	// write to file
	File.FileWrite(urls)
	return true
}
