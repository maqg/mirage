package api

import "fmt"

func APIAddHost(paras *ApiParas) *ApiResponse {
	fmt.Printf("running in APIAddHost\n")

	resp := new(ApiResponse)
	resp.Error = 0

	return resp
}

func APIDeleteHost(paras *ApiParas) *ApiResponse {
	fmt.Printf("running in APIDeleteHost\n")
	resp := new(ApiResponse)
	resp.Error = 0

	return resp
}
