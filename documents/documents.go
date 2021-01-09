package documents

import (
	"log"

	"github.com/uber/gonduit"
	"github.com/uber/gonduit/requests"
)

type ConstraintsRequest struct {
	Statuses []string `json:"statuses"`
	Query    string   `json:"query"`
}

type AttachmentsRequest struct {
	Content bool `json:"content"`
}

type PhidDocumentRequest struct {
	Constraints      ConstraintsRequest `json:"constraints"`
	Attachments      AttachmentsRequest `json:"attachments"`
	requests.Request                    // Includes __conduit__ field needed for authentication.
}

type ContentRawResponse struct {
	Raw string `json:"raw"`
}

type ContentResponse struct {
	Title   string             `json:"title"`
	Path    string             `json:"path"`
	Content ContentRawResponse `json:"content"`
}

type AttachmentsResponse struct {
	Content ContentResponse `json:"content"`
}

type PhidDocumentResponse struct {
	Id          int                 `json:"id"`
	Phid        string              `json:"phid"`
	Attachments AttachmentsResponse `json:"attachments"`
}

type PhidDocumentDataResponse struct {
	Data []PhidDocumentResponse `json:"data"`
}

func LookForDocument(client *gonduit.Conn, documentQuery *string) PhidDocumentDataResponse {
	conActive := "active"

	constraints := &ConstraintsRequest{
		Statuses: []string{conActive},
		Query:    *documentQuery,
	}
	attachments := &AttachmentsRequest{
		Content: true,
	}

	req := &PhidDocumentRequest{
		Constraints: *constraints,
		Attachments: *attachments,
	}
	var res PhidDocumentDataResponse

	err1 := client.Call("phriction.document.search", req, &res)
	if err1 != nil {
		log.Fatal("Error: ", err1)
	}
	return res
}
