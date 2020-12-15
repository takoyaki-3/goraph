package main

import (
	"fmt"
	"github.com/takoyaki-3/goraph/loader"
	csvloader "github.com/takoyaki-3/goraph/loader/csv"
)

func main(){
	
	fmt.Println("start")
	g := loader.Load("kanto.graph.pbf")
	fmt.Println("loaded")
	csvloader.WriteEdge("edge.csv",g)
	csvloader.WriteLatLon("latlon.csv",g)
	fmt.Println("writed")
}