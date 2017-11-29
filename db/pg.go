package db

import
(
	"github.com/jackc/pgx"
	"domain_info_collector/transfer"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type single_normal_search struct{
	callerIp string
	searchType string
	domainServer string
	date string
}

func SendSearchInfoToPostgresql(searchList []*transfer.SearchInfo)(err error){
	info_list := query_data(searchList)

	config := initConfig()
	conn, err := pgx.Connect(*config)
	if err != nil {
		fmt.Println("conn err")
		return err
	}
	tx, _ := conn.Begin()
	defer  tx.Rollback()

	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO DsDirSearch (BindId,MessageId,CallerIP,CallerPort,ObjDN,Filter,RequiredAttributes,TimeCreated,Index,EntriesVisited,EntriesReturned,TimeEnded,KernelTime,UserTime,ProcessID,ThreadID,ProcessorID,DomainServer) VALUES")
	for i,search := range searchList{
		buffer.WriteString("('")
		buffer.WriteString(search.BindId)
		buffer.WriteString("','")
		buffer.WriteString(search.MessageId)
		buffer.WriteString("','")
		buffer.WriteString(search.CallerIP)
		buffer.WriteString("','")
		buffer.WriteString(search.CallerPort)
		buffer.WriteString("','")
		buffer.WriteString(search.ObjDN)
		buffer.WriteString("','")
		buffer.WriteString(search.Filter)
		buffer.WriteString("','")
		buffer.WriteString(search.RequiredAttributes)
		buffer.WriteString("','")
		buffer.WriteString(search.TimeCreated)
		buffer.WriteString("','")
		buffer.WriteString(search.Index)
		buffer.WriteString("','")
		buffer.WriteString(search.EntriesVisited)
		buffer.WriteString("','")
		buffer.WriteString(search.EntriesVisited)
		buffer.WriteString("','")
		buffer.WriteString(search.TimeEnded)
		buffer.WriteString("','")
		buffer.WriteString(search.KernelTime)
		buffer.WriteString("','")
		buffer.WriteString(search.UserTime)
		buffer.WriteString("','")
		buffer.WriteString(search.ProcessId)
		buffer.WriteString("','")
		buffer.WriteString(search.ThreadId)
		buffer.WriteString("','")
		buffer.WriteString(search.ProcessorId)
		buffer.WriteString("','")
		buffer.WriteString(search.DomainServer)
		buffer.WriteString("')")
		if i+1 != len(searchList){
			buffer.WriteString(",")
		}
	}
	sql := buffer.String()

	fmt.Println(sql)
	_,err = tx.Exec(sql)
	if err != nil {
		fmt.Println("exec err1")
		return err
	}

	sql = generate_write_sql(info_list)

	_,err = tx.Exec(sql)
	if err != nil {
		fmt.Println("exec err2")
		return err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("commit err")
		return err
	}
	return nil
}

func SendLdapSearchInfoToPostgresql(ldapList []*transfer.LdapRequest)(err error){
	config := initConfig()
	conn, err := pgx.Connect(*config)
	if err != nil {
		fmt.Println("conn err")
		return err
	}
	tx, _ := conn.Begin()
	defer  tx.Rollback()

	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO LdapRequest (BindId,MessageId,RemoteSocketIP,RemoteSocketPort,Udptcp,SearchType,Errmsg,TimeCreated,EncryptionType,RequestType,TimeEnded,KernelTime,UserTime,ProcessID,ThreadID,ProcessorID,DomainServer) VALUES")
	for i,search := range ldapList{
		buffer.WriteString("('")
		buffer.WriteString(search.BindId)
		buffer.WriteString("','")
		buffer.WriteString(search.MessageId)
		buffer.WriteString("','")
		buffer.WriteString(search.RemoteSocketIP)
		buffer.WriteString("','")
		buffer.WriteString(search.RemoteSocketPort)
		buffer.WriteString("','")
		buffer.WriteString(search.Udptcp)
		buffer.WriteString("','")
		buffer.WriteString(search.SearchType)
		buffer.WriteString("','")
		buffer.WriteString(search.ErrMsg)
		buffer.WriteString("','")
		buffer.WriteString(search.TimeCreated)
		buffer.WriteString("','")
		buffer.WriteString(search.EncryptionType)
		buffer.WriteString("','")
		buffer.WriteString(search.RequestType)
		buffer.WriteString("','")
		buffer.WriteString(search.TimeEnded)
		buffer.WriteString("','")
		buffer.WriteString(search.KernelTime)
		buffer.WriteString("','")
		buffer.WriteString(search.UserTime)
		buffer.WriteString("','")
		buffer.WriteString(search.ProcessId)
		buffer.WriteString("','")
		buffer.WriteString(search.ThreadId)
		buffer.WriteString("','")
		buffer.WriteString(search.ProcessorId)
		buffer.WriteString("','")
		buffer.WriteString(search.DomainServer)
		buffer.WriteString("')")
		if i+1 != len(ldapList){
			buffer.WriteString(",")
		}
	}
	sql := buffer.String()

	fmt.Println(sql)
	_,err = tx.Exec(sql)
	if err != nil {
		fmt.Println("exec err")
		return err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("commit err")
		return err
	}


	return nil
}

func query_data(origin_list []*transfer.SearchInfo) ([]*single_normal_search){
	var info_list []*single_normal_search

	for _,ins := range origin_list{
		var single single_normal_search
		if ignore(*ins){
			continue
		} else {
			single.callerIp = ins.CallerIP
			single.domainServer = ins.DomainServer
			single.date = strings.Split(ins.TimeCreated,"T")[0]
		}
		if isSearchingUser(*ins){
			single.searchType = "User"
		} else if isSearchingGroup(*ins){
			single.searchType = "Group"
		} else if isSearchingByObjectUid(*ins){
			single.searchType = "ObjectSid"
		} else {
			single.searchType = "Other"
		}
		info_list = append(info_list, &single)
	}


	return info_list
}

func ignore(search transfer.SearchInfo) bool {
	//callerip正则表达式
	ignoreReg := "(Internal|SAM|LSA|KCC|NTDSAPI|\\[)"

	if match,_ := regexp.MatchString(ignoreReg, search.CallerIP);match{
		return true
	}
	return false
}

func isSearchingUser(search transfer.SearchInfo) bool{
	//filter正则表达式
	filterReg := ".*(objectClass=user|displayName=|sAMAccountName=).*$"

	//requiredattributes正则表达式
	requiredAttributesReg := "(userAccountControl|memberOf|primaryGroupID|mail|userPrincipalName|facsimileTelephoneNumber|sn|givenName|pwdLastSet)"

	//objdn正则表达式
	objdnReg := "(OU=公司成员,OU=gridsum-members,DC=gridsum,DC=com)"

	if match,_ := regexp.MatchString(filterReg, search.Filter);match{
		return true
	} else if match,_ := regexp.MatchString(requiredAttributesReg, search.RequiredAttributes);match{
		return true
	} else if match,_ := regexp.MatchString(objdnReg, search.ObjDN);match{
		return true
	}

	return false
}

func isSearchingGroup(search transfer.SearchInfo) bool {
	//filter正则表达式
	filterReg := "(objectClass=group)"

	//requiredattributes正则表达式 先行判断了memberOf，故此处不必担心
	requiredAttributesReg := "(member)"

	//objdn正则表达式
	objdnReg := "(ou=group,ou=gridsum-members,dc=gridsum,dc=com)"

	if match,_ := regexp.MatchString(filterReg, search.Filter);match{
		return true
	} else if match,_ := regexp.MatchString(requiredAttributesReg, search.RequiredAttributes);match{
		return true
	} else if match,_ := regexp.MatchString(objdnReg, search.ObjDN);match{
		return true
	}

	return false
}

func isSearchingByObjectUid(search transfer.SearchInfo) bool {
	//objectUid正则表达式
	objectUidReg := "(objectSid)"

	if match,_ := regexp.MatchString(objectUidReg, search.Filter);match{
		return true
	}

	return false
}

func generate_write_sql(list []*single_normal_search) string {

	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO searchtype (callerip,searchtype,domainserver,date) VALUES ")

	for i,search := range list{
		buffer.WriteString("('")
		buffer.WriteString(search.callerIp)
		buffer.WriteString("','")
		buffer.WriteString(search.searchType)
		buffer.WriteString("','")
		buffer.WriteString(search.domainServer)
		buffer.WriteString("','")
		buffer.WriteString(search.date)
		buffer.WriteString("')")
		if i+1 != len(list){
			buffer.WriteString(",")
		}
	}
	sql := buffer.String()

	return sql
}

func initConfig() (config *pgx.ConnConfig){
	config = &pgx.ConnConfig{
		Host: "127.0.0.1",
		Port: 5433,
		User: "postgres",
		Database: "postgres",
		Password: "root",
	}
	return config
}
