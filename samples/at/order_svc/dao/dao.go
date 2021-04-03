package dao

import (
	"github.com/transaction-wg/seata-golang/pkg/at/exec"
	"github.com/transaction-wg/seata-golang/pkg/context"
)

const (
	insertSoMaster = `INSERT INTO seata_order.so_a (id,name) VALUES (?,?)`
)

type Dao struct {
	*exec.DB
}

//现实中涉及金额可能使用长整形，这里使用 float64 仅作测试，不具有参考意义

type So_a struct {
	Id int
	Name string
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
		_, err = tx.Exec(insertSoMaster, soid,soMaster.Name)
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

