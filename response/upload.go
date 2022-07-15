package Response

// response struct
type Message struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    DataMessage `json:"data"`
}

type DataMessage struct {
	Files []FileMessage `json:"files"`
	Meta  MetaMessage   `json:"meta"`
	Type  string        `json:"type"`
}

type FileMessage struct {
	Name       string `json:"name"`
	Size       int    `json:"size"`
	Type       int    `json:"type"`
	Driver     string `json:"driver"`
	Updated_at string `json:"updated_at"`
	Thumbnail  string `json:"thumbnail"`
	Url        string `json:"url"`
	Size_str   string `json:"size_str"`
	Time_str   string `json:"time_str"`
}

type MetaMessage struct {
	Driver string `json:"driver"`
	README string `json:"readme"`
	Total  int    `json:"total"`
	Upload bool   `json:"true"`
}
