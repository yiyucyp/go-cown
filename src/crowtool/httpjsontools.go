package crowtool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ExampleStruct struct {
	ErrorCode int    `json:"errcode"`
	ErrorMsg  string `json:"errmsg"`
}

func ExampleJosn() {
	resp, err := http.Get("http://ww.mydogo.com/api/kaokeacc/get_taokeacc.php")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
	var vv ExampleStruct
	json.Unmarshal(body, &vv)
	fmt.Print(vv)
}
