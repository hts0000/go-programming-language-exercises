练习 7.7

解释为什么帮助信息在它的默认值是20.0没有包含°C的情况下输出了°C

在`tempconv/conv.go`文件中，`flag.CommandLine.Var(&f, name, usage)`函数的类型为`func (f *FlagSet) Var(value Value, name string, usage string)`，其中Value是一个接口类型
```golang
type Value interface {
	String() string
	Set(string) error
}
```
对于我们定义的`celsiusFlag`类型，实现了Set()方法和String()方法（从Celsius类型中组合得到），也就是说我们实现了Value这个接口类型，因此可以传入`&f`作为参数。

同时根据name（从命令行获取参数的标志），usage（使用说明），构建了一个存储值为`celsiusFlag`类型的flag。

`Flag`类型和`func (f *FlagSet) Var(value Value, name string, usage string)`源码如下
```golang
type Flag struct {
	Name     string // name as it appears on command line
	Usage    string // help message
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message
}

func (f *FlagSet) Var(value Value, name string, usage string) {
	// Remember the default value as a string; it won't change.
	flag := &Flag{name, usage, value, value.String()} // here
	...
}
```
可以看到在构建Flag时，DefValue使用了value.String()方法，也就是Celsius的String方法
```golang
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
```

因此在帮助信息中会输出带有°C的字符