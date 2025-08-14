package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	App   *App   `json:"app"`
	MySQL *MySQL `json:"mysql"`
	Log   *Log   `json:"log"`
}

func (c *Config) String() string {
	v, _ := json.Marshal(c)
	return string(v)
}

type App struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (a *App) Address() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

type MySQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Debug    bool   `json:"debug"`
	db       *gorm.DB
	lock     sync.Mutex
}

// 初始化数据库
func (m *MySQL) DB() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.db == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FShanghai&allowNativePasswords=true", m.Username, m.Password, m.Host, m.Port, m.Database)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("数据库连接失败:%v", err))
		}
		if m.Debug {
			db = db.Debug()
		}
		m.db = db
	}
	return m.db
}

type Log struct {
	Level string
	Log   *zerolog.Logger
	mu    sync.Mutex
}

func (l Log) Logger() *zerolog.Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.Log == nil {
		log := zerolog.New(os.Stdout)
		l.Log = &log
	}

	return l.Log
}
