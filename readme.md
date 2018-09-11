webs 

GO语言写的最简单的静态文件服务器，方便开发调试用

使用方法：


直接运行：
===========

1. 在上方 Release 板块直接下载编译好的版本 ( 下载 webs.exe )

（A） 复制到静态文件目录,直接运行 webs.exe

（B） 以系统工具运行（Win系统为例）：

    将webs.exe放在 c:/windows/system32 目录下，

    在需要静态文件服务的目录，打开cmd，运行webs即可


（C） 打开浏览器，访问 http://localhost:9090/


自己编译（可选）
===============

运行：

go run webs.go

编译：

go build webs.go


