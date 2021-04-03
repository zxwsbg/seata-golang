package dao

import (
	"github.com/transaction-wg/seata-golang/pkg/at/exec"
	"github.com/transaction-wg/seata-golang/pkg/context"
)

const (
	allocateInventorySql = `INSERT INTO seata_product.so_b (id,name) VALUES (?,?)`
)

type Dao struct {
	*exec.DB
}

type AllocateInventoryReq struct {
	Id int
	Name string
}

var ID = 0
func (dao *Dao) AllocateInventory(ctx *context.RootContext, reqs []*AllocateInventoryReq) error {
	tx, err := dao.Begin(ctx)
	if err != nil {
		return err
	}
	for _, req := range reqs {
		soid := ID
		ID++
		_, err := tx.Exec(allocateInventorySql, soid, req.Name)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
