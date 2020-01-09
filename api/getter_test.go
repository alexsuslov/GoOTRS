package api

import (
	"github.com/alexsuslov/godotenv"
	"io/ioutil"
	"log"
	"go.uber.org/thriftrw/ptr"
	"testing"
)

func TestGetter(t *testing.T) {
	if err :=godotenv.Load("../.env"); err!= nil{
		panic(err)
	}
	if err := Init(); err!= nil{
		panic(err)
	}
	type args struct {
		id      string
		options []Options
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
	}{
		{
			"get tiket 744485",
			args{
				"744485",
				[]Options{{ptr.Bool(true), ptr.Bool(true), ptr.Bool(true)}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DEBUGGING = true
			gotBody, err := Getter(tt.args.id, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Getter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer gotBody.Close()
			data, _ := ioutil.ReadAll(gotBody)
			log.Println(string(data))
		})
	}
}
