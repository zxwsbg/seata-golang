package dao

import (
	"github.com/transaction-wg/seata-golang/pkg/at/exec"
	"github.com/transaction-wg/seata-golang/pkg/context"
)

const (
	allocateInventorySql = `INSERT INTO seata_product.so_b (id,name,money) VALUES (?,?,?)`
	updateSoMaster = `UPDATE seata_product.so_b set money=? WHERE id=?`
)

type Dao struct {
	*exec.DB
}

type AllocateInventoryReq struct {
	Id int
	Name string
	Money int
}

var ID = 0
func (dao *Dao) AllocateInventory(ctx *context.RootContext, reqs []*AllocateInventoryReq) error {
	tx, err := dao.Begin(ctx)
	if err != nil {
		return err
	}
	for _, req := range reqs {
		soid := ID%10
		ID++
		req.Money = 70
		_, err := tx.Exec(updateSoMaster, req.Money,soid)
		//_, err := tx.Exec(allocateInventorySql, soid, req.Name,req.Money)
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
