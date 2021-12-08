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

func (Update *Update) AddUpdateDynamicField(name string, value interface{}) {
	if Update.DynamicFields ==nil{
		Update.DynamicFields = []DynamicField{}
	}
	updated := false
	for k, dField := range Update.DynamicFields{
		if dField.Name==name{
			Update.DynamicFields[k] = DynamicField{
				name,
				value,
			}
			updated = true
			break
		}
	}
	if !updated{
		Update.DynamicFields = append(Update.DynamicFields)
	}

	Update.DynamicFields = append(Update.DynamicFields, DynamicField{name, value})
}
