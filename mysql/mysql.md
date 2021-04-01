# mysql 操作解析


## 常规操作
- 数据库版本 5.7

### 创建用户
create user 'username'@'host' identified by 'password';
### 更改用户密码
set password for 'username'@'host' = password('new_password');
### 删除用户
drop user 'username'@'localhost'

### 查看用户权限
show grants for 'username'@'host';
### 授权
grant all privileges on DATABASE_NAME.TABLE_NAME to 'username'@'host';
flush privileges;

### 授权用户对某个数据库的所有表拥有所有操作权限
grant all privileges on `db_name`.* to 'username'@'host';
### 对某个数据库某个表进行某些操作的授权
grant select, insert, update on 'db_name'.'table_name' to 'username'@'host';

### 撤销授权
revoke privilege on DATABASE_NAME.TABLE_NAME from 'username'@'host';


### 查看版本号
select version();

### 查看事物隔离级别
select @@transaction_isolation;
show variables like 'transaction_isolation';

### 修改事物隔离级别
- 读未提交: set global transaction isolation level read uncommitted; 
- 读已提交: set global transaction isolation level read committed;  
- 可重复读: set global transaction isolation level repeatable read;
- 串行化: set global transaction isolation level serializable;

| \ | 读脏 | 不可重复读  | 幻读  |  
|:---:|:---:|:---:|:---:|
| 读未提交| 可能 | 可能 | 可能 | 
| 读已提交| 不可能 |  可能 | 可能 |  
| 可重复读| 不可能 | 不可能 | 可能  |  
| 串行化| 不可能 | 不可能 | 不可能 |  
| 针对语句|insert,update,delete | update,delete  | insert |  



## golang 源码解析
    - database/sql
    - github.com/go-sql-driver/mysql
    - gorm 
