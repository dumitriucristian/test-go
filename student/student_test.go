package student

import (
	"fmt"
	"testing"
)

func TestStudent(t *testing.T) {
	data := []byte(`
			{
				"Info":{
					"id":123,
					"name":"Smith Joh",
					"enrolled": true
				},
				"classes": [
						{
							"coursename":"Intro to Golang", 
							"coursenum":101,
							"coursehours": 4
						},
						{
							"coursename":"English Literature", 
							"coursenum":102,
							"coursehours": 4
						},
						{
							"coursename":"World History", 
							"coursenum":103,
							"coursehours": 4
						}
				]
			}
		`)
	var some Student
	err := some.Convert(data)
	if err != nil {
		fmt.Println("Conversion error:", err)
	}
	fmt.Println(some)

}
