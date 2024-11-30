# es-go
elasticsearch for go

下方为本地部署ES

```bash
docker run -p 9200:9200 -e "discovery.type=single-node" -e "xpack.security.enabled=false" -e "xpack.security.http.ssl.enabled=false" --name es816 --network es_network_for_816 docker.elastic.co/elasticsearch/elasticsearch:8.16.1
```

```bash
docker network create es_network_for_816

docker network inspect es_network_for_816
```


从上面inspect到的Docker网络IP，替换下面的172.19.0.1，下方的172.19.0.1仅是我的Docker内网IP
组网能够避免网络切换带来的麻烦，不使用Kibana可以不这样，禁用HTTPS也同理，本地避免麻烦
```bash
docker run -p 5601:5601 -e "ELASTICSEARCH_HOSTS=http://172.19.0.1:9200" -e "I18N_LOCALE=zh-CN" -e "TZ=Asia/Shanghai" --name kb816 --network es_network_for_816 docker.elastic.co/kibana/kibana:8.16.1
```

文档链接
https://veimq752wpr.larksuite.com/docx/GKMide91So8UBhx6qW2uXbQXs5g

