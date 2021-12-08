package model

import "github.com/sirupsen/logrus"

type Article struct {
	ArticleID              ID           `json:"ArticleID"`
	ArticleNumber          int          `json:"ArticleNumber"`
	Attachments            []Attachment `json:"Attachment"`
	Bcc                    string       `json:"Bcc"`
	Body                   string       `json:"Body"`
	Cc                     string       `json:"Cc"`
	ChangeBy               ID           `json:"ChangeBy"`
	ChangeTime             string       `json:"ChangeTime"`
	Charset                string       `json:"Charset"`
	CommunicationChannelID string       `json:"CommunicationChannelID"`
	ContentCharset         string       `json:"ContentCharset"`
	ContentType            string       `json:"ContentType"`
	CreateBy               ID           `json:"CreateBy"`
	CreateTime             ID           `json:"CreateTime"`
	From                   string       `json:"From"`
	InReplyTo              string       `json:"InReplyTo"`
	IncomingTime           string       `json:"IncomingTime"`
	IsVisibleForCustomer   string       `json:"IsVisibleForCustomer"`
	MessageID              string       `json:"MessageID"`
	MimeType               string       `json:"MimeType"`
	References             string       `json:"References"`
	ReplyTo                string       `json:"ReplyTo"`
	SenderType             string       `json:"SenderType"`
	SenderTypeID           string       `json:"SenderTypeID"`
	Subject                string       `json:"Subject"`
	TicketID               ID           `json:"TicketID"`
	TimeUnit               int          `json:"TimeUnit"`
	To                     string       `json:"To"`
}

type ArticleUpdate struct {
	CommunicationChannelID string `json:"CommunicationChannelID"`
	Subject                string `json:"Subject"`
	Body                   string `json:"Body"`
	MimeType               string `json:"MimeType"`
	Charset                string `json:"Charset"`

	ArticleID            ID           `json:"ArticleID,omitempty"`
	ArticleNumber        int          `json:"ArticleNumber,omitempty"`
	Attachments          []Attachment `json:"Attachment,omitempty"`
	Bcc                  string       `json:"Bcc,omitempty"`
	Cc                   string       `json:"Cc,omitempty"`
	ChangeBy             ID           `json:"ChangeBy,omitempty"`
	ChangeTime           string       `json:"ChangeTime,omitempty"`
	ContentCharset       string       `json:"ContentCharset,omitempty"`
	ContentType          string       `json:"ContentType,omitempty"`
	CreateBy             ID           `json:"CreateBy,omitempty"`
	CreateTime           ID           `json:"CreateTime,omitempty"`
	From                 string       `json:"From,omitempty"`
	InReplyTo            string       `json:"InReplyTo,omitempty"`
	IncomingTime         string       `json:"IncomingTime,omitempty"`
	IsVisibleForCustomer string       `json:"IsVisibleForCustomer,omitempty"`
	MessageID            string       `json:"MessageID,omitempty"`
	References           string       `json:"References,omitempty"`
	ReplyTo              string       `json:"ReplyTo,omitempty"`
	SenderType           string       `json:"SenderType,omitempty"`
	SenderTypeID         string       `json:"SenderTypeID,omitempty"`
	TicketID             ID           `json:"TicketID,omitempty"`
	TimeUnit             int          `json:"TimeUnit,omitempty"`
	To                   string       `json:"To,omitempty"`

	article *Article
	logger  bool
}
/**
 *	System
 */
// log
func (au *ArticleUpdate) log(FieldName string, from, to interface{}) *ArticleUpdate {
	if au.logger {
		logrus.
			WithField("from", from).
			WithField("to", to).
			Infof("set %s", FieldName)
	}
	return au
}


/**
 *	Required fields
 */

// SetCommunicationChannelID set CommunicationChannelID
func (au *ArticleUpdate) SetCommunicationChannelID(to string) *ArticleUpdate {
	var from string
	if au.article != nil {
		if au.article.CommunicationChannelID == to {
			return au
		}
		from = au.article.CommunicationChannelID
	}
	au.article.CommunicationChannelID = to

	return au.log("CommunicationChannelID", from, to)

}


// SetSubject Set Subject
func (au *ArticleUpdate) SetSubject(to string) *ArticleUpdate {
	var from string
	if au.article != nil {
		if au.article.Subject == to {
			return au
		}
		from = au.article.Subject[:25]
	}
	au.article.Subject = to

	return au.log("subject", from, to)

}


// SetBodyHTML set Body as HTML
func (au *ArticleUpdate) SetBodyHTML(to string) *ArticleUpdate {
	var from string
	if au.article != nil {
		if au.article.Body == to {
			return au
		}
		from = au.article.Body[:25]
	}

	au.article.Body = to
	au.article.MimeType = "text/html"
	au.article.Charset = "utf-8"

	return au.log("body", from, to)
}

//SetBodyText Set Body as Text
func (au *ArticleUpdate) SetBodyText(to string) *ArticleUpdate {
	var from string
	if au.article != nil {
		if au.article.Body == to {
			return au
		}
		from = au.article.Body[:25]
	}

	au.article.Body = to
	au.article.MimeType = "text/plaine"
	au.article.Charset = "utf-8"

	return au.log("body", from, to)
}

/**
 *	Optional fields
 */

