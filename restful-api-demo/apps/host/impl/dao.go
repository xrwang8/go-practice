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
	stmt, err := tx.PrepareContext(ctx, InsertResourceSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
	)
	if err != nil {
		return err
	}

	return nil

}
