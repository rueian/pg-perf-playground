package multicolumns

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/go-pg/pg"
	"github.com/satori/go.uuid"
)

type Ctx struct {
	b  *testing.B
	db *pg.DB
}

func BenchmarkMultiColumns(b *testing.B) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "postgres",
		Addr:     "postgres:5432",
	})
	defer db.Close()

	ctx := &Ctx{
		b:  b,
		db: db,
	}

	benchmarkWithRows(ctx, 10000)
	benchmarkWithRows(ctx, 1000000)
}

func benchmarkWithRows(ctx *Ctx, rows int) {
	e(CreateTable(ctx.db))
	e(CreateData(ctx.db, rows))
	prefix := fmt.Sprintf(`rows=%d/`, rows)
	benchmarkWithIndex(ctx, prefix, "btree", "(c1_int,c2_int)",
		NewQueryModeWhereIn("c1_int", "c2_int", c1IntValuer, c2IntValuer),
		NewQueryModeForeach("c1_int", "c2_int", c1IntValuer, c2IntValuer),
	)
	benchmarkWithIndex(ctx, prefix, "btree", "(c1_int,c2_uuid)",
		NewQueryModeWhereIn("c1_int", "c2_uuid", c1IntValuer, c2StrValuer),
		NewQueryModeForeach("c1_int", "c2_uuid", c1IntValuer, c2StrValuer),
	)
	benchmarkWithIndex(ctx, prefix, "btree", "(c1_str,c2_uuid)",
		NewQueryModeWhereIn("c1_str", "c2_uuid", c1StrValuer, c2StrValuer),
		NewQueryModeForeach("c1_str", "c2_uuid", c1StrValuer, c2StrValuer),
	)
	benchmarkWithIndex(ctx, prefix, "btree", "(c1_str,c2_str)",
		NewQueryModeWhereIn("c1_str", "c2_str", c1StrValuer, c2StrValuer),
		NewQueryModeForeach("c1_str", "c2_str", c1StrValuer, c2StrValuer),
	)
	benchmarkWithIndex(ctx, prefix, "btree", "(c2_int)",
		NewQueryModeWhereIn("c1_int", "c2_int", c1IntValuer, c2IntValuer),
		NewQueryModeForeach("c1_int", "c2_int", c1IntValuer, c2IntValuer),
	)
	benchmarkWithIndex(ctx, prefix, "btree", "(c2_uuid)",
		NewQueryModeWhereIn("c1_int", "c2_uuid", c1IntValuer, c2StrValuer),
		NewQueryModeForeach("c1_int", "c2_uuid", c1IntValuer, c2StrValuer),
	)
	benchmarkWithIndex(ctx, prefix, "btree", "(c2_str)",
		NewQueryModeWhereIn("c1_int", "c2_str", c1IntValuer, c2StrValuer),
		NewQueryModeForeach("c1_int", "c2_str", c1IntValuer, c2StrValuer),
	)
	benchmarkWithIndex(ctx, prefix, "hash", "(c2_int)",
		NewQueryModeWhereIn("c1_int", "c2_int", c1IntValuer, c2IntValuer),
		NewQueryModeForeach("c1_int", "c2_int", c1IntValuer, c2IntValuer),
	)
	benchmarkWithIndex(ctx, prefix, "hash", "(c2_uuid)",
		NewQueryModeWhereIn("c1_int", "c2_uuid", c1IntValuer, c2StrValuer),
		NewQueryModeForeach("c1_int", "c2_uuid", c1IntValuer, c2StrValuer),
	)
	benchmarkWithIndex(ctx, prefix, "hash", "(c2_str)",
		NewQueryModeWhereIn("c1_int", "c2_str", c1IntValuer, c2StrValuer),
		NewQueryModeForeach("c1_int", "c2_str", c1IntValuer, c2StrValuer),
	)
}

func benchmarkWithIndex(ctx *Ctx, prefix, method, def string, modes ...*QueryMode) {
	e(CreateIndex(ctx.db, method, def))
	prefix = prefix + fmt.Sprintf(`%s%s/`, method, def)
	benchmarkWithBatch(ctx, prefix, 1, modes...)
	benchmarkWithBatch(ctx, prefix, 100, modes...)
}

