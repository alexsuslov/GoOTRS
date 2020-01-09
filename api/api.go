package api

import (
	"fmt"
	"github.com/alexsuslov/godotenv"
	"log"
	"strings"
)

var DEBUGGING bool
var unsafeTLS bool
var host string
var apiWS string
var user string
var pass string

func Print(opts ...interface{}) {
	if DEBUGGING {
		log.Println(opts...)
	}
}

func Init() error {
	unsafeTLS = "YES" == godotenv.GetPanic("OTRS_API_UNSAFE_TLS")
	host = godotenv.GetPanic("OTRS_API_HOST")
	apiWS = godotenv.GetPanic("OTRS_API_WS_NAME")
	user = godotenv.GetPanic("OTRS_API_USER")
	pass = godotenv.GetPanic("OTRS_API_PASSWORD")
	//todo: check login
	return nil
}


type Options struct{
	DynamicFields *bool
	AllArticles *bool
	Attachments *bool
}


// GetURL get url
func GetURL(method string, id string, options ...Options) string {

	opts := ""
	for  _,option := range(options){
		if option.DynamicFields!= nil {
			opts = opts+"&DynamicFields=1"
		}
		if option.AllArticles!= nil {
			opts = opts+"&AllArticles=1"
		}
		if option.Attachments!= nil {
			opts = opts+"&Attachments=1"
		}
	}

	GetUserPass := fmt.Sprintf("UserLogin=%s&Password=%s", user, pass)

	return strings.Join([]string{
		host,
		"/otrs/nph-genericinterface.pl/Webservice/",
		apiWS,
		"/",
		method,
		"/",
		id,
		"?",
		GetUserPass,
		opts,
	}, "")
}
