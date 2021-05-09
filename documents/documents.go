package documents

import (
	"log"
	"gopkg.in/yaml.v2"
    "fmt"
    "bytes"
	"os"
	"bufio"

	"github.com/uber/gonduit"
	"github.com/uber/gonduit/requests"
)

type WikiSlugRequest struct {
	Slug    string   `json:"slug"`
	Content    string   `json:"content"`
	requests.Request                    // Includes __conduit__ field needed for authentication.

}


type ConstraintsRequest struct {
	Statuses []string `json:"statuses"`
	Paths []string `json:"paths"`
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
type WikiPhidDocumentResponse struct {
	Phid        string              `json:"phid"`
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


func LookForDocumentByPath(client *gonduit.Conn, documentPath *string) PhidDocumentDataResponse {
	conActive := "active"

	constraints := &ConstraintsRequest{
		Statuses: []string{conActive},
		Paths: []string{*documentPath},
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

// Edit Document
func EditDoc(client *gonduit.Conn, content *string, slug *string) WikiPhidDocumentResponse {
	fmt.Println("📃 EditDoc content: ", *content)
	fmt.Println("📃 EditDoc slug: ", *slug)
	req := &WikiSlugRequest{
		Slug:    *slug,
		Content: *content,
	}
	
	var res WikiPhidDocumentResponse

	err1 := client.Call("phriction.edit", req, &res)
	if err1 != nil {
		log.Fatal("Error: ", err1)
	}
	return res
}

func ReadTemplate(pathTemplate string) string {

    f, err := os.Open(pathTemplate)

    if err != nil {
        fmt.Println(err)
     }

    defer f.Close()

    scanner := bufio.NewScanner(f)
	var content string 
	content = ""
    for scanner.Scan() {

        content = content + scanner.Text() + "\n"
		
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
	return content

}

type Customer struct {
    Slug string `yaml:"slug"`
    Domain string `yaml:"domain"`
    Team string `yaml:"team"`
	Name string 
}


type Customers map[string]*Customer
func (s *Customers) Unmarshal(data []byte) error {
    err := yaml.NewDecoder(bytes.NewReader(data)).Decode(s)
    if err != nil {
        return err
    }
    for k, v := range *s {
        v.Name= k
    }
    return nil
}

// Load Mappings
func LoadMappings() Customers {
    var yamlFile string
	yamlFile = ReadTemplate("support/mapping.yaml")

    var customers Customers = map[string]*Customer{}
    err :=customers.Unmarshal([]byte(yamlFile))
    if err != nil {
        log.Fatal("failed to decode: %v", err)
    }

    /*for k, v := range customers {
        fmt.Printf("%v -> name: %v, \n", k, v.Slug, v.Team)
    }*/
	return customers

}