package listRegions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-gemfire-regions-list")

const (
	ivGFHost       = "host"
	ivGFPort       = "port"
	ivGFListURI    = "uri"
	ivGFListMethod = "method"
	ovGFRegions    = "regions"
)

// MyActivity is a stub for your Activity implementation
type ListActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ListActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *ListActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func callHTTP(host string, port int, method string, uri string) string {
	fmt.Println("Starting the application...")
	//fmt.Println(port)
	p := strconv.Itoa(port)
	httpurl := "http://" + host + ":" + p + uri

	if method != "POST" {
		response, err := http.Get(httpurl)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			return ""
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			//fmt.Println(string(data))
			return string(data)
		}
	} else {
		jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post(httpurl, "application/json", bytes.NewBuffer(jsonValue))

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}

	}

	return ""
	//fmt.Println("Terminating the application...")
	//return data
}

// Eval implements activity.Activity.Eval
func (a *ListActivity) Eval(context activity.Context) (done bool, err error) {

	// do eval

	host := context.GetInput(ivGFHost).(string)
	port := context.GetInput(ivGFPort).(int)
	if port == 0 {
		port = 21
	}
	method := context.GetInput(ivGFListMethod).(string)
	uri := context.GetInput(ivGFListURI).(string)

	res := callHTTP(host, port, method, uri)

	context.SetOutput(ovGFRegions, string(res))

	return true, nil
}
