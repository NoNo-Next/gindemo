package inits

import "gindemo/go/auth"

var CasdoorEndpoint = "http://localhost:8000"
var ClientId = "c7e359e2afc556df1d66"
var ClientSecret = "88fcc9704ad4b1cd68430806ac03b7e3b80ad464"
var JwtSecret = "CasdoorSecret"
var CasdoorOrganization = "casbin-forum"

func init() {
	auth.InitConfig(CasdoorEndpoint, ClientId, ClientSecret, JwtSecret, CasdoorOrganization , "gindemo")
}
