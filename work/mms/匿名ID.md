# 匿名ID
user_id 和 anonymous_id 关系是一对多
## 产生
- APP获取ip信息时，服务器返回一个唯一的anonymousID，前端保存本地，用于以后每一个请求头
## 落地
- 登陆时落地
    - 登陆时获取请求头部的anonymousID, 落地到shopping.user表
- 普通用户登陆时
    - 向shopping.uid2anonyid表中写user_id 和 anonymous_id 对应关系（一对多）
    
    
## 使用
- 1. guest迁移购物车
    - 从cart.momo_cart表中找到，anonymousKey对应的购物车
    - 从cart.momo_cart表中找到，user_id对应的购物车
    - 然后进行两个购物之前的数据迁移 
- 2. 根据anonymousKey查询用户信息(根据ByAnonymousKey去重)
- 3. 根据anonymousKey删除用户信息
- 4. 根据anonymousKey发送noti
- 5. 根据anonymousKey修改playerID
    - user-api.UpdatePlayerID
- 6. 折扣discount的使用和锁定采用， anonymousKey关联

