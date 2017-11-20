package transfer

import (
	"github.com/beevik/etree"
	"strings"
	"bytes"
	"time"
)

type SearchInfo struct {
	BindId string
	MessageId string
	CallerIP string
	CallerPort string
	ObjDN string
	Filter string
	RequiredAttributes string
	TimeCreated string
	Index string
	EntriesVisited string
	EntriesReturned string
	TimeEnded string
	KernelTime string
	UserTime string
	ProcessId string
	ThreadId string
	ProcessorId string
	DomainServer string
	}

type LdapRequest struct {
	BindId string
	MessageId string
	RemoteSocketIP string
	RemoteSocketPort string
	EncryptionType string
	Udptcp string
	SearchType string
	ErrMsg string
	RequestType string
	TimeCreated string
	TimeEnded string
	KernelTime string
	UserTime string
	ProcessId string
	ThreadId string
	ProcessorId string
	DomainServer string
}

func Parse_xml(dir string , host string) ([]*SearchInfo,[]*LdapRequest, error) {
	var searchList []*SearchInfo
	var ldapRequestList []*LdapRequest

	doc := etree.NewDocument()

	//yesterday := get_yesterday()
	//"\\\\idcshare.op.internal.gridsumdissector.com\\idcshare\\wangyiqi\\test05.XML"

	if err := doc.ReadFromFile(dir); err != nil {
		panic(err)
	}

	root := doc.SelectElement("Events")
	for _,event := range root.SelectElements("Event") {
		renderingInfo := event.SelectElement("RenderingInfo")
		opcode := renderingInfo.SelectElement("Opcode").Text()
		eventName := renderingInfo.SelectElement("EventName").Text()

		switch eventName {
			case "DsDirSearch":
				eventData := event.SelectElement("EventData")
				system := event.SelectElement("System")

				var BindId string
				var MessageId string
				var CallerIP string
				var CallerPort string
				var ObjDN string
				var Filter string
				var RequiredAttributes string
				var Index string
				var EntriesVisited string
				var EntriesReturned string

				if opcode == "Start" {
					datas := eventData.SelectElements("Data")
					for _,data := range datas{
						dataValue := data.Attr[0].Value
						switch dataValue {
						case "messageId":
							MessageId = data.Text()
						case "BindId":
							BindId = data.Text()
						case "Caller":
							dataset := strings.Split(data.Text(),":")
							CallerIP = dataset[0]
							if len(dataset) == 2 {
								CallerPort = strings.Split(data.Text(), ":")[1]
							}
						case "ObjDN":
							ObjDN = data.Text()
						case "Filter":
							Filter = data.Text()
						case "RequiredAttributes":
							RequiredAttributes = data.Text()
						}
					}


					search := &SearchInfo{
						BindId:             BindId,
						MessageId:          MessageId,
						CallerIP:             CallerIP,
						CallerPort:  CallerPort,
						ObjDN:              ObjDN,
						Filter:             Filter,
						RequiredAttributes: RequiredAttributes,
						TimeCreated:        system.SelectElement("TimeCreated").SelectAttrValue("SystemTime", "unknow"),
						DomainServer: host,
					}
					searchList = append(searchList, search)
				} else {
					datas := eventData.SelectElements("Data")
					for _,data := range datas {
						dataValue := data.Attr[0].Value
						switch dataValue {
						case "Index":
							Index = data.Text()
						case "EntriesVisited":
							EntriesVisited = data.Text()
						case "EntriesReturned":
							EntriesReturned = data.Text()
						}
					}

					ins_len := len(searchList)
					execution := system.SelectElement("Execution")

					searchList[ins_len-1].Index = Index
					searchList[ins_len-1].EntriesVisited = EntriesVisited
					searchList[ins_len-1].EntriesReturned = EntriesReturned
					searchList[ins_len-1].TimeEnded = system.SelectElement("TimeCreated").SelectAttrValue("SystemTime", "unknow")
					searchList[ins_len-1].KernelTime = execution.SelectAttrValue("KernelTime", "unknow")
					searchList[ins_len-1].UserTime = execution.SelectAttrValue("UserTime", "unknow")
					searchList[ins_len-1].ProcessId = execution.SelectAttrValue("ProcessID", "unknow")
					searchList[ins_len-1].ThreadId = execution.SelectAttrValue("ThreadID", "unknow")
					searchList[ins_len-1].ProcessorId = execution.SelectAttrValue("ProcessorID", "unknow")
				}
			case "LdapRequest":
				eventData := event.SelectElement("EventData")
				system := event.SelectElement("System")

				var BindId string
				var MessageId string
				var RemoteSocketIP string
				var RemoteSocketPort string
				var Udptcp string
				var EncryptionType string
				var SearchType string
				var ErrMsg string
				var RequestType string

				if opcode == "Start" {
					datas := eventData.SelectElements("Data")
					for _,data := range datas{
						dataValue := data.Attr[0].Value
						switch dataValue {
						case "messageId":
							MessageId = data.Text()
						case "BindId":
							BindId = data.Text()
						case "RemoteSocketString":
							dataset := strings.Split(data.Text(),":")
							RemoteSocketIP = dataset[0]
							if len(dataset) == 2 {
								RemoteSocketPort = strings.Split(data.Text(), ":")[1]
							}
						case "EncryptionType":
							EncryptionType = data.Text()
						case "udptcp":
							Udptcp = data.Text()
						}
					}

					ldap := &LdapRequest{
						BindId:             BindId,
						MessageId:          MessageId,
						RemoteSocketIP: RemoteSocketIP,
						RemoteSocketPort: RemoteSocketPort,
						Udptcp:             Udptcp,
						EncryptionType: EncryptionType,
						TimeCreated:        system.SelectElement("TimeCreated").SelectAttrValue("SystemTime", "unknown"),
						DomainServer: host,
					}
					ldapRequestList = append(ldapRequestList, ldap)

				} else {
					datas := eventData.SelectElements("Data")
					for _,data := range datas{
						dataValue := data.Attr[0].Value
						switch dataValue {
						case "SearchType":
							SearchType = data.Text()
						case "ErrMsg":
							ErrMsg = data.Text()
						case "RequestType":
							RequestType = data.Text()
						}
					}
					ldapRequestList[len(ldapRequestList)-1].SearchType = SearchType
					ldapRequestList[len(ldapRequestList)-1].ErrMsg = ErrMsg
					ldapRequestList[len(ldapRequestList)-1].RequestType = RequestType
					ldapRequestList[len(ldapRequestList)-1].TimeEnded = system.SelectElement("TimeCreated").SelectAttrValue("SystemTime", "unknown")

					ins_len := len(ldapRequestList)
					execution := system.SelectElement("Execution")

					ldapRequestList[ins_len-1].KernelTime = execution.SelectAttrValue("KernelTime", "unknow")
					ldapRequestList[ins_len-1].UserTime = execution.SelectAttrValue("UserTime", "unknow")
					ldapRequestList[ins_len-1].ProcessId = execution.SelectAttrValue("ProcessID", "unknow")
					ldapRequestList[ins_len-1].ThreadId = execution.SelectAttrValue("ThreadID", "unknow")
					ldapRequestList[ins_len-1].ProcessorId = execution.SelectAttrValue("ProcessorID", "unknow")
				}
		}

	}


	return searchList,ldapRequestList,nil
}

func get_yesterday() string {
	timestamp := time.Now().Unix()-86400
	now := time.Unix(timestamp, 0)
	date := strings.Split(strings.Split(now.String(), " ")[0], "-")
	var buffer bytes.Buffer
	for _,str := range date{
		buffer.WriteString(str)
	}
	yesterday := buffer.String()
	return yesterday
}

