# baiduai

- 代码以GPL协议放出 
- Author:ghostwwl@gmail.com 
- ghostlib 库看 https://github.com/ghostwwl/go-code/blob/master/ghostlib.go 呢 
- 代码里我的app_key 和 app_cret 别用的毕竟免费有限制 
- go没咋学 边看手册边写 有不对的 请直接邮件 

### 使用说明呢:

```
package main


import (
	"ghostlib"
	"fmt"
	"io/ioutil"
	"baiduai"
)

/**
 * 语音合成
 */
func test_text2voice(){
	voice := baiduai.NewVoice()
	T := `
晚上跑完步有点口渴，到路边摊买橘子，挑了几个卖相好的。
老板又拿起几个有斑点的说：“帅哥，这种长得不好看的其实更甜。”
我颇有感悟的说：“是因为橘子觉得自己长的不好看，所以努力让自己变得更甜吗？”
老板微微一愣道：“不是，我想早点卖完回家。”
`
	T = `
王辉。中国美术家协会会员。1980年生于湖北。2003年毕业于湖北美术学院。
中国画专业。学士学位。2006年毕业于广西艺术学院。中国画专业。硕士学位。
现任职于四川绵阳师范学院美术学院。中国画教师。作品多次入选全国性学术展览并获奖。
	`
	flag, result := voice.GetVoice(T)
	if !flag {
		ghostlib.Msg(string(result), 3)
	} else {
		fmt.Printf("\nlen:%v", len(result))
		err := ioutil.WriteFile("/data1/bd_v.mp3", result, 0766)
		if err != nil {
			ghostlib.Msg("写入结果文件[/data1/bd_v.mp3]出错", 3)
		}
	}
}

/**
 * 语音识别
 */
func test_voice2txt(){
	engine := baiduai.NewVoice()
	// bd_voice.wav 输入的是 单声道16k采样
	r, err := ioutil.ReadFile("/data1/bd_voice.wav")
	if nil != err{
		panic(err)
	}

	txtresult, err := engine.GetText(r)
	if nil != err {
		fmt.Printf("\n%v\n", err)
	}
	fmt.Printf("%v\n", txtresult)
}

/**
 * 评论观点抽取 或 情感识别  垃圾 很容易没结果
 */
func test_ctag(){
	T := "京东商城的商品质量很好,价格也非常便宜,京东商城的快递很迅速,快递员服务态度也非常好,在京东商城购物省时省力省钱"
	T = "辛苦快递小哥了！抢购的！相信京东没喝应该是正品！个人觉得盒子好小，毕竟只有400Ml幸好有纸质包装袋，可以装一下，才敢送出手！不管怎么说这个牌子广告响！值吧！就是赠品一个都没送，还买了3箱"
	engine := baiduai.NewText()
	x := engine.GetCommentTag(T)
	//fmt.Println(x)

	fmt.Printf("\n---------\nsrc:%s\n------\n", T)
	for _, r := range(x) {
		r := r.(map[string]interface{})
		fmt.Printf("abstract: %s\n", r["abstract"].(string))
		fmt.Printf("fea: %s\n", r["fea"].(string))
		fmt.Printf("adj: %s\n-------\n", r["adj"].(string))
	}
}

/**
 * 分词
 */
func test_splitword(){
	T := `
王辉。中国美术家协会会员。1980年生于湖北。2003年毕业于湖北美术学院。
中国画专业。学士学位。2006年毕业于广西艺术学院。中国画专业。硕士学位。
现任职于四川绵阳师范学院美术学院。中国画教师。作品多次入选全国性学术展览并获奖。
	`
	engine := baiduai.NewText()
	x := engine.SplitWords(T)


	namebuf := x["namebuf"]		// 人名
	subphrbuf := x["subphrbuf"]	// 短语
	wordsepbuf := x["wordsepbuf"]	// 标准粒度
	wpcompbuf := x["wpcompbuf"]	// 混排粒度
	newwordbuf := x["pnewword"].(map[string]interface{})["newwordbuf"] // 新词

	fmt.Printf("src:%v", T)
	fmt.Printf("\n人名: %v", namebuf)
	fmt.Printf("\n短语: %v", subphrbuf)
	fmt.Printf("\n新词: %v", newwordbuf)
	fmt.Printf("\n标准: %v", wordsepbuf)
	fmt.Printf("\n混排: %v", wpcompbuf)

}

/**
 * 词性标注
 */
func test_wordpos(){
	T := `
