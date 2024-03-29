package gitler

import (
	"errors"
	"fmt"
	"strings"
)

// Tag
type Tag struct {
	Hash      string
	Author    string
	Email     string
	Version   string
	Timestamp int
}

func (tag *Tag) String() string {
	return fmt.Sprintf("\n+ Hash: %s\n| Author: %s <%s>\n|"+
		"| Timestamp: %d\n| Email: %s\n", tag.Hash, tag.Author,
		tag.Timestamp, tag.Email)
}

// Parse Git Tag
func ParseTag(path, line string, filter string) (*Tag, error) {
	if contains(strings.Split(filter, "|"), line) && // contains filter
		strings.Contains(line, "tag: ") { // also is also a git tag
		// Contains the filters, so let's parse
		tagTmp := strings.Split(line, "tag: ")
		tagStr := strings.ReplaceAll(tagTmp[1], ")'", "")
		tagStr = strings.Split(tagStr, ",")[0]
		fmt.Println("Tag: ", tagStr)
		// TODO: parse it
		tag := &Tag{}
		return tag, nil
	}

	tag := &Tag{}
	return tag, errors.New("No tag for filter or available")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(e, a) {
			return true
		}
	}
	return false
}
