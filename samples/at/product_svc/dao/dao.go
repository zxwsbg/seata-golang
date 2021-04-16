package dao

import (
	"github.com/transaction-wg/seata-golang/pkg/at/exec"
	"github.com/transaction-wg/seata-golang/pkg/context"
	"sync"
)

const (
	insertSoMaster = `INSERT INTO seata_product.so_b (id,name,money) VALUES (?,?,?)`
	updateSoMaster = `UPDATE seata_product.so_b set money=? WHERE id=?`
	selectSoMaster2 = `SELECT * FROM seata_product.so_b WHERE id=?`
	selectSoMaster = `SELECT * FROM seata_product.so_b WHERE id>=? and id<=?`
	selectForUpdateSoMaster = `SELECT * FROM seata_product.so_b WHERE id>=? and id<=? FOR UPDATE`
	deleteSoMaster = `DELETE FROM seata_product.so_b WHERE id=?`

)

type Dao struct {
	*exec.DB
}

var mtx sync.Mutex

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
		mtx.Lock()
		soid := ID
		ID++
		mtx.Unlock()
		req.Money = 195
		//var res *sql.Rows
		//res,err = tx.Query(selectSoMaster2,soid)
		//log.Info("Result",res)
		//res.Close()
		//_, err := tx.Exec(insertSoMaster,soid,"sbw",req.Money)
		//_, err := tx.Exec(updateSoMaster,req.Money,soid)
		_, err := tx.Exec(deleteSoMaster,soid)
		if err != nil {
			tx.Rollback()
			return err
		}
		//var res *sql.Rows
		//res, err = tx.Query(selectForUpdateSoMaster,soid,soid+5)
		//log.Info("Result",res)
		//res.Close()
		//if err != nil {
		//	tx.Rollback()
		//	return err
		//}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
