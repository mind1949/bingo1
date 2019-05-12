---
title: "mysql语句极简入门"
date: 2018-3-20
tags:
   - 数据库
   - mysql
---
## 安装mysql

```ruby
brew install mysql
```

## mysql服务

### 启动

```Mysql
mysql.server start
```

### 停止

```Mysql
mysql.server stop
```

### 了解更多命令

```Mysql
mysql.server --help
```

## mysql用户

### root用户

#### 密码设置

##### 初次设置密码

```Mysql
mysqladmin -u root password 'new-password'
```

##### 重设密码

```Mysql
mysqladmin -u root -p'old-password' password 'new-password'
```

### 普通用户

#### 登录与退出

S1:启动mysql服务

```Mysql
mysql.server start
```

S2:登录

```Mysql
mysql>mysql -h 'hostname' -u 'username' -p
```

S3:按回车键,输入密码

S4:退出

```
Control+d
```

#### 创建&权限设置

##### 创建

S1:登录root用户

```mysql
mysql -h hostname -u root -p
```

S2:创建用户

```mysql
mysql>CREATE USER 'username'@'hostname' IDENTIFIED BY 'password';
```

#### 权限设置

```Mysql
mysql>GRANT ALL PRIVILEGES ON *.* TO `username`@hostname IDENTIFIED BY "password";
```

#### 设置用户密码

### 查询mysql用户

#### 查询mysql用户

```mysql
mysql>select user from mysql.user;
+---------------+
| User          |
+---------------+
| Mind1949      |
| mysql.session |
| mysql.sys     |
| nm-mind1949   |
| nm-username   |
| oa-username   |
| root          |
+---------------+
```

```mysql
mysql>select user,host from mysql.user;
+---------------+-----------+
| User          | Host      |
+---------------+-----------+
| Mind1949      | localhost |
| mysql.session | localhost |
| mysql.sys     | localhost |
| nm-mind1949   | localhost |
| nm-username   | localhost |
| oa-username   | localhost |
| root          | localhost |
+---------------+-----------+
```

#### 查询用户(非重复)

```mysql
select distinct user from mysql.user
```

## 数据库

### 创建数据库

#### 常规方法

```Mysql
mysql>create database database_name;
```

#### 通过sql类型文件创建数据库表

```shell
mysql -D database_name -h hostname -u root -p < .sql类型文件
```

### 查看创建的数据库

```Mysql
mysql>show databases;
```

### 选择要使用/操作的数据库

#### 常用方法

```mysql
mysql>use database_name;
```

#### 登录mysql数据库时,指定操作数据库

```Mysql
mysql -d database_name -h host_name -u root -p
```

### 删除数据库

```mysql
drop database database_name;
```

## 数据库的表

### 表的创建

S1:选择使用数据库

S2:在该数据库中创建表

```Mysql
mysql>create table table_name (column1 datatype,column2 datatype); #注意最后一个column后面无逗号
```

### 表的读取

S1:选择进入要查看的数据库

S2:查看此数据库中的表

```mysql
mysql>show tables;
```

### 表的重命名

S1:选择该表所属的数据库

S2:重命名

```Mysql
mysql>alter table old_table_name renmae new_table_name;
```

### 表的删除

S1:选择该表所属的数据库

S2:删除

```mysql
mysql>drop table table_name;
```

## 表的列

### 列的添加

S1:进入对应数据库

S2:添加列

```Mysql
mysql>alter table table_name add column_name datatype;
```

### 列的读取

S1:进入对应的数据库

S2:读取

```
mysql>desc或者describe table_name;
```

### 列的修改

> 修改列的name,datatype等

```Mysql
mysql>alter table table_name change old_column_name new_column_name datatype;
```

### 列的删除

```Mysql
mysql>alter table table_name drop column_name;
```

## 表的数据

### 数据的添加

#### 对所有字段(列)同时添加数据

```mysql
mysql>insert into table_name values(value1,value2,...);
```

#### 对特定字段添加数据

```
mysql>insert into table_name (column1_name,column2_name..) values(value1,value2...);
```

### 数据的查询

#### 从某表中查询所有数据

```Mysql
mysql>select * from table_name;
```

#### 查询某表中指定列的数据

```Mysql
mysql>selct column1_name,column2_name from table_name;
```

#### 按照特定条件查询表中数据

```Mysql
mysql>select column_name from table_name where clause;
```

##### 示例:

```Mysql
mysql> select * from network3 where name='xpleaf';
mysql> select * from network3 where sex='male' and address='QingYuan';
mysql> select * from network3 where age > 40;
mysql> select * from network3 where age < 40 and age >= 31;
mysql> select * from network3 where name like "%leaf";
```

### 数据的更新

```mysql
mysql>update table_name set column1_name="new_value" where clause;
```

#### 示例:

```Mysql
mysql>create table itable (id int not null,name char(12) not null);
```

```Mysql
mysql>insert into itable (id,name) values(1,'tony_teacher');
```

```Mysql
mysql>update itable set name='tony';
```

### 数据的删除

```Mysql
mysql>delete from table_name [where clause];#通过where限定删除哪一行
```

## 分号有无得判断

> 所有登录mysql后写的语句都加分号

