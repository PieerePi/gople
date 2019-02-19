## 练习 1.1：修改echo程序输出os.Args[0]，即命令的名字。

基于echo1.go

```go
	fmt.Println(os.Args[0])
```

## 练习 1.2：修改echo程序，输出参数的索引和值，每行一个。

基于echo1.go

```go
	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		fmt.Println(i, v)
	}
```

## 练习 1.3：尝试测量可能低效的程序和使用strings.Join的程序在执行时间上的差异。（1.6节有time包，11.4节展示如何撰写系统性的性能评估测试。）

`go test gople\ch1\e1.3\strop_test.go -bench=. -slen=10 -scount=1000`

slen指定了单个字符串的长度；scount指定了总的字符串的个数。

```go
package strop_test

import (
	"flag"
	"fmt"
	"strings"
	"testing"
)

// for BenchmarkStrOp1, the higher the value, the worse the performance
// try 100 1000 10000
var scount = flag.Int("scount", 1000, "total string count")

// try 1 10 100
var slen = flag.Int("slen", 10, "single string length")

var longStrArray []string

func init() {
	var str = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	flag.Parse()
	longStrArray = make([]string, *scount)
	for i := range longStrArray {
		for j := 0; j < *slen; j++ {
			longStrArray[i] += str[i%10]
		}
	}
	fmt.Println("strop_test init done")
}

// BenchmarkStrOp1 for test
func BenchmarkStrOp1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var result string // result should be declared here, in the loop
		for _, v := range longStrArray {
			result += v + " "
		}
	}
}

// BenchmarkStrOp2 join test
func BenchmarkStrOp2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(longStrArray, " ")
	}
}
```

## 练习1.4：修改dup2程序，输出出现重复行的文件的名称。

## 练习1.5：改变利萨茹程序的画板颜色为绿色黑底来增加真实性。使用color.RGBA{0xRR,0xGG,0xGG,0xff}创建一种Web颜色#RRGGBB，每一对十六进制数字表示组成一个像素红、绿、蓝分量的亮度。

## 练习1.6：通过在画板中添加更多颜色，然后通过有趣的方式改变SetColorIndex的第三个参数，修改利萨茹程序来产生多种色彩的图片。

## 练习1.7：函数io.Copy(dst,src)从src读，并且写入dst。使用它代替ioutil.ReadAll来复制响应内容到os.Stdout，这样不需要装下整个响应数据流的缓冲区。确保检查io.Copy返回的错误结果。

## 练习1.8：修改fetch程序添加一个http://前缀（假如该URL参数缺失协议前缀）。可能会用到strings.HasPrefix。

## 练习1.9：修改fetch来输出HTTP的状态码，可以在resp.Status中找到它。

## 练习1.10：找出一个产生大量数据的网站。连续再次运行fetchall，看报告的时间是否会有大的变化，调查缓存情况。每一次获取的内容一样吗？修改fetchall将内容输出到文件，这样可以检查它是否一致。

## 练习1.11：使用更长的参数列表来尝试fetchall，例如使用alexa.com排名前100万的网站。如果一个网站没有响应，程序的行为是怎样的？（8.9节会通过复制这个例子来描述响应的机制。）

## 练习1.12：修改萨利茹服务器以通过URL参数读取参数值。例如，你可以通过调整它，使得像http://localhost:8000/?cycles=20这样的网址将其周期设置为20，以替代默认的5。使用strconv.Atoi函数来将字符串参数转化为整形。可以通过go doc strconv.Atoi来查看文档。