王辉。中国美术家协会会员。1980年生于湖北。2003年毕业于湖北美术学院。
中国画专业。学士学位。2006年毕业于广西艺术学院。中国画专业。硕士学位。
现任职于四川绵阳师范学院美术学院。中国画教师。作品多次入选全国性学术展览并获奖。
	`
	engine := baiduai.NewText()
	x := engine.WordPos(T)
	for _, r := range(x) {
		r := r.(map[string]interface{})
		// 非符号的词
		if "w" != r["type"].(string) {
			fmt.Printf("word: %s\n", r["word"].(string))
			fmt.Printf("type: %s\n", r["type"].(string))
			fmt.Printf("kind: %s\n-------\n", engine.GetWordKind(r["type"].(string)))

		}
	}
}

/**
 * 身份证识别
 */
func test_ocridcard(){
	engine := baiduai.NewOcr()
	//r, err := ioutil.ReadFile("/data1/s2.jpg")
	r, err := ioutil.ReadFile("/data1/s1.png")
	if nil != err{
		panic(err)
	}
	x := engine.OcrIdCard(r, true)
	//fmt.Printf("%v\n", x)
	result := x["words_result"].(map[string]interface{})
	for kk, vv := range(result){
		fmt.Printf("--------\n%v:%v\n", kk, vv.(map[string]interface{})["words"])
	}
}

/**
 * 银行卡识别
 */
func test_ocrbankcard() {
	engine := baiduai.NewOcr()
	r, err := ioutil.ReadFile("/data1/b2.jpg")
	if nil != err{
		panic(err)
	}
	x := engine.OcrBankCard(r)
	fmt.Printf("%v\n", x)
	result := x["result"].(map[string]interface{})
	for kk, vv := range(result){
		fmt.Printf("--------\n%v:%v\n", kk, vv)
	}
}

/**
 * 文字识别
 */
func test_ocrgeneral(){
	engine := baiduai.NewOcr()
	r, err := ioutil.ReadFile("/data1/11.jpg")
	if nil != err{
		panic(err)
	}
	x := engine.OcrGeneral(r)
	//fmt.Printf("%v\n", x)
	result := x["words_result"].([]interface{})
	for kk, vv := range(result){
		fmt.Printf("--------\n%v:%v\n", kk, vv.(map[string]interface{})["words"])
	}
}

/**
 * 人脸检测
 */
func test_facedetect(){
	engine := baiduai.NewFace()
	r, err := ioutil.ReadFile("/data1/f1.jpg")
	if nil != err{
		panic(err)
	}
	x := engine.FaceDetect(r)
	fmt.Printf("%v\n", x)
}

/**
 * 人脸相似度匹配
 */
func test_facematch(){
	engine := baiduai.NewFace()
	r, err := ioutil.ReadFile("/data1/f2.jpg")
	if nil != err{
		panic(err)
	}
	r1, err := ioutil.ReadFile("/data1/f3.jpg")
	if nil != err{
		panic(err)
	}
	r2, err := ioutil.ReadFile("/data1/f4.jpg")
	if nil != err{
		panic(err)
	}
	r3, err := ioutil.ReadFile("/data1/f5.jpg")
	if nil != err{
		panic(err)
	}
	x := engine.FaceMatch(r, r1, r2, r3)
	//fmt.Printf("%v\n", x)

	results := x["results"].([]interface{})
	for i, r := range(results) {
		r := r.(map[string]interface{})
		fmt.Printf("\n--------第%v组比对----------\n", i)
		index_i := ghostlib.ToInt64(r["index_i"]) + 1
		index_j := ghostlib.ToInt64(r["index_j"]) + 1
		fmt.Printf("img[%v]与img[%v] 相似度:%v%%\n", index_i, index_j, r["score"])
	}
}

func main(){
	//test_text2voice()
	//test_voice2txt()
	//test_ctag()
	//test_splitword()
	//test_wordpos()

	//test_ocridcard()
	//test_ocrbankcard()
	//test_ocrgeneral()

	//test_facedetect()
	test_facematch()
}



```

输出如下:

```
/usr/local/go/bin/go run /data/ghostwwl/project/Go/src/test_bdai.go

--------第0组比对----------
img[1]与img[2] 相似度:83.063133239746%

--------第1组比对----------
img[1]与img[3] 相似度:94.308540344238%

--------第2组比对----------
img[1]与img[4] 相似度:74.360145568848%

--------第3组比对----------
img[2]与img[3] 相似度:90.899742126465%

--------第4组比对----------
img[2]与img[4] 相似度:81.979782104492%

--------第5组比对----------
img[3]与img[4] 相似度:86.821075439453%

Process finished with exit code 0

```


>
1. 百度ai主页：[http://ai.baidu.com](http://ai.baidu.com).
2. 百度ai接口文档：[http://ai.baidu.com/docs](http://ai.baidu.com/docs).
3. 怎么申请 自己去看首页
4. 百度ai的开放接口golang sdk呢


--------------------------------

**免费接口每天有调用次数限制**





