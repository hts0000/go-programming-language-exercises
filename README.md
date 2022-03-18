# Exercises

# Chapter 1 Tutorial
## Exercise 1.1 (P8)
Modify the echo program to also print os.Args[0], the name of the command that invoked it.

修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
- 完成

## Exercise 1.2 (P8)
Modify the echo program to print the index and value of each of its arguments, one per line.
修改echo程序，使其打印每个参数的索引和值，每个一行。
- 完成

## Exercise 1.3 (P8)
Experiment to measure the difference in running time between our potentially inefficient versions and the one that usesstrings.Join. (Section 1.6 illustrates part of the time package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)

做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
- 完成

## Exercise 1.4 (P8)
Modifydup2to print the names of all files in which each duplicated line occurs.

修改dup2，出现重复的行时打印文件名称。

## Exercise 1.5 (P8)
Change the Lissajous program’s color palette to green on black, for added authenticity. To create the web color#RRGGBB, usecolor.RGBA{0xRR, 0xGG, 0xBB, 0xff}, where each pair of hexadecimal digits represents the intensity of the red, green, or blue component of the pixel.

修改前面的Lissajous程序里的调色板，由黑色改为绿色。我们可以用color.RGBA{0xRR, 0xGG, 0xBB, 0xff}来得到#RRGGBB这个色值，三个十六进制的字符串分别代表红、绿、蓝像素。

## Exercise 1.6 (P8)
Modify theLissajousprogram to produce images in multiple colors by adding more values topaletteand then displaying them by changing the third argument ofSetColorIndexin some interesting way.

修改Lissajous程序，修改其调色板来生成更丰富的颜色，然后修改SetColorIndex的第三个参数，看看显示结果吧。

## Exercise 1.7 (P8)
The function callio.Copy(dst,src)reads from src and writes to dst. Use it instead ofioutil.ReadAllto copy the response body toos.Stdoutwithout requiring a buffer large enough to hold the entire stream. Be sure to check the error result ofio.Copy.

函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，使用这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout，避免申请一个缓冲区（例子中的b）来存储。记得处理io.Copy返回结果中的错误。

## Exercise 1.8 (P8)
Modifyfetchto add the prefix http:// to each argument URL if it is missing. You might want to usestrings.HasPrefix.

修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。

## Exercise 1.9 (P8)
Modifyfetchto also print the HTTP status code, found inresp.Status.

修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。

## Exercise 1.10 (P8)
Find a web site that produces a large amount of data. Investigate caching by running fetchall twice in succession to see whether the reported time changes much. Do you get the same content each time? Modify fetchall to print its output to a file so it can be examined.

找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个URL执行两遍请求，查看两次时间是否有较大的差别，并且每次获取到的响应内容是否一致，修改本节中的程序，将响应结果输出，以便于进行对比。

## Exercise 1.11 (P8)
Try fetchall with longer argument lists, such as samples from the top million web sites available atalexa.com. How does the program behave if a web site just doesn’t respond? (Section 8.9 describes mechanisms for coping in such cases.)

在fetchall中尝试使用长一些的参数列表，比如使用在alexa.com的上百万网站里排名靠前的。如果一个网站没有回应，程序将采取怎样的行为？（Section8.9 描述了在这种情况下的应对机制）。

## Exercise 1.12 (P8)
Modify the Lissajous server to read parameter values from the URL. For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets the number of cycles to 20 instead of the default 5. Use the strconv.Atoi function to convert the string parameter into an integer. You can see its documentation with go doc strconv.Atoi.

修改Lissajour服务，从URL读取变量，比如你可以访问 http://localhost:8000/?cycles=20 这个URL，这样访问可以将程序里的cycles默认的5修改为20。字符串转换为数字可以调用strconv.Atoi函数。你可以在godoc里查看strconv.Atoi的详细说明。

# Chapter 2 Program Structure
## Exercise 2.1
Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.

向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，Kelvin 绝对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。

## Exercise 2.2
Write a general-purpose unit-conversion program analogous to cf that reads numbers from its command-line arguments or from the standard input if there are no arguments, and converts each number into units like temperature in Celsius and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the like.

写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。

## Exercise 2.3
Rewrite PopCount to use a loop instead of a single expression. Compare the performance of the two versions. (Section 11.4 shows how to compare the performance of different implementations systematically.)

重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。（11.4节将展示如何系统地比较两个不同实现的性能。）

## Exercise 2.4
Write a version of PopCount that counts bits by shifting its argument through 64 bit positions, testing the right most bit each time. Compare its performance to the table-lookup version.

用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。比较和查表算法的性能差异。

## Exercise 2.5
The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version of PopCountthat counts bits by using this fact, and assess its performance.

表达式x&(x-1)用于将x的最低的一个非零的bit位清零。使用这个算法重写PopCount函数，然后比较性能。

# Chapter 3 Basic Data Types
## Exercise 3.1
If the function f returns a non-finite float64 value, the SVG file will contain invalid `<polygon>` elements (although many SVG renderers handle this gracefully). Modify the program to skip invalid polygons.

如果f函数返回的是无限制的float64值，那么SVG文件可能输出无效的多边形元素（虽然许多SVG渲染器会妥善处理这类问题）。修改程序跳过无效的多边形。

## Exercise 3.2
Experiment with visualizations of other functions from the math package. Can you produce an egg box, moguls, or a saddle?

试验math包中其他函数的渲染图形。你是否能输出一个egg box、moguls或a saddle图案?

