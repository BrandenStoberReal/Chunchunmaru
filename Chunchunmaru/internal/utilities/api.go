package utilities

type ApiBaseInfoReply struct {
	AppVersion string  `json:"appVersion"`
	Uptime     float64 `json:"uptime"`
	Os         string  `json:"os"`
	Arch       string  `json:"arch"`
}

type ApiTemplateInfoReply struct {
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
