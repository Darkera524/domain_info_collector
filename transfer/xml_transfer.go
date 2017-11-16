package transfer

import (
	"github.com/beevik/etree"
)

type SearchInfo struct {
	BindId string
	MessageId string
	Caller string
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
	}

type LdapRequest struct {
	BindId string
	MessageId string
	RemoteSocketString string
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
}

func Parse_xml() ([]*SearchInfo,[]*LdapRequest, error) {
	var searchList []*SearchInfo
	var ldapRequestList []*LdapRequest

	doc := etree.NewDocument()

	if err := doc.ReadFromFile("\\\\idcshare.op.internal.gridsumdissector.com\\idcshare\\wangyiqi\\test04.xml"); err != nil {
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
				var Caller string
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
							Caller = data.Text()
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
						Caller:             Caller,
						ObjDN:              ObjDN,
						Filter:             Filter,
						RequiredAttributes: RequiredAttributes,
						TimeCreated:        system.SelectElement("TimeCreated").SelectAttrValue("SystemTime", "unknow"),
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
				var RemoteSocketString string
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
							RemoteSocketString = data.Text()
						case "EncryptionType":
							EncryptionType = data.Text()
						case "udptcp":
							Udptcp = data.Text()
						}
					}

					ldap := &LdapRequest{
						BindId:             BindId,
						MessageId:          MessageId,
						RemoteSocketString: RemoteSocketString,
						Udptcp:             Udptcp,
						EncryptionType: EncryptionType,
						TimeCreated:        system.SelectElement("TimeCreated").SelectAttrValue("SystemTime", "unknown"),
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