## Exercise 3.3
Color each polygon based on its height, so that the peaks are colored red (#ff0000) and the valleys blue (#0000ff).

根据高度给每个多边形上色，那样峰值部将是红色（#ff0000），谷部将是蓝色（#0000ff）。

## Exercise 3.4
Following the approach of the Lissajous example in Section 1.7, construct a web server that computes surfaces and writes SVG data to the client. The server must set the Content-Type header like this:

参考1.7节Lissajous例子的函数，构造一个web服务器，用于计算函数曲面然后返回SVG数据给客户端。服务器必须设置Content-Type头部：

## Exercise 3.5
Implement a full-color Mandelbrot set using the function image.NewRGBA and the type color.RGBA or color.YCbCr.

实现一个彩色的Mandelbrot图像，使用image.NewRGBA创建图像，使用color.RGBA或color.YCbCr生成颜色。

## Exercise 3.6
Supersampling is a technique to reduce the effect of pixelation by computing the color value at several points within each pixel and taking the average. The simplest method is to divide each pixel into four ‘‘subpixels.’’ Implement it.

升采样技术可以降低每个像素对计算颜色值和平均值的影响。简单的方法是将每个像素分成四个子像素，实现它。

## Exercise 3.7
Another simple fractal uses Newton’s method to find complex solutions to a function such as z4−1 = 0. Shade each starting point by the number of iterations required to get close to one of the four roots. Color each point by the root it approaches.

另一个生成分形图像的方式是使用牛顿法来求解一个复数方程，例如$z^4-1=0$。每个起点到四个根的迭代次数对应阴影的灰度。方程根对应的点用颜色表示。

## Exercise 3.8
Rendering fractals at high zoom levels demands great arithmetic precision. Implement the same fractal using four different representations of numbers: complex64, complex128, big.Float, and big.Rat. (The latter two types are found in the math/big package. Float uses arbitrary but bounded-precision floating-point; Rat uses unbounded-precision rational numbers.) How do they compare in performance and memory usage? At what zoom levels do rendering artifacts become visible?

通过提高精度来生成更多级别的分形。使用四种不同精度类型的数字实现相同的分形：complex64、complex128、big.Float和big.Rat。（后面两种类型在math/big包声明。Float是有指定限精度的浮点数；Rat是无限精度的有理数。）它们间的性能和内存使用对比如何？当渲染图可见时缩放的级别是多少？

## Exercise 3.9
Write a web server that renders fractals and writes the image data to the client. Allow the client to specify thex,y, and zoom values as parameters to the HTTP request.

编写一个web服务器，用于给客户端生成分形的图像。运行客户端通过HTTP参数指定x、y和zoom参数。

## Exercise 3.10
Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.

编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。

## Exercise 3.11
Enhance comma so that it deals correctly with floating-point numbers and an optional sign.

完善comma函数，以支持浮点数处理和一个可选的正负号的处理。

## Exercise 3.12
Write a function that reports whether two strings are anagrams of each other, that is, they contain the same letters in a different order.

编写一个函数，判断两个字符串是否是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。

## Exercise 3.13
Write const declarations for KB, MB, up through YB as compactly as you can.

编写KB、MB的常量声明，然后扩展到YB。

# Chapter 4 Composite Types
## Exercise 4.1
Write a function that counts the number of bits that are different in two SHA256 hashes. (See PopCount from Section2.6.2.)

编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)

## Exercise 4.2
Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print the SHA384 or SHA512 hash instead.

编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。

## Exercise 4.3
Rewrite rverse to use an array pointer instead of a slice.

重写reverse函数，使用数组指针代替slice。

## Exercise 4.4
Write a version of rotate that operates in a single pass.

编写一个rotate函数，通过一次循环完成旋转。

## Exercise 4.5
Write an in-place function to eliminate adjacent duplicates in a []string slice.

写一个函数在原地完成消除[]string中相邻重复的字符串的操作。

