
##
Go版配置中心，基于HTTP


## 参数列表

|序号|参数|参数名|备注|
|:---|:---|:---|:---|
|1|AppId|应用Id||
|2|Cluster|集群||
|3|Profile|配置||
|4|NotifyUrl|通知Url|通知新配置|
|5|Sign|签名||

> Sign = SHA512（AppID|Cluster|Profile|NotifyURL|AppKey）

## 通知URL参数

|序号|参数|参数名|备注|
|:---|:---|:---|:---|
|1|AppId|应用Id||
|2|Cluster|集群||
|3|Profile|配置||
|4|Config|新配置|AES/CBC加密|
|5|UpdateTime|更新时间||
|6|UpdateBy|更新人||
|7|Sign|签名||

> AES = AES/CBC/PKCPS7Padding

> AppKey即为Config的AES密码，同时为参数签名的一部分 

> Config = Base64(AES(Plain_Config,AppKey))

> Sign = SHA512（AppID|Cluster|Profile|Config|NotifyURL|UpdateTime|UpdateBy|AppKey）