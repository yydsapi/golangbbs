# 快速启动

获取并运行 golangbbs:

1. 源代码模式
	
	git clone https://github.com/kdhly/golangbbs.git 
	cd golangbbs  
	go build  
	chmod a+x golangbbs  
	./golangbbs  

	## components:

	go get github.com/Sirupsen/logrus  
	go get github.com/gin-contrib/sessions  
	go get github.com/gin-gonic/gin  
	go get github.com/go-sql-driver/mysql  
	go get github.com/mattn/go-sqlite3  
	go get github.com/quasoft/memstore  
	etc  


2. 可执行文件

	git clone https://github.com/kdhly/golangbbs.git   
	cd golangbbs/dist
	解压缩 golangbbs/dist/dist.rar   

	windows(x64):  
	cd windows  
	复制golangbbs/sqlite.db 和 golangbbs/documents 到这个目录  
	在cmd中以管理员身份执行golangbbs.exe  

	linux(x64):  
	cd linux  
	复制golangbbs/sqlite.db 和 golangbbs/documents 到这个目录    
	chmod a+x golangbbs  
	./golangbbs  
	
	mac(x64):  
	cd mac  
	copy golangbbs/sqlite.db and golangbbs/documents to this folder  
	chmod a+x golangbbs  
	./golangbbs  

	arm:  
	cd arm  
	复制golangbbs/sqlite.db 和 golangbbs/documents 到这个目录    
	chmod a+x golangbbs  
	./golangbbs  

然后打开页面: http://127.0.0.1

# 配置文件: bbs_config_main_i18n_xxx.json

如果只在本机访问，请将IPAddr改为：127.0.0.1；默认数据库是sqlite3,如果您想使用mysql,您可以将golangbbs/dist/mysql目录中的.sql导入到mysql数据库中,将DbType改为"mysql"。

# 结构和功能:
golang + gin + vue + element ui + i18n;  
插件: ckedit + html5player + blueimp Gallery (music and photo manager)  
数据库: mysql or sqlite;  
1. 所有资源本地化,自动切换语言;  
2. 防止脚本注入攻击、文本过滤、防止入侵攻击、XSS攻击;  
3. 开放紧凑的便携式设计;  
4. 本系统功能较为全面和实用，有比较完整的共享功能和附件、多媒体功能，可作为知识库和个人电子记事本及个人媒体中心或简易bbs;  
5. 你可以用它建立一个私人博客;  
6. 菜单和博客类别在线修改;  
7. 用户信息在线修改;  
8. 文章及回复在线修改;  
9. 基本支持各种移动浏览器; 

#### 使用需知:  
1. 5级及以上用户（评分>2999）可发布图片、附件、媒体，默认管理用户名：limon，密码：password;  
2. 可以指定多个管理员，用户 allow > 100 的为管理员;  
3. 您可以在bbs_config_main_i18n.json文件中自己调整一些参数,比如找回密码，必须配置正确的邮件服务器地址，SMTP用户名密码等;  
4. BbsUploadPath或当前目录的documents必须要有读写权限;  
5. 要显示的图像目录在BbsUploadPath+“/Picture/photos”中，默认显示列表中不包含以“my”开头的目录名，但仍可以在URL中显示;   
6. 已知兼容版本为go 1.10或更高版本，gin v1.4或更高版本，数据库：mysql 5.7或更高版本或sqlite 3;  
#### 问题:
1. 目前还存在一些bug和非模块化内容，但并不影响使用;  
2. 旧浏览器不受支持，因为Vue使用的ES6 Promise对象功能是旧浏览器无法模拟的，您可以查找一些插件来与之兼容;  

## 联系作者
mail: 1269866868@qq.com,yydsapi@gmail.com  
QQ群：920788836  
或访问: https://yydsapi.com/library/listitem?page=library&category=bL5tJwTz_r9SiKLWL_bc

## tips:
如果您有好的项目或建议，我们可以帮助您实现它.  
## 特别感谢:
@fhst, @kdhly,以及项目结构和插件中使用的所有功能模块,和其他未列出的功能模块;  

## Give me a star
如果你喜欢或计划使用这个项目，请Give me a star，谢谢！

## Donation
如果本项目让你感觉不错，你可以通过以下链接捐款，以更好地支持本项目或团队的发展: <br /><br />
![10](/static/img/donation/alipay.jpg)   <br /><br /> <br />

![10](/static/img/donation/weixin.jpg)    <br /><br /> <br />

##### 或者https://paypal.me/golangbbs

## Screenshots ：<br /><br />
#### mainpage 
>![11](/static/img/screenshots/mainpage.jpg)  <br /><br />
#### publish 
>![11](/static/img/screenshots/publish.jpg)  <br /><br />
#### show content 1
>![11](/static/img/screenshots/show1.jpg)  <br /><br />
#### show content 2
>![11](/static/img/screenshots/show2.jpg)  <br /><br />
#### display photo 
>![11](/static/img/screenshots/photoshow1.jpg)  <br /><br />
#### manage1 
>![11](/static/img/screenshots/manage1.jpg)  <br /><br />
#### manage2 
>![11](/static/img/screenshots/manage2.jpg)  <br /><br />
