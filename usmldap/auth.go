// This file can be removed once we copied or merge the
// code into appropriate module

package main

import (
       "fmt"
       "github.com/user/usmldap"
       "strings"
       )
      

/* the following functions ldapAuth and getLdapUser to be added into
auth module which nisanth has written

func ldapAuth(user, passwd string) bool {
     var conf usmldap.Config

     conf, err := usmldap.GetLDAPServerDetails("ldapconfig.json")
     if err != nil {
     	glog.Errorf("Unable to fetch server details")
	return false
     }

     // Fetching usm-ldap config values
     ldapserver := conf.Servers[0].Address
     port := conf.Servers[0].Port
     base := conf.Servers[0].Base
     url := usmldap.GetUrl(ldapserver, port)

     // Authenticating user
     // err = usmldap.AuthSSL(url, base, user, passwd)
     err = usmldap.Authenticate(url, base, user, passwd)
     if err != nil {
        glog.Errorf("Authentication unsuccess\nLogin error: %s", err)
	return false
     } else {
        glog.Infof("Ldap user login success!")
	return true
     }
     return false
}


func getLdapUser(rw http.ResponseWriter) bool {
     var conf usmldap.Config

     conf, err := usmldap.GetLDAPServerDetails("ldapconfig.json")
     if err != nil {
     	glog.Errorf("Unable to fetch server details")
	return false
     }

     // Fetching usm-ldap config values
     ldapserver := conf.Servers[0].Address
     port := conf.Servers[0].Port
     base := conf.Servers[0].Base
     url := usmldap.GetUrl(ldapserver, port)

     var bytes []byte
     users := usmldap.GetUsers(url, base)
     if bytes, err = json.Marshal(users); err != nil {
	glog.Errorf("Unable marshal the list of Users", err)
	util.HandleHHttpError(rw, err)
	return false
     }

     rw.Write(bytes)
     return true
}

the above getLdapUser function will provide the following output:
[{"UserId":"admin","UidNumber":"1842200000","CN":"Administrator","SN":"Administrator",
"GivenName":"","DisplayName":"","Mail":""},{"UserId":"tasir","UidNumber":"1842200001",
"CN":"timothy asir","SN":"asir","GivenName":"timothy","DisplayName":"timothy asir",
"Mail":""}, ...]
*/

// the following main code is for testing the usmldap api
func main() {

     var user, passwd string
     var conf usmldap.Config

     user = "testuser1"
     passwd = "testuser1"

     conf, err := usmldap.GetLDAPServerDetails("ldapconfig.json")
     if err != nil {
     	fmt.Printf("Unable to fetch server details!\n")
	return
     }

     // Fetching usm-ldap config values
     ldapserver := conf.Servers[0].Address
     port := conf.Servers[0].Port
     base := conf.Servers[0].Base
     url := usmldap.GetUrl(ldapserver, port)


     // Authenticating user
     err = usmldap.AuthSSL(url, base, user, passwd)
     // err = usmldap.Authenticate(url, base, user, passwd)
     if err != nil {
     	fmt.Printf("Authentication unsuccess\nLogin error: %s", err)
     } else {
	fmt.Printf("Login success!\n")
     }


     // Listing users
     users := usmldap.GetUsers(url, base)
     for _, user := range users {
	for _, attr := range user.Attributes() {
	    fmt.Printf("%s=[%s]\n", attr.Name(), strings.Join(attr.Values(), ", "))
	}
	fmt.Printf("\n")
     }
}
