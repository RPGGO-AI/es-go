package common

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net"
	"net/http"
	"time"
)

func GetClient() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://127.0.0.1:9200",
		},
		//Username: "foo",
		//Password: "bar",
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   1000,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		},
	}

	es, _ := elasticsearch.NewClient(cfg)
	log.Println(elasticsearch.Version)
	log.Println(es.Info())
	return es
}
