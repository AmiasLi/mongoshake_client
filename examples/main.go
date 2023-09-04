package main

import (
	"fmt"
	"github.com/AmiasLi/mongoshake_client"
)

func main() {

	// new client with url of mongoshake Incremental API
	// Progress API only support when full sync, not support incremental sync after restart mongoshake
	client, _ := mongoshake.New("http://192.168.124.65:9100")

	config, _ := client.GetConfig()
	fmt.Println(config)

	repl, _ := client.GetRepl()
	fmt.Println(repl)

	executor, _ := client.GetExecutor()
	fmt.Println(executor)

}
