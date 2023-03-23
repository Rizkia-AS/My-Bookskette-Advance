// package db
package unused

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// )

// type Store struct {
// 	*Queries
// 	db *sql.DB
// }

// func NewStore(db *sql.DB) *Store {
// 	return &Store{
// 		Queries: New(db),
// 		db:      db,
// 	}
// }

// func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx,err := store.db.BeginTx(ctx,nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)

// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err : %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit()

// }



// // NOTE : sepertinya kita tidak memerlukan TX untuk aplikasi ini karena, jika kita ingin memasukan buku kita hanya perlu menggunakan satu operasi database saja yaitu InsertBook


// // type SaveBookTxParams struct {

// // }

// // type SaveBookTxResult struct {

// // }

// // func (store *Store) SaveBookTx(ctx context.Context, arg SaveBookTxParams) (SaveBookTxResult, error) {
// // 	var result SaveBookTxResult



// // }