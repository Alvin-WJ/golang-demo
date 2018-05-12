package account

import "fmt"

type Account struct {
	Id string
	Name string
}

const default_account_name= "Dummy"

func Create(id string, name string) *Account{
	fmt.Println("insert account successfully")
	return &Account{Id: id, Name: name}
}

func Query(id string){
	fmt.Println("Query result : account id:'%s'", id)
}

func (account Account) ToString(){
	fmt.Println("Id:" + account.Id + ", Name: " + account.Name)
}



//---------------set,get-----------------

func (account Account) SetId(id string){
	account.Id = id
}

func (account Account) SetName(name string){
	account.Name = name
}
