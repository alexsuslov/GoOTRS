package api

import "github.com/prometheus/client_golang/prometheus"

// GetterTotal GetterTotal
var GetterTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "api_getter_total",
		Help: "The total number of getter api requests",
	})

// GetterError Getter Error
var GetterError = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "api_getter_errors",
		Help: "The total number of getter api errors",
	})

func getterErrorInc(e error){
	if e!= nil{
		GetterError.Inc()
	}
}


// SetterTotal SetterTotal
var SetterTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "api_setter_total",
		Help: "The total number of setter api requests",
	})

// SetterError SetterError
var SetterError = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "api_setter_errors",
		Help: "The total number of setter api errors",
	})


func setterErrorInc(e error){
	if e!= nil{
		SetterError.Inc()
	}
}
