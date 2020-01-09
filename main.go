package main

import (
	"flag"
	"fmt"
	"github.com/alexsuslov/GoOTRS/api"
	"github.com/alexsuslov/godotenv"
	"io/ioutil"
)

var help string
var debugger bool

var get string
var DynamicFields bool
var AllArticles bool
var Attachments bool

func init() {
	flag.BoolVar(&debugger, "debugger", false, "enable debugger")

	flag.StringVar(&get, "get", "", "get tiket from OTRS")
	flag.BoolVar(&DynamicFields, "DynamicFields", false, "get  DynamicFields from tiket")
	flag.BoolVar(&AllArticles, "AllArticles", false, "get AllArticles from tiket")
	flag.BoolVar(&Attachments, "Attachments", false, "get Attachments from tiket")
	flag.Parse()
}

func main(){

	if err := godotenv.Load(); err!= nil{
		panic(err)
	}
	if err := api.Init(); err!= nil {
		panic(fmt.Errorf("Init:%v", err))
	}

	api.DEBUGGING=debugger

	if get!= ""{
		body, err := api.Getter(get, DynamicFields, AllArticles, Attachments)
		if err != nil{
			panic(err)
		}
		defer body.Close()
		data, err:=ioutil.ReadAll(body)
		fmt.Print(string(data))
	}

}
