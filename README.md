# github.com/duskNNNN/alist-uploader

> of course, `alist` server must be rebulid

## rebuild server
1. git clone `alist` v2 repo

   ```shell
   git clone -b v2 https://github.com/alist-org/alist.git && cd alist
   ```
   
2. modify `server/middlewares/path.go`

  ```shell
  nano server/middlewares/path.go
  # add package
  import ("strconv")
  
  # add to 17 line
  if req.Path == "" {
  	req.Path = c.PostForm("path")
      req.Password = c.PostForm("password")
      req.PageNum, _ = strconv.Atoi(c.PostForm("page_num"))
      req.PageSize, _ = strconv.Atoi(c.PostForm("page_size"))
  }
  ```

  ![image-20220715102552296](https://mqin.duskhouse.cn:9000/?/images/2022/07/15/lFk5leIV4q/image-20220715102552296.png)

3. compile front file

   ```shell
   cd ..
   git clone https://github.com/alist-org/alist-web.git
   cd alist-web
   yarn
   yarn build
   sed -i -e "s/\/CDN_URL\//\//g" dist/index.html
   sed -i -e "s/assets/\/assets/g" dist/index.html
   rm -f dist/index.html-e
   cp -r dist/* ../alist/public/
   cd ../alist
   ```

4. rebuild

   ```shell
   go build -o alist -tags=jsoniter alist.go
   ```

## idea
1. choose files which will be uploader

2. access to API `/api/public/uploader`

3. access to API `/api/public/path`

4. through files prefix or suffix to distinguish which files were uploaded you just now

5. urls will be add to file_url.txt in now path

## Limitation

+ when upload folders, the folders will not automatically create, only upload files
+ it's recommended that the upload folders' file shouldn't more than 1000
  + if your files are too much, please upload for a lots of times  
  + you can create a new visitor path for upload every time