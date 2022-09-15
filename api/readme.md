创建数据库
```sql
CREATE SCHEMA `ginadmin` DEFAULT CHARACTER SET utf8mb4 ;
```

创建表结构
```
go run main.go -m
```

插入管理员数据,默认账号密码都是`admin`,密码使用HmacSHA256加密，密钥在`service.base.go`的`SystemSecretKey`中，建议修改为自己的密钥(Hex格式，长度32字符)
```sql
INSERT INTO account (id,acct, passwd, otp_secret, account_status, create_at , create_by) VALUES (1,'admin', '9c427838783fc42c5747f9d741c5df2ecd23971feb2075accbd0980efc972a32', '', 'NORMAL', 1662131046 , 1);
INSERT INTO unicom.`user` (id,alias,user_type,update_at,update_by) VALUES (1,'测试管理员','admin',0,0);
```

启动api
```
go run main.go
```

配置文件:conf.yaml
```yaml
debug: true  #是否调试模式
port: :8080  #api监听端口
file_srv_dir: ./files #文件服务目录
file_srv_base_url: http://127.0.0.1:8081/files #文件服务地址，与/nginx.conf配置保持一致
dbstr: root:localdbpasswd@tcp(127.0.0.1:3306)/ginadmin?charset=utf8mb4&parseTime=True&loc=Local&writeTimeout=0  #mysql连接串，用户密码，地址，数据库名称可以根据需要修改
redis_str: 127.0.0.1:6379  #redis地址，用于用户登录验证等功能
```

项目结构
```
├─controllers   控制器
├─dao           数据库交互层
├─dto           api入参与返回结构
├─models        数据库模型
├─service       业务逻辑
├─test          单元测试
└─utils         工具库
```