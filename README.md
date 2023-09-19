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
```



