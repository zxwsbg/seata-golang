package dao

import (
	"github.com/transaction-wg/seata-golang/pkg/client/at/exec"
	"github.com/transaction-wg/seata-golang/pkg/client/context"
)

const (
	selectSoMaster = `SELECT * FROM seata_product.so_b where id=?`
	selectSoMasterForUpdate = `SELECT * FROM seata_product.so_b where id=? for update`
	updateSoMaster = `UPDATE seata_product.so_b set money=? WHERE id=?`
	insertSoMaster = `INSERT INTO seata_product.so_b (id,name,money) VALUES (?,?,?)`
	deleteSoMaster = `DELETE FROM seata_product.so_b WHERE id=?`
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
		soid := ID
		ID++
		req.Money = 70
		//_, err := tx.Query(selectSoMaster, soid)
		//_, err := tx.Exec(updateSoMaster,req.Money,soid)
		//_, err = tx.Exec(insertSoMaster, soid,"sbw",req.Money)
		_, err := tx.Exec(selectSoMaster,soid)
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
