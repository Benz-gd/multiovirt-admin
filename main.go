package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"multiovirt-admin/controllers"
	"multiovirt-admin/dao/mysql"
	"multiovirt-admin/dao/postgresql"
	"multiovirt-admin/dao/redis"
	"multiovirt-admin/logger"
	"multiovirt-admin/pkg/snowflake"
	"multiovirt-admin/routes"
	"multiovirt-admin/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//0、获取当前路径
	workDir, _ := os.Getwd()
	//1、加载配置
	if err := settings.Init(workDir); err != nil {
		fmt.Printf("initial setting error: %v\n", err)
		return
	}

	//2、初始化日志

	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("initial logger error: %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Info("logger init success!")

	//3、初始化mysql连接
	if _, err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("initial mysql error: %v\n", err)
		zap.L().Error("init mysql error!", zap.Error(err))
		return
	} else {
		zap.L().Info("mysql init success!")
	}
	defer mysql.DBClose()

	//4、初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("initial redis error: %v\n", err)
		//return
	} else {
		zap.L().Info("redis init success!")
	}
	defer redis.Close()

	//4、初始化PG
	if _, err := postgresql.InitPostgreSQL(settings.Conf.PostgreSQLConfig); err != nil {
		fmt.Printf("initial postgresql error: %v\n", err)
	} else {
		zap.L().Info("postgresql init success!")
	}
	defer postgresql.DBClose()

	//5、初始化用户ID
	if _, err := snowflake.Init(settings.Conf.SnowFlake); err != nil {
		fmt.Printf("initial snowflake error: %v\n", err)
	} else {
		fmt.Println("initial snowflake success!")
	}

	//6、初始化validator
	if err := controllers.Init(settings.Conf.Locale); err != nil {
		zap.L().Error("init trans error", zap.Error(err))
	}

	//7、注册路由
	r := routes.InitRoutes()

	//8、启动服务优雅关机
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen error: ", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
