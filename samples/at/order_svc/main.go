package main

import (
	"database/sql"
	"net/http"
	"time"
)

import (
	"github.com/gin-gonic/gin"
)

import (
	"github.com/transaction-wg/seata-golang/pkg"
	"github.com/transaction-wg/seata-golang/pkg/at/exec"
	"github.com/transaction-wg/seata-golang/pkg/config"
	"github.com/transaction-wg/seata-golang/pkg/context"
	"github.com/transaction-wg/seata-golang/samples/at/order_svc/dao"
)

const configPath = "/home/sbw/go/src/seata-golang/samples/at/order_svc/conf/client.yml"

func main() {
	r := gin.Default() // router 创建一个路由handler
	config.InitConf(configPath)
	pkg.NewRpcClient()
	exec.InitDataResourceManager()

	sqlDB, err := sql.Open("mysql", config.GetATConfig().DSN)
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(4 * time.Hour)

	db, err := exec.NewDB(config.GetATConfig(), sqlDB)
	if err != nil {
		panic(err)
	}
	d := &dao.Dao{
		DB: db,
	}

	r.POST("/createSo", func(c *gin.Context) {
		type req struct {
			Req []*dao.So_a
		}
		var q req
		if err := c.ShouldBindJSON(&q); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		rootContext := &context.RootContext{Context: c}
		rootContext.Bind(c.Request.Header.Get("Xid"))

		_, err := d.CreateSO(rootContext, q.Req)

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "fail",
			})
		} else {
			c.JSON(200, gin.H{
				"success": true,
				"message": "success",
			})
		}
	})
	r.Run(":8002")
}
