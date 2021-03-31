package main

import ( 	
	"fmt"
	"encoding/json"
)

type Status struct {
    Status  string
    Node_id string
}

type Meta struct {
    To         string
    From       string
    Id         string
    EntryCount int64
    Size       int64
    Depricated bool
}

type Mydata struct {
    Metadata  Meta `json:"Meta Data"`
    Status []Status
}


// type Meta struct {
	// Information string `json: "1. Information"`
	// SYMBOL string `json: "2. Symbol"`
	// LAST_REFRESHED string `json: "3. Last Refreshed"`
	// OUTPUT_SIZE string `json: "4. Output Size"`
	// TIME_ZONE string `json: "5. Time Zone"`
// }

func main() {

	var dat Mydata

	byt := []byte(`{
		"Meta Data":{
			"id":"2377f625-619b-4e20-90af-9a6cbfb80040",
			"from":"2014-12-30T07:23:42.000Z",
			"to":"2015-01-14T05:11:51.000Z",
			"entryCount":801,
			"size":821472,
			"deprecated":false
		},
		"status":[{
			 "node_id":"de713614-be3d-4c39-a3f8-1154957e46a6",
			 "status":"PUBLISHED"
		}],
		
	}`)

	if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat.Metadata)
}