package pkg

import (
	"fmt"
	"golangbot/pkg/maps"
)

func Mymap() {
	fmt.Println("\n+++ Map +++++")
	//map
	maps.Map_create()
	maps.Map_assign()
	maps.Map_assign2()
	maps.Map_fetch()
	maps.In_map()
	maps.Map_range()
	maps.Map_delete_element()
	maps.Map_length()
	maps.Map_reference_types()
	maps.Map_equality()
}
