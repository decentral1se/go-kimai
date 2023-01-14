package main

import (
	"decentral1se/go-kimai/client"
	"decentral1se/go-kimai/client/customer"
	"fmt"
	"log"
	"os"

	httpTransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func main() {
	user := os.Getenv("GO_KIMAI_USER")
	token := os.Getenv("GO_KIMAI_TOKEN")
	authUser := httpTransport.APIKeyAuth("X-AUTH-USER", "header", user)
	authToken := httpTransport.APIKeyAuth("X-AUTH-TOKEN", "header", token)
	auth := httpTransport.Compose(authUser, authToken)

	domain := os.Getenv("GO_KIMAI_DOMAIN")
	transport := httpTransport.New(domain, "", []string{"https"})
	cl := client.New(transport, strfmt.Default)

	csParams := customer.NewGetAPICustomersParams()
	csParams.WithOrder("ASC")
	csParams.WithOrderBy("name")

	resp, err := cl.Customer.GetAPICustomers(csParams, auth)
	if err != nil {
		log.Fatal(err)
	}

	for _, cs := range resp.Payload {
		fmt.Println(*cs.Name)
	}
}
