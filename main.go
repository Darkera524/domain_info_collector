package main

import
(
	"domain_info_collector/transfer"
	"domain_info_collector/db"
	"fmt"
	"path/filepath"
	"os"
	"io/ioutil"
	"time"
)

func main() {
	hostlist := host_list()
	var interval int64 = 300
	var ticker = time.NewTicker(time.Duration(interval) * time.Second)

	for {
		for _, host := range hostlist {
			main_flow(host)
		}
		<-ticker.C
	}
}

func main_flow(host string){
	insdir := getDir(host)
	searchlist, ldapSearchList, _ := transfer.Parse_xml(insdir,host)
	err := db.SendSearchInfoToPostgresql(searchlist)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.SendLdapSearchInfoToPostgresql(ldapSearchList)
	if err != nil {
		fmt.Println(err.Error())
	}
}

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


func host_list() []string {
	var hostlist []string
	path := "\\\\idcshare.op.internal.gridsumdissector.com\\idcshare\\wangyiqi\\xml\\"
	dir_list, e := ioutil.ReadDir(path)
	if e != nil {
		fmt.Println("read dir error")
		return nil
	}
	for _, v := range dir_list {
		if (v.IsDir()){
			hostlist = append(hostlist, v.Name())
		}
	}

	return hostlist
}

func getDir(path string) string {
	ins := ""
	real_path := "\\\\idcshare.op.internal.gridsumdissector.com\\idcshare\\wangyiqi\\xml\\" + path + "\\"
	err := filepath.Walk(real_path, func(path string, f os.FileInfo, err error) error {
		if ( f == nil ) {return err}
		if f.IsDir() {return nil}
		ins = path
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	fmt.Println(ins)
	return ins
}