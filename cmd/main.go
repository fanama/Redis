package main

import (
	"fmt"
	"net/url"

	"github.com/fanama/Redis/cmd/types"
	"github.com/fanama/Redis/package/axios"
)

func main() {

	url1 := "https://catfact.ninja/fact"
	url2 := "https://reqbin.com/echo/post/json"
	postData := url.Values{}

	// url.Values{"id": {"8"}}

	resp1 := new(types.Response)
	resp2 := new(types.LoginInfos)

	axios.Get(url1, &resp1)

	axios.Post(url2, &resp2, postData)

	fmt.Println(resp1.Fact)
	fmt.Println(resp2.Success)

}
