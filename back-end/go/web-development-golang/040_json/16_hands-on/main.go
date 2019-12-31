package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type code struct {
	Code    int    `json:"code"`
	Descrip string `json:"descrip"`
}

func main() {
	var data []code

	rcvd := `[{"code":200,"descrip":"StatusOK"},{"code":301,"descrip":"StatusMovedPermanently"},{"code":302,"descrip":"StatusFound"},{"code":303,"descrip":"StatusSeeOther"},{"code":307,"descrip":"StatusTemporaryRedirect"},{"code":400,"descrip":"StatusBadRequest"},{"code":401,"descrip":"StatusUnauthorized"},{"code":402,"descrip":"StatusPaymentRequired"},{"code":403,"descrip":"StatusForbiden"},{"code":404,"descrip":"StatusNotFound"},{"code":405,"descrip":"StatusMethodNotAllowed"},{"code":418,"descrip":"StatusTeapot"},{"code":500,"descrip":"StatusInternalServerError"}]`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}

	for _, v := range data {
		fmt.Println(v.Code, "-", v.Descrip)
	}
}
