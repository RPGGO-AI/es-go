package main

import (
	"es-go/common"
	"fmt"
	"io"
)

// 新版es sdk与旧版有出入，es官网尚未完全匹配
// 本地部署的es版本为8.16.1
func main() {
	es := common.GetClient()
	//create, err := es.Indices.Create("my_index2222")
	//if err != nil {
	//	return
	//}
	//fmt.Println(create.StatusCode)
	//res, err := es.Indices.Create("test-index").
	//	Request(&create.Request{
	//		Mappings: &types.TypeMapping{
	//			Properties: map[string]types.Property{
	//				"price": types.NewIntegerNumberProperty(),
	//			},
	//		},
	//	}).Do(nil)

	wanna_check_indexs := []string{"maple-lang-index"}
	exists, err := es.Indices.Exists(wanna_check_indexs)
	if err != nil {
		panic(err)
	}
	fmt.Println(exists.String())
	res, err := es.Indices.Get(wanna_check_indexs)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	//[200 OK] {"maple-lang-index":{"aliases":{},"mappings":{"properties":{"text":{"type":"semantic_text","inference_id":"maple-inference-endpoint","model_settings":{"task_type":"text_embedding","dimensions":1536,"similarity":"dot_product","element_type":"float"}}}},"settings":{"index":{"routing":{"allocation":{"include":{"_tier_preference":"data_content"}}},"number_of_shards":"1","provided_name":"maple-lang-index","creation_date":"1732875912394","number_of_replicas":"1","uuid":"phVnhNihRQGljamdr4rCaA","version":{"created":"8518000"}}}}}
	fmt.Println(res)

}