## Exercise 4.6
Write an in-place function that squashes each run of adjacent Unicode spaces (seeunicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.

编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回

## Exercise 4.7
Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string, in place. Can you do it without allocating new memory?

修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？

## Exercise 4.8
Modify charcount to count letters, digits, and so on in their Unicode categories, using functions like unicode.IsLetter.

修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。

## Exercise 4.9
Write a program wordfreq to report the frequency of each word in an input text file.Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead of lines.

编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。

## Exercise 4.10
Modify issues to report the results in age categories, say less than a month old, less than a year old, and more than a year old.

修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。

## Exercise 4.11
Build a tool that lets users create, read, update, and delete GitHub issues from the command line, invoking their preferred text editor when substantial text input is required.

编写一个工具，允许用户在命令行创建、读取、更新和关闭GitHub上的issue，当必要的时候自动打开用户默认的编辑器用于输入文本信息。

## Exercise 4.12
The popular web comic xkcd has a JSON interface. For example, a request to https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd that, using this index, prints the URL and transcript of each comic that matches a search term provided on the command line.

流行的web漫画服务xkcd也提供了JSON接口。例如，一个 https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述。下载每个链接（只下载一次）然后创建一个离线索引。编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。

## Exercise 4.13
The JSON-based web service of the Open Movie Database lets you search https://omdbapi.com/ for a movie by name and download its poster image. Write a tool poster that downloads the poster image for the movie named on the command line.

使用开放电影数据库的JSON服务接口，允许你检索和下载 https://omdbapi.com/ 上电影的名字和对应的海报图像。编写一个poster工具，通过命令行输入的电影名字，下载对应的海报。

## Exercise 4.14
Create a web server that queries GitHub once and then allows navigation of the list of bug reports, milestones, and users.

创建一个web服务器，查询一次GitHub，然后生成BUG报告、里程碑和对应的用户信息。

# Chapter 5 Functions
## Exercise 5.1
Change the findlinks program to traverse the n.FirstChild linked list using recursive calls to visit instead of a loop.

修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。

## Exercise 5.2
Write a function to populate a mapping from element names—p, div, span, and so on—to the number of elements with that name in an HTML document tree.

编写函数，记录在HTML树中出现的同名元素的次数。

## Exercise 5.3
Write a function to print the contents of all text nodes in an HTML document tree. Do not descend into`<script>`or`<style>`elements, since their contents are not visible in a web browser.

编写函数输出所有text结点的内容。注意不要访问`<script>`和`<style>`元素，因为这些元素对浏览者是不可见的。

## Exercise 5.4
Extend the visit function so that it extracts other kinds of links from the document, such as images, scripts, and style sheets.

扩展visit函数，使其能够处理其他类型的结点，如images、scripts和style sheets。

## Exercise 5.5
Implement countWordsAndImages. (See Exercise 4.9 for word-splitting.)

实现countWordsAndImages。（参考练习4.9如何分词）

## Exercise 5.6
Modify the corner function in gopl.io/ch3/surface (§3.2) to use named results and a bare return statement.

修改gopl.io/ch3/surface（§3.2）中的corner函数，将返回值命名，并使用bare return。

## Exercise 5.7
Develop startElement and endElement into a general HTML pretty-printer. Print comment nodes, text nodes, and the attributes of each element (`<a href='...'>`). Use short forms like `<img/>` instead of `<img>` `</img>` when an element has no children. Write a test to ensure that the output can be parsed successfully. (See Chapter 11.)

完善startElement和endElement函数，使其成为通用的HTML输出器。要求：输出注释结点，文本结点以及每个元素的属性（`< a href='...'>`）。使用简略格式输出没有孩子结点的元素（即用`<img/>`代替`<img></img>`）。编写测试，验证程序输出的格式正确。（详见11章）

## Exercise 5.8
Modify forEachNode so that the pre and post functions return a boolean result indicating whether to continue the traversal. Use it to write a function ElementByID with the following signature that finds the first HTML element with the specified id attribute. The function should stop the traversal as soon as a match is found.

修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNoded的遍历。使用修改后的代码编写ElementByID函数，根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。

## Exercise 5.9
Write a function expand(s string, f func(string) string) string that replaces each substring ‘‘$foo’’ within s by the text returned by f("foo").

编写函数expand，将s中的"foo"替换为f("foo")的返回值。

## Exercise 5.10
Rewrite topoSort to use maps instead of slices and eliminate the initial sort. Verify that the results, though nondeterministic, are valid topological orderings.

重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。

## Exercise 5.11
The instructor of the linear algebra course decides that calculus is now a prerequisite. Extend the topoSort function to report cycles.  

现在线性代数的老师把微积分设为了前置课程。完善topSort，使其能检测有向图中的环。

## Exercise 5.12
The startElement and endElement functions in gopl.io/ch5/outline2(§5.5) share a global variable, depth. Turn them into anonymous functions that share a variable local to the outline function.

gopl.io/ch5/outline2（5.5节）的startElement和endElement共用了全局变量depth，将它们修改为匿名函数，使其共享outline中的局部变量。

## Exercise 5.13
Modify crawl to make local copies of the pages it finds, creating directories as necessary. Don’t make copies of pages that come from a different domain. For example, if the original page comes from golang.org, save all files from there, but exclude ones from vimeo.com.

修改crawl，使其能保存发现的页面，必要时，可以创建目录来保存这些页面。只保存来自原始域名下的页面。假设初始页面在golang.org下，就不要保存vimeo.com下的页面。

## Exercise 5.14
Use the breadthFirst function to explore a different structure. For example, you could use the course dependencies from the topoSort example (a directed graph), the file system hierarchy on your computer (a tree), or a list of bus or subway routes downloaded from your city government’s web site (an undirected graph).

使用breadthFirst遍历其他数据结构。比如，topoSort例子中的课程依赖关系（有向图）、个人计算机的文件层次结构（树）；你所在城市的公交或地铁线路（无向图） 。

## Exercise 5.15
Write variadic functions max and min, analogous to sum. What should these functions do when called with no arguments? Write variants that require at least one argument.

编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。

## Exercise 5.16
Write a variadic version of strings.Join.

编写多参数版本的strings.Join。

## Exercise 5.17
Write a variadic function ElementsByTagName that, given an HTML node tree and zero or more names, returns all the elements that match one of those names. Here are two example calls:

编写多参数版本的ElementsByTagName，函数接收一个HTML结点树以及任意数量的标签名，返回与这些标签名匹配的所有元素。下面给出了2个例子：

## Exercise 5.18
Without changing its behavior, rewrite the fetch function to use defer to close the writable file.

不修改fetch的行为，重写fetch函数，要求使用defer机制关闭文件。

## Exercise 5.19
Use panic and recover to write a function that contains no return statement yet returns a non-zero value.

使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。

# Chapter 6 Methods
## Exercise 6.1
Implement these additional methods:

为bit数组实现下面这些方法

## Exercise 6.2
Define a variadic (*IntSet).AddAll(...int) method that allows a list of values to be added, such as s.AddAll(1, 2, 3).

定义一个变参方法(*IntSet).AddAll(...int)，这个方法可以添加一组IntSet，比如s.AddAll(1,2,3)。

## Exercise 6.3
(*IntSet).UnionWithcomputes the union of two sets using |, the word-parallel bitwise OR operator. Implement methods for IntersectWith, DifferenceWith, and Symmetric Difference for the corresponding set operations. (The symmetric difference of two sets contains the elements present in one set or the other but not both.)

(*IntSet).UnionWith会用|操作符计算两个集合的并集，我们再为IntSet实现另外的几个函数IntersectWith（交集：元素在A集合B集合均出现），DifferenceWith（差集：元素出现在A集合，未出现在B集合），SymmetricDifference（并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A）。

## Exercise 6.4
Add a method Elems that returns a slice containing the elements of the set, suitable for iterating over with a range loop.

实现一个Elems方法，返回集合中的所有元素，用于做一些range之类的遍历操作。

## Exercise 6.5
The type of each word used by IntSet is uint64, but 64-bit arithmetic may be inefficient on a 32-bit platform. Modify the program to use the uint type, which is the most efficient unsigned integer type for the platform. Instead of dividing by 64, define a constant holding the effective size of uint in bits, 32 or 64. You can use the perhaps too-clever expression 32 << (^uint(0) >> 63) for this purpose.

我们这章定义的IntSet里的每个字都是用的uint64类型，但是64位的数值可能在32位的平台上不高效。修改程序，使其使用uint类型，这种类型对于32位平台来说更合适。当然了，这里我们可以不用简单粗暴地除64，可以定义一个常量来决定是用32还是64，这里你可能会用到平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)

