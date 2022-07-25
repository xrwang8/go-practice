package impl

import (
	"context"
	"fmt"
	"github.com/go-practice/restful-api-demo/apps/host"
)

// 把Host对象保存到数据内, 数据的一致性
func (hs *HostService) save(ctx context.Context, ins *host.Host) error {

	tx, err := hs.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)

	}
	// 通过Defer处理事务提交方式
	// 1. 无报错，则Commit 事务
	// 2. 有报错, 则Rollback 事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				hs.l.Error("rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				hs.l.Error("commit error, %s", err)
			}
		}
	}()
	rstmt, err := tx.PrepareContext(ctx, InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()

	_, err = rstmt.ExecContext(ctx,
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
	)
	if err != nil {
		return err
	}

	return nil

}
