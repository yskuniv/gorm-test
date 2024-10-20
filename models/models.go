package models

type Human struct {
	ID   uint
	Name string
	Age  uint
}

type Account struct {
	HumanID uint
	Human   Human
	Amount  uint
}
