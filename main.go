package main

import (
	"encoding/json"
	"fmt"
)

func findSingletons(value interface{}) {
	switch value.(type) {
	case []interface{}:
		if len(value.([]interface{})) == 1 {
			fmt.Println("1 length array found!", value)
		}
		for _, v := range value.([]interface{}) {
			findSingletons(v)
		}
	case map[string]interface{}:
		for _, v := range value.(map[string]interface{}) {
			findSingletons(v)
		}
	}
}

func removeSingletonsFromJSON(input string) {
	jsonFromInput := json.RawMessage(input)
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonFromInput), &jsonMap)

	if err != nil {
		panic(err)
	}

	findSingletons(jsonMap)

	fmt.Printf("JSON value of without singletons:%s\n", jsonMap)
}

func main() {
	jsonParsed := []byte(`{"path": [{"secret/foo": [{"capabilities": ["read"]}]}]}`)
	removeSingletonsFromJSON(string(jsonParsed))
	fmt.Println(`Should have output {"path": {"secret/foo": {"capabilities": "read"}}}`)
}
