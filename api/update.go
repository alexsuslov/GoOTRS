package api

import (
	"fmt"
	"github.com/alexsuslov/GoOTRS/model"
	"github.com/ddliu/go-httpclient"
	"io"
	"io/ioutil"
	"net/url"
)

var UPDATE = "ticket/update"

func (OTRS OTRS) Update(id string, req model.Update, o url.Values) (Body io.ReadCloser, err error) {

	point := UPDATE + "/" + id
	URL, err := OTRS.url(&point)
	if err != nil {
		return nil, err
	}

	URL = URL + "?" + o.Encode()

	res, err := httpclient.
		WithHeader("User-Agent", "ServiceChain").
		PostJson(URL, req)

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		data, err := ioutil.ReadAll(res.Body)
		if err == nil {
			defer res.Body.Close()
		}
		err = fmt.Errorf("%v:%v", res.Status, string(data))
		return nil, err
	}
	return res.Body, nil
}

/**
{"Error":{
	"ErrorCode":"TicketUpdate.MissingParameter",
	"ErrorMessage":"TicketUpdate: UserLogin, CustomerUserLogin or SessionID is required!"
	}
}

{"TicketNumber":"102004616","TicketID":"2006804"}

*/
