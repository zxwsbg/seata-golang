package main

import (
	"github.com/gin-gonic/gin"
)

import (
	"github.com/transaction-wg/seata-golang/pkg"
	"github.com/transaction-wg/seata-golang/pkg/config"
	"github.com/transaction-wg/seata-golang/pkg/tcc"
	"github.com/transaction-wg/seata-golang/pkg/tm"
	"github.com/transaction-wg/seata-golang/samples/tcc/service"
)

func main() {
	r := gin.Default()

	config.InitConfWithDefault("testService")
	pkg.NewRpcClient() //新建了一个rpc的client, 同时创建一个remoteRpcClient
	tcc.InitTCCResourceManager() //开启handlecommit 和 rollback 的协程

	tm.Implement(service.ProxySvc)
	tcc.ImplementTCC(service.TccProxyServiceA)
	tcc.ImplementTCC(service.TccProxyServiceB)
	tcc.ImplementTCC(service.TccProxyServiceC)

	r.GET("/commit", func(c *gin.Context) {
		service.ProxySvc.TCCCommitted(c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/rollback", func(c *gin.Context) {
		service.ProxySvc.TCCCanceled(c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8004")
}
