当我们定义了一个对象时，一般会创建一个方法方便外部初始化一个实例。

在 C 族语言中, 可以使用不同数量的参数提供相同函数的多个版本; 
在像 PHP 这样的语言中, 可以给参数一个默认值，并在调用方法时忽略它们.
但是在 Golang 中, 这两种方式你哪个也用不了.

下面简绍一种优雅的方案：Functional Options  
1.在 `file/options.go` 中定义了 `Options` 的 struct。  
2.在 `file/new.go` 中定义了 `New(filepath string, setters ...Option)` 方法  
3.`main.go` 中有示例代码  
  
注意 `file/new.go` 内部实现， 首先使用默认值创建一个 `Options`，之后
使用 setters 循环修改设置项。  

```go
package file

func New(filepath string, setters ...Option) (*os.File, error) {
	// Default Options
	args := &Options{
		UID:         os.Getuid(),
		GID:         os.Getgid(),
		Contents:    "",
		Permissions: 0666,
		Flags:       os.O_CREATE | os.O_EXCL | os.O_WRONLY,
	}

	for _, setter := range setters {
		setter(args)
	}

	...
}
```

通过 `Functional Options` 的接口设计，即使在之后添加设置项，我们只需要修改很少的代码即可。