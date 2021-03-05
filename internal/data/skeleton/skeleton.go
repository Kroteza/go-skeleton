package skeleton

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"go-skeleton/pkg/errors"
	testingEntity "go-skeleton/internal/entity/testing"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt
	}

	// Key - value map for query selector
	statement struct {
		key   string
		query string
	}
)

// Query constants
const (
	getAllData = `GetAllUsers`
	qGetAllData = `SELECT * FROM testing_tsepty`

	getDataByID = `GetDataByID`
	qGetDataByID = `SELECT * FROM testing_tsepty
					WHERE ID = ?`
				//	IFNULL(Nama, "kosong")AS Nama IFNULL(Nama, 0)AS Nama
)

// Add queries to key value order to be initialized as prepared statements
var (
	readStmt = []statement{
		{getAllData, qGetAllData},
		{getDataByID, qGetDataByID},
	}
)
//test

// New returns data instance
func New(db *sqlx.DB) Data {
	d := Data{
		db: db,
	}

	d.initStmt()
	return d
}

// Initialize queries as prepared statements
func (d *Data) initStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)

	)


	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize statement key %v, err : %v", v.key, err)
		}
	}

	d.stmt = stmts
}

// GetAllData ...
func (d Data) GetAllData(ctx context.Context) ([]testingEntity.Testing, error) {
	var (
		singleTesting testingEntity.Testing
		testingArray  []testingEntity.Testing
	)

	rows, err := d.stmt[getAllData].QueryxContext(ctx)
	//ini dia buat rows untuk manggil query getAllData untuk disimpen disitu sementara sebelum dimasukin ke testingArray

	for rows.Next(){
		if err:= rows.StructScan(&singleTesting); err != nil {
			//dia akan scan apabila ada row yang masih belum di scan dan kemudian dimasukan ke singleTesting
			return testingArray, errors.Wrap(err, "[DATA][SKELETON][GetAllData] ")
		}
		testingArray = append(testingArray, singleTesting)
		//data yang di scan dan dimasukan ke singleTesting tadi, akan di append ke testingArray
	}
	return testingArray, err
}

	//dia ga perlu pakai array untuk ambilnya
	//kemdudian di querynya kan ada "?", nah itu harus diisi di data
	//(ctx context.Context, ID string) (testingEntity.Testing, error)

// GetDataByID ...
func (d Data) GetDataByID(ctx context.Context, ID string) (testingEntity.Testing, error) {
	var singleTesting testingEntity.Testing

	err := d.stmt[getDataByID].QueryRowxContext(ctx, ID).StructScan(&singleTesting)
	if err != nil{
		return singleTesting, errors.Wrap(err, "[DATA][SKELETON][GetDataByID] ")
	}
	
return singleTesting, err

}