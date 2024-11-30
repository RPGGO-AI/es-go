package main

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net"
	"net/http"
	"time"
)

// link: https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/connecting.html
func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		//Username: "foo",
		//Password: "bar",
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	es, _ := elasticsearch.NewClient(cfg)
	// 或者，默认本地连接使用下方无参数链接
	//es, _, := elasticsearch.NewDefaultClient()
	log.Println(elasticsearch.Version)
	log.Println(es.Info())
}
