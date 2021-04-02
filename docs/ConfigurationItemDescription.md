# Configuration File
### Introduction
The default configuration file of Doraemon is app.conf,which is located at cmd/alert-gateway/conf.

### Basic Configuration
*[Reference Beego configuration](https://beego.me/mvc/controller/config.md)*
- appname: The name of App
- httpport: Set the port the app listens on.By default this is 8080
- runmode: dev/prod
- copyrequestbody: Toggle copying of raw request body in context.By default this is false except for GET, HEAD or file uploading.
- autorender: Enable auto render. By default this is True. This value should be set to false for API applications, as there is no need to render templates
- EnableDocs: Enable Docs. By default this is False
- sessionon: Enable session. By default this is False

### DataBase Configuration
- DBTns: the Tns of database,for example: tcp(127.0.0.1:3306)
- DBName: the name of database
- DBUser: the username of database
- DBPasswd: the password of database
- DBLoc: the location,for example:Asia%2FShanghai

### Related Interface
- SmsUrl: The endpoint of sms
- LanxinUrl: The endpoint of LANXIN
- CallUrl: The endpoint of CALL
- DutyGroupUrl: The interface of duty group from which we can get the user list of this duty group
- BrokenUrl: The interface of getting the list of fault machines

### Domain Name of Web-UI
- WebUrl: The domain name of Web-UI,for example: "http://www.doraemon.com:8080"
