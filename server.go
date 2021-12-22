package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mark-rushakoff/ldapserver"
)

/////////////
// Sample searches you can try against this simple LDAP server:
//
// ldapsearch -H ldap://localhost:3389 -x -b 'dn=test,dn=com'
// ldapsearch -H ldap://localhost:3389 -x -b 'dn=test,dn=com' 'cn=ned'
// ldapsearch -H ldap://localhost:3389 -x -b 'dn=test,dn=com' 'uidnumber=5000'
/////////////

///////////// Run a simple LDAP server
func main() {
	s := ldapserver.NewServer()

	// register Bind and Search function handlers
	handler := ldapHandler{}

	var bindFunc ldapserver.BindFunc = handler.Bind
	var searchFunc ldapserver.SearchFunc = handler.Search
	s.Bind = bindFunc
	s.Search = searchFunc

	// start the server
	listen := ":389"
	log.Printf("Starting example LDAP server on %s", listen)
	if err := s.ListenAndServeTLS(listen, "ssl.crt", "ssl.key"); err != nil {
		log.Fatal("LDAP Server Failed: %s", err.Error())
	}
}

type ldapHandler struct {
}

///////////// Allow anonymous binds only

func (h ldapHandler) Bind(bindDN, bindSimplePw string, conn net.Conn) (ldapserver.LDAPResultCode, error) {
	fmt.Println("####### process bind request!!!!")
	fmt.Println("bindDN=", bindDN)
	return ldapserver.LDAPResultSuccess, nil
	// if bindDN == "LDAP_READER" && bindSimplePw == "123123" {
	// 	return ldapserver.LDAPResultSuccess, nil
	// }
	// return ldapserver.LDAPResultInvalidCredentials, nil
}

///////////// Return some hardcoded search results - we'll respond to any baseDN for testing
func (h ldapHandler) Search(boundDN string, searchReq ldapserver.SearchRequest, conn net.Conn) (ldapserver.ServerSearchResult, error) {
	fmt.Println("####### process search request!!!!")
	fmt.Println("boundDN=", boundDN)
	fmt.Println("searchReq.basedn=", searchReq.BaseDN)
	fmt.Println("searchReq.Filter=", searchReq.Filter)
	entries := []*ldapserver.Entry{
		&ldapserver.Entry{"cn=ned,ou=Users," + searchReq.BaseDN, []*ldapserver.EntryAttribute{
			&ldapserver.EntryAttribute{"sAMAccountName", []string{"ned"}},
			&ldapserver.EntryAttribute{"displayName", []string{"ned"}},
			&ldapserver.EntryAttribute{"cn", []string{"ned"}},

			&ldapserver.EntryAttribute{"givenName", []string{"ned"}},
			&ldapserver.EntryAttribute{"uidNumber", []string{"5000"}},
			&ldapserver.EntryAttribute{"uid", []string{"ned"}},
			&ldapserver.EntryAttribute{"sn", []string{"ned"}},

			&ldapserver.EntryAttribute{"mail", []string{"active@test.com"}},

			&ldapserver.EntryAttribute{"objectClass", []string{"posixAccount", "top", "inetOrgPerson", "sambaSamAccount", "account", "person", "customOC", "organizationalPerson"}},
		}},
	}
	return ldapserver.ServerSearchResult{entries, []string{}, []ldapserver.Control{}, ldapserver.LDAPResultSuccess}, nil
}
