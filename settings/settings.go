package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init(workDir string) (err error) {
	fmt.Printf("workdir is: %s\n", workDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir)
	if err = viper.ReadInConfig(); err != nil {
		return err
		fmt.Printf("Fatal error config file: %s \n", err)
	}
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper unmarshal faild err: %v\n", err)
	}
	//实时监控配置变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper unmarshal faild err: %v\n", err)
		}
		fmt.Println("configuration has been modified! ")
	})
	return
}
