# go-short-url

> 使用beego开发短链接服务

## 个人服务部署

修改conf > app.conf 下的

sinaSource  需要在新浪开发者中心申请


## 使用方式

访问

```text
http://127.0.0.1:8080/v1/short?url=需要转换的url
```

返回结果

```json
{
    "code": 200,
    "message": "",
    "data": {
        "long_url": "http://www.java1234.com/a/javabook/",
        "short_url": "http://t.cn/8sd9Xa4"
    }
}
```


