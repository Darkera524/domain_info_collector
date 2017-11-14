package main

import
(
	"domain_info_collector/transfer"
	"fmt"
)

func main(){
	searchList, _, _ := transfer.Parse_xml()
	for _,search := range searchList{
		fmt.Println(search.MessageId)
		fmt.Println(search.BindId)
		fmt.Println(search.Caller)
		fmt.Println(search.ObjDN)
		fmt.Println(search.Filter)
		fmt.Println(search.RequiredAttributes)
		fmt.Println(search.TimeCreated)
		fmt.Println(search.Opcode)
	}
}