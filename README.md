# cqsdk-golang
酷q的go版本sdk

## 注意
此sdk只支持在docker下运行的酷q，win环境下运行的酷q调用api的时候会造成酷q闪退！  
编译环境不限。

# 使用
```shell
> go get -u github.com/juzi5201314/cqsdk-golang
```
由于酷q插件必须是x86的，而cgo不支持跨平台编译。  
所以如果编译环境并非windows x86，请看下面，否则跳过此段。  
>前往go官网或者go中文网下载go的32位源码包(编译环境是什么系统就下什么系统的，这里以编译环境是windows10 x64的为例)。  
>前往[下载](http://tdm-gcc.tdragon.net/download)tmd-gcc 32位版，并安装。  
>解压go源码，如当前环境本来就存在64位的go环境，可以如我这样将32位的解压在比如C:\Go32。  
>将tmd-gcc安装目录(一般是在C:\TMD-GCC-32)添加到环境变量。  
1，打开命令行，cd到插件源码目录(以[demo](https://github.com/juzi5201314/cqsdk-golang/tree/master/demo)为例)。  
2，执行
```shell
> C:\Go32\bin\go.exe build -buildmode=c-shared -o cn.you.test.dll
```
3, 将cn.you.test.dll动态库与cn.you.test.json放到酷q的app目录，打包，完成。
