package main

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/decoder"
	"math"
	"reflect"
)

// Very simple validator
func notZZ(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return errors.New("notZZ only validates strings")
	}
	if st.String() == "ZZ" {
		return errors.New("value cannot be ZZ")
	}
	return nil
}


func cov(new float64, existingData float64, cov float64) (bool, float64) {
	c:= new - existingData
	if math.Abs(c) < cov {
		return false, existingData
	} else {
		return false, new
	}

}


func main()  {
	fmt.Println(decoder.SensorNames.ME)
	a := decoder.SensorNames.ME

	//def _check_cov(cls, _new, _existing_data, cov):
		//if abs(_new - _existing_data) < cov:
		//return [False, _existing_data]
		//else:
	//return [True, _new]

	fmt.Println(string(a))
	fmt.Println(reflect.TypeOf(a))



}