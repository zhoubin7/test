package main

import "fmt"

type Accounts struct {
	Name , PhoneNum string		//账户姓名,手机号
}

func writeChannel(ac chan Accounts){
	ac1 := Accounts{
		Name:"zhoubin",
		PhoneNum:"13269898751",
	}
	ac2 := Accounts{
		Name:"lilong",
		PhoneNum:"18311053771",
	}
	ac <- ac1
	ac <- ac2
	close(ac)
}

func main(){
	ac := make(chan Accounts)
    go writeChannel(ac)
	for v := range ac {
		fmt.Println(v)
	}
}

