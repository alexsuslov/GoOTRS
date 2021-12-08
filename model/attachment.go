package model

type Attachment struct {
	Content            []byte `json:"Content"`
	ContentAlternative []byte `json:"ContentAlternative"`
	ContentID          string `json:"ContentID"`
	ContentType        string `json:"ContentType"`
	Disposition        string `json:"Disposition"`
	FileID             ID     `json:"FileID"`
	Filename           string `json:"Filename"`
	FilesizeRaw        int    `json:"FilesizeRaw"`
}
