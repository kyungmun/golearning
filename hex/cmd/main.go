package main

import (
	"fmt"
	"goproject/hex/internal/adapters/core/arithmetic"
)

func main() {

	// dbaseDriver := os.Getenv("DB_DRIVER")
	// dsourceName := os.Getenv("DB_NAME")
	// dbaseAdapter, err := db.NewAdapter(dbaseDriver, dsourceName)
	// if err != nil {
	// 	log.Fatalf("failed to initial dbase connection : %v", err)
	// }
	// defer dbaseAdapter.CloseDBConnection()

	// core := arithmetic.NewAdapter()

	// appAdapter := api.NewAdapter(dbaseAdapter, core)

	// result, err := appAdapter.GetAddition(1, 2)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(result)

	result, err := arithmetic.NewAdapter().Addition(3, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
