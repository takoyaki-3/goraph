//package go_graph
package main

import (
	"fmt"
	"github.com/takoyaki-3/goraph/loader"
	csvloader "github.com/takoyaki-3/goraph/loader/csv"
)

func main(){

	fmt.Println("start")
	graph := loader.Load("kanto.graph.pbf")
	fmt.Println("loaded")
	csvloader.WriteEdge("edge.csv",graph)
	csvloader.WriteLatLon("latlon.csv",graph)
	fmt.Println("writed")
}