# remReplace
Golang编写的简单脚本，实现对文件夹下所有文件中以rem为单位的css属性的值进行批量倍增/倍减。用于修改html font-size值后对以rem为单位的属性值就行批量重新计算，通过在解决font-size设置错误时才会使用。

# 使用方法
1. 通过git clone或者Download zip将本项目下载到你的本地
2. 进入项目目录，通过```go build main.go```编译为可执行文件。
3. 执行可执行文件，格式如```./main -m x /XXX```，执行后即可对/XXX下的所有文件中以rem为单位的css属性的值进行重新计算，计算规则为新值 = 旧值 * x。