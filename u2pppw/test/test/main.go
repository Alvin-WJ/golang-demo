package main

import (
	"fmt"
	"u2pppw/test/beans/account"
	"u2pppw/test/beans/invoice"
	"u2pppw/test/beans/payment"
	"u2pppw/test/initialize"
)

var wholeVariable1 = 3
var wholeVariable2 = true
var Test_ID = "Test-123"

func main() {

	initialize.Configurationloading()
	initialize.ConfigurationloadingPrint()

	fmt.Println("-------------------")
	accountTemp := account.Create(Test_ID, "Alvin")
	account.Query(accountTemp.Id)
	accountTemp.ToString()

	fmt.Println("-------------------")
	invoiceTemp := invoice.Create(Test_ID, 20, accountTemp)
	account.Query(invoiceTemp.Id)

	fmt.Println("-------------------")
	paymentTemp := payment.Create(Test_ID,20, accountTemp, invoiceTemp)
	account.Query(paymentTemp.Id)
	fmt.Println(paymentTemp.Status)

}