# Chapter 7 Interfaces
## Exercise 7.1
Using the ideas from ByteCounter, implement counters for words and for lines. You will find bufio.ScanWordsuseful.

使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。

## Exercise 7.2
Write a function CountingWriter with the signature below that, given an io.Writer, returns a newWriter that wraps the original, and a pointer to an int64 variable that at any moment contains the number of bytes written to the newWriter.

写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，返回一个把原来的Writer封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针。

## Exercise 7.3
Write a String method for the *tree type in gopl.io/ch4/treesort(§4.4) that reveals the sequence of values in the tree.

为在gopl.io/ch4/treesort（§4.4）中的*tree类型实现一个String方法去展示tree类型的值序列。

## Exercise 7.4
The strings.NewReader function returns a value that satisfies the io.Reader interface (and others) by reading from its argument, a string. Implement a simple version of NewReader yourself, and use it to make the HTML parser (§5.2) take input from a string.

strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）。实现一个简单版本的NewReader，用它来构造一个接收字符串输入的HTML解析器（§5.2）

## Exercise 7.5
The LimitReader function in the io package accepts an io.Reader r and a number of bytesn, and returns another Reader that reads from r but reports an end-of-file condition after n bytes. Implement it.

io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。实现这个LimitReader函数：

## Exercise 7.6
Add support for Kelvin temperatures to tempflag.

对tempFlag加入支持开尔文温度。

## Exercise 7.7
Explain why the help message contains °C when the default value of 20.0 does not.

解释为什么帮助信息在它的默认值是20.0没有包含°C的情况下输出了°C。

## Exercise 7.8
Many GUIs provide a table widget with a stateful multi-tier sort: the primary sort key is the most recently clicked column head, the secondary sort key is the second-most recently clicked column head, and so on. Define an implementation of sort.Interface for use by such a table. Compare that approach with repeated sorting using sort.Stable.

很多图形界面提供了一个有状态的多重排序表格插件：主要的排序键是最近一次点击过列头的列，第二个排序键是第二最近点击过列头的列，等等。定义一个sort.Interface的实现用在这样的表格中。比较这个实现方式和重复使用sort.Stable来排序的方式。

## Exercise 7.9
Use the html/template package (§4.6) to replaceprintTrackswith a function that displays the tracks as an HTML table. Use the solution to the previous exercise to arrange that each click on a column head makes an HTTP request to sort the table.

使用html/template包（§4.6）替代printTracks将tracks展示成一个HTML表格。将这个解决方案用在前一个练习中，让每次点击一个列的头部产生一个HTTP请求来排序这个表格。

## Exercise 7.10
Thesort.Interfacetype can be adapted to other uses. Write a functionIsPalindrome(s sort.Interface) boolthat reports whether the sequencesis a palindrome, in other words, reversing the sequence would not change it. Assume that the elements at indicesiandjare equal if!s.Less(i, j) && !s.Less(j, i).     

sort.Interface类型也可以适用在其它地方。编写一个IsPalindrome(s sort.Interface) bool函数表明序列s是否是回文序列，换句话说反向排序不会改变这个序列。假设如果!s.Less(i, j) && !s.Less(j, i)则索引i和j上的元素相等。

## Exercise 7.11
Add additional handlers so that clients can create, read, update, and delete database entries. For example, a request of the form/update?item=socks&price=6will update the price of an item in the inventory and report an error if the item does not exist or if the price is invalid. (Warning: this change introduces concurrent variable updates.)

增加额外的handler让客户端可以创建，读取，更新和删除数据库记录。例如，一个形如 /update?item=socks&price=6 的请求会更新库存清单里一个货品的价格并且当这个货品不存在或价格无效时返回一个错误值。（注意：这个修改会引入变量同时更新的问题）

## Exercise 7.12
Change the handler for/listto print its output as an HTML table, not text. You may find thehtml/templatepackage (§4.6) useful.

