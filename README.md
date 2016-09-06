# Personal-Dictionry

## Why create this Project?

我发现在如今的互联网时代，很多人对于一个词的认知需要的往往不是一个很官方的解释，而是通过互联网查阅以及阅读一定量的文章／博客之后，从而得出带有个人色彩的认知与理解。
这不是一个简简单单的 `百度百科` 或者 `官方词典` 能够解决的问题，所以我想做一个属于个人的词典的 Web 应用，在里面你可以创造属于自己的词典。

## How

- DB: Redis or Support Map[string]string DB.
    + Redis : Hash[Key(UserID)[Map[Key(Dictionry-Key)][Value]]]
- Back-End: Of course Golang.
- Font-End: Vue.js or React do like a List.

## TODO

- [弹性云](http://www.shujuba.net/help/help.asp)
- [MQ](http://baike.baidu.com/link?url=_oODT10riIrKTGHD3evOZNNg5SOHpxTf_KKkstiavuHmXUY9Fkvz0_nDH599K49yxEll1FVtSXeclScPJ76fVHY49WfPxeHwfoju5sZqNc_)
- [面包屑](http://baike.baidu.com/link?url=rM6FBTTTRKhNcn5JgAjD2ctjbew9xXM6JGh4MPxHlwyWz8oOLjJy_ZfKleYzHs93TTY_ss6Z5P7QUUPr0wzbtq)
- [锚点链接]()
- [mysql时区问题]()
    + sql-base.go --> root:passwd@tcp(localhost:3306)/dbname?charset=utf8& `loc=Local`
- [结伴发布]()
    + 指的是开发与运维在一起上线应用（主要因为开发环境和生产环境不一致而产生的），效率低下。
- [Zabbix](http://baike.baidu.com/link?url=Hy2yXjgxcG55hPGLdRIdpz1M3Ck28Bo8D2jh4H0q6Tpdj8j8X5Qot4L548yR0fstqU6IcCPspa-c3PaHXhkCYq)
    + zabbix（音同 zæbix）是一个基于WEB界面的提供分布式系统监视以及网络监视功能的企业级的开源解决方案。


## Dependence

```
$ go get github.com/russross/blackfriday
```

## Usage

```
// Please set runmode in app.conf first when you use it in Server!
$ bee run # if you syncdb already.
$ ./Personal-Dictionry -syncdb # if you run this app at first time.
```

## Problem

- 不支持 markdown 的多行代码
- 不支持空格