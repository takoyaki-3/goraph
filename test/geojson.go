package main

import (
	"github.com/takoyaki-3/goraph/loader/geojson"
	csvloader "github.com/takoyaki-3/goraph/loader/csv"
)

func main(){
	g := geojson.Load("N02-19_RailroadSection.geojson")
	csvloader.WriteEdge("edge.csv",g)
	csvloader.WriteLatLon("latlon.csv",g)
}