修改/list的handler让它把输出打印成一个HTML的表格而不是文本。html/template包（§4.6）可能会对你有帮助。

## Exercise 7.13
Add aStringmethod toExprto pretty-print the syntax tree. Check that the results, when parsed again, yield an equivalent tree.

为Expr增加一个String方法来打印美观的语法树。当再一次解析的时候，检查它的结果是否生成相同的语法树。

## Exercise 7.14
Define a new concrete type that satisfies theExprinterface and provides a new operation such as computing the minimum value of its operands. Since theParsefunction does not create instances of this new type, to use it you will need to construct a syntax tree directly (or extend the parser).

定义一个新的满足Expr接口的具体类型并且提供一个新的操作例如对它运算单元中的最小值的计算。因为Parse函数不会创建这个新类型的实例，为了使用它你可能需要直接构造一个语法树（或者继承parser接口）。

## Exercise 7.15
Write a program that reads a single expression from the standard input, prompts the user to provide values for any variables, then evaluates the expression in the resulting environment. Handle all errors gracefully.

编写一个从标准输入中读取一个单一表达式的程序，用户及时地提供对于任意变量的值，然后在结果环境变量中计算表达式的值。优雅的处理所有遇到的错误。

## Exercise 7.16
Write a web-based calculator program.

编写一个基于web的计算器程序。

## Exercise 7.17
Extendxmlselectso that elements may be selected not just by name, but by their attributes too, in the manner of CSS, so that, for instance, an element like`<div id="page" class="wide">`could be selected by a matchingidorclassas well as its name.

扩展xmlselect程序以便让元素不仅可以通过名称选择，也可以通过它们CSS风格的属性进行选择。例如一个像这样

## Exercise 7.18
Using the token-based decoder API, write a program that will read an arbitrary XML document and construct a tree of generic nodes that represents it. Nodes are of two kinds:CharDatanodes represent text strings, andElementnodes represent named elements and their attributes. Each element node has a slice of child nodes. You may find the following declarations helpful.

使用基于标记的解码API，编写一个可以读取任意XML文档并构造这个文档所代表的通用节点树的程序。节点有两种类型：CharData节点表示文本字符串，和 Element节点表示被命名的元素和它们的属性。每一个元素节点有一个子节点的切片。

# Chapter 8 Goroutines and Channels
## Exercise 8.1
Modifyclock2to accept a port number, and write a program,clockwall, that acts as a client of several clock servers at once, reading the times from each one and displaying the results in a table, akin to the wall of clocks seen in some business offices. If you have access to geographically distributed computers, run instances remotely; otherwise run local instances on different ports with fake time zones.

修改clock2来支持传入参数作为端口号，然后写一个clockwall的程序，这个程序可以同时与多个clock服务器通信，从多个服务器中读取时间，并且在一个表格中一次显示所有服务器传回的结果，类似于你在某些办公室里看到的时钟墙。如果你有地理学上分布式的服务器可以用的话，让这些服务器跑在不同的机器上面；或者在同一台机器上跑多个不同的实例，这些实例监听不同的端口，假装自己在不同的时区。像下面这样：

## Exercise 8.2
Implement a concurrent File Transfer Protocol (FTP) server. The server should interpret commands from each client such ascdto change directory,lsto list a directory,getto send the contents of a file, andcloseto close the connection. You can use the standardftpcommand as the client, or write your own.

实现一个并发FTP服务器。服务器应该解析客户端发来的一些命令，比如cd命令来切换目录，ls来列出目录内文件，get和send来传输文件，close来关闭连接。你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个。

## Exercise 8.3
In netcat3, the interface value conn has the concrete type *net.TCPConn, which represents a TCP connection. A TCP conne ction consists of two halves that may be closed independently using its CloseRead and CloseWrite methods. Modify the main goroutine of netcat3 to close only the write half of the connection so that the program will continue to print the final ech oes from the reverb1 server even after the stand ard input has been closed. (Doing this for the reverb2 server is harder ; see Exercise 8.4.)

在netcat3例子中，conn虽然是一个interface类型的值，但是其底层真实类型是*net.TCPConn，代表一个TCP连接。一个TCP连接有读和写两个部分，可以使用CloseRead和CloseWrite方法分别关闭它们。修改netcat3的主goroutine代码，只关闭网络连接中写的部分，这样的话后台goroutine可以在标准输入被关闭后继续打印从reverb1服务器传回的数据。（要在reverb2服务器也完成同样的功能是比较困难的；参考练习 8.4。）

## Exercise 8.4
Modify the reverb2 server to use async.WaitGroupper connection to count the number of activeechogoroutines. When it falls to zero, close the write half of the TCP connection as described in Exercise 8.3. Verify that your modifiednetcat3client from that exercise waits for the final echoes of multiple concurrent shouts, even after the standard input has been closed.

修改reverb2服务器，在每一个连接中使用sync.WaitGroup来计数活跃的echo goroutine。当计数减为零时，关闭TCP连接的写入，像练习8.3中一样。验证一下你的修改版netcat3客户端会一直等待所有的并发“喊叫”完成，即使是在标准输入流已经关闭的情况下。

## Exercise 8.5
Take an existing CPU-bound sequential program, such as the Mandelbrot program of Section 3.3 or the 3-D surface computation of Section 3.2, and execute its main loop in parallel using channels for communication. How much faster does it run on a multiprocessor machine? What is the optimal number of goroutines to use?

使用一个已有的CPU绑定的顺序程序，比如在3.3节中我们写的Mandelbrot程序或者3.2节中的3-D surface计算程序，并将他们的主循环改为并发形式，使用channel来进行通信。在多核计算机上这个程序得到了多少速度上的改进？使用多少个goroutine是最合适的呢？

## Exercise 8.6
Add depth-limiting to the concurrent crawler. That is, if the user sets-depth=3, then only URLs reachable by at most three links will be fetched.

为并发爬虫增加深度限制。也就是说，如果用户设置了depth=3，那么只有从首页跳转三次以内能够跳到的页面才能被抓取到。

## Exercise 8.7
Write a concurrent program that creates a local mirror of a web site, fetching each reachable page and writing it to a directory on the local disk. Only pages within the original domain (for instance,golang.org) should be fetched. URLs within mirrored pages should be altered as needed so that they refer to the mirrored page, not the original.

完成一个并发程序来创建一个线上网站的本地镜像，把该站点的所有可达的页面都抓取到本地硬盘。为了省事，我们这里可以只取出现在该域下的所有页面（比如golang.org开头，译注：外链的应该就不算了。）当然了，出现在页面里的链接你也需要进行一些处理，使其能够在你的镜像站点上进行跳转，而不是指向原始的链接。

## Exercise 8.8
Using a select statement, add a timeout to the echo server from Section 8.3 so that it disconnects any client that shouts nothing within 10 seconds.

使用select来改造8.3节中的echo服务器，为其增加超时，这样服务器可以在客户端10秒中没有任何喊话时自动断开连接。

## Exercise 8.9
Write a version ofduthat computes and periodically displays separate totals for each of therootdirectories.

编写一个du工具，每隔一段时间将root目录下的目录大小计算并显示出来。

## Exercise 8.10
HTTP requests may be cancelled by closing the optionalCancelchannel in thehttp.Requeststruct. Modify the web crawler of Section 8.6 to support cancellation.

HTTP请求可能会因http.Request结构体中Cancel channel的关闭而取消。修改8.6节中的web crawler来支持取消http请求。（提示：http.Get并没有提供方便地定制一个请求的方法。你可以用http.NewRequest来取而代之，设置它的Cancel字段，然后用http.DefaultClient.Do(req)来进行这个http请求。）

## Exercise 8.11
Following the approach ofmirroredQueryin Section 8.4.4, implement a variant offetchthat requests several URLs concurrently. As soon as the first response arrives, cancel the other requests.

紧接着8.4.4中的mirroredQuery流程，实现一个并发请求url的fetch的变种。当第一个请求返回时，直接取消其它的请求。

## Exercise 8.12
Make the broadcaster announce the current set of clients to each new arrival. This requires that theclientsset and theenteringandleavingchannels record the client name too.

使broadcaster能够将arrival事件通知当前所有的客户端。这需要你在clients集合中，以及entering和leaving的channel中记录客户端的名字。

## Exercise 8.13
Make the chat server disconnect idle clients, such as those that have sent no messages in the last five minutes. Hint: callingconn.Close()in another goroutine unblocks activeReadcalls such as the one done byinput.Scan().

使聊天服务器能够断开空闲的客户端连接，比如最近五分钟之后没有发送任何消息的那些客户端。提示：可以在其它goroutine中调用conn.Close()来解除Read调用，就像input.Scanner()所做的那样。

## Exercise 8.14
Change the chat server’s network protocol so that each client provides its name on entering. Use that name instead of the network address when prefixing each message with its sender’s identity.

修改聊天服务器的网络协议，这样每一个客户端就可以在entering时提供他们的名字。将消息前缀由之前的网络地址改为这个名字。

## Exercise 8.15
Failure of any client program to read data in a timely manner ultimately causes all clients to get stuck. Modify the broadcaster to skip a message rather than wait if a client writer is not ready to accept it. Alternatively, add buffering to each client’s outgoing message channel so that most messages are not dropped; the broadcaster should use a non-blocking send to this channel.

如果一个客户端没有及时地读取数据可能会导致所有的客户端被阻塞。修改broadcaster来跳过一条消息，而不是等待这个客户端一直到其准备好读写。或者为每一个客户端的消息发送channel建立缓冲区，这样大部分的消息便不会被丢掉；broadcaster应该用一个非阻塞的send向这个channel中发消息。

# Chapter 9 Concurrency with Shared Variables
## Exercise 9.1
Add a functionWithdraw(amount int) boolto thegopl.io/ch9/bank1program. The result should indicate whether the transaction succeeded or failed due to insufficient funds. The message sent to the monitor goroutine must contain both the amount to withdraw and a new channel over which the monitor goroutine can send the boolean result back toWithdraw.

给gopl.io/ch9/bank1程序添加一个Withdraw(amount int)取款函数。其返回结果应该要表明事务是成功了还是因为没有足够资金失败了。这条消息会被发送给monitor的goroutine，且消息需要包含取款的额度和一个新的channel，这个新channel会被monitor goroutine来把boolean结果发回给Withdraw。

## Exercise 9.2
Rewrite thePopCountexample from Section 2.6.2 so that it initializes the lookup table usingsync.Oncethe first time it is needed. (Realistically, the cost of synchronization would be prohibitive for a small and highly optimized function likePopCount.)

重写2.6.2节中的PopCount的例子，使用sync.Once，只在第一次需要用到的时候进行初始化。（虽然实际上，对PopCount这样很小且高度优化的函数进行同步可能代价没法接受。）

## Exercise 9.3
Extend theFunctype and the(*Memo).Getmethod so that callers may provide an optionaldonechannel through which they can cancel the operation (§8.9). The results of a cancelledFunccall should not be cached.

扩展Func类型和(*Memo).Get方法，支持调用方提供一个可选的done channel，使其具备通过该channel来取消整个操作的能力（§8.9）。一个被取消了的Func的调用结果不应该被缓存。

## Exercise 9.4
Construct a pipeline that connects an arbitrary number of goroutines with channels. What is the maximum number of pipeline stages you can create without running out of memory? How long does a value take to transit the entire pipeline?

创建一个流水线程序，支持用channel连接任意数量的goroutine，在跑爆内存之前，可以创建多少流水线阶段？一个变量通过整个流水线需要用多久？（这个练习题翻译不是很确定）

## Exercise 9.5
Write a program with two goroutines that send messages back and forth over two unbuffered channels in ping-pong fashion. How many communications per second can the program sustain?

写一个有两个goroutine的程序，两个goroutine会向两个无buffer channel反复地发送ping-pong消息。这样的程序每秒可以支持多少次通信？

## Exercise 9.6
Measure how the performance of a compute-bound parallel program (see Exercise 8.5) varies withGOMAXPROCS. What is the optimal value on your computer? How many CPUs does your computer have?

测试一下计算密集型的并发程序（练习8.5那样的）会被GOMAXPROCS怎样影响到。在你的电脑上最佳的值是多少？你的电脑CPU有多少个核心？

# Chapter 10 Packages and the Go Tool
## Exercise 10.1
Extend thejpegprogram so that it converts any supported input format to any output format, usingimage.Decodeto detect the input format and a flag to select the output format.

扩展jpeg程序，以支持任意图像格式之间的相互转换，使用image.Decode检测支持的格式类型，然后通过flag命令行标志参数选择输出的格式。

## Exercise 10.2
Define a generic archive file-reading function capable of reading ZIP files (archive/zip) and POSIX tar files (archive/tar). Use a registration mechanism similar to the one described above so that support for each file format can be plugged in using blank imports.

设计一个通用的压缩文件读取框架，用来读取ZIP（archive/zip）和POSIX tar（archive/tar）格式压缩的文档。使用类似上面的注册技术来扩展支持不同的压缩格式，然后根据需要通过匿名导入选择导入要支持的压缩格式的驱动包。

## Exercise 10.3
Usingfetch http://gopl.io/ch1/helloworld?go-get=1, find out which service hosts the code samples for this book. (HTTP requests from go get include the go-get parameter so that servers can distinguish them from ordinary browser requests.)

从 http://gopl.io/ch1/helloworld?go-get=1 获取内容，查看本书的代码的真实托管的网址（go get请求HTML页面时包含了go-get参数，以区别普通的浏览器请求）。

## Exercise 10.4
Construct a tool that reports the set of all packages in the workspace that transitively depend on the packages specified by the arguments. Hint: you will need to run go listtwice, once for the initial packages and once for all packages. You may want to parse its JSON output using theencoding/jsonpackage (§4.5).

创建一个工具，根据命令行指定的参数，报告工作区所有依赖包指定的其它包集合。提示：你需要运行go list命令两次，一次用于初始化包，一次用于所有包。你可能需要用encoding/json（§4.5）包来分析输出的JSON格式的信息。

# Chapter 11 Testing
## Exercise 11.1
Write tests for thecharcountprogram in Section 4.3.

为4.3节中的charcount程序编写测试。

## Exercise 11.2
Write a set of tests forIntSet(§6.5) that checks that its behavior after each operation is equivalent to a set based on built-in maps. Save your implementation for benchmarking in Exercise 11.7.

为（§6.5）的IntSet编写一组测试，用于检查每个操作后的行为和基于内置map的集合等价，后面Exercise11.7将会用到。

## Exercise 11.3
TestRandomPalindromesonly tests palindromes. Write a randomized test that generates and verifies non-palindromes.

TestRandomPalindromes测试函数只测试了回文字符串。编写新的随机测试生成器，用于测试随机生成的非回文字符串。

## Exercise 11.4
ModifyrandomPalindrometo exerciseIsPalindrome’s handling of punctuation and spaces.

修改randomPalindrome函数，以探索IsPalindrome是否对标点和空格做了正确处理。

## Exercise 11.5
ExtendTestSplitto use a table of inputs and expected outputs.

用表格驱动的技术扩展TestSplit测试，并打印期望的输出结果。

## Exercise 11.6
Write benchmarks to compare thePopCountimplementation in Section 2.6.2 with your solutions to Exercise 2.4 and Exercise 2.5. At what point does the table-based approach break even?

为2.6.2节的练习2.4和练习2.5的PopCount函数编写基准测试。看看基于表格算法在不同情况下对提升性能会有多大帮助。

## Exercise 11.7
Write benchmarks forAdd,UnionWith, and other methods of*IntSet(§6.5) using large pseudo-random inputs. How fast can you make these methods run? How does the choice of word size affect performance? How fast isIntSetcompared to a set implementation based on the built-in map type?

为*IntSet（§6.5）的Add、UnionWith和其他方法编写基准测试，使用大量随机输入。你可以让这些方法跑多快？选择字的大小对于性能的影响如何？IntSet和基于内建map的实现相比有多快？

# Chapter 12 Reflection
## Exercise 12.1
Extend Display so that it can display maps whose keys are structs or arrays.

扩展Display函数，使它可以显示包含以结构体或数组作为map的key类型的值。

## Exercise 12.2
Make display safe to use on cyclic data structures by bounding the number of steps it takes before abandoning the recursion. (In Section 13.3, we’ll see another way to detect cycles.)

增强display函数的稳健性，通过记录边界的步数来确保在超出一定限制后放弃递归。（在13.3节，我们会看到另一种探测数据结构是否存在环的技术。）

## Exercise 12.3
Implement the missing cases of the encode function. Encode booleans as t and nil, floating-point numbers using Go’s not ation, and complex numbers like 1+2i as #C(1.0 2.0). Interfaces can be encoded as a pair of a type name and a value, for ins tance ("[]int" (1 2 3)), but beware that this not ation is ambiguous: the reflect.Type.String method may retur n the same string for different types.

实现encode函数缺少的分支。将布尔类型编码为t和nil，浮点数编码为Go语言的格式，复数1+2i编码为#C(1.0 2.0)格式。接口编码为类型名和值对，例如（"[]int" (1 2 3)），但是这个形式可能会造成歧义：reflect.Type.String方法对于不同的类型可能返回相同的结果。

## Exercise 12.4
Modify encode to pretty-print the S-expression in the sty le shown above .

修改encode函数，以上面的格式化形式输出S表达式。

## Exercise 12.5
Adapt encode to emit JSON ins tead of S-expressions. Test your encoder using the stand ard decoder, json.Unmarshal.

修改encode函数，用JSON格式代替S表达式格式。然后使用标准库提供的json.Unmarshal解码器来验证函数是正确的。

## Exercise 12.6
Adapt encode so that, as an optimization, it does not encode a field whose value is the zero value of its type.

修改encode，作为一个优化，忽略对是零值对象的编码。

## Exercise 12.7
Create a streaming API for the S-expression decoder, fol low ing the sty le of json.Decoder (§4.5).

创建一个基于流式的API，用于S表达式的解码，和json.Decoder(§4.5)函数功能类似。

## Exercise 12.8
The sexpr.Unmarshal function, like json.Marshal, requires the complete input in a byte slice before it can beg in de coding. Define a sexpr.Decoder type that, like json.Decoder, allows a sequence of values to be decoded from an io.Reader. Change sexpr.Unmarshal to use this new type.

sexpr.Unmarshal函数和json.Unmarshal一样，都要求在解码前输入完整的字节slice。定义一个和json.Decoder类似的sexpr.Decoder类型，支持从一个io.Reader流解码。修改sexpr.Unmarshal函数，使用这个新的类型实现。

## Exercise 12.9
Write a token-b ased API for decoding S-expressions, following the style of xml.Decoder (§7.14). You will need five types of tokens: Symbol, String, Int, StartList, and EndList.

编写一个基于标记的API用于解码S表达式，参考xml.Decoder（7.14）的风格。你将需要五种类型的标记：Symbol、String、Int、StartList和EndList。

## Exercise 12.10
Extend sexpr.Unmarshal to handle the boole ans, floating-p oint numbers, and int erfaces encoded by your solut ion to Exercis e 12.3. (Hint: to decode int erfaces, you will need a mapping fro m the name of each supported type to its reflect.Type.)

扩展sexpr.Unmarshal函数，支持布尔型、浮点数和interface类型的解码，使用 练习 12.3： 的方案。（提示：要解码接口，你需要将name映射到每个支持类型的reflect.Type。）

## Exercise 12.11
Write the corresponding Pack function. Given a struct value, Pack should return a URL incorporating the parameter values from the struct.

编写相应的Pack函数，给定一个结构体值，Pack函数将返回合并了所有结构体成员和值的URL。

## Exercise 12.12
Extend the field tag not ation to express parameter validity requirements. For example, a string might need to be a valid email address or credit-card number, and an integer might need to be a valid US ZIP code. Modify Unpack to check these requirements.

扩展成员标签以表示一个请求参数的有效值规则。例如，一个字符串可以是有效的email地址或一个信用卡号码，还有一个整数可能需要是有效的邮政编码。修改Unpack函数以检查这些规则。

## Exercise 12.13
Modify the S-expression encoder (§12.4) and decoder (§12.6) so that they honor the sexpr:"..." field tag in a similar manner to encoding/json (§4.5).

修改S表达式的编码器（§12.4）和解码器（§12.6），采用和encoding/json包（§4.5）类似的方式使用成员标签中的sexpr:"..."字串。

# Chapter 13 Low-Level Programming
## Exercise 13.1
Define a deep comparison function that considers numbers (of any typ e) equal if they differ by less than one part in a billion.

定义一个深比较函数，对于十亿以内的数字比较，忽略类型差异。

## Exercise 13.2
Write a function that reports whether its argument is a cyclic data structure.

编写一个函数，报告其参数是否为循环数据结构。

## Exercise 13.3
Use sync.Mutex to make bzip2.writer safe for concurrent use by multiple goroutines.

使用sync.Mutex以保证bzip2.writer在多个goroutines中被并发调用是安全的。

## Exercise 13.4
Depending on C libraries has its drawbacks. Provide an alternative pure-Go implementation of bzip.NewWriter that uses the os/exec package to run /bin/bzip2 as a subprocess.

因为C库依赖的限制。 使用os/exec包启动/bin/bzip2命令作为一个子进程，提供一个纯Go的bzip.NewWriter的替代实现（译注：虽然是纯Go实现，但是运行时将依赖/bin/bzip2命令，其他操作系统可能无法运行）。
