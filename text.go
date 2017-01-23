package baiduai

import (
	"time"
	"ghostlib"
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
	"encoding/json"
	"errors"
)


type AiText struct {
	AiClient
	WordKindMap map[string]string
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
		11: "生活",
		12: "购物",
	}
	// 很明显百度的词性标注基本和 北大计算所词性标注集简表 一样一样的
	c.WordKindMap = map[string]string{
		"Dg" : "副语素",
		"Ng" : "名语素",
		"Tg" : "时语素",
		"Vg" : "动语素",
		"a" : "形容词",
		"ad" : "副形词",
		"an" : "名形词",
		"b" : "区别词",
		"c" : "连词",
		"d" : "副词",
		"e" : "叹词",
		"f" : "方位词",
		"g" : "语素",
		"h" : "前接成分",
		"i" : "成语",
		"j" : "简称略语",
		"k" : "后接成分",
		"l" : "习用于",
		"m" : "数词",
		"n" : "名词",
		"nr" : "人名",
		"ns" : "地名",
		"nt" : "机构团体",
		"nz" : "其他专名",
		"o" : "拟声词",
		"p" : "介词",
		"q" : "量词",
		"r" : "代词",
		"s" : "处所词",
		"t" : "时间词",
		"u" : "助词",
		"v" : "动词",
		"vd" : "副动词",
		"vn" : "名动词",
		"w" : "标点符号",
		"x" : "非语素字",
		"y" : "语气词",
		"z" : "状态词",
	}

	return c
}

/**
 * 评论观点抽取 或 情感识别  垃圾 很容易没结果
 */
func (this *AiText) GetCommentTag(intxt string) ([]interface{}) {
	post_arg := map[string]interface{}{
		"comment":  ghostlib.UrlEncode(ghostlib.ConvertStrEncode(intxt, "utf-8", "gbk")),
		"type":  "4",	// 类别
		"entity":  "NULL", //实体名
	}
	post_json, err := json.Marshal(post_arg)
	if nil != err{
		panic(err)
	}

	real_uri := this.getInterFaceUri(COMTAG_API_URI)
	req, err := http.NewRequest("POST", real_uri, bytes.NewReader(post_json))
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

	error_msg, ok := map_result["error_msg"]; if ok {
		panic(error_msg)
	}

	return map_result["tags"].([]interface{})
}

/**
 * 获取中文词性说明
 */
func (this *AiText) GetWordKind(intype string) (string){
	out_type, ok := this.WordKindMap[intype]
	if ok {
		return out_type
	}
	return ""
}


/**
 * 分词
 */
func (this *AiText) SplitWords(intxt string) (map[string]interface{}) {
	post_arg := map[string]interface{}{
		"query":  ghostlib.UrlEncode(ghostlib.ConvertStrEncode(intxt, "utf-8", "gbk")),
	}
	post_json, err := json.Marshal(post_arg)
	if nil != err{
		panic(err)
	}

	real_uri := this.getInterFaceUri(SEG_API_URI)
	req, err := http.NewRequest("POST", real_uri, bytes.NewReader(post_json))
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

	error_msg, ok := map_result["error_msg"]; if ok {
		panic(error_msg)
	}

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
	post_arg := map[string]interface{}{
		"query":  ghostlib.UrlEncode(ghostlib.ConvertStrEncode(intxt, "utf-8", "gbk")),
	}
	post_json, err := json.Marshal(post_arg)
	if nil != err{
		panic(err)
	}

	real_uri := this.getInterFaceUri(WORDPOS_API_URI)
	req, err := http.NewRequest("POST", real_uri, bytes.NewReader(post_json))
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

	error_msg, ok := map_result["error_msg"]; if ok {
		panic(error_msg)
	}

	return map_result["result_out"].([]interface{})
}
