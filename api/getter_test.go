package api

import (
	"github.com/alexsuslov/godotenv"
	"io"
	"log"
	"net/url"
	"os"
	"testing"
)

func TestOTRS_Getter(t *testing.T) {
	godotenv.Load(("../.env"))
	type fields struct {
		Host               string
		Webservice         string
		InsecureSkipVerify bool
	}
	type args struct {
		id string
		o  url.Values
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantBody []byte
		wantErr  bool
	}{
		{
			"get ticket from otrs",
			fields{
				os.Getenv("OTRS_HOST"),
				os.Getenv("OTRS_WS_NAME"),
				false,
			},
			args{
				"2007831",
				url.Values{
					"UserLogin":             []string{os.Getenv("OTRS_USERNAME")},
					"Password":              []string{os.Getenv("OTRS_PASSWORD")},
					"DynamicFields":         []string{"0"},
					"AllArticles":           []string{"0"},
					"Attachments":           []string{"0"},
					"GetAttachmentContents": []string{"0"},
				},
			},
			nil,
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
			gotBody, err := OTRS.Getter(tt.args.id, tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("Getter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			data, err := io.ReadAll(gotBody)
			if err != nil {
				panic(err)
			}
			gotBody.Close()

			log.Println(string(data))

			//if !reflect.DeepEqual(data, tt.wantBody) {
			//	t.Errorf("Getter() gotBody = %v, want %v", gotBody, tt.wantBody)
			//}
		})
	}
}
