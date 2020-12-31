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
grant privileges on DATABASE_NAME.TABLE_NAME to 'username'@'host';
flush privileges;

### 授权用户对某个数据库的所有表拥有所有操作权限
grant all privileges on 'db_name'.* to 'username'@'host';
### 对某个数据库某个表进行某些操作的授权
grant select, insert, update on 'db_name'.'table_name' to 'username'@'host';

### 撤销授权
revoke privilege on DATABASE_NAME.TABLE_NAME from 'username'@'host';


## golang 源码解析
    - database/sql
    - github.com/go-sql-driver/mysql
    - gorm 
