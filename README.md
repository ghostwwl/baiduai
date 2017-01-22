# baiduai

### 代码以GPL协议放出 
### Author:ghostwwl@gmail.com 
### ghostlib 库看 https://github.com/ghostwwl/go-code/blob/master/ghostlib.go 呢 
### 代码里我的app_key 和 app_cret 别用的毕竟免费有限制 
### go没咋学 边看手册边写 有不对的 请直接邮件 

### 使用说明呢:

```
package main

import (
	"ghostlib"
	"fmt"
	"io/ioutil"
	"baiduai"
)

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
			fmt.Printf("type: %s\n-------\n", r["type"].(string))
		}
	}
}

func main(){
	//test_text2voice()
	//test_voice2txt()
	test_ctag()
	//test_splitword()
	//test_wordpos()
}

```


>
1. 百度ai主页：[http://ai.baidu.com](http://ai.baidu.com).
2. 百度ai接口文档：[http://ai.baidu.com/docs](http://ai.baidu.com/docs).
3. 怎么申请 自己去看首页
4. 百度ai的开放接口golang sdk呢


--------------------------------

**免费接口每天有调用次数限制**





