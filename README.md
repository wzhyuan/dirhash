# dirhash
This program list all file at given path.
record all file sha1 and size into a file.
support skip given path and file/support regular expression.

usage：
go run dirhash.go -D  set working path   -P set skip cata or file regular  -L  set log file


test usage:
go run read.go  "log file path"

This test check if the size and sha1 of file  same with logfile that dishash give out. And print the result "OK" or "failed".

 
小工具需求说明

1. 用golang开发，代码放到github上，用github进行问题跟踪
2. 对整个目录下的所有文件进行遍历，获取所有文件的大小和计算文件的sha1哈希值，记录在一个文件里面
3. 建议结果文件格式：每一行一个文件，用逗号隔开，前面是文件名称，后面是哈希值，文件大小
4. 需要可以指定忽略哪些目录、文件，需要支持通配符
5. 代码实现简洁，运行性能高得分高
6. 要求通过测试代码自我证明代码能够可靠运行并正确实现上述功能
