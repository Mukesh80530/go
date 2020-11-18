package main

import "fmt"

func main(){
	grades := make(map[string] float32)
	// grades["name"] = "mukesh"
	grades["jess"] = 33
	grades["sam"] = 78
	grades["muk"] = 88
	fmt.Println(grades)

	TimsGrade := grades["jess"]
	fmt.Println(TimsGrade)

	// delete(grades,"jess")
	// fmt.Println(grades)

	for k,v := range grades {
		fmt.Println(k, ":", v)
	}
}