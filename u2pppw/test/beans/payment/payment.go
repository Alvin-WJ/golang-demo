package payment

import "fmt"
import "u2pppw/test/beans/account"
import "u2pppw/test/beans/invoice"



type Payment struct {
	Id string
	Amount int
	Account *account.Account
	Invoice *invoice.Invoice
	Status int
}

const (
	created = iota
	processing
	_
	success
	error
)

func Create(id string, amount int, account *account.Account, invoice *invoice.Invoice) *Payment{
	fmt.Println("insert payment successfully")
	var payment = Payment{Id: id, Amount: amount, Account: account, Invoice: invoice}
	payment.Status = success
	return &payment
}

func Query(id string){
	fmt.Println("payment id:%s", id)
}