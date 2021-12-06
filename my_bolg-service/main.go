package main

import (
	"bolg/global"
	"bolg/internal/model"
	"bolg/internal/routers"
	"bolg/pkg/logger"
	s "bolg/pkg/setting"
	"bolg/pkg/tracer"
	"bolg/pkg/validator"
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	port      string
	runMode   string
	config    string
	isVersion bool
)

func init() {
	//命令行参数
	flagErr := setupFlag()
	if flagErr != nil {
		log.Fatalf("init.setupFlag err: %v", flagErr)
	}

	//初始化配置
	settintErr := setupSetting()
	if settintErr != nil {
		log.Fatalf("init.setupsetting err: %v", settintErr)
	}
	//初始化db
	dbErr := setupDBEngine()
	if dbErr != nil {
		log.Fatalf("init.setupDBEngine err: %v", dbErr)
	}

	//初始化日志
	logErr := setupLogger()
	if logErr != nil {
		log.Fatalf("init.setupLogger err: %v", logErr)
	}

	//验证器
	validErr := setupValidator()
	if validErr != nil {
		log.Fatalf("init.setupValidator err: %v", validErr)
	}

	traceErr := setupTracer()
	if traceErr != nil {
		log.Fatalf("init.setupTracer err: %v", traceErr)
	}

}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "[ChenJingBo]", log.LstdFlags).WithCaller(2)

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}

func setupSetting() error {
	setting, err := s.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupValidator() error {
	global.Validator = validator.NewCustomValidator()
	global.Validator.Engine()
	binding.Validator = global.Validator
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "39.97.119.197:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    10 * global.ServerSetting.ReadTimeout,
		WriteTimeout:   10 * global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	ctx := context.Background()
	global.Logger.Info(ctx, "blog service")

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.listenAndServer err :%v", err)
		}
	}()

	quit := make(chan os.Signal)
	//Ctrl-C产生SIGINT信号，Ctrl-\产生SIGQUIT信号，Ctrl-Z产生SIGTSTP信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) //SIGTERM 终止进程
	<-quit
	log.Println("Shutting down server...")

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Sever forced to shutdown:", err)
	}
	log.Println("Server exiting")

}

//SIGHUP	1	终端挂起或控制进程终止。当用户退出Shell时，由该进程启动的所有进程都会收到这个信号，默认动作为终止进程。
//SIGINT	2	键盘中断。当用户按下组合键时，用户终端向正在运行中的由该终端启动的程序发出此信号。默认动作为终止进程。
//SIGQUIT	3	键盘退出键被按下。当用户按下或组合键时，用户终端向正在运行中的由该终端启动的程序发出此信号。默认动作为退出程序。
//SIGFPE	8	发生致命的运算错误时发出。不仅包括浮点运算错误，还包括溢出及除数为0等所有的算法错误。默认动作为终止进程并产生core文件。
//SIGKILL	9	无条件终止进程。进程接收到该信号会立即终止，不进行清理和暂存工作。该信号不能被忽略、处理和阻塞，它向系统管理员提供了可以杀死任何进程的方法。
//SIGALRM	14	定时器超时，默认动作为终止进程。
//SIGTERM	15	程序结束信号，可以由 kill 命令产生。与SIGKILL不同的是，SIGTERM 信号可以被阻塞和终止，以便程序在退出前可以保存工作或清理临时文件等。
