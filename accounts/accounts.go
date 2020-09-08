package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount dsds
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
