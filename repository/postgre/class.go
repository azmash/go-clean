package postgre

import (
	"context"
	"time"
)

func (db *DB) SetMembership(ctx context.Context, shopID int64, newStatus int) error {

	query := `INSERT INTO shop_membership (shop_id, status, created_on)
	VALUES($1,$2,$3) 
	ON CONFLICT (shop_id) 
	DO 
	   UPDATE SET status = $2`

	_, err := db.DB.ExecContext(ctx, query, shopID, newStatus, time.Now())

	return err
}
