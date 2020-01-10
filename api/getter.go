package api

import (
	"fmt"
	"io"
	"net/url"
)

//Getter Getter
func Getter(id string, options ...Options) (Body io.ReadCloser, err error) {
	GetterTotal.Inc()
	defer getterErrorInc(err)

	Url := fmt.Sprintf(GetURL("ticket/get", id, options...))
	u, err := url.Parse(Url)
	if err!= nil{
		return
	}
	return  Request("POST", u, nil)
}
