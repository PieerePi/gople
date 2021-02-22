# 第七章 接口

- 练习7.1：使用类似ByteCounter的想法，实现单词和行的计数器。实现时考虑使用bufio.ScanWords。

- 练习7.2：实现一个满足如下签名的CountingWriter函数，输入一个io.Writer，输出一个封装了输入值的新Writer，以及一个指向int64的指针，该指针对应的值是新的Writer写入的字节数。

```go
func CountingWriter(w io.Writer) (io.Writer, *int64)
```

- 练习7.3：为`gopl.io/ch4/treesort`中的*tree类型（见4.4节）写一个String方法，用于展示其中的值序列。

- 练习7.4：strings.NewReader函数输入一个字符串，返回一个从字符串读取数据且满足io.Reader接口（也满足其他接口）的值。请自己实现该函数，并且通过它来让HTML分析器（参与5.2节）支持以字符串作为输入。

- 练习7.5：io包中的LimitReader函数接受io.Reader r和字节数n，返回一个Reader，该返回值从r读取数据，但在读取n字节后报告文件结束。请实现该函数。

```go
func LimitReader(r io.Reader, n int64) io.Reader
```

- 练习7.6：在tempflag中支持热力学温度。

- 练习7.7：请解释为什么默认值20.0没写°C，而帮助消息中却包含°C。

- 练习7.8：很多图形界面提供了一个表格控件，它支持有状态的多层排序：先按照最近单击的列来排序，接着是上一次单击的列，依次类推。请定义sort.Interface接口实现来满足如上需求。试比较这个方法与多次使用sort.Stable排序的异同。

- 练习7.9：利用html/template（见4.6节）来替换printTracks函数，使用HTML表格来显示音乐列表。结合上一个练习，来实现通过单击列头来发送HTTP请求，进而对表格排序。

- 练习7.10：sort.Interface也可以用于其他用途。试与一个函数IsPalindrome(s sort.Interface)bool来判断一个序列是否是回文，即序列反转后是否保持不变。可以假定对于下标分别为i、j的元素，如果!s.Less(i,j) && !s.Less(j,i)，那么两个元素相等。

- 练习7.11：增加额外的处理程序，来支持创建、读取、更新和删除数据库条目。比如，/update?item=socks&price=6这样的请求将更新仓库中物品的价格，如果商品不存在或者价格无效就返回错误。（注意：这次修改会引入并发变量修改）

- 练习7.12：修改/list处理程序，改为输出HTML表格，而不是纯文本。可以考虑使用html/template包（参考4.6节）

- 练习7.13：给Expr增加一个String方法用来美化输出语法树。要求生成的语法树重新解析后是完全一致的树。

- 练习7.14：定义一个新的满足Expr接口的具体类，提供一个新操作，比如计算它的操作数的最小值。因为Parse函数无法实例化新创建的类型，所以测试时需要直接构造语法树（当然，也可以扩充一下解析函数）。

- 练习7.15：写一个程序从标准输入读取一个表达式，提示用户输入表达式中变量的值，最后计算表达式的值。请妥善处理各种异常。

- 练习7.16：写一个基于Web的计算器程序。

- 练习7.17：扩展xmlselect，让我们不仅可以用名字，还可以用CSS风格的属性来做选择。比如一个`<div id="page" class="wide">`元素，不仅可以通过名字，还可以通过id和class来做匹配。

- 练习7.18：使用基于标记的解析API，写一个程序来读入一个任意的XML文档，构造出一棵树来展现XML中的主要节点。节点包括两种类型：CharData节点表示文本字符串，Element节点表示元素及其属性。每个元素节点包含它的子节点数组。

可以参考如下类型定义：

```go
import "encoding/xml"

type Node interface {} // CharData 或 *Element

type CharData string

type Element struct {
    Type     xml.Name
    Attr     []xml.Attr
    Children []Node
}
```
