package main

import (
	"fmt"

	config "github.com/LavaJover/storage-sso-service/sso-service/internal/config"
)

func main(){
	fmt.Println(config.MustLoad())
}