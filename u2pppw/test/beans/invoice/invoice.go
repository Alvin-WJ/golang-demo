package invoice

import "fmt"
import (
	"u2pppw/test/beans/account"
)



type Invoice struct {
	Id string
	Amount int
	Account *account.Account
}


func Create(id string, amount int, account *account.Account) *Invoice{
	fmt.Println("insert invoice successfully")
	return &Invoice{Id: id, Amount: amount, Account: account}
}

func Query(id string){
	fmt.Println("invoice id:%s", id)
}