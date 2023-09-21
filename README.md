# netdisk
for learning


## Build 
```shell
make build
```
Check the `./out/netdisk-server.exe`.

## Run
Be sure that there is `/etc/config.yaml` in the wd.

#### Config
```yaml
server:
  ip: 0.0.0.0
  port: 8080
  name: netdisk
  token:
    expireDuration: 24h
    issuer: xxx

cos:
  cosBucketAddr: 
  secretId: 
  secretKey:
  innerPath: repository/
  chunkSize: 16 # MB 

dataSource:
  sourceName: mysql
  username: root
  password: xxxxxxx
  address: x.x.x.x:3306
  database: netdisk

email:
  username: xxxxxxxx@xxxx.com
  authCode: xxxxxxxxxxxxxx
    host: smtp.xxxx.com
    port: 465
  
  redis:
  address: x.x.x.x:6379
  password: xxxxxxxx
  db: 0
  poolSize: 20
  expiration: 2m
```



