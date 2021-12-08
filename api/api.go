package api

import (
	"net/url"
	"strings"
)

var GENERIC_PATH = "otrs/nph-genericinterface.pl/Webservice"

type OTRS struct {
	Host               string
	Webservice         string
	InsecureSkipVerify bool
}

func New(Host, Webservice string, InsecureSkipVerify bool) *OTRS {
	return &OTRS{
		Host,
		Webservice,
		InsecureSkipVerify,
	}
}

type Opts struct {
	DynamicFields bool
	AllArticles   bool
	Attachments   bool
}

func (OTRS OTRS) url(point *string) (string, error) {

	points := []string{OTRS.Host, GENERIC_PATH, OTRS.Webservice}

	if point != nil {
		points = append(points, *point)
	}

	path := strings.Join(points, "/")

	URL, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	return URL.String(), nil
}
