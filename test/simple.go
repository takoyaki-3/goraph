package main

import (
	"fmt"
	"os"
	"log"
	"github.com/takoyaki-3/goraph"
	"github.com/takoyaki-3/goraph/geometry"
	"github.com/takoyaki-3/goraph/geometry/h3"
	"github.com/takoyaki-3/goraph/loader"
	"github.com/takoyaki-3/goraph/search"
)

func main(){

	fmt.Println("start")
	g := loader.Load("kanto.graph.pbf")
	fmt.Println("loaded")

	h3indexes := h3.MakeH3Index(g,9)
	
	var q search.Query

	q.To = h3.Find(g,h3indexes,goraph.LatLon{35.654803,139.542766},9)
	q.From = h3.Find(g,h3indexes,goraph.LatLon{35.686354,139.673279},9)

	rv := search.Search(g,q)

	rawJSON := geometry.MakeLineString(g,rv.Nodes)
	file, err := os.Create("out.geojson")
	if err != nil {
		log.Fatal(err)  //ファイルが開けなかったときエラー出力
	}
	defer file.Close()
	file.Write(([]byte)(rawJSON))

	fmt.Println(rv.Cost)
}