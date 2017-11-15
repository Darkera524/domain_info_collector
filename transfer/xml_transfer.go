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
	Opcode string
}

type LdapRequest struct {
	BindId string
	MessageId string
	RemoteSocketString string
	Udptcp string
	SearchType string
	ErrMsg string
	TimeCreated string
}

func Parse_xml() ([]*SearchInfo,[]*LdapRequest, error) {
	var searchList []*SearchInfo
	var ldapRequestList []*LdapRequest

	doc := etree.NewDocument()

	if err := doc.ReadFromFile("\\\\idcshare.op.internal.gridsumdissector.com\\idcshare\\wangyiqi\\test05.xml"); err != nil {
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
						Opcode:             "start",
					}
					searchList = append(searchList, search)
				} else {
					/*search := &SearchInfo{
						BindId:             eventData.SelectElement("BindId").Text(),
						MessageId:          eventData.SelectElement("messageId").Text(),
						Caller:             eventData.SelectElement("Caller").Text(),
						ObjDN:              eventData.SelectElement("ObjDN").Text(),
						Filter:             eventData.SelectElement("Filter").Text(),
						RequiredAttributes: eventData.SelectElement("RequiredAttributes").Text(),
						TimeCreated:        system.SelectElement("TimeCreated").SelectAttrValue("SystemTime", "unknow"),
						Opcode:             "start",
					}
					searchList = append(searchList, search)*/
				}
			case "LdapRequest":
				/*eventData := event.SelectElement("EventData")
				system := event.SelectElement("System")

				var BindId string
				var MessageId string
				var RemoteSocketString string
				var Udptcp string
				var SearchType string
				var ErrMsg string

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
						case "udptcp":
							Udptcp = data.Text()
						}
					}

					ldap := &LdapRequest{
						BindId:             BindId,
						MessageId:          MessageId,
						RemoteSocketString: RemoteSocketString,
						Udptcp:             Udptcp,
						SearchType:         "",
						ErrMsg:             "",
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
						}
					}
					ldapRequestList[len(ldapRequestList)-1].SearchType = SearchType
					ldapRequestList[len(ldapRequestList)-1].ErrMsg = ErrMsg
				}*/
		}

	}


	return searchList,ldapRequestList,nil
}
