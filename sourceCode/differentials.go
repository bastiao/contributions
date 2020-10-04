package sourceCode

import (
	"log"

	"github.com/uber/gonduit"
	"github.com/uber/gonduit/requests"
)

type differentialRequest struct {
	Status           string `json:"status"`
	IDs              string `json:"id"`
	requests.Request        // Includes __conduit__ field needed for authentication.
}

type DifferentialResponse struct {
	ID             string `json:"phid"`
	URI            string `json:"uri"`
	StatusName     string `json:"statusName"`
	RepositoryPHID string `json:"repositoryPHID"`
	Branch         string `json:"branch"`
}

func LookForDifferential(client *gonduit.Conn) []DifferentialResponse {
	status := "status-open"
	req := &differentialRequest{
		Status: status,
	}
	var res []DifferentialResponse

	err1 := client.Call("differential.query", req, &res)
	if err1 != nil {
		log.Fatal("Error: ", err1)
	}
	return res
}

func LookForDifferentialById(client *gonduit.Conn, revision string) []DifferentialResponse {
	status := "status-open"
	req := &differentialRequest{
		Status: status,
	}
	var res []DifferentialResponse

	err1 := client.Call("differential.query", req, &res)
	if err1 != nil {
		log.Fatal("Error: ", err1)
	}
	return res
}
