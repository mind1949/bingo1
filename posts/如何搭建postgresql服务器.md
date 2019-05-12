---
layout: post
title: "如何搭建postgresql服务器"
date: 2018-11-26
tags:
   - 数据库
   - postgresql
---

## 起因
> 为实现站库分离,这样应用服务器被我搞崩了,也不用担心原本的数据

## 步骤:

* 安装postgresql

`$ sudo apt-get install postgresql libpq-dev postgresql-contrib`

* 创建postgresql用户

```
$ su postgres
$ createuser user random_name;
$ \psql
$ alter user random_name with encrypted password;
$ alter user random_name with createdb;
```

* 修改配置文件允许远程访问
`postgresql.conf`
```
listen_addresses = '*'
```
`pg_hba.conf`
```
host all random_name 0.0.0.0/0 trust
```

* 重启postgresql

`$ service postgresql restart`

## 注意
> 在要想远程访问数据库要在阿里云的安全组中开放postgresql的5432端口

