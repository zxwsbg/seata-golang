package dao

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/transaction-wg/seata-golang/pkg/util/log"
)

import (
	"github.com/transaction-wg/seata-golang/pkg/client/at/exec"
	"github.com/transaction-wg/seata-golang/pkg/client/context"
)

const (
	insertSoMaster = `INSERT INTO seata_order.so_a (id,name,money) VALUES (?,?,?)`
	updateSoMaster = `UPDATE seata_order.so_a set money=? WHERE id=?`
	deleteSoMaster = `DELETE FROM seata_order.so_a WHERE id=?`
	selectSoMaster = `SELECT * FROM seata_order.so_a where id>=? and id<=?+5`
	selectSoMasterForUpdate = `SELECT * FROM seata_order.so_a where id>=? and id<=?+5 for update`
)

type Dao struct {
	*exec.DB
}

type So_a struct {
	Id int
	Name string
	Money int
}

var Tbl_ID = 0

func (dao *Dao) CreateSO(ctx *context.RootContext, soMasters []*So_a) ([]uint64, error) {
	result := make([]uint64, 0, len(soMasters))
	tx, err := dao.Begin(ctx)
	if err != nil {
		return nil, err
	}
	for _, soMaster := range soMasters {
		soid := Tbl_ID
		Tbl_ID += 1
		soMaster.Money = 100
		soMaster.Name = "wbs"
		_, err = tx.Exec(updateSoMaster,soMaster.Money,soid)
		//_, err = tx.Exec(insertSoMaster, soid,soMaster.Name,soMaster.Money)
		//_, err = tx.Exec(deleteSoMaster, soid)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		var res *sql.Rows
		res, err = tx.Query(selectSoMaster,soid,soid)
		log.Info("Result",res)
		if res!=nil {
			res.Close()
		}
		////_, err = tx.Exec(insertSoMaster, soid,soMaster.Name,soMaster.Money)
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

func NextID() uint64 {
	id, _ := uuid.NewUUID()
	return uint64(id.ID())
}
