package main

type Idcard struct {
	Name    string
	Address string
	Phone   int
	Age     int
}

func Purefunc(name, address string, phone, age int) Idcard {
	return Idcard{
		Name:    name,
		Address: address,
		Phone:   phone,
		Age:     age,
	}
}
