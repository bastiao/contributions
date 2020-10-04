package sourceCode

import (
	"github.com/uber/gonduit"
	"github.com/uber/gonduit/requests"
)

type PhidLookupRequest struct {
	Names            []string `json:"names"`
	requests.Request          // Includes __conduit__ field needed for authentication.
}
type PhidLookupResponse map[string]*struct {
	URI      string `json:"uri"`
	FullName string `json:"fullName"`
	Status   string `json:"status"`
}

func GetSourceCode(client *gonduit.Conn, idRepo string) PhidLookupResponse {
	req := &PhidLookupRequest{
		Names: []string{idRepo},
	}
	var res PhidLookupResponse

	client.Call("phid.lookup", req, &res)

	return res
}
