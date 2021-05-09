package main

import (
	"fmt"
	"regexp"
	"strings"
	"html/template"
	"bytes"
	"bufio"
    "os"
	"github.com/bastiao/contributions/config"
	"github.com/bastiao/contributions/documents"
	"github.com/uber/gonduit"
	"github.com/uber/gonduit/core"
)

// Verify if a string match a specific criterias with filters such as str1|str2
func matchSequenceStr(strEntry *string, filter *string) bool {
	for _, line := range strings.Split(*filter, "|") {
		if strings.Contains(*strEntry, line) {
			return true
		}
	}
	return false

}

// Management options for documents
func ShowDocuments(confFile *config.ConfGoPath, documentList *bool, documentQuery *string,
	documentFilter *string, documentFilterBoolean *string, documentTitle *string,
	documentShowAll *bool, documentsRawRegexContent *string) {
	fmt.Println("\n⭐ Starting pha-go with documents command.")
	fmt.Println("Options: ")
	fmt.Println("📃 Document list: ", *documentList)
	fmt.Println("🏧 Document query (--query): ", *documentQuery)
	fmt.Println("🏧 Document filter: (--filter)", *documentFilter)
	fmt.Println("🏧 Document match (--match): ", *documentFilterBoolean)
	fmt.Println("🏧 Document title: (--title)", *documentTitle)
	fmt.Println("🏧 Document show all (--show-all): ", *documentShowAll)
	fmt.Println("🏧 Document Raw Regex(--raw-regex): ", *documentsRawRegexContent)

	client, err := gonduit.Dial(confFile.PhaConf.Endpoint,
		&core.ClientOptions{APIToken: confFile.PhaConf.Token})
	_ = err
	client.Connect()

	docs := documents.LookForDocument(client, documentQuery)
	for _, v := range docs.Data {

		reTitle := regexp.MustCompile(*documentTitle)
		matchTitle := reTitle.FindStringSubmatch(v.Attachments.Content.Title)

		if len(matchTitle) > 1 && matchSequenceStr(&matchTitle[1], documentFilterBoolean) {
			if len(*documentsRawRegexContent) == 0 {
				fmt.Println("\t🐊 Customer: ", v.Attachments.Content.Title)
			}
			for _, line := range strings.Split(strings.TrimSuffix(v.Attachments.Content.Content.Raw, "\n"), "\n") {
				if *documentShowAll {
					if len(*documentsRawRegexContent) > 1 {
						reRaw := regexp.MustCompile(*documentsRawRegexContent)
						matchRaw := reRaw.FindStringSubmatch(line)
						if len(matchRaw) > 1 {
							fmt.Println("\t🐊 Path: ", v.Attachments.Content.Path)
							fmt.Println("\tℹ ", matchRaw[1])
						}
					} else {
						fmt.Println(line)
					}

				}

				re := regexp.MustCompile("^## (.*?)$")
				re2 := regexp.MustCompile("^\\=\\=\\ (.*?) \\=\\=$")

				match := re.FindStringSubmatch(line)

				if len(match) > 1 && strings.Contains(match[1], *documentFilter) == false {
					break

				} else {
					match2 := re2.FindStringSubmatch(line)
					if len(match2) > 1 && strings.Contains(match2[1], *documentFilter) == false {
						break
					}
				}

			}
		}
	}
}
func ConvertTemplate(a string,b map[string]interface{}) string {
	tmpl := template.Must(template.New("email.tmpl").Parse(a))
	buf := &bytes.Buffer{}
        err := tmpl.Execute(buf, b)
	if err != nil {
		panic(err)
        }
	s := buf.String()
	return s
}


// Edit Documents
func EditDocuments(confFile *config.ConfGoPath, documentsId *string, documentsAsk *bool) {
	fmt.Println("\n⭐ Starting pha-go with edit documents command.")
	fmt.Println("Options: ")
	fmt.Println("📃 Document id: ", *documentsId)

	
	client, err := gonduit.Dial(confFile.PhaConf.Endpoint,
		&core.ClientOptions{APIToken: confFile.PhaConf.Token})
	_ = err
	client.Connect()

	var customers documents.Customers
	customers = documents.LoadMappings()
	data := make(map[string]interface{}, 9)
	data["task"] = "Email"
	data["slug"] = "none"
	data["date"] = "2021-05-08"
	data["customer"] = "Helpdesk"
	data["support"] = "Support"
	data["problem"] = "This problem\nneeds a solution."
	data["solution"] = "This problem\nHappens sometimes."
	data["time"] = "1h/1h/1h"
	data["severity"] = "Low"
	if (*documentsAsk==true){
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter date (e.g. 2021-05-08): ")
    	data["date"], _ = reader.ReadString('\n')
		fmt.Print("Enter Task (e.g. Email, SMS, Task): ")
		var taskStr string
    	taskStr, _ = reader.ReadString('\n')
		var taskStrArr []string
		taskStrArr = strings.SplitN(taskStr, "\n",2)
		if (taskStrArr != nil && taskStrArr[0]!=""){
			data["task"] = taskStrArr[0]
		}


		fmt.Print("Enter customer (e.g. Who?): ")
		var customerStr string
    	customerStr, _ = reader.ReadString('\n')
		var customerStrArr []string
		customerStrArr = strings.SplitN(customerStr, "\n",2)

		data["customer"] = customerStrArr[0]
		fmt.Print("Enter support (e.g. Support): ")
    	data["support"], _ = reader.ReadString('\n')
		fmt.Print("Enter problem (e.g. My issue): ")
    	data["problem"], _ = reader.ReadString('\n')
		fmt.Print("Enter solution: ")
    	data["solution"], _ = reader.ReadString('\n')
		fmt.Print("Enter time  (e.g. Time spent/to response/to solve): ")
    	data["time"], _ = reader.ReadString('\n')
		fmt.Print("Enter severity (Low, Medium, High): ")
    	data["severity"], _ = reader.ReadString('\n')
	}
	
	
	for k, v := range customers {
		if (k == *documentsId) {
			fmt.Printf("%v -> name: %v, \n", k, v.Slug, v.Team)
			docs := documents.LookForDocumentByPath(client, &v.Slug)
			for _, v2 := range docs.Data {
				
				//fmt.Println("📃 Document id: ", v2.Attachments.Content.Content)
				var templateStr string 
				templateStr = documents.ReadTemplate("support/template.txt")
				templateStr = ConvertTemplate(templateStr,data)
				var sStr []string
				sStr = strings.SplitN(templateStr, "\n",2)
	
				templateStr = sStr[1] +"\n\n" + v2.Attachments.Content.Content.Raw
				//fmt.Println("📃 Document : ", templateStr)
				documents.EditDoc(client, &templateStr, &v.Slug)
			}
		}
        
    }
	
	

}
