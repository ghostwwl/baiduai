package baiduai
//
//
///*****************************************
// * FileName : text.go
// * Author   : ghostwwl
// * Note     : 百度ai的中文文本相关
// *  	中文分词 https://aip.baidubce.com/rpc/2.0/nlp/v1/wordseg 1000次/天
// * 		中文词向量表示 https://aip.baidubce.com/rpc/2.0/nlp/v1/wordembedding 1000次/天
// * 		短文本相似度 https://aip.baidubce.com/rpc/2.0/nlp/v1/simnet 1000次/天
// *		中文DNN语言模型 https://aip.baidubce.com/rpc/2.0/nlp/v1/dnnlm_cn 1000次/天
// *		评论观点抽取 https://aip.baidubce.com/rpc/2.0/nlp/v1/comment_tag 1000次/天
// *		词性标注 https://aip.baidubce.com/rpc/2.0/nlp/v1/wordpos 1000次/天
// *****************************************/
//
//import (
//	"net/http"
//	"time"
//	"bytes"
//	"ghostlib"
//	"encoding/json"
//	"io/ioutil"
//	"github.com/bitly/go-simplejson"
//	"fmt"
//	"errors"
//)
//
//const (
//	SEG_API_URI = "https://aip.baidubce.com/rpc/2.0/nlp/v1/wordseg"
//	WORDPOS_API_URI = "https://aip.baidubce.com/rpc/2.0/nlp/v1/wordpos"
//	WORDEMBED_API_URI = "https://aip.baidubce.com/rpc/2.0/nlp/v1/wordembedding"
//	DNNL_API_URI = "https://aip.baidubce.com/rpc/2.0/nlp/v1/dnnlm_cn"
//	SIMNET_API_URI = "https://aip.baidubce.com/rpc/2.0/nlp/v1/simnet"
//	COMTAG_API_URI = "https://aip.baidubce.com/rpc/2.0/nlp/v1/comment_tag"
//)
//
//
//type AiOCR struct {
//	client       *http.Client
//	access_token string
//	error_code map[float64]string
//	type_code map[float64]string
//}
//
//func NewOCR() *AiOCR {
//	c := new(AiOCR)
//	c.client = &http.Client{Timeout: API_TIMEOUT * time.Second}
//	c.error_code = map[float64]string{
//		100 : "无效参数",
//		110 : "Token过期失效",
//	}
//	c.type_code = map[float64]string {
//		1: "酒店",
//		2: "KTV",
//		3: "丽人",
//		4: "美食（默认）",
//		5: "旅游",
//		6: "健康",
//		7: "教育",
//		8: "商业",
//		9: "房产",
//		10: "汽车",
//		11:	"生活",
//		12:	"购物",
//	}
//
//
//	return c
//}
//
//func (this *AiOCR) getToken() (bool, string) {
//	post_arg := map[string]interface{}{
//		"client_id":     BAIDU_AI_KEY,
//		"client_secret": BAIDU_AI_CRET,
//		"grant_type":    "client_credentials",
//	}
//
//	resp, _ := this.client.PostForm(TOKEN_API_URI, ghostlib.InitPostData(post_arg))
//	defer resp.Body.Close()
//	data, _ := ioutil.ReadAll(resp.Body)
//
//	json_result, err := simplejson.NewJson(data)
//	if err != nil {
//		panic(err)
//	}
//
//	//fmt.Printf("%v", string(data))
//
//	map_result := make(map[string]interface{})
//	map_result, _ = json_result.Map()
//	access_token, ok := map_result["access_token"]
//
//	if ok {
//		this.access_token = ghostlib.ToString(access_token)
//		return true, this.access_token
//	}
//	return false, ""
//}
