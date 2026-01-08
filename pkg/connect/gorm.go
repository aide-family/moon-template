package connect

import (
	"fmt"
	"net/url"
	"strings"

	klog "github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/aide-family/magicbox/log/gormlog"
	"github.com/aide-family/magicbox/pointer"
	"github.com/aide-family/sovereign/pkg/config"
	"github.com/aide-family/sovereign/pkg/merr"
)

func init() {
	globalRegistry.RegisterORMFactory(config.ORMConfig_MYSQL, NewORMConfigBuilderFromMySQL)
	globalRegistry.RegisterORMFactory(config.ORMConfig_SQLITE, NewORMConfigBuilderFromSQLite)
}

func NewORMConfigBuilderFromMySQL(c *config.ORMConfig, logger *klog.Helper) (*ORMConfig, error) {
	mysqlConf := &config.MySQLOptions{}
	if pointer.IsNotNil(c.GetOptions()) {
		if err := anypb.UnmarshalTo(c.GetOptions(), mysqlConf, proto.UnmarshalOptions{Merge: true}); err != nil {
			return nil, merr.ErrorInternalServer("unmarshal mysql config failed: %v", err)
		}
	}
	params := url.Values{}
	for key, value := range mysqlConf.Parameters {
		params.Add(key, value)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", mysqlConf.Username, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database, params.Encode())
	ormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	if strings.EqualFold(c.GetUseSystemLogger(), "true") {
		ormConfig.Logger = gormlog.New(logger.Logger())
	}

	return &ORMConfig{
		Dialector: mysql.Open(dsn),
		Config:    ormConfig,
		IsDebug:   strings.EqualFold(c.GetDebug(), "true"),
	}, nil
}

func NewORMConfigBuilderFromSQLite(c *config.ORMConfig, logger *klog.Helper) (*ORMConfig, error) {
	sqliteConf := &config.SQLiteOptions{}
	if pointer.IsNotNil(c.GetOptions()) {
		if err := anypb.UnmarshalTo(c.GetOptions(), sqliteConf, proto.UnmarshalOptions{Merge: true}); err != nil {
			return nil, merr.ErrorInternalServer("unmarshal sqlite config failed: %v", err)
		}
	}
	ormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	if strings.EqualFold(c.GetUseSystemLogger(), "true") {
		ormConfig.Logger = gormlog.New(logger.Logger())
	}
	return &ORMConfig{
		Dialector: sqlite.Open(sqliteConf.Dsn),
		Config:    ormConfig,
		IsDebug:   strings.EqualFold(c.GetDebug(), "true"),
	}, nil
}

type ORMConfig struct {
	Dialector gorm.Dialector
	Config    *gorm.Config
	IsDebug   bool
}

func (c *ORMConfig) BuildDB() (*gorm.DB, error) {
	db, err := gorm.Open(c.Dialector, c.Config)
	if err != nil {
		return nil, err
	}

	if c.IsDebug {
		return db.Debug(), nil
	}

	return db, nil
}

func CloseDB(db *gorm.DB) error {
	mdb, err := db.DB()
	if err != nil {
		return fmt.Errorf("get db connection failed: %w", err)
	}
	return mdb.Close()
}

func NewDB(c *config.ORMConfig, logger *klog.Helper) (*gorm.DB, error) {
	factory, ok := globalRegistry.GetORMFactory(c.GetDialector())
	if !ok {
		return nil, merr.ErrorInternalServer("orm factory not registered")
	}
	ormConfig, err := factory(c, logger)
	if err != nil {
		return nil, err
	}
	return ormConfig.BuildDB()
}
