package main

// Parent struct
type ParentData struct {
	UserId     string `json:"userId"`
	CategoryId string `json:"categoryId"`
	StatusId   string `json:"statusId"`
	BoxNo      string `json:"boxNo"`
}

// Child struct
type ChildData struct {
	FileId      string `json:"fileId"`
	FileCount   string `json:"fileCount"`
	ReceiptDate string `json:"receiptDate"`
	TimeFrameId string `json:"timeFrameId"`
	WayId       string `json:"wayId"`
}

// Progress struct
type ProgressData struct {
	Parent ParentData  `json:"parent"`
	Child  []ChildData `json:"child"`
}

// Status update struct
type StatusUpdate struct {
	ReceiptionNo string `json:"receiptionNo"`
	CategoryId   string `json:"categoryId"`
	StatusId     string `json:"statusId"`
	BoxNo        string `json:"boxNo"`
}
