package multicolumns

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-pg/pg"
	"github.com/satori/go.uuid"
)

func CreateTable(db *pg.DB) error {
	_, err := db.Exec(fmt.Sprintf(`
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		DROP TABLE IF EXISTS t;
		CREATE TABLE IF NOT EXISTS t (
			c1_int integer,
			c1_str varchar,
			c2_int integer,
			c2_str varchar,
			c2_uuid uuid
		);
	`))
	return err
}

func CreateData(db *pg.DB, rows int) (err error) {
	rand.Seed(time.Now().Unix())

	batch := 10000
	if batch > rows {
		batch = rows
	}
	for start := 0; start < rows; start += batch {
		remain := rows - start
		if remain > batch {
			remain = batch
		}
		stmt := "INSERT INTO t VALUES "
		for remain := batch; remain > 0; remain-- {
			c1 := uint32(rand.Int31n(10))
			c2 := uint32(rand.Int31n(100000000))
			buf := make([]byte, 16)
			binary.LittleEndian.PutUint32(buf, c2)
			if c2uuid, err := uuid.FromBytes(buf); err == nil {
				if remain == 1 {
					stmt += fmt.Sprintf(`(%d, '%d', %d, '%d', '%s');`, c1, c1, c2, c2, c2uuid.String())
				} else {
					stmt += fmt.Sprintf(`(%d, '%d', %d, '%d', '%s'),`, c1, c1, c2, c2, c2uuid.String())
				}
			}
		}

		_, err = db.Exec(stmt)
		if err != nil {
			fmt.Println(stmt)
			break
		}
	}
	return
}

func CreateIndex(db *pg.DB, method, def string) error {
	_, err := db.Exec(fmt.Sprintf(`
		DROP INDEX IF EXISTS t_idx;
		CREATE INDEX t_idx ON t USING %s %s;
		ANALYZE t;
	`, method, def))
	return err
}
