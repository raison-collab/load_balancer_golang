package endpoints

type PostTask struct {
	Bash     string `json:"bash"`
	Ram      uint16 `json:"ram"`
	Disk     uint   `json:"disk"`
	CPU      uint   `json:"cpu"`
	Priority uint8  `json:"priority"`
}

type IdentifyTaskData struct {
	Token string `json:"token"`
	Hash  string `json:"hash"`
}

type ErrorJsonMessage struct {
	Message    string `json:"message"`
	Error      string `json:"error"`
	StatusCode uint16 `json:"status_code"`
}

type DataJsonMessage struct {
	Message      string           `json:"message"`
	StatusCode   uint16           `json:"status_code"`
	Data         interface{}      `json:"data"`
	IdentifyData IdentifyTaskData `json:"identify_data"`
}
