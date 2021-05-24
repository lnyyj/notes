
# shopping项目重构
## 1. 概述
## 2. 问题
## 3. 目标
- 总体目标QPS达到1000
 
## 4. 任务明细
- 1. 订单状态图分析（目前分析了前半部分，后半部还没有） 
- 2. 对app提供同步接口定义 
- 3. 支付相关异步通知接口定义 
- 4. noti相关接口定义 
- 5. 订单摘要研发（计价接口）
    - 5.1. 订单摘要-访问库表及字段调研
    - 5.2. 订单摘要-流程设计
    - 5.3. 订单摘要-完成研发
    - 5.4. 订单摘要-自测（需要测试同学配合-场景配置）
    - 备注：返回商品详情数据（附加任务）
- 6. 订单使用优惠券列表 
    - 6.1. 订单使用优惠券列表-访问库表和服务调研
    - 6.2. 订单使用优惠券列表-流程设计和完成研发
    - 6.3. 订单使用优惠券列表-订单使用优惠券自测（需要测试同学配合-场景配置）    
- 7. app下单接口(逻辑很长) 
    - 7.1. app下单-访问库表及字段调研
    - 7.2. app下单-流程设计
    - 7.3. app下单-完成研发
    - 7.4. app下单-自测（需要测试同学配合-场景配置）
- 8. 获取订单支付方式（目前是代码里写死的）
    - 8.1. 支付方式-流程设计
    - 8.2. 支付方式-研发
    - 8.3. 支付方式-自测（需要测试同学配合-场景配置）
- 9. app通知支付成功接口（主要是删除购物车数据）
    - 9.1. app通知支付成功接口-研发
    - 9.2. app通知支付成功接口-自测
- 10. 获取订单详情(v_43， 只有web使用) 
    - 10.1. 获取订单详情v_43-访问库表及字段调研 
    - 10.2. 获取订单详情v_43-研发
    - 10.3. 获取订单详情v_43-自测
- 11. Paypal支付请求
    - Paypal支付请求-流程调研
    - Paypal支付请求-接口研发
    - Paypal支付请求-自测
- 12. Paypal支付完成后通知服务器
    - Paypal支付完成后通知服务器-流程调研
    - Paypal支付完成后通知服务器-研发
    - Paypal支付完成后通知服务器-自测
- 13. adyen支付请求
    - adyen支付请求-流程调研
    - adyen支付请求-研发
    - adyen支付请求-测试
- 14. 支付成功事件处理（两种场景，authorized/captured）
    - 14.1. 支付成功事件-写入库表及字段调研
    - 14.2. 支付成功事件-研发
    - 14.3. 支付成功事件-自测
    - 备注: paypal 的authorized状态在callback过滤了
- 15. 支付失败事件 
    - 15.1. 支付失败事件-写入库表及字段调研
    - 15.2. 支付失败事件-研发
    - 15.3. 支付失败事件-自测
- 16. 购买事件处理（event/purchased）
    - 16.1. 购买事件处理-使用场景调研
    - 16.2. 购买事件处理-研发
    - 16.3. 购买事件处理-自测
- 17. 其他事件 (*h)
    - 17.1. event/group_update使用场景调研（接口是否已经废弃）
- 18. 退款处理接口（成功/失败/重试等场景）
    - 18.1. 退款接口调用（工单）
    - 18.2. 退款接口使用场景/调用链/库表字段调研
    - 18.3. 退款接口-研发
    - 18.4. 退款接口-测试（需要模拟写入退款相关数据）
    - 18.5. 退款结果通知事件-流程/库表调研
    - 18.6. 退款结果通知事件-研发
    - 18.7. 退款结果通知事件-测试
- ~~19. 模拟物流数据 (*h)~~
    - ~~19.1. 构造物流数据（模拟履约）~~
    - ~~19.2. 创建物流通知消息（模拟履约）~~
    - ~~19.3. 物流消息的处理接口（平台）~~
    - ~~使订单的状态可以顺利的从 收货中 ---> 收获完成~~
- 20. noti接口
- 21. 订单列表适配 (*h)
    - 21.1. 前端订单列表适配
    - 21.2. 后台管理界面订单列表适配
    - 整个下单流程的状态/信息显示验证

### 
    
    
## 5. 设计方案
- 第一阶段选用方案2

### 方案一
继续在python老代码上改动， 将不需要的服务调用全部移除

### 方案二
- 保留现有DB结构和kafka通信链路，使用golang重写目前在使用的几个接口
    - shopping/cal_price/v4_4 ---- 订单摘要
    - shopping/user_coupons_for_order -- 获取订单使用的优惠券
    - shopping/cart/checkout/entries -- 下单
    - shopping/orders/{order_id}/pay_cm ---获取支付方式
    - shopping/orders/{order_id}/paid_from_client  ----告诉后台支付成功
    - event/payment_notified_cm 支付成功事件处理
    - event/payment_failed_cm 支付失败事件处理
    - event/purchased 购买事件处理
    - refund/refund 退款接口处理
    - event/refund_notified 退款结果处理
    - event/* 其他事件处理
    - noti相关接口    

### 方案三
- 微服务库内聚
- 微服务DB隔离，不能夸服务访问其他的库
- 业务流程重理



## 后台时序图

```plantuml
@startuml
... 加购流程 ...
... 下单流程 ... 
APP -> shopping: shopping/cal_price/v4_4 订单摘要

alt 已登陆并且使用优惠券
    shopping -> shopping: 判断有可供消费的优惠券
    shopping -> shopping: 订单使用的优惠券
    shopping -> shopping: 判读用户优惠券 + 地区
else 未登陆并且使用优惠券
    shopping -> shopping: 判读用户优惠券 + 地区
end 

shopping -> pricectrl: 商城优惠调用MallDiscounts
shopping -> rebateSrv: 请求Wallet信息
shopping -> forex: 获取外汇信息
shopping -> shopping: calc_price

APP <- shopping: 订单摘要返回
APP <-> payment: payment/adyen/recurring_details 已支付卡列表

APP -> address: 获取地址列表address-api/addresses/list
opt 添加地址
APP -> address: 添加地址: address-api/upsert
end

APP -> shopping: 获取订单使用的优惠券shopping/user_coupons_for_order
opt 兑换优惠券
APP -> coupon: coupon/redeem
end 

APP -> shopping: 下单 shopping/cart/checkout/entries
shopping -> shopping: 地址合法性判断（黑名单等）
shopping -> shopping: 计算价格calc_price

shopping -> APP: 下单 成功

APP -> shopping: shopping/orders/{order_id}/pay_cm 
shopping -> APP: 返回支付方式

APP <-> payment: adyen卡支付payment/adyen/authorise
APP -> shopping: 支付通知shopping/orders/{order_id}/paid_from_client
shopping -> shopping: 移除购物车等

adyen --> payment_callbak: 支付成功通知(http)
payment_callbak -> orders: 支付成功通知(mrpc)
orders -> shopping: 支付成功通知(http)
shopping -> shopping: 修改状态
...

@enduml
```                   


### 计价时序图
### 下单时序图
### 退款时序图
### 物流时序图




### Q&A
- Q: 对orders服务的依赖
- A:
    - 1.下单接口
    - 2. 支付事件通知
    - 3. noti事件通知
