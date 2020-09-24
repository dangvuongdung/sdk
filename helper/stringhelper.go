package helper

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"time"
)

func RandInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

func IndexStringAtArray(str string, strs []string) int {
	for k, v := range strs {
		if str == v {
			return k
		}
	}
	return -1
}

func RemoveStringAtArray(str string, strs []string) []string {
	rets := make([]string, 0)
	for idxT, strT := range strs {
		if str == strT {
			rets = append(rets, strs[0:idxT]...)
			rets = append(rets, strs[idxT+1:]...)
			return rets
		}
	}
	rets = append(rets, strs...)
	return strs
}

func RemoveIndexAtArray(idx int, strs []string) []string {
	rets := make([]string, 0)
	for idxT, _ := range strs {
		if idx == idxT {
			rets = append(rets, strs[0:idxT]...)
			rets = append(rets, strs[idxT+1:]...)
			return rets
		}
	}
	rets = append(rets, strs...)
	return strs
}

// InterfaceSlice receives a slice which is a interface
// and converts it into slice of interface
func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		log.Println("InterfaceSlice() given a non-slice type")
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}


func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func JoinAddress(address1 string, address2 string) string {
	if address2 != "" {
		return fmt.Sprintf("%s %s", address1, address2)
	}
	return address1
}

