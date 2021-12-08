package model

import (
	"bytes"
	"encoding/json"
	"io"
)

type Update struct {
	UserLogin     *string        `json:"UserLogin,omitempty"`
	Password      *string        `json:"Password,omitempty"`
	TicketID      *string        `json:"TicketID,omitempty"`
	TicketNumber  *string        `json:"TicketNumber,omitempty"`
	Ticket        *Ticket        `json:"Ticket,omitempty"`
	Article       *ArticleUpdate `json:"Article,omitempty"`
	DynamicFields []DynamicField `json:"DynamicField,omitempty"`
	Attachments   []Attachment   `json:"Attachment,omitempty"`
}

func (Update Update) ReaderCloser() io.ReadCloser {
	data, _ := json.Marshal(Update)
	return io.NopCloser(bytes.NewReader(data))
}
