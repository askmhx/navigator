# Go版配置中心，基于HTTP

### 造轮子同样让人愉悦

说明：
> SORTED_PLAIN_STRING是指参数名（不含sign本身）sort之后取其值，用竖线拼接成的字符串。
>
> 例如：
> 
> AppId="ABC"
> 
> Cluster="C1"
> 
> Profile="PRO"
> 
> NotifyUrl="https://www.baidu.com"
> 
> 得出SORTED_PLAIN_STRING=ABC|C1|https://www.baidu.com|PRO

## 1. AppServer获取配置

#### 请求方向：AppServer -> ConfigServer

### 请求参数（Form K/V）

|序号|参数|参数名|备注|
|:---|:---|:---|:---|
|1|AppId|应用Id||
|2|Cluster|集群||
|3|Profile|配置||
|5|Sign|签名||

> Sign = SHA512（SORTED_PLAIN_STRING|AppKey）

### 返回参数（JSON）


|序号|参数|参数名|备注|
|:---|:---|:---|:---|
|1|AppId|应用Id||
|2|Cluster|集群||
|3|Profile|配置||
|4|Config|新配置|AES/CBC加密|
|5|CreateAt|更新时间||
|6|CreateBy|更新人||
|7|Sign|签名||

> AES = AES/CBC/PKCPS7Padding

> AppKey即为Config的AES密码，同时为参数签名的一部分

> Config = Base64(AES(Plain_Config,AppKey))

> Sign = SHA512（SORTED_PLAIN_STRING|AppKey）

## 2. ConfigServer推送配置

#### 请求方向：ConfigServer -> AppServer

### 请求参数（Form K/V）

|序号|参数|参数名|备注|
|:---|:---|:---|:---|
|1|AppId|应用Id||
|2|Cluster|集群||
|3|Profile|配置||
|4|Config|新配置|AES/CBC加密|
|5|CreateAt|更新时间||
|6|CreateBy|更新人||
|7|Sign|签名||

> 签名及加密方式同AppServer获取配置时返回参数

### 返回参数（JSON）

|序号|参数|参数名|备注|
|:---|:---|:---|:---|
|1|AppId|应用Id||
|2|Status|是否处理|yes/no|
|3|Time|接收时间|yyyyMMddHHmmss|
|4|Sign|签名||

> Sign = SHA512（SORTED_PLAIN_STRING|AppKey）

> 只通知一次，爱咋咋