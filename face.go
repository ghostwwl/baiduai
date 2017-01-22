package baiduai

import (
	"net/http"
	"time"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"ghostlib"
	"encoding/base64"
	"strings"
)


type AiFace struct {
	AiClient
}

func NewFace() *AiFace {
	c := new(AiFace)
	c.client = &http.Client{Timeout: API_TIMEOUT * time.Second}
	c.error_code = map[float64]string{
		100 : "无效参数",
		110 : "Token过期失效",
	}
	c.type_code = map[float64]string {
		216015 : "模块关闭",
		216100 : "非法参数",
		216101 : "参数数量不够",
		216102 : "业务不支持",
		216103 : "参数太长",
		216110 : "ID不存在",
		216111 : "非法用户ID",
		216200 : "空的图片",
		216201 : "图片格式错误",
		216202 : "图片大小错误",
		216300 : "DB错误",
		216400 : "后端系统错误",
		216401 : "内部错误",
		216402 : "没有找到人脸",
		216500 : "未知错误",
		216611 : "用户不存在",
		216613 : "用户查找不到",
		216614 : "图片信息不完整",
		216615 : "处理图片信息失败",
		216616 : "图片已存在",
		216617 : "添加用户失败",
		216618 : "群组里没有用户",
		216630 : "识别错误",
	}


	return c
}

/**
 * 人脸检测
 */
func (this *AiFace) FaceDetect(imgbytes []byte) (map[string]interface{}) {
	// 真实使用时这里要判断过期时间 避免重复获取token
	if this.access_token == "" {
		doflag, _ := this.getToken()
		if !doflag {
			panic("获取access token 失败")
		}
	}

	post_arg := map[string]interface{}{
		"image": base64.StdEncoding.EncodeToString(imgbytes),
		"max_face_num": "10",	// 最多处理人脸数目，默认值1
		"face_fields":  "age,beauty,expression,faceshape,gender,glasses,landmark,race,qualities",
		// 包括age、beauty、expression、faceshape、gender、glasses、landmark、race、qualities信息，逗号分隔，默认只返回人脸框、概率和旋转角度
	}

	real_uri := fmt.Sprintf("%s?access_token=%s", FACEDETECT_API_URI, this.access_token)
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
 * 人脸比对
 */
func (this *AiFace) FaceMatch(imgbytes... []byte) (map[string]interface{}) {
	// 真实使用时这里要判断过期时间 避免重复获取token
	if this.access_token == "" {
		doflag, _ := this.getToken()
		if !doflag {
			panic("获取access token 失败")
		}
	}

	if len(imgbytes) < 2 {
		panic("最少需要两张图比对")
	}

	var images_blist []string
	for _, ibyte := range(imgbytes) {
		images_blist = append(images_blist, base64.StdEncoding.EncodeToString(ibyte))
	}


	post_arg := map[string]interface{}{
		"images": strings.Join(images_blist, ","),
	}

	real_uri := fmt.Sprintf("%s?access_token=%s", FACEMATCH_API_URI, this.access_token)
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
