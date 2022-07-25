package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	Api "github.com/duskNNNN/alist-uploader/api"
	Judge "github.com/duskNNNN/alist-uploader/judge"
)

func main() {
	list := os.Args
	// judge params
	if len(list) != 5 {
		log.Fatal("command [alist_url] [alist_path] [alist_password] [file_path]")
	}
	// get params
	// set dst alist site url
	alist_url := list[1]
	alist_path := list[2]
	alist_password := list[3]
	src_path, _ := filepath.Abs(list[4])
	// judge this path whether exist
	if !Judge.PathJudgeIsExists(src_path) {
		log.Fatal("file not exist")
	}
	// judge this path is folder or file
	isDir := Judge.PathJudgeIsFolder(src_path)
	// all files name
	var filesName []string
	// upload params
	params := map[string]string{
		"path":     alist_path,
		"password": alist_password,
	}
	if isDir {
		filepath.Walk(src_path, func(path string, info os.FileInfo, err error) error {
			// because the api limitation,only upload file, can't create folder auto
			if !Judge.PathJudgeIsFolder(path) {
				now_filename := filepath.Base(path)
				filesName = append(filesName, now_filename)
				now_file, err1 := os.Open(path)
				if err != nil {
					log.Println(err1)
				}
				if !Api.ApiUploadFile(alist_url, params, now_file, now_filename) {
					log.Println("upload file error")
				}
				time.Sleep(3 * time.Second)
			}
			return nil
		})
	} else {
		filename := filepath.Base(src_path)
		filesName = append(filesName, filename)
		fileContent, err := os.Open(src_path)
		if err != nil {
			log.Println(err)
		}
		if !Api.ApiUploadFile(alist_url, params, fileContent, filename) {
			log.Println("upload file error")
		}
	}
	// path params
	path_params := make(map[string]string)
	path_params["path"] = alist_path
	path_params["password"] = alist_password
	path_params["page_num"] = "1"
	path_params["page_size"] = "1000"
	if !Api.ApiPathGetUrl(alist_url, &path_params, filesName) {
		log.Fatal("get path error")
	}
}
