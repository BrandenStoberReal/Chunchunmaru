package utilities

// ApiBaseInfoReply OUTPUT: Defines data the server sends to the client regarding general server information.
type ApiBaseInfoReply struct {
	AppVersion string  `json:"appVersion"`
	Uptime     float64 `json:"uptime"`
	Os         string  `json:"os"`
	Arch       string  `json:"arch"`
}

// ApiTemplateInfoReply OUTPUT: Defines data the server sends to the client regarding template information.
type ApiTemplateInfoReply struct {
	FileNames      []string `json:"fileNames"`
	Count          int      `json:"count"`
	TotalDiskUsage int64    `json:"totalDiskUsage"`
}

type ApiQueryInfoReply struct {
	TotalQueries int `json:"totalQueries"`
}

type IpInfoStruct struct {
	Ip      string `json:"ip"`
	Queries int    `json:"queries"`
}

type UserAgentInfoStruct struct {
	UserAgent string `json:"userAgent"`
	Queries   int    `json:"queries"`
}

// ApiUploadTemplateData INPUT: Defines data the client needs to send to the server to create a new template.
type ApiUploadTemplateData struct {
	FileName      string `json:"fileName"`
	ContentBase64 string `json:"contentBase64"`
}

// ApiDeleteTemplateData INPUT: Defines data the client needs to send to the server to delete an existing template.
type ApiDeleteTemplateData struct {
	FileName string `json:"fileName"`
}
