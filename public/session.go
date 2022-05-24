package public

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/spf13/viper"
	"os"
)

func GetSessionStore() (sessions.RedisStore, error) {
	work, _ := os.Getwd()
	//设置文件名和文件后缀
	viper.SetConfigName("redis_map")
	viper.SetConfigType("toml")
	//配置文件所在的文件夹
	viper.AddConfigPath(work + ConfPath)
	viper.ReadInConfig()
	store, err := sessions.NewRedisStore(10, "tcp", viper.GetStringSlice("list.default.proxy_list")[0],
		viper.GetString("list.default.password"), []byte("secret"))
	return store, err
}
