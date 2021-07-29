webs
==== 

GO语言写的最简单的静态文件服务器，方便开发调试用

使用方法：

运行方式：
--------

方式①： 在某一目录下运行 webs , 默认监听 9090 端口，打开  http://localhost:9090/

方式②： 设置端口号，运行 webs -p 7070 ， 打开 http://localhost:7070/

方式③： 设置目录（支持当前目录下的下级目录）。

例如设置托管webs所在当前目录下的dist子文件夹，运行webs -d ./dist  

* 查看版本 webs -v

直接运行：
---------

1. 在上方 Release 板块直接下载编译好的版本 ( 下载 webs.exe )

（A） 复制到静态文件目录,直接运行 webs.exe

（B） 以系统工具运行（Win系统为例）：

    将webs.exe放在 c:/windows/system32 目录下，

    在需要静态文件服务的目录，打开cmd，运行webs即可


（C） 打开浏览器，访问 http://localhost:9090/


自己编译（可选）
---------------

运行：

go run webs.go

编译：

go build webs.go


