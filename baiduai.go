package baiduai

import (
	"net/http"
	"ghostlib"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
)

const (
	BAIDU_AI_APPID = 0	// 去ai.baidu.com申请反正不要钱	
	BAIDU_AI_KEY  = "0"
	BAIDU_AI_CRET = "0"

	API_TIMEOUT = 120
)

const (
	TOKEN_API_URI = "https://openapi.baidu.com/oauth/2.0/token"
	// 文本相关API
	SEG_API_URI = "https://aip.baidubce.com/rpc/2.0/nlp/v1/wordseg"
	WORDPOS_API_URI = "https://aip.baidubce.com/rpc/2.0/nlp/v1/wordpos"
	WORDEMBED_API_URI = "https://aip.baidubce.com/rpc/2.0/nlp/v1/wordembedding"
	DNNL_API_URI = "https://aip.baidubce.com/rpc/2.0/nlp/v1/dnnlm_cn"
	SIMNET_API_URI = "https://aip.baidubce.com/rpc/2.0/nlp/v1/simnet"
	COMTAG_API_URI = "https://aip.baidubce.com/rpc/2.0/nlp/v1/comment_tag"

	// 语音相关API
	TXT2VOICE_API_URI = "http://tsn.baidu.com/text2audio" // 语音合成
	VOICE2TXT_API_URI = "http://vop.baidu.com/server_api"  // 语音识别

	// OCR相关API
	IDCARD_API_URI = "https://aip.baidubce.com/rest/2.0/ocr/v1/idcard"
	BANKCARD_API_URI = "https://aip.baidubce.com/rest/2.0/ocr/v1/bankcard"
	GENERALOCR_API_URI = "https://aip.baidubce.com/rest/2.0/ocr/v1/general"

	// 人脸检测相关API
	FACEDETECT_API_URI = "https://aip.baidubce.com/rest/2.0/face/v1/detect"
	FACEMATCH_API_URI = "https://aip.baidubce.com/rest/2.0/faceverify/v1/match"


)


type AiClient struct {
	client       *http.Client
	access_token string
	error_code map[float64]string
	type_code map[float64]string
}

func (this *AiClient) getToken() (bool, string) {
	post_arg := map[string]interface{}{
		"client_id":     BAIDU_AI_KEY,
		"client_secret": BAIDU_AI_CRET,
		"grant_type":    "client_credentials",
	}

	resp, _ := this.client.PostForm(TOKEN_API_URI, ghostlib.InitPostData(post_arg))
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	json_result, err := simplejson.NewJson(data)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%v", string(data))

	map_result := make(map[string]interface{})
	map_result, _ = json_result.Map()
	access_token, ok := map_result["access_token"]

	if ok {
		this.access_token = ghostlib.ToString(access_token)
		return true, this.access_token
	}
	return false, ""
}
