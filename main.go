package main

import (
	"fmt"
	"encoding/base64"
	
)



func main() {

	fmt.Println(base64.StdEncoding.EncodeToString([]byte("sarath:123")))

	

}

