package main

import "gindemo/go/auth"

func main() {
	init1()

	token, err := auth.GetOAuthToken("code", "state")
	if err != nil {
		panic(err)
	}

	claims, err := auth.ParseJwtToken(token.AccessToken)
	if err != nil {
		panic(err)
	}

	claims.AccessToken = token.AccessToken

}
var CasdoorEndpoint = "http://localhost:8000"
var ClientId = "c7e359e2afc556df1d66"
var ClientSecret = "88fcc9704ad4b1cd68430806ac03b7e3b80ad464"
var JwtSecret = "CasdoorSecret"
var CasdoorOrganization = "casbin-forum"

func init1() {
	auth.InitConfig(CasdoorEndpoint, ClientId, ClientSecret, JwtSecret, CasdoorOrganization , "gindemo")
}