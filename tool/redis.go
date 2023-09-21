package tool

import "context"



func RedisSetkeyVal(key, val string) bool {
	if _, err := rdb.Set(context.Background(), key, val, Conf.Redis.expiration).Result(); err != nil {
		return false
	}
	return true
}

func RedisGetVal(key string) (string, bool) {
	if val, err := rdb.Get(context.Background(), key).Result(); err != nil {
		return "", false
	} else {
		return val, true
	}
}