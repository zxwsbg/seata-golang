package dao

import (
	"github.com/transaction-wg/seata-golang/pkg/at/exec"
	"github.com/transaction-wg/seata-golang/pkg/context"
	"sync"
)

const (
	insertSoMaster = `INSERT INTO seata_order.so_a (id,name,money) VALUES (?,?,?)`
	updateSoMaster = `UPDATE seata_order.so_a set money=? WHERE id=?`
	selectSoMaster2 = `SELECT * FROM seata_order.so_a WHERE id=? `
	selectSoMaster = `SELECT * FROM seata_order.so_a WHERE id>=? and id<=? `
	selectForUpdateSoMaster = `SELECT * FROM seata_order.so_a WHERE id=? FOR UPDATE`
	deleteSoMaster = `DELETE FROM seata_product.so_b WHERE id=?`
)

type Dao struct {
	*exec.DB
}
var mtx sync.Mutex

//现实中涉及金额可能使用长整形，这里使用 float64 仅作测试，不具有参考意义

type So_a struct {
	Id int
	Name string
	Money int
}

var Tbl_ID = 5

func (dao *Dao) CreateSO(ctx *context.RootContext, soMasters []*So_a) ([]uint64, error) {
	result := make([]uint64, 0, len(soMasters))
	tx, err := dao.Begin(ctx)
	if err != nil {
		return nil, err
	}
	for _, soMaster := range soMasters {
		mtx.Lock()
		soid := Tbl_ID
		Tbl_ID += 1
		mtx.Unlock()
		soMaster.Money = 5
		//_,err = tx.Query(selectSoMaster2,soid)
		//_, err = tx.Exec(insertSoMaster,soid,"sbw",soMaster.Money)
		//_, err = tx.Exec(updateSoMaster,soMaster.Money,soid)
		_, err := tx.Exec(deleteSoMaster,soid)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		//var res *sql.Rows
		//res,err = tx.Query(selectForUpdateSoMaster,soid)
		//log.Info("Result",res)
		//res.Close()
		//if err != nil {
		//	tx.Rollback()
		//	return nil, err
		//}
		result = append(result, uint64(soid))
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return result, nil
}

