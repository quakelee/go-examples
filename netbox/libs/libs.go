package libs

import (
	"fmt"
	"net/http"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/quakelee/go-netbox/netbox/client"
)

const authHeaderName = "Authorization"
const authHeaderFormat = "Token %s"

// NetboxConnect is a function to create netbox api connection
func NetboxConnect(host string, port int, token, scheme string, skipverify bool) (*client.NetBox, error) {
	var (
		tc *http.Client
		rt *runtimeclient.Runtime
	)
	if scheme == "https" {
		if skipverify {
			tc, _ = runtimeclient.TLSClient(runtimeclient.TLSClientOptions{
				InsecureSkipVerify: true,
			})
		}
		rt = runtimeclient.NewWithClient(fmt.Sprintf("%s:%d", host, port), client.DefaultBasePath, []string{scheme}, tc)
	} else {
		rt = runtimeclient.New(fmt.Sprintf("%s:%d", host, port), client.DefaultBasePath, []string{scheme})
	}

	rt.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, token))
	return client.New(rt, nil), nil
}
