## 练习5.1：改变findlinks程序，使用递归调用visit（而不是循环）遍历n.FirstChild链表。

## 练习5.2：写一个函数，用于统计HTML文档树内所有的元素个数，如p、div、span等。

## 练习5.3：写一个函数，用于输出HTML文档树中所有文本节点的内容。但不包括`<script>`或`<style>`元素，因为这些内容在Web浏览器中是不可见的。

## 练习5.4：扩展visit函数，使之能够获得到其他各类的链接地址，比如图片、脚本或样式表的链接。

## 练习5.5：实现函数countWordsAndImages（参照练习4.9中的单词分隔）。

## 练习5.6：修改gopl.io/ch3/surface（参考3.2节）中的函数corner，以使用命名的结果以及祼返回语句。

## 练习5.7：开发startElement和endElement函数并应用到一个通用的HTML输出代码中。输出注释节点、文本节点和所有元素属性（`<a href='...'>`）。当一个元素没有子节点时，使用简短的形式，比如<img/>而不是<img></img>。写一个测试程序保证输出语句可以正确解析（参与第11章）。

## 练习5.8：修改forEachNode使得pre和post函数返回一个布尔型的结果来确定遍历是否继续下去。使用它写一个函数ElementByID，该函数使用下面的函数签名并且找到第一个符合id属性的HTML元素。函数在找到符合条件的元素时应该尽快停止遍历。

```go
func ElementByID(doc *html.Node, id string) *html.Node
```

## 练习5.9：写一个函数expand(s string, f func(string)string)string，该函数替换参数s中每一个子字符串"$foo"为f("foo")的返回值。

## 练习5.10：重写topSort以使用map代替slice并去掉开头的排序。结果不是唯一的，验证这个结果是合法的拓扑排序。

## 练习5.11：现在有“线性代数”（linear algebra）这门课程，它的先决课程是微积分（calculus）。扩展topSort以函数输出结果。

## 练习5.12：5.5节（gopl.io/ch5/outline2）的startElement和endElement函数共享一个全局变量depth。把它们变为匿名函数以共享outline函数的一个局部变量。

## 练习5.13：修改crawl函数保存找到的页面，根据需要创建目录。不要保存不同域名下的页面。比如，如果本来的页面来自golang.org，那么就把它们保存下来但是不要保存vimeo.com下的页面。

## 练习5.14：使用广度优先遍历搜索一个不同的拓扑结构。比如，你可以借鉴拓扑排序的例子（有向图）里的课程依赖关系，计算机文件系统的分层结构（树形结构），或者从当前城市的官方网站上下载公共汽车或者地铁线路图（无向图）。

## 练习5.15：模仿sum写两个变长函数max和min。当不带任何参数调用这些函数时应该怎么应对？编写类似函数的变种，要求至少需要一个参数。

## 练习5.16：写一个变长版本的strings.Join函数。

## 练习5.17：写一个变长函数ElementsByTagname，已知一个HTML节点树和零个或多个名字，返回所有符合给出名字的元素。下面有两个示例调用：

```go
func ElementsByTagname(doc *html.Node, name ...string) []*html.Node
images := ElementsByTagname(doc, "img")
headings := ElementsByTagname(doc, "h1", "h2", "h3", "h4")
```

## 练习5.18 不改变原本的行为，重写fetch函数以使用defer语句关闭打开的可写的文件。

## 练习5.19：使用panic和recover写一个函数，它没有return语句，但是能够返回一个非零的值。
