package store

import (
	"fmt"
)

type Repo struct {
	store *Store
}

func (r *Repo) SinglePay(id int, sum int) error {
	fmt.Printf("single tranzaction: user %d change balance: %d", id, sum)
	return nil
}

func (r *Repo) P2PPay(IDGetter, IDSetter int, sum int) error {
	fmt.Printf("user %d send user %d %d vallents", IDGetter, IDSetter, sum)
	return nil
}

func (r *Repo) GetBalance(ID int) error {
	fmt.Printf("there is balance for %d", ID)
	return nil
}
