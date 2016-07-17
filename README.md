# dirhash
This program list all file at given path.
record all file sha1 and size into a file.
support skip given path and file/support regular expression.

usage：
go run dirhash.go -D  set working path   -P set skip cata or file regular  -L  set log file


test usage:
go run read.go  "log file path"

This test check if the size and sha1 of file  same with logfile that dishash give out. And print the result "OK" or "failed".

 
