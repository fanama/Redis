# Go-axios

## Install Go-axios

in your project use the following command : 
- > go get github.com/fanama/Redis/package/axios

## Use axios

### 1. Get request (JSON)

1. get the Url

- > url := "https://catfact.ninja/fact"

2. prepare an interface to store the result

```go

type Fact struct {
	Story  string `json:"fact"`
	Length int    `json:"lenght"`
}


```

2. Prepare a variable to store the result
 ```go
get_Response := new(Fact)
```

3. Use the get request withe the parameters

 ```go
     axios.Get(url_Get, get_Response)
```

4. log the data

 ```go
	fmt.Println(get_Response)
```

### full code : example

```go
package main

import (
	"fmt"

	"github.com/fanama/Redis/package/axios"
)

type Fact struct {
	Story  string `json:"fact"`
	Length int    `json:"lenght"`
}

func main() {

	url_Get := "https://catfact.ninja/fact"

	get_Response := new(Fact)

	axios.Get(url_Get, get_Response)

	fmt.Println(get_Response.Story)
}
```