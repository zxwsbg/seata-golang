package dao

import (
	"github.com/transaction-wg/seata-golang/pkg/at/exec"
	"github.com/transaction-wg/seata-golang/pkg/context"
)

const (
	insertSoMaster = `INSERT INTO seata_order.so_a (id,name,money) VALUES (?,?,?)`
	updateSoMaster = `UPDATE seata_order.so_a set money=? WHERE id=?`
)

type Dao struct {
	*exec.DB
}

//现实中涉及金额可能使用长整形，这里使用 float64 仅作测试，不具有参考意义

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
		soid := Tbl_ID%10
		Tbl_ID += 1
		soMaster.Money = 30
		_, err = tx.Exec(updateSoMaster,soMaster.Money,soid)
		//_, err = tx.Exec(insertSoMaster, soid,soMaster.Name,soMaster.Money)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		result = append(result, uint64(soid))
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return result, nil
}

