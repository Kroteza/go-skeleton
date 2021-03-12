package testing

import (
	"time"
)

// Testing model
type Testing struct {
	ID   			string `db:"ID" json:"id"`
	Nama   			string `db:"Nama" json:"nama"`
	Age   			int `db:"Age" json:"age"`
	Balance   		int `db:"Balance" json:"balance"`
	BalanceAfterTax int `db:"Balance_After_Tax" json:"balance_after_tax"`
	Description   	string `db:"Description" json:"description"`
	DateAdded   	time.Time `db:"Date_Added" json:"date_added"`
}
