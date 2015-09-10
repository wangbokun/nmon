
#nmon 进程守护
#史上最轻量级进程守护程序,仅仅2.1k程序

- 以下内容周期性读取，配置文件，启动相应进程
-  配置文件字段 server,user, key, cmd含义

-----
  server服务器的IP，user是程序需要用哪个用户启动,Type需要监控进程的类型（jps或普通程序）， key关键字（需唯一）， cmd启动语句（不要以&结束）
  
 cmd启动命令，不能包含文件重定向，本程序会自动将其放入后台，丢弃所有的标准输出和标准错误输出
 
 

##说明
 - nmon -v
 - nmon  -h
##使用方法：
  1. crontab  配置
    10/* * * * *   sh $dir/nmon &> /tmp/nmon.tmp
  2. mkdir /etc/nmon && cp nmon.conf.example  /etc/nmon/nmon.conf
