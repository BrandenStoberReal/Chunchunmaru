package utilities

type ApiBaseReply struct {
	AppVersion string  `json:"appVersion"`
	Uptime     float64 `json:"uptime"`
}

type ApiTemplateReply struct {
	Count     int      `json:"count"`
	FileNames []string `json:"fileNames"`
}

type ApiUploadTemplateData struct {
	FileName      string `json:"fileName"`
	ContentBase64 string `json:"contentBase64"`
}

type ApiDeleteTemplateData struct {
	FileName string `json:"fileName"`
}
