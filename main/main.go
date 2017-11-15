package main

import
(
	"domain_info_collector/transfer"
	"domain_info_collector/db"
	"fmt"
)

func main(){
	searchlist, _, _ := transfer.Parse_xml()
	/*for _,ldap := range ldaplist{
		fmt.Println(ldap.MessageId)
		fmt.Println(ldap.BindId)
		fmt.Println(ldap.RemoteSocketString)
		fmt.Println(ldap.Udptcp)
		fmt.Println(ldap.SearchType)
		fmt.Println(ldap.ErrMsg)
		fmt.Println(ldap.TimeCreated)
		//fmt.Println(search.Opcode)
	}*/
	err := db.SendSearchInfoToPostgresql(searchlist)
	if err != nil{
		fmt.Println(err.Error())
	}
	/*for _,ldap := range searchlist{
		fmt.Println(ldap.MessageId)
		fmt.Println(ldap.BindId)
		fmt.Println(ldap.Caller)
		fmt.Println(ldap.ObjDN)
		fmt.Println(ldap.TimeCreated)
		fmt.Println(ldap.RequiredAttributes)
		fmt.Println(ldap.Filter)
		fmt.Println(ldap.Opcode)
	}*/
}