func benchmarkWithBatch(ctx *Ctx, prefix string, batch int, modes ...*QueryMode) {
	prefix = prefix + fmt.Sprintf(`batch=%d/`, batch)
	var pairs []*Pair
	for i := 0; i < batch; i++ {
		c2 := uint32(rand.Int31n(100000000))
		buf := make([]byte, 16)
		binary.LittleEndian.PutUint32(buf, c2)
		c2uuid, err := uuid.FromBytes(buf)
		if err != nil {
			panic(err)
		}

		pairs = append(pairs, &Pair{
			c1:     uint32(rand.Int31n(10)),
			c2:     c2,
			c2uuid: c2uuid,
		})
	}

	for _, mode := range modes {
		benchmarkWithQueryMode(ctx, prefix, pairs, mode)
	}
}

func benchmarkWithQueryMode(ctx *Ctx, prefix string, pairs []*Pair, mode *QueryMode) {
	ctx.b.Run(prefix+fmt.Sprintf(`mode=%s`, mode.name), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := mode.run(ctx, pairs)
			if err != nil {
				panic(err)
			}
		}
	})
}

type T struct {
	tableName struct{}  `sql:"t"`
	C1Int     int       `sql:"c1_int"`
	C1Str     string    `sql:"c1_str"`
	C2Int     int       `sql:"c2_int"`
	C2Str     string    `sql:"c2_str"`
	C2UUID    uuid.UUID `sql:"c2_uuid,type:uuid"`
}

type Pair struct {
	c1     uint32
	c2     uint32
	c2uuid uuid.UUID
}

type valuer func(p *Pair) interface{}

var c1IntValuer valuer = func(p *Pair) interface{} {
	return p.c1
}

var c1StrValuer valuer = func(p *Pair) interface{} {
	return strconv.Itoa(int(p.c1))
}

var c2IntValuer valuer = func(p *Pair) interface{} {
	return p.c2
}

var c2StrValuer valuer = func(p *Pair) interface{} {
	return p.c2uuid.String()
}

type QueryMode struct {
	name string
	run  func(ctx *Ctx, pairs []*Pair) (res []*T, err error)
}

func NewQueryModeWhereIn(c1, c2 string, c1Valuer, c2Valuer valuer) *QueryMode {
	condition := fmt.Sprintf("(%s, %s) IN (?)", c1, c2)
	return &QueryMode{
		name: "WhereIn",
		run: func(ctx *Ctx, pairs []*Pair) (res []*T, err error) {
			var where []interface{}
			for _, pair := range pairs {
				where = append(where, []interface{}{c1Valuer(pair), c2Valuer(pair)})
			}
			err = ctx.db.Model(&res).Where(condition, pg.InMulti(where...)).Select()
			return res, err
		},
	}
}

func NewQueryModeForeach(c1, c2 string, c1Valuer, c2Valuer valuer) *QueryMode {
	query := fmt.Sprintf(`SELECT * FROM t WHERE %s = $1::%s AND %s = $2::%s`, c1, columnType(c1), c2, columnType(c2))
	return &QueryMode{
		name: "Foreach",
		run: func(ctx *Ctx, pairs []*Pair) (res []*T, err error) {
			stmt, err := ctx.db.Prepare(query)
			defer stmt.Close()
			if err != nil {
				return nil, err
			}
			for _, pair := range pairs {
				var tmp []*T
				_, err := stmt.Query(&tmp, c1Valuer(pair), c2Valuer(pair))
				if err != nil {
					return nil, err
				}
				res = append(res, tmp...)
			}
			return res, nil
		},
	}
}

func columnType(name string) string {
	if strings.HasSuffix(name, "int") {
		return "int"
	}
	if strings.HasSuffix(name, "uuid") {
		return "uuid"
	}
	return "text"
}

func e(err error) {
	if err != nil {
		panic(err)
	}
}
