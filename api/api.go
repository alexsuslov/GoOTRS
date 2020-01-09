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


// GetURL get url
func GetURL(method string, id string, options ...bool) string {

	opts := ""
	for  i,_:=range(options){
		switch(i){
		case 0:
			if options[i]{
				opts = "&DynamicFields=1"
			}
		case 1:
			if options[i]{
				opts = opts+"&AllArticles=1"
			}
		case 2:
			if options[i] {
				opts = opts + "&Attachments=1"
			}
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
