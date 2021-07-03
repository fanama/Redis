package main

import (
	"fmt"
	"net/url"

	"github.com/fanama/Redis/cmd/types"
	"github.com/fanama/Redis/package/axios"
)

func main() {

	url1 := "https://catfact.ninja/fact"
	url11 := "https://httpbin.org/ip"
	url2 := "https://reqbin.com/echo/post/json"
	postData := url.Values{}

	// url.Values{"id": {"8"}}

	resp1 := new(types.Response)
	resp11 := new(types.IP)
	resp2 := new(types.TestPost)

	axios.Get(url1, &resp1)
	axios.Get(url11, &resp11)

	axios.Post(url2, &resp2, postData)

	fmt.Println("resp1 : ", resp1)
	fmt.Println("resp11 : ", resp11)
	fmt.Println("resp2 : ", resp2)

}
