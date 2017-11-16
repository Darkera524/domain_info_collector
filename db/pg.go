package db

import
(
	"github.com/jackc/pgx"
	"domain_info_collector/transfer"
	"bytes"
	"fmt"
)

func SendSearchInfoToPostgresql(searchList []*transfer.SearchInfo)(err error){
	config := initConfig()
	conn, err := pgx.Connect(*config)
	if err != nil {
		fmt.Println("conn err")
		return err
	}
	tx, _ := conn.Begin()
	defer  tx.Rollback()

	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO DsDirSearch (BindId,MessageId,Caller,ObjDN,Filter,RequiredAttributes,TimeCreated,Index,EntriesVisited,EntriesReturned,TimeEnded,KernelTime,UserTime,ProcessID,ThreadID,ProcessorID) VALUES")
	for i,search := range searchList{
		buffer.WriteString("('")
		buffer.WriteString(search.BindId)
		buffer.WriteString("','")
		buffer.WriteString(search.MessageId)
		buffer.WriteString("','")
		buffer.WriteString(search.Caller)
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
		buffer.WriteString("')")
		if i+1 != len(searchList){
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
	buffer.WriteString("INSERT INTO LdapRequest (BindId,MessageId,RemoteSocketString,Udptcp,SearchType,Errmsg,TimeCreated,EncryptionType,RequestType,TimeEnded,KernelTime,UserTime,ProcessID,ThreadID,ProcessorID) VALUES")
	for i,search := range ldapList{
		buffer.WriteString("('")
		buffer.WriteString(search.BindId)
		buffer.WriteString("','")
		buffer.WriteString(search.MessageId)
		buffer.WriteString("','")
		buffer.WriteString(search.RemoteSocketString)
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

func initConfig() (config *pgx.ConnConfig){
	config = &pgx.ConnConfig{
		Host: "127.0.0.1",
		Port: 5433,
		User: "postgres",
		Database: "domain",
		Password: "root",
	}
	return config
}
