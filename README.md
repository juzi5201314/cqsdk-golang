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
>前往go官网或者go中文网下载go的weindows 32位源码包(编译环境是什么系统就下什么系统的，这里以编译环境是windows10 x64的为例)。  

>前往[下载](http://tdm-gcc.tdragon.net/download)tmd-gcc 32位版，并安装。 

>解压go源码，如当前环境本来就存在64位的go环境，可以如我这样将32位的解压在比如C:\Go32。  

>将tmd-gcc安装目录(一般是在C:\TMD-GCC-32)添加到环境变量(其他系统同理，下载个32位的gcc)。  

1，打开命令行，cd到插件源码目录(以[demo](https://github.com/juzi5201314/cqsdk-golang/tree/master/demo)为例)。  
2，执行
```shell
> C:\Go32\bin\go.exe build -buildmode=c-shared -o cn.you.test.dll
```
3, 将cn.you.test.dll动态库与cn.you.test.json放到酷q的app目录，打包，完成。

# sdk
提供  
cqsdk-golang.Gb18030(utf8 string) string  
此方法将utf8编码的字符串转换成酷q需要的gb18030编码。sdk内置api已自动转换，调用api时直接填入utf8字符串即可，此方法使用较少。  

cqsdk-golang.Utf8(gb18030 string) string 
此方法将gb18030编码的字符串解码成utf8编码的字符串。由于酷q事件传入的是gb18030编码的，所以接收const char \*的时候除了使用C.GoString来将C类型转换成go的字符串之后，还需要gb18030解码.  
具体操作请看[demo](https://github.com/juzi5201314/cqsdk-golang/tree/master/demo)
