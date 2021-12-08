package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/alexsuslov/GoOTRS/api"
	"github.com/alexsuslov/GoOTRS/model"
	"github.com/alexsuslov/godotenv"
	"github.com/sirupsen/logrus"
	"io"
	"net/url"
	"os"
)

var version string
var help string
var debugger bool

var config string
var get string
var update string
var DynamicFields bool
var AllArticles bool
var Attachments bool

func init() {
	// Config
	flag.StringVar(&config, "config", ".env", "gotrs config env")
	// Update
	flag.StringVar(&update, "update", "", "update ticket in OTRS")
	//GET
	flag.StringVar(&get, "get", "", "get ticket by id")
	// GET options
	flag.BoolVar(&DynamicFields, "DynamicFields", false, "get  DynamicFields from ticket")
	flag.BoolVar(&AllArticles, "AllArticles", false, "get AllArticles from ticket")
	flag.BoolVar(&Attachments, "Attachments", false, "get Attachments from ticket")

	flag.Parse()
}

func main() {
	if err := godotenv.Load(config); err != nil {
		logrus.Warningf("no %s file", config)
	}

	otrs := api.New(
		os.Getenv("OTRS_HOST"),
		os.Getenv("OTRS_WS_NAME"),
		os.Getenv("OTRS_UNSAFE_TLS")=="YES",
	)

	vs := url.Values{}
	vs.Set("UserLogin", os.Getenv("OTRS_USER"))
	vs.Set("Password", os.Getenv("OTRS_PASSWORD"))


	// Getter
	if get != "" {
		if DynamicFields{
			vs.Set("DynamicFields", "1")
		}

		if AllArticles{
			vs.Set("AllArticles", "1")
		}

		if Attachments{
			vs.Set("Attachments", "1")
		}

		body, err := otrs.Getter(get, vs)
		Done(body, err)
		return
	}

	// Setter
	if update != "" {
		Update:=model.Update{}
		err := json.NewDecoder(os.Stdin).Decode(&Update)
		if err!= nil{
			panic(err)
		}
		body, err := otrs.Update(*Update.TicketID, Update, vs)
		Done(body, err)
		return
	}

	// help
	fmt.Printf("GoOTRS is a Golang wrapper for accessing OTRS using the REST API. Version %v \n", version)
}

func Done(body io.ReadCloser, err error) {
	if err != nil {
		panic(err)
	}
	defer body.Close()
	if _, err := io.Copy(os.Stdout, body); err != nil {
		panic(err)
	}
}
