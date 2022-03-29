package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	//str := "H4sIAAAAAAAAAE2MQQqDMBBF7zLrLkSsqFcpJcRk2gRkEmYSQcS7d0QXXT7ef3+HKsgmepig65/N0A5jBw+QkLi4WtQITC/IdjMOlyWHRGgYXbD8RR2eQh/W6FBuLGxJsvbK187kFOnED0ckL1qQ/8NQZ3gfP5/kk+SMAAAA"
	str := "H4sIAAAAAAAAAFXOTQ6CQAwF4Lt0LYYZEJFzuNKYyQBVGucHpyXRGO/ubIy66+uXpu8JC2MyNEIHqqzqWm12uoUV8BSTDItkYuiOMNM9b4fJpguaW/rOc6Qg3+goXHPqF6aAzOaMyIYCi3XOYxD+xSGysGHyi7MSE5xWMCdklNyG0bnczLs+XzjLYnwc6Uw4mv6R/fP3n0YrmFGXWhdlVah2r3VX1125WeumVc22Us0BXm8dWfQW9wAAAA=="
	fmt.Println("antes: " + str)
	data, _ := base64.StdEncoding.DecodeString(str)
	fmt.Print("durante: ")
	fmt.Println(data)
	rdata := bytes.NewReader(data)
	r,_ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)
	fmt.Println("depois: "+ string(s))
}
