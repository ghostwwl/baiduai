package baiduai

import (
	"net/http"
	"time"
	"fmt"
	//"errors"
	"encoding/json"
	"io/ioutil"
	"ghostlib"
	"encoding/base64"
)


type AiOcr struct {
	AiClient
}

func NewOcr() *AiOcr {
	c := new(AiOcr)
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
 * 身份证识别
 */
func (this *AiOcr) OcrIdCard(imgbytes []byte, isFront bool) (map[string]interface{}) {
	// 真实使用时这里要判断过期时间 避免重复获取token
	if this.access_token == "" {
		doflag, _ := this.getToken()
		if !doflag {
			panic("获取access token 失败")
		}
	}

	post_arg := map[string]interface{}{
		"image": base64.StdEncoding.EncodeToString(imgbytes),
		"id_card_side": "front",	// front 正面  back 背面
		"detect_direction":  "false", // 是否检测图像朝向[true/false]，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。
	}
	if !isFront {
		post_arg["id_card_side"] = "back"
	}

	real_uri := fmt.Sprintf("%s?access_token=%s", IDCARD_API_URI, this.access_token)
	resp, _ := this.client.PostForm(real_uri, ghostlib.InitPostData(post_arg))

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		panic(err)
	}

	map_result := make(map[string]interface{})
	json.Unmarshal(data, &map_result)

	error_msg, ok := map_result["error_msg"]; if ok {
		panic(error_msg)
	}

	return map_result
}

/**
 * 身份证识别
 */
func (this *AiOcr) OcrBankCard(imgbytes []byte) (map[string]interface{}) {
	// 真实使用时这里要判断过期时间 避免重复获取token
	if this.access_token == "" {
		doflag, _ := this.getToken()
		if !doflag {
			panic("获取access token 失败")
		}
	}

	post_arg := map[string]interface{}{
		"image": base64.StdEncoding.EncodeToString(imgbytes),
	}
	real_uri := fmt.Sprintf("%s?access_token=%s", BANKCARD_API_URI, this.access_token)
	resp, _ := this.client.PostForm(real_uri, ghostlib.InitPostData(post_arg))

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		panic(err)
	}

	map_result := make(map[string]interface{})
	json.Unmarshal(data, &map_result)

	error_msg, ok := map_result["error_msg"]; if ok {
		panic(error_msg)
	}

	return map_result
}

/**
 * 通用ocr识别
 */
func (this *AiOcr) OcrGeneral(imgbytes []byte) (map[string]interface{}) {
	// 真实使用时这里要判断过期时间 避免重复获取token
	if this.access_token == "" {
		doflag, _ := this.getToken()
		if !doflag {
			panic("获取access token 失败")
		}
	}

	post_arg := map[string]interface{}{
		"image": base64.StdEncoding.EncodeToString(imgbytes),
		"recognize_granularity": "big",	// 是否定位单字符位置，big：不定位单字符位置，默认值；small：定位单字符位置
		"mask":  "", // 是否检测图像朝向[true/false]，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。
		"language_type" : "CHN_ENG", // CHN_ENG：中英文混合； ENG：英文； POR：葡萄牙语； FRE：法语； GER：德语； ITA：意大利语； SPA：西班牙语； RUS：俄语； JAP：日语
		"detect_direction" : "false", // 是否检测图像朝向[true/false]，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。
		"detect_language" : "false", // 是否检测语言，默认不检测。当前支持（中文、英语、日语、韩语）
		"classify_dimension" : "lottery", // 分类维度（根据OCR结果进行分类）
		"vertexes_location" : "false", //是否返回文字外接多边形顶点位置
	}

	real_uri := fmt.Sprintf("%s?access_token=%s", GENERALOCR_API_URI, this.access_token)
	resp, _ := this.client.PostForm(real_uri, ghostlib.InitPostData(post_arg))

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		panic(err)
	}

	map_result := make(map[string]interface{})
	json.Unmarshal(data, &map_result)

	error_msg, ok := map_result["error_msg"]; if ok {
		panic(error_msg)
	}

	return map_result
}
