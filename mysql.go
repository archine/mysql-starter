package starter

import (
	"fmt"
	"github.com/archine/ioc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type M struct {
	Db *gorm.DB
}

type config struct {
	Mysql mysqlConfig `mapstructure:"mysql"`
}

type mysqlConfig struct {
	LogLevel   string        `mapstructure:"log_level"` // Log level: debug or error
	URL        string        `mapstructure:"url"`       // Mysql url, IP:PORT
	UserName   string        `mapstructure:"username"`
	Password   string        `mapstructure:"password"`
	Database   string        `mapstructure:"database"`
	MaxIdle    int           `mapstructure:"max_idle"`
	MaxConnect int           `mapstructure:"max_connect"`
	IdleTime   time.Duration `mapstructure:"max_idle_time"`
}

func (m *M) CreateBean() ioc.Bean {
	var conf config
	v := ioc.GetBeanByName("viper.Viper").(*viper.Viper)
	v.SetDefault("mysql.log_level", "error")
	v.SetDefault("mysql.max_idle", 10)
	v.SetDefault("mysql.max_connect", 50)
	v.SetDefault("mysql.max_idle_time", 30*time.Second)
	if err := v.Unmarshal(&conf); err != nil {
		log.Fatalf("Failed to create a mysql bean, %s", err.Error())
	}
	dsn := conf.Mysql.URL
	if conf.Mysql.Database != "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Mysql.UserName, conf.Mysql.Password, conf.Mysql.URL, conf.Mysql.Database)
	}
	gormConf := &gorm.Config{
		PrepareStmt: true,
	}
	if conf.Mysql.LogLevel == "error" {
		gormConf.Logger = logger.Default.LogMode(logger.Error)
	} else {
		gormConf.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), gormConf)
	if err != nil {
		log.Fatalf("Failed to create a mysql bean, %s", err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to create a mysql bean, %s", err.Error())
	}
	sqlDB.SetMaxIdleConns(conf.Mysql.MaxIdle)
	sqlDB.SetMaxOpenConns(conf.Mysql.MaxConnect)
	sqlDB.SetConnMaxLifetime(conf.Mysql.IdleTime)
	return &M{Db: db}
}
