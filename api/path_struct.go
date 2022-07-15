package Api

type PathReq struct {
	Path     string `json:"path"`
	Password string `json:"password"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
}
