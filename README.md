火币网的API移植

# 使用
面向用户的调用需要设置Access-Key与Private-Key, 面积行情的调用不用设置
``` Go

ak := "xxxx-xxxx-xxxx-xxxx"
pk := "xxxx-xxxx-xxxx-xxxx"

// 获取用户基本信息
session = huobi.NewSession(ak, pk)
accountInfo := session.GetAccountInfo()


// 获取LTC市场深度
depthInfo := huobi.GetDepth(1, huobi.CNYLTC)
fmt.Println(depthInfo)
```

# 注意
lab下对应的是websocket的调用, 现在只是在实验代码可能会调整比较频繁

服务器json都没有Unmarshal, 留给用户自行反序列化
可以使用 [gabs](https://github.com/Jeffail/gabs)


# 相关参考
[huobi_wiki](https://github.com/huobiapi/API_Docs/wiki)

