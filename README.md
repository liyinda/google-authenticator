> ### 引言
> 只要是单一的密码方式登录系统，就存在被爆破的风险，更何况用户没有密码复杂度要求的情况下，往往会存在大量弱口令。<br>
双因子认证（2FA）主要是通过两种不同的认证方式登录系统，避免由于密码被泄露导致的安全问题。
google-authenticator需要用户先通过手机下载TOTP工具，比如：阿里云APP中的"虚拟MFA"，用户通过手机的MFA终端扫描google-authenticator中生成的二维码，
在对应的登录系统时会调用google-authenticator的认证API，在完成认证码的验证后方可登录。



# google-authenticator

Google authenticator是基于TOTP(Time-based One-time Password Algorithm)原理实现的双因子认证方案。通过
一致算法保持手机端和服务端相同，并每30秒改变认证码。<br>
该程序一共分为两部分：<br>
1) **服务端**: 为谷歌二维码的管理后台，管理员通过服务管理后台生成二维码，用户通过手机下载APP（TOTP、阿里云APP中的"虚拟MFA"）扫描生成的二维码，在用户的手机端就会有对应的6位数字的动态令牌，在登录系统A时需要输入对应的6位数字。
2) **API端**: 系统A与API端交互，通过请求API端验证用户的6位数字的动态令牌是否正确，如果正确返回ok，如果错误返回error。

## 目录
* [环境](#环境)
* [下载](#下载)
* [编译](#编译)
  * [build binary](#build-binary)
  * [build docker image](#build-docker-image)
* [运行](#运行)
  * [run binary](#run-binary)
  * [run docker image](#run-docker-image)
  * [view google-authenticator](#view google-authenticator)
* [二维码管理后台&google-authenticator手机端](#二维码管理后台&google-authenticator手机端)
* [管理后台与google-authenticator对接](#管理后台与google-authenticator对接)
  * [Redmine](#redmine)
  * [Zabbix](#zabbix)



## 环境

* [Sqlite3](https://www.sqlite.org/)
* [Golang 1.15.15](https://golang.org/)


## 下载

Binary can be downloaded from [Releases](https://github.com/liyinda/google-authenticator/releases) page.

## 编译

### build binary

``` shell
cd cmd/http-server
go build 
```
### build docker image
``` shell
make docker
DOCKER 部署方式作者会尽快补充
docker pull 空:latest
```

## 运行
### run binary
``` shell
1）下载编译好的二进制代码包
...

2）测试API接口访问是否正常
curl "http://127.0.0.1:7000/get?issuser=[用户名]&code=[google验证码]"
如返回ok表示返回正常
如返回error表示返回异常
```
### run docker
```
DOCKER 部署方式作者会尽快补充
docker pull 空:latest
```
### view google-authenticator
```
通过浏览器访问谷歌二维码管理后台，http://localhost:7000
注意：不要使用http://127.0.0.1:7000，因为有跨越问题vue无法发送cookie
```

## 二维码管理后台&google-authenticator手机端

### 二维码管理后台
#### 修改二维码管理后台的admin密码
``` shell
...

```

#### 创建用户二维码
![image](https://github.com/liyinda/google-authenticator/blob/master/jpg/create-qrcode.jpg)

#### 展示用户二维码
![image](https://github.com/liyinda/google-authenticator/blob/master/jpg/show-qrcode.jpg)

#### 删除用户二维码
![image](https://github.com/liyinda/google-authenticator/blob/master/jpg/delete-qrcode.jpg)

### 手机下载google-authenticator客户端
iphone手机和android手机都有对应的客户端，请大家自行下载

#### google-authenticator app
![image](https://github.com/liyinda/google-authenticator/blob/master/jpg/google-authenticator.jpg)

#### 阿里云APP中的"虚拟MFA"
![image](https://github.com/liyinda/google-authenticator/blob/master/jpg/aliyun-mfa.jpg)

## 用户后台与google-authenticator对接

### Redmine
vi app/views/account/login.html.erb
``` shell 
添加
14 <tr>
15     <td style="text-align:right;"><label for="code">Google验证码:</label></td>
16     <td style="text-align:left;"><%= text_field_tag 'code', nil, :tabindex => '3' %></td>
17 </tr>

```

vi app/controllers/account_controller.rb
``` shell 
添加
1   require "open-uri"

192   def password_authentication
193 
194     uri = 'http://[google-authenticator服务端地址]/get?issuser=' + params[:username] + '&code=' + params[:code]
195     html_response = nil
196     open(uri) do |http|
197     html_response = http.read
198     end
199 
200     if html_response == 'ok'
201 
202     user = User.try_to_login(params[:username], params[:password], false)
203     if user.nil?
204       invalid_credentials
205     elsif user.new_record?
206       onthefly_creation_failed(user, {:login => user.login, :auth_source_id => user.auth_source_id })
207     else
208       # Valid user
209       if user.active?
210         successful_authentication(user)
211       else
212         handle_inactive_user(user)
213       end
214     end
215 
216     else
217         redirect_to(:action => 'login')
218     end
219 
220 
221   end

```

![image](https://github.com/liyinda/google-authenticator/blob/master/jpg/redmine.jpg)

### Zabbix
vi include/views/general.login.php
``` shell 
添加
55         ->addItem([new CLabel(_('Password'), 'password'), (new CTextBox('password'))->setType('password')])
56         ->addItem([
57                 new CLabel(_('Google Code'), 'code'),
58                 (new CTextBox('code'))->setAttribute('', ''),
59                 $error
60         ])

```

vi index.php 
``` shell 
添加
65 if (isset($_REQUEST['enter']) && $_REQUEST['enter'] == _('Sign in')) {
66         // try to login
67         $autoLogin = getRequest('autologin', 0);
68         //print_r($_REQUEST);
69         $authflag=file_get_contents("http://[google-authenticator服务端地址]/get?issuser=".getRequest('name', '')."&code=".getRequest('code', ''));
70         //echo "http://[google-authenticator服务端地址]/get?issuser=".getRequest('name', '')."&code=".getRequest('code', '');
71         if ($authflag=='ok'){}else{
72             echo 'Google验证码错误'; header('Refresh: 2; url=http://zabbix.org/');exit;
73         }
74         //echo getRequest('code', '');

```


![image](https://github.com/liyinda/google-authenticator/blob/master/jpg/zabbix.jpg)


更多后台对接改造等您实现
