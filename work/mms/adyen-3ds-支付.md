
# 3ds支付
[toc]

## 概述
- adyen平台下的一种支付校验模式

## 需求
- 需要对卡支付进行3ds支付校验

## 接入方案
- 最新用法方案
    - 概述: 支付时进行3ds支付， 对平台来说研发逻辑基本一致。具体参数可能有区别
    - 相关SDK版本信息:
        - 平台: V67 last
        - web: 
        - IOS: 
        - Android: 
    - 1. 新方案-嵌入方式： 
        - 优点: 前端集成快速, UI是多合一的解决方案
        - 缺点: 对UI的控制力度小
    - 2. 新方案-组件方式:
        - 优点: 可以每一种支付方式有单独UI组件, 可以在一定程度在自定义（需要前端同学确定自定义力度）
        - 缺点: 前端集成相比于嵌入方式要复杂
    - 3. 新方案-纯API接入方式:
        - 优点: 完全可以自定义UI
        - 缺点: 集成最慢
- 旧用法方案
    - 概述: 授权支付方案, 到账时间决定于capture时间，是否自动。从授权到支付的中间有时间窗口，非原子性。
    - 4. 3ds授权支付方案
        - 缺点: 授权后钱仍然在用户卡上，存在欺诈。 （***待确定？？？？***）

### 最新方案
- 主要有3个步骤
    - 1. 获取支付方式
    - 2. 支付方式3ds验证
    - 3. 拿验证结果请求支付
    
#### 流程图

```plantuml
@startuml
title 新方案流程图
start

:客户下单完成点击支付;

repeat
repeat :点击支付;
backward: 客户端渲染错误;
repeatwhile (请求支付方式失败?)
:选择卡支付/并输入卡信息;
:创建adyen支付;
if (不需要验证？) then (不需要)
    if (判断错误码) then(支付成功)
    :做订单状态处理;
    stop
    else
    :将错误信息渲染给用户;
    stop
    endif
elseif (需要3ds2验证？) then (需要进行3ds2)
else if (需要3ds1跳转验证) then (需要跳转验证)
else
:其他异常;
stop
endif 
  :拿到验证结果;
:请求adyen进行支付详情验证;
repeatwhile (验证失败)

end 
@enduml
```

#### 最新方案

```plantuml
@startuml

... 已经完成了下单的过程 ...
客户端 -> CM平台: 询问可以支付的方式
CM平台 <-> adyen: 询问公司账户可以支持的支付方式
CM平台 -> 客户端: 返回支付方式
客户端 -> 客户端: 收集卡信息并点击支付
客户端 -> CM平台: 必要支付方信息(卡号等信息)
CM平台 -> adyen: 将支付方和收款方信息提交adyen支付请求接口(POST /payments)
adyen -> CM平台: 返回 response
CM平台 -> CM平台: 判断response是否包含action对象

note left
官方要求
先判断action是否存在
再判断resultCode
如果action存在则不判断resultCode
end note

alt 没有action对象（表示本次交易不需要3D认证或者不支持3D认证）
CM平台 -> CM平台: 根据resultCode记录支付结果
CM平台 -> 客户端: 将response透传下去
客户端 -> 客户端: 使用resultCode将付款结果呈现给购物者
else  有action对象并且type字段为threeDS（表示本次交易需要3D认证）
CM平台 -> CM平台: 记录支付单当前状态为3ds2支付验证中
CM平台 -> 客户端: 通知客户端做3ds2支付验证
客户端 <-> adyen: 客户端组件SDK使用3D Secure 2组件执行身份验证流程,并获取认证接口
客户端 -> CM平台: 向CM平台提交认证结果
CM平台 -> CM平台: 记录支付单当前状态为3ds2支付验证结果
CM平台 -> adyen: 接受客户端提交的认证结果提交认证结果请求（POST /payments/details）
adyen -> CM平台: response 包含resultCode 透传给客户端将付款结果呈现给购物者
CM平台 -> CM平台: 记录支付结果
CM平台 -> 客户端: 透传上一步结果
else 有action对象并且type为redirect（表示需要重定向进行认证）
CM平台 -> CM平台: 记录支付单当前状态为3ds1支付验证中
CM平台 -> 客户端: 通知客户端做3ds1跳转验证
客户端 <-> adyen: 客户端组件使用createFromAction处理重定向, 并获取验证结果redirectResult
客户端 -> CM平台: 客户端请求CM平台提供的returnUrl,并附加了以Base64编码的redirectResult
CM平台 -> CM平台: 记录支付单当前状态为3ds2支付验证结果
CM平台 -> adyen: 接受认证结果提交认证结果请求（POST /payments/details）
adyen -> CM平台: response 包含resultCode 透传给客户端将付款结果呈现给购物者
CM平台 -> CM平台: 记录支付结果
CM平台 -> 客户端: 透传上一步结果
end 

...
== 异步通知 ==

loop 支付结果通知
alt 未处理订单
adyen --> CM平台: 支付结果通知
CM平台 -> CM平台: 处理订单/支付单
CM平台 -> 客户端: 处理订单/支付单
end
end

... 支付完成 ...

@enduml

```

