package Api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	File "github.com/duskNNNN/alist-uploader/file"
	Response "github.com/duskNNNN/alist-uploader/response"
	"github.com/duskNNNN/alist-uploader/utils"
)

// get all url for target path
func ApiPathGetUrl(alist_url string, params *map[string]string, file_names []string) bool {
	alist_pathGet_url := alist_url + "/api/public/path"
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	var alist_path string
	for key, val := range *params {
		if key == "path" {
			alist_path = val
		}
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
	HttpClient := &http.Client{
		Timeout: 3600 * time.Second,
	}
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
	var temp_thumbnail string
	var temp_url string
	if pathGetMessage.Message != "success" {
		return false
	} else {
		for _, v := range pathGetMessage.Data.Files {
			for _, vname := range file_names {
				if v.Name == vname {
					// if v.Url != "" {
					// 	urls = append(urls, v.Url)
					// } else {
					// 	urls = append(urls, v.Thumbnail)
					// }
					// urlencode
					urls = append(urls, v.Name)
					temp_url = "url:" + alist_url + "/d" + alist_path + "/" + utils.PathEscape(v.Name)
					urls = append(urls, temp_url)
					temp_thumbnail = "thumbnail:" + alist_url + alist_path + "/" + v.Thumbnail
					urls = append(urls, temp_thumbnail)
					break
				}
			}
		}
	}
	// write to file
	File.FileWrite(urls)
	return true
}
