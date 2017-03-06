package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func loadModule(name string) bool {

	protos := make(map[string]interface{}, 50)

	filePath := "apiconfig/" + name + ".json"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file " + filePath + "error")
		return false
	}

	data, err := ioutil.ReadFile(filePath)

	file.Close()

	err = json.Unmarshal(data, &protos)
	if err != nil {
		fmt.Println("Transfer json bytes error")
		fmt.Println(err)
		return false
	}

	keys := make([]string, len(protos))
	i := 0
	for k, _ := range protos {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for _, key := range keys {

		apiKey := "octlink.mirage.center" + "." + name + "." + string(key)

		fmt.Printf("\n")
		fmt.Printf("\t// --------------------\n")
		fmt.Printf("\t// for %s\n", key)
		fmt.Printf("\t// --------------------\n")
		fmt.Printf("\tservice = new(ApiService)\n")
		fmt.Printf("\tservice.Name = \"%s\"\n", key)
		fmt.Printf("\tservice.Handler = %s\n", key)
		fmt.Printf("\tGApiServices[\"%s\"] = service\n", apiKey)
	}

	return true
}

func main() {

	modules := []string{"user", "host", "account"}

	fmt.Printf("// Auto generated by System, don't modify this file\n")
	fmt.Printf("package api\n")
	fmt.Printf("\n")
	fmt.Printf("func InitApiService() {\n")
	fmt.Printf("\n")
	fmt.Printf("\tvar service *ApiService\n")
	fmt.Printf("\n")
	fmt.Printf("\tGApiServices = make(map[string]*ApiService, 10000)\n")

	for i := 0; i < len(modules); i++ {
		state := loadModule(modules[i])
		if state != true {
			break
		}
	}

	fmt.Printf("}\n")
}