- resultCode 说明：
    - Authorised（付款成功）-> 通知购物者付款已成功
    - Cancelled（付款被取消）-> 询问购物者是否要继续订购，或要求他们选择其他付款方式
    - Error（出错，原因包含在resultReason字段中）-> 通知购物者处理他们的付款时出错
    - Refused（拒绝，原因包含在resultReason字段中）-> 要求购物者使用其他付款方式再次尝试付款

### 旧用法方案接入时序图

```plantuml
@startuml

... 已经完成下单流程 ...
客户端 -> CM平台: 询问可以支付的方式
CM平台 <-> adyen: 询问公司账户可以支持的支付方式
CM平台 -> 客户端: 返回支付方式
客户端 -> CM平台: 1. 提供必要支付信息(卡号等信息)
CM平台 -> adyen: 2. 提交支付请求接口(POST /authorise v64)
adyen -> CM平台: 3. 返回 authorise resultCode
CM平台 -> 客户端: 5. 透传 resultCode

客户端 -> 客户端: 6. 判断resultCode

alt Authorised(成功)
客户端 -> 客户端: 支付成功跳转
else IdentifyShopper
客户端 -> 客户端: 7. 设备端调用SDK收集指纹信息
客户端 -> CM平台: 8. 上送指纹等必要信息
CM平台 -> adyen: 9. 设备指纹信息上送（POST / authorise3ds2）    
adyen -> CM平台: 10. 返回 authorise3ds2 resultCode
CM平台 -> 客户端: 11.  从第6步继续判断
else  ChallengeShopper
客户端 -> 客户端: 12. 客户端调用SDK,会弹出一个iframe框，进行二次验证
客户端 -> CM平台: 13. 二次验证结果,上次服务器
CM平台 -> adyen: 14. 二次验证结果请求
adyen -> CM平台: 15. 返回 authorise3ds2 resultCode
CM平台 -> 客户端: 16.  从第6步继续判断
else RedirectShopper
客户端 -> 客户端: 跳转到adyen进行支付

else AuthenticationNotRequired

else Cancelled Error Refused （最终状态）
CM平台 -> CM平台: 更新支付单状态
客户端 -> 客户端: 关闭订单并展示相应的界面

else  AuthenticationFinished Pending Received （中间状态）
CM平台 -> CM平台: 忽略
客户端 -> 客户端: 忽略

else PresentToShopper

end 

...
== 异步通知 ==

loop 支付结果通知
alt 未处理订单
adyen --> CM平台: 支付结果通知
CM平台 -> CM平台: 处理订单/支付单
CM平台 -> 客户端: 处理订单/支付单
end
end

@enduml

```

## 参考文档
- 在线支付概述: https://docs.adyen.com/online-payments
- drop-in接入: https://docs.adyen.com/online-payments/drop-in-web
- components接入: https://docs.adyen.com/online-payments/components-web
- 纯API接入方式: https://docs.adyen.com/online-payments/api-only
- 3ds2接入文档: https://docs.adyen.com/online-payments/3d-secure/native-3ds2
- 老方案接入: https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce/3d-secure
- API接口文档: https://docs.adyen.com/api-explorer/#/CheckoutService/v67/overview

## Q&A
1. /payments 请求/应答reference
    1. request会传两个
        - order.pspReference :  交易单ID
        - reference : 订单ID
    2. response 有3个
        - 返回请求两个
        - pspReference adyen下唯一交易单, 
    
2. 是否会产生循环跳转支付？
答： 就目前的文档阐述而言，是不会发生的。（调试时验证下） 

3. 3d1/3d2响应吗范围
答: 参考官网 




