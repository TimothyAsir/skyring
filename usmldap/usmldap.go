
package usmldap

// #cgo LDFLAGS: -lldap
// #include <stdio.h>
// #include <ldap.h>
import "C"

import (
	"errors"
	"encoding/json"
	"fmt"
	"github.com/user/openldap"
	"io/ioutil"
	"strings"
)


type Directory struct {
     Address string
     Port int
     Base string
}

type Config struct {
    Servers []Directory
}


type UserData struct {
     UserId string
     UidNumber string
     CN string
     SN string
     GivenName string
     DisplayName string
     Mail string
}


func AuthSSL(ldap_server string, base string,
     username string, passwd string) error {

     var ld *C.LDAP
     var cred C.struct_berval
     
     userBase := "uid=" + username + "," + base
     C.ldap_initialize( &ld, C.CString(ldap_server) )
     cred.bv_val = C.CString(passwd)
     cred.bv_len = 256

     rc, err := C.ldap_sasl_bind_s( ld, C.CString(userBase),
     	 nil, &cred, nil, nil, nil )
     if err != nil {
	fmt.Printf("Authentication error\n")
	fmt.Printf("Error: %s\n", err)
	return err
   	}
     if rc != C.LDAP_SUCCESS {
    	fmt.Printf("Invalid password\n")
	fmt.Printf("Error: %s\n", rc)
	return errors.New("Authentication error!")
	}
     fmt.Printf("\n\nAuthenticated Successfully!\n\n")
     // C.ldap_unbind_ext( unsafe.Pointer(ld) )
     return nil
     }


func Authenticate(url string, base string,
                  user string, passwd string) error {

     ldap, err := openldap.Initialize(url)

     if err != nil {
     	fmt.Printf("Failed to connect the server!\n")
	return err
     }

     ldap.SetOption(openldap.LDAP_OPT_PROTOCOL_VERSION, openldap.LDAP_VERSION3)
     userConnStr := fmt.Sprintf("uid=%s,%s", user, base)

     err = ldap.Bind(userConnStr, passwd)
     if err != nil {
     	return err
     }
     defer ldap.Close()
     return err
}


func GetLDAPServerDetails(file string) (Config, error) {

     var conf Config

     content, err := ioutil.ReadFile(file)
     if err != nil {
     	// TODO: log
        fmt.Print("Failed to read LDAP Json Config file Error:",err)
	return conf, err
     }

     err = json.Unmarshal(content, &conf)
     if err != nil {
         fmt.Print("Parsing Error:", err)
	 return conf, err
     }
     return conf, nil
}


func GetUrl(ldapserver string, port int) string {
     return fmt.Sprintf("ldap://%s:%d/", ldapserver, port)
}


func GetUsers(url string, base string) []UserData {

     var users []UserData
     var user UserData

     ldap, err := openldap.Initialize(url)

     if err != nil {
     	// TODO: Log the error: failed to connect the LDAP/AD server
	return nil
     }

     scope := openldap.LDAP_SCOPE_SUBTREE
     // LDAP_SCOPE_BASE, LDAP_SCOPE_ONELEVEL, LDAP_SCOPE_SUBTREE  
     // filter := "cn=*group*"
     filter := "(objectclass=*)"
     attributes := [] string {"Uid", "UidNumber", "CN", "SN",
     		"Givenname", "Displayname", "mail"}

     rv, err := ldap.SearchAll(base, scope, filter, attributes)

     if err != nil {
     	// TODO: log the error
	fmt.Println(err)
	return nil
     }

     /* TODO: Log base, search, search attribute details */
     // fmt.Printf("# num results : %d\n", rv.Count())
     // fmt.Printf("# base : %s\n", rv.Base())

     for _, entry := range rv.Entries() {
         for _, attr := range entry.Attributes() {
	   switch attr.Name() {
	   case "Uid":
    	         user.UserId = strings.Join(attr.Values(), ", ")
	   case "UidNumber":
	         user.UidNumber = strings.Join(attr.Values(), ", ")
	   case "CN":
	         user.CN = strings.Join(attr.Values(), ", ")
	   case "SN":
	         user.SN = strings.Join(attr.Values(), ", ")
	   case "Givenname":
	         user.GivenName = strings.Join(attr.Values(), ", ")
	   case "Displayname":
	         user.DisplayName = strings.Join(attr.Values(), ", ")
	   case "Mail":
	         user.Mail = strings.Join(attr.Values(), ", ")
	   default:
                 // TODO: Log it saying we do not support this property yet
            }
	}
	users = append(users, user)  
     }
     return users
}


func getUsersOfGroup(url string, base string, search string) []openldap.LdapEntry {
     // “CN=GRoup,OU=Users,DC=Domain,DC=com”
     ldap, err := openldap.Initialize(url)

     if err != nil {
     	// TODO: Log the error
	return nil
     }

     scope := openldap.LDAP_SCOPE_SUBTREE
     // LDAP_SCOPE_BASE, LDAP_SCOPE_ONELEVEL, LDAP_SCOPE_SUBTREE  
     attributes := [] string {"uid", "uidNumber", "gidNumber",
     		"cn", "sn", "givenname", "displayname", "mail"}

     rv, err := ldap.SearchAll(base, scope, search, attributes)

     if err != nil {
     	// TODO: log the error
	fmt.Println(err)
	return nil
     }

     /* TODO: Log base, search, search attribute details */
     fmt.Printf("# num results : %d\n", rv.Count())
     fmt.Printf("# base : %s\n", rv.Base())

     //fmt.Printf(rv.Entries())
     return rv.Entries()
}
