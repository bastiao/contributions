package documents

import (
	"log"

	"github.com/uber/gonduit"
	"github.com/uber/gonduit/requests"
)

type PhidDocumentRequest struct {
	Constraints string `json:"constraints"`
	Attachments string `json:"attachments"`

	requests.Request // Includes __conduit__ field needed for authentication.
}
type PhidDocumentResponse map[string]*struct {
	Phid string `json:"phid"`
	Id   string `json:"id"`
}

func LookForDocument(client *gonduit.Conn) []PhidDocumentResponse {
	constraints := "{\"statuses\": [ \"active\"]}"
	attachments := "{\"content\": true}"

	req := &PhidDocumentRequest{
		Constraints: constraints,
		Attachments: attachments,
	}
	var res []PhidDocumentResponse

	err1 := client.Call("phriction.document.search", req, &res)
	if err1 != nil {
		log.Fatal("Error: ", err1)
	}
	return res
}
