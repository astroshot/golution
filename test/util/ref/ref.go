package ref

import (
	"fmt"
	"reflect"
	// "sync"
	"testing"
	// "time"
)

type Student struct {
	Id   int64
	Name string
}

func TestRef(t *testing.T) {

	s := &Student{
		Id:   1,
		Name: "jack",
	}

	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.ValueOf(s))
}

// func main() {
// s := &Student{
// Id:   1,
// Name: "jack",
// }

// fmt.Println(reflect.TypeOf(s))
// fmt.Println(reflect.ValueOf(s))
// }
