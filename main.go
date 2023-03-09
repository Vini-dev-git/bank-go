package main

import (
	"bank/accounts"
	"bank/custumers"
	"fmt"
)

func main() {

	client01 := custumers.Holder{"Cliente01", "123.345.678-00", "Developer Go"}
	clientAccount01 := accounts.CheckinAccount{client01, 1234, 123456, 5000}

	fmt.Println(clientAccount01)

}
