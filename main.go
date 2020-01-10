package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/alexsuslov/GoOTRS/api"
	"github.com/alexsuslov/godotenv"
	"io"
	"io/ioutil"
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
	// debugger
	flag.BoolVar(&debugger, "debugger", false, "enable debugger")
	// Config
	flag.StringVar(&config, "config", ".env", "gotrs config env")
	// Update
	flag.StringVar(&update, "update", "", "update tiket in OTRS")
	//GET
	flag.StringVar(&get, "get", "", "get tiket from OTRS")
	// GET options
	flag.BoolVar(&DynamicFields, "DynamicFields", false, "get  DynamicFields from tiket")
	flag.BoolVar(&AllArticles, "AllArticles", false, "get AllArticles from tiket")
	flag.BoolVar(&Attachments, "Attachments", false, "get Attachments from tiket")
	
	flag.Parse()
}

func main(){
	if err := godotenv.Load(config); err!= nil{
		panic(err)
	}
	if err := api.Init(); err!= nil {
		panic(fmt.Errorf("Init:%v", err))
	}

	api.DEBUGGING=debugger

	// Getter
	if get!= ""{
		body, err := api.Getter(get, api.Options{&DynamicFields, &AllArticles, &Attachments})
		Done(body, err)
		return
	}

	// Setter
	if update!= ""{
		reader := bufio.NewReader(os.Stdin)
		body, err := api.Update(update, ioutil.NopCloser(reader))
		Done(body, err)
		return
	}
	fmt.Printf("GoOTRS is a Golang wrapper for accessing OTRS using the REST API. Version %v \n", version)

}

func Done(body io.ReadCloser, err error){
	if err != nil{
		panic(err)
	}
	defer body.Close()
	if _, err := io.Copy(os.Stdout, body); err!= nil{
		panic(err)
	}
}
