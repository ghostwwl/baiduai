package baiduai
/*****************************************
 * FileName : voice.go
 * Author   : ghostwwl
 * Note     : 百度ai的语音相关api
 *            将 中文 --> 语音 用的百度语音合成
 *            将 语音 --> 中文 用百度语音识别
 *****************************************/


import (
	"encoding/json"
	"ghostlib"
	"io/ioutil"
	"net/http"
	"time"

	"errors"
	"encoding/base64"
	"fmt"
	"bytes"
	"strings"
)

type VoiceError struct {
	Err_no  int    `json:"err_no"`
	Err_msg string `json:"err_msg"`
	Sn      string `json:"sn"`
	Idx     int    `json:"idx"`
}

type AiVoice struct {
	AiClient
}

func NewVoice() *AiVoice {
	c := new(AiVoice)
	c.client = &http.Client{Timeout: API_TIMEOUT * time.Second}
	c.error_code = map[float64]string{
		500 : "不支持输入",
		501 : "输入参数不正确",
		502 : "token验证失败",
		503 : "合成后端错误",
		3300: "输入参数不正确",
		3301: "识别错误",
		3302: "验证失败",
		3303: "语音服务器后端问题",
		3304: "请求 GPS 过大，超过限额",
		3305: "产品线当前日请求数超过限额",
	}
	return c
}

/**
 * @param voicebyte  音频文件内容
 */
func (this *AiVoice) GetText(voicebyte []byte) (string, error) {
	// 真实使用时这里要判断过期时间 避免重复获取token
	if this.access_token == "" {
		doflag, _ := this.getToken()
		if !doflag {
			return "", errors.New("获取access token 失败")
		}
	}

	/**
	 * 这里有个坑货 注意压缩格式 和采样率
	 * 如果格式对 但是采样率不对 会乱出结果
	 * 如果都格式不对 会返回 3300[输入参数错误]
	 */
	post_arg := map[string]interface{}{
		"format":  "wav",	// 压缩格式支持：pcm（不压缩）、wav、opus、speex、amr、x-flac
		"rate":  16000,	// 原始语音的录音格式目前只支持评测 8k/16k 采样率 16bit 位深的单声道语音
		"channel" : 1,
		"lan": "zh",	// 语种选择，中文=zh、粤语=ct、英文=en，不区分大小写，默认中文
		"token":  this.access_token,
		"cuid": "12:34:56:78",
		"len":  len(voicebyte),
		"speech":  base64.StdEncoding.EncodeToString(voicebyte),
	}
	post_json, err := json.Marshal(post_arg)
	if nil != err{
		return "", err
	}

	req, err := http.NewRequest("POST", VOICE2TXT_API_URI, bytes.NewReader(post_json))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, _ := this.client.Do(req)

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	map_result := map[string]interface{}{}
	json.Unmarshal(data, &map_result)

	err_no := map_result["err_no"].(float64)
	if 0 == err_no {
		result := make([]string, 0)
		for _, r := range map_result["result"].([]interface{}) {
			result = append(result, ghostlib.ToString(r))
		}
		return strings.Join(result, ""), nil
	} else {
		return "", errors.New(this.error_code[err_no])
	}
}

/**
 * @param intxt 要合成语音的文字
 */
func (this *AiVoice) GetVoice(intxt string) (bool, []byte) {
	// 真实使用时这里要判断过期时间 避免重复获取token
	if this.access_token == "" {
		doflag, _ := this.getToken()
		if !doflag {
			return false, []byte("获取access token 失败")
		}
	}

	post_arg := map[string]interface{}{
		"tex":  intxt,
		"lan":  "zh",
		"tok":  this.access_token,
		"ctp":  1,
		"cuid": "12:34:56:78", // 用户唯一标识,用来区分用户,web 端参考填写机器 mac地址或 imei 码,长度为 60 以内
		"spd":  3,             // 语速,取值 0-9,默认为 5
		"pit":  3,             // 音调,取值 0-9,默认为 5
		"vol":  5,             // 音量,取值 0-9,默认为 5
		"per":  0,             // 发音人选择, 0为女声，1为男声，3为情感合成-度逍遥，4为情感合成-度丫丫，默认为普通女声
	}

	resp, _ := this.client.PostForm(TXT2VOICE_API_URI, ghostlib.InitPostData(post_arg))


	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	contentType := resp.Header.Get("Content-type")
	fmt.Printf("%v", contentType)
	switch contentType {
	case "audio/mp3":
		return true, data
	case "application/json":
		var errobj VoiceError
		if err := json.Unmarshal(data, &errobj); nil != err {
			return false, []byte(ghostlib.ToString(err))
			//panic(err.Error())
		} else {
			return false, []byte(errobj.Err_msg)
		}
	}

	return false, nil
}


