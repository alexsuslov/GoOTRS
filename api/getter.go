package api

import (
	"encoding/json"
	"fmt"
	"github.com/alexsuslov/GoOTRS/model"
	"github.com/ddliu/go-httpclient"
	"io"
	"io/ioutil"
	"net/url"
)

var GETTER = "ticket/get"

/*
// https://otrscommunityedition.com/doc/api/otrs/6.0/Perl/Kernel/GenericInterface/Operation/Ticket/TicketGet.pm.html
o url.Values

	UserLogin            => 'some agent login',                            # UserLogin or CustomerUserLogin or SessionID is
																					   #   required
	CustomerUserLogin    => 'some customer login',
	SessionID            => 123,

	Password             => 'some password',                               # if UserLogin or customerUserLogin is sent then
																		   #   Password is required
	TicketID             => '32,33',                                       # required, could be coma separated IDs or an Array
	DynamicFields        => 0,                                             # Optional, 0 as default. Indicate if Dynamic Fields
																		   #     should be included or not on the ticket content.
	Extended             => 1,                                             # Optional, 0 as default
	AllArticles          => 1,                                             # Optional, 0 as default. Set as 1 will include articles
																		   #     for tickets.
	ArticleSenderType    => [ $ArticleSenderType1, $ArticleSenderType2 ],  # Optional, only requested article sender types
	ArticleOrder         => 'DESC',                                        # Optional, DESC,ASC - default is ASC
	ArticleLimit         => 5,                                             # Optional
	Attachments          => 1,                                             # Optional, 0 as default. If it's set with the value 1,
																		   # attachments for articles will be included on ticket data
	GetAttachmentContents = 1                                              # Optional, 1 as default. 0|1,
	HTMLBodyAsAttachment => 1                                              # Optional, If enabled the HTML body version of each article
                                                                                   #    is added to the attachments list
*/

//Getter Getter
func (OTRS OTRS) Getter(id string, o url.Values) (Body io.ReadCloser, err error) {

	point := GETTER + "/" + id
	URL, err := OTRS.url(&point)
	if err != nil {
		return nil, err
	}

	res, err := httpclient.
		WithHeader("User-Agent", "ServiceChain").
		Get(URL, o)

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

type GetterResponse struct {
	Error   model.Error    `json:"Error"`
	Tickets []model.Ticket `json:"Ticket"`
}

func (OTRS OTRS) GetTickets(id string, o url.Values, resp *GetterResponse) error {
	body, err := OTRS.Getter(id, o)
	if err != nil {
		return err
	}
	defer body.Close()

	return json.NewDecoder(body).Decode(resp)
}
