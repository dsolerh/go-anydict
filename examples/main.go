package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"githug.com/dsolerh/go-anydict"
)

func main() {
	data := `{
      "name": "jon",
      "job-description": {
         "title": "awesome developer",
         "desc": "does some black magic and it runs in his computer"
      }
   }`

	var dict anydict.Dict
	_ = json.Unmarshal([]byte(data), &dict)

	if name, err := anydict.String[string](dict, "name"); err != nil || name != "jon" {
		panic("Something's not working as it should")
	}

	jobDescription, err := anydict.Value[anydict.Dict](dict, "job-description")
	expectedJobDescription := anydict.Dict{
		"title": "awesome developer",
		"desc":  "does some black magic and it runs in his computer",
	}
	if err != nil || !reflect.DeepEqual(jobDescription, expectedJobDescription) {
		panic("I mess it up")
	}

	fmt.Println("It works just fine")
}
