package utils

import (
	"strconv"
	"github.com/JIeeiroSst/go-app/log"
)


func StringToInt(num string) int {
	numInt, err := strconv.Atoi(num)
	if err!=nil{
		log.InitZapLog().Error("server running error")
	}
	log.InitZapLog().Error("server running success")
	return numInt
}