# qc-gin-use
gin + mysql + redis 


## run
```
go run .

```

## build
```
运行go build得到了一个名叫app的二进制文件，这个在linux上如何部署
1. 将app文件上传到linux服务器的某个目录下，比如/home/user/app。

2. 给app文件赋予可执行权限，使用命令chmod +x /home/user/app。
在后台运行app文件，使用命令nohup /home/user/app &。这样可以让app文件在后台持续运行，即使你退出终端也不会影响。

3.查看app文件的运行日志，使用命令tail -f nohup.out。这样可以实时查看app文件的输出信息，比如错误或调试信息。

```



## go base
[go基础](https://github.com/canwhite/qc-go-use)