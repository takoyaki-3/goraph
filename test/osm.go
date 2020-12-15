//package go_g
package main

import (
	"fmt"
	"os"
	"log"
	"github.com/takoyaki-3/goraph"
	"github.com/takoyaki-3/goraph/geometry"
	"github.com/takoyaki-3/goraph/geometry/h3"
	"github.com/takoyaki-3/goraph/loader"
	"github.com/takoyaki-3/goraph/loader/osm"
	"github.com/takoyaki-3/goraph/search"
)

func main(){

	fmt.Println("start")
	g := osm.Load("./kanto-latest.osm.pbf")
	fmt.Println("loaded")
	loader.Write("kanto.g.pbf",g)
	fmt.Println("writed")

	var q search.Query

	h3indexes := h3.MakeH3Index(g,9)

	q.To = h3.Find(g,h3indexes,goraph.LatLon{35.654803,139.542766},9)
	q.From = h3.Find(g,h3indexes,goraph.LatLon{35.686354,139.673279},9)

	fmt.Println(q)

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