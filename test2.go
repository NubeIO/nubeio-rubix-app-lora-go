package main
import (
	"fmt"
)

type ReturnValue struct {
	Status string
	CustomStruct interface{}
}

func GetReturn(status string, class interface{}){
	var result = ReturnValue {Status : status, CustomStruct: class}

	fmt.Println(result)

	msg, ok := result.CustomStruct.(Message1)
	if ok {
		fmt.Printf("Message1 is %s\n", msg.message)
	}
}

type Message1 struct {
	message string
}

type Message2 struct {
	message string
}

func main(){
	var m1 = Message1 {message: "Hello1"}
	GetReturn("success",  m1)

	var m2 = Message2 {message: "Hello2"}
	GetReturn("success",  m2)
}

//package main
//
//import (
//	"errors"
//	"fmt"
//	"github.com/NubeIO/nubeio-rubix-app-lora-go/decoder"
//	"math"
//	"reflect"
//)
//
//// Very simple validator
//func notZZ(v interface{}, param string) error {
//	st := reflect.ValueOf(v)
//	if st.Kind() != reflect.String {
//		return errors.New("notZZ only validates strings")
//	}
//	if st.String() == "ZZ" {
//		return errors.New("value cannot be ZZ")
//	}
//	return nil
//}
//
//
//func cov(new float64, existingData float64, cov float64) (bool, float64) {
//	c:= new - existingData
//	if math.Abs(c) < cov {
//		return false, existingData
//	} else {
//		return false, new
//	}
//
//}
//
//
//
//
//
//func main()  {
//	fmt.Println(decoder.SensorNames.ME)
//	a := decoder.SensorNames.ME
//
//	//def _check_cov(cls, _new, _existing_data, cov):
//		//if abs(_new - _existing_data) < cov:
//		//return [False, _existing_data]
//		//else:
//	//return [True, _new]
//
//	fmt.Println(string(a))
//	fmt.Println(reflect.TypeOf(a))
//
//
//
//}