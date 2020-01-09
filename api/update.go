package api

import (
	"fmt"
	"io"
	"net/url"
)

// Update update
func Update(id string, Req io.ReadCloser) (Body io.ReadCloser, err error) {
	SetterTotal.Inc()
	defer setterErrorInc(err)

	Url := fmt.Sprintf(GetURL("ticket/update", id))
	u, err := url.Parse(Url)
	return  Request("POST", u, Req)
}
