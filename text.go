package baiduai

import (
	"time"
	"ghostlib"
	"fmt"
	"io/ioutil"
)

/*****************************************
 * FileName : text.go
 * Author   : ghostwwl
 * Note     : 百度ai的中文文本相关
 *  	中文分词 https://aip.baidubce.com/rpc/2.0/nlp/v1/wordseg 1000次/天
 * 		中文词向量表示 https://aip.baidubce.com/rpc/2.0/nlp/v1/wordembedding 1000次/天
 * 		短文本相似度 https://aip.baidubce.com/rpc/2.0/nlp/v1/simnet 1000次/天
 *		中文DNN语言模型 https://aip.baidubce.com/rpc/2.0/nlp/v1/dnnlm_cn 1000次/天
 *		评论观点抽取 https://aip.baidubce.com/rpc/2.0/nlp/v1/comment_tag 1000次/天
 *		词性标注 https://aip.baidubce.com/rpc/2.0/nlp/v1/wordpos 1000次/天
 *****************************************/

import (
	"net/http"
	"bytes"
	"ghostlib"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"errors"
	"time"
)


type AiText struct {
	AiClient
}

func NewText() *AiText {
	c := new(AiText)
	c.client = &http.Client{Timeout: API_TIMEOUT * time.Second}
	c.error_code = map[float64]string{
		100 : "无效参数",
		110 : "Token过期失效",
	}
	c.type_code = map[float64]string {
		1: "酒店",
		2: "KTV",
		3: "丽人",
		4: "美食（默认）",
		5: "旅游",
		6: "健康",
		7: "教育",
		8: "商业",
		9: "房产",
		10: "汽车",
		11:	"生活",
		12:	"购物",
	}


	return c
}

/**
 * 评论观点抽取 或 情感识别  垃圾 很容易没结果
 */
func (this *AiText) GetCommentTag(intxt string) ([]interface{}) {
	// 真实使用时这里要判断过期时间 避免重复获取token
	if this.access_token == "" {
		doflag, _ := this.getToken()
		if !doflag {
			panic("获取access token 失败")
		}
	}

	post_arg := map[string]interface{}{
		"comment":  ghostlib.UrlEncode(ghostlib.ConvertStrEncode(intxt, "utf-8", "gbk")),
		"type":  "4",	// 类别
		"entity":  "NULL", //实体名
	}
	post_json, err := json.Marshal(post_arg)
	if nil != err{
		panic(err)
	}

	real_uri := fmt.Sprintf("%s?access_token=%s", COMTAG_API_URI, this.access_token)

	req, err := http.NewRequest("POST", real_uri, bytes.NewReader(post_json))
	//req.Header.Add("User-Agent", "baidu-aip-php-sdk-1.0.0.1")
	resp, _ := this.client.Do(req)

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		panic(err)
	}

	udata := ghostlib.ConvertStrEncode(string(data), "gbk", "utf-8")

	if "null" == udata {
		panic(errors.New("貌似返回空"))
	}

	map_result := make(map[string]interface{})
	json.Unmarshal([]byte(udata), &map_result)

	return map_result["tags"].([]interface{})
}

/**
 * 分词
 */
func (this *AiText) SplitWords(intxt string) (map[string]interface{}) {
	// 真实使用时这里要判断过期时间 避免重复获取token
	if this.access_token == "" {
		doflag, _ := this.getToken()
		if !doflag {
			panic("获取access token 失败")
		}
	}

	post_arg := map[string]interface{}{
		"query":  ghostlib.UrlEncode(ghostlib.ConvertStrEncode(intxt, "utf-8", "gbk")),
	}
	post_json, err := json.Marshal(post_arg)
	if nil != err{
		panic(err)
	}

	real_uri := fmt.Sprintf("%s?access_token=%s", SEG_API_URI, this.access_token)

	req, err := http.NewRequest("POST", real_uri, bytes.NewReader(post_json))
	//req.Header.Add("User-Agent", "baidu-aip-php-sdk-1.0.0.1")
	resp, _ := this.client.Do(req)

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		panic(err)
	}

	udata := ghostlib.ConvertStrEncode(string(data), "gbk", "utf-8")

	if "null" == udata {
		panic(errors.New("貌似返回空"))
	}

	//fmt.Printf("%s", udata)
	map_result := make(map[string]interface{})
	json.Unmarshal([]byte(udata), &map_result)

	return map_result["scw_out"].(map[string]interface{})

	// 结果里重要的内容说明
	//namebuf := x["namebuf"]		// 人名 [有就会识别出]
	//subphrbuf := x["subphrbuf"]	// 短语
	//wordsepbuf := x["wordsepbuf"]	// 标准粒度
	//wpcompbuf := x["wpcompbuf"]	// 混排粒度
	//newwordbuf := x["pnewword"].(map[string]interface{})["newwordbuf"] // 新词

}

/**
 * 词性标注
 */
func (this *AiText) WordPos(intxt string) ([]interface{}) {
	// 真实使用时这里要判断过期时间 避免重复获取token
	if this.access_token == "" {
		doflag, _ := this.getToken()
		if !doflag {
			panic("获取access token 失败")
		}
	}

	post_arg := map[string]interface{}{
		"query":  ghostlib.UrlEncode(ghostlib.ConvertStrEncode(intxt, "utf-8", "gbk")),
	}
	post_json, err := json.Marshal(post_arg)
	if nil != err{
		panic(err)
	}

	real_uri := fmt.Sprintf("%s?access_token=%s", WORDPOS_API_URI, this.access_token)

	req, err := http.NewRequest("POST", real_uri, bytes.NewReader(post_json))
	//req.Header.Add("User-Agent", "baidu-aip-php-sdk-1.0.0.1")
	resp, _ := this.client.Do(req)

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		panic(err)
	}

	udata := ghostlib.ConvertStrEncode(string(data), "gbk", "utf-8")

	if "null" == udata {
		panic(errors.New("貌似返回空"))
	}

	//fmt.Printf("%s", udata)
	map_result := make(map[string]interface{})
	json.Unmarshal([]byte(udata), &map_result)

	return map_result["result_out"].([]interface{})
}