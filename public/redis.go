package public

import (
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
	"os"
)

var pool *redis.Pool

func init() {
	work, _ := os.Getwd()
	//设置文件名和文件后缀
	viper.SetConfigName("redis_map")
	viper.SetConfigType("toml")
	//配置文件所在的文件夹
	viper.AddConfigPath(work + ConfPath)
	viper.ReadInConfig()

	pool = &redis.Pool{
		MaxIdle:     viper.GetInt("list.default.max_idle"),   // 最大空闲连接数
		MaxActive:   viper.GetInt("list.default.max_active"), // 和数据库的最大连接数，0 表示没有限制
		IdleTimeout: 100,                                     // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化连接的代码
			return redis.Dial("tcp", viper.GetStringSlice("list.default.proxy_list")[0],
				redis.DialPassword(viper.GetString("list.default.password")))
		},
	}
}

func RedisConfPipline(pip ...func(c redis.Conn)) error {
	// 从 pool 中取出一个连接
	c := pool.Get()
	defer c.Close()
	for _, f := range pip {
		f(c)
	}
	c.Flush()
	return nil
}

func RedisConfDo(commandName string, args ...interface{}) (interface{}, error) {
	c := pool.Get()
	defer c.Close()
	return c.Do(commandName, args...)
}
