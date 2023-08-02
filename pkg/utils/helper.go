package utils

import (
	"log"
	"strconv"
	"strings"
)

func StrToInt(src string) int{
	i, err := strconv.Atoi(src)
	
	if err != nil {
		log.Fatal(err)
	}

	return i
}

func JoinURL(src ...string) string{
	return strings.Join(src, "/")
}