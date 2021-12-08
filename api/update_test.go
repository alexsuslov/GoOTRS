package api

import (
	"github.com/alexsuslov/GoOTRS/model"
	"github.com/alexsuslov/godotenv"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

func rc(s string) io.ReadCloser {
	return io.NopCloser(strings.NewReader(s))
}

func TestOTRS_Update(t *testing.T) {
	godotenv.Load(("../.env"))

	type fields struct {
		Host               string
		Webservice         string
		InsecureSkipVerify bool
	}

	f := fields{
		os.Getenv("OTRS_HOST"),
		os.Getenv("OTRS_WS_NAME"),
		false,
	}

	type args struct {
		id  string
		upd model.Update
		o   url.Values
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		//wantBody io.ReadCloser
		wantErr bool
	}{

		{
			"update 2007831 addArticle",
			f,
			args{

				"2007831",
				model.Update{
					Article: &model.ArticleUpdate{
						CommunicationChannelID: "1",
						Subject:                "test" + time.Now().String(),
						Body:                   "Body1",
						MimeType:               "text/plain",
						Charset:                "utf-8",
					},
				},
				url.Values{
					"UserLogin": []string{os.Getenv("OTRS_USERNAME")},
					"Password":  []string{os.Getenv("OTRS_PASSWORD")},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OTRS := OTRS{
				Host:               tt.fields.Host,
				Webservice:         tt.fields.Webservice,
				InsecureSkipVerify: tt.fields.InsecureSkipVerify,
			}
			gotBody, err := OTRS.Update(tt.args.id, tt.args.upd, tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			data, err := io.ReadAll(gotBody)
			if err != nil {
				panic(err)
			}
			defer gotBody.Close()
			log.Println(string(data))
			//if !reflect.DeepEqual(gotBody, tt.wantBody) {
			//	t.Errorf("Update() gotBody = %v, want %v", gotBody, tt.wantBody)
			//}
		})
	}
}

var updateStateID = `{
	"Ticket": { "StateID": "0" }
}`

var addArticle = `{
	"Article": {
		"Subject": "test Subject",
		"Body": "test Body",
		"MimeType": "text/plain",
		"Charset": "utf-8"
	}
}`

var addDynamicField = `{
	"DynamicField":[
		{
			"Name": "IntegrationHistory",
			"Value" "test",
		}
	]

}`
