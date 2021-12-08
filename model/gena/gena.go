package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"text/template"
)

func main() {
	f, err := os.OpenFile("generated.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	// head
	t, err := template.
		New("t").
		Parse(head)
	if err != nil {
		panic(err)
	}
	err = t.Execute(f, map[string]string{
		"Name":"model",
	})

	if err != nil {
		panic(err)
	}

	// create ticket update
	config, err := getConfig(TicketUpdate)
	if err != nil {
		panic(err)
	}
	// structure
	fields := map[string]string{}
	s := "type TicketUpdate struct{\n"
	for k, v := range config {
		vs := strings.Split(v, ",")
		fields[k]=vs[0]
		note := fmt.Sprintf("`json:\"%s\"`", k)
		if strings.Contains(v, "optional") {
			note = fmt.Sprintf("`json:\"%s,omitempty\"`", k)
		}
		s += fmt.Sprintf("\t%s %s %s\n ", k, vs[0], note)
	}
	s += `
	logger bool
	prev *Ticket
}


`

	//f.WriteString(s)

	//t, err = template.New("t").Parse(logTmpl)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = t.Execute(f, map[string]string{
	//	"Name": "TicketUpdate",
	//})
	//if err != nil {
	//	panic(err)
	//}


	t, err = template.
		New("t").
		Parse(setterTmpl)
	if err != nil {
		panic(err)
	}

	for k,v :=range fields {
		t.Execute(f, map[string]string{
			"Obj": "TicketUpdate",
			"Name": k,
			"Type": v,
		})
	}


	f.Close()
	//logrus.WithField("TicketUpdate", s).Info("debug")
}

var TicketUpdate = `
    Title:          string
    QueueID:        ID,             optional
    Queue:          string,         optional
    LockID:         ID,             optional
    Lock:           string,         optional
    TypeID:         ID,             optional
    Type:           string,         optional
    ServiceID:      ID,             optional
    Service:        string,         optional
    SLAID:          ID,             optional
    SLA:            string,         optional
    StateID:        ID,             optional
    State:          string,         optional
    PriorityID:     ID,             optional
    Priority:       string,         optional
    OwnerID:        ID,             optional
    Owner:          string,         optional
    ResponsibleID:  ID,             optional
    Responsible:    string,         optional
    CustomerUser:   string,         optional
    # PendingTime:    PendingTime,    optional
`

//type Fields map[string]string

type Config map[string]string

func getConfig(s string) (Config, error) {
	config := Config{}
	return config, yaml.Unmarshal([]byte(s), &config)
}

var head = `
// Code generated. DO NOT EDIT!

package {{.Name}}


`

var logTmpl = `
/**
 *	System
 */
// log
func ({{.Name}} *{{.Name}}) log(FieldName string, from, to interface{}) *{{.Name}} {
	if {{.Name}}.logger {
		logrus.
			WithField("from", from).
			WithField("to", to).
			Infof("set %s", FieldName)
	}
	return {{.Name}}
}
`

var setterTmpl = `
// Set{{.Name}} set {{.Name}}
func ({{.Obj}} *{{.Obj}}) Set{{.Name}}(to {{.Type}}) *{{.Obj}} {
	var from {{.Type}}
	if {{.Obj}}.prev != nil {
		if {{.Obj}}.prev.{{.Name}} == to {
			return {{.Obj}}
		}
		from = {{.Obj}}.prev.{{.Name}}
	}
	{{.Obj}}.prev.{{.Name}} = to

	return {{.Obj}}.log("{{.Name}}", from, to)

}


`