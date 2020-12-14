//package go_graph
package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"github.com/takoyaki-3/goraph"
	"github.com/takoyaki-3/goraph/geometry/h3"
	"github.com/takoyaki-3/goraph/loader"
	"github.com/takoyaki-3/goraph/search"
	"encoding/csv"
)

func main(){

	fmt.Println("start")
	graph := loader.Load("kanto.graph.pbf")
	fmt.Println("loaded")

	h3indexes := h3.MakeH3Index(graph,9)
	
	bases := []int64{}
	bases = append(bases,h3.Find(graph,h3indexes,goraph.LatLon{35.654803,139.542766},9))
	bases = append(bases,h3.Find(graph,h3indexes,goraph.LatLon{35.686354,139.673279},9))

	nodes := search.Voronoi(graph,bases)


	// 書き出し
	wf, err := os.Create("./out.csv")
	if err != nil {
		log.Println(err)
	}
	defer wf.Close()

	w := csv.NewWriter(wf) // utf8
	w.Write([]string{"base","lat","lon"})
	for k,v := range nodes{
		lat := strconv.FormatFloat(graph.LatLons[k].Lat, 'f', -1, 64)
		lon := strconv.FormatFloat(graph.LatLons[k].Lon, 'f', -1, 64)
		w.Write([]string{strconv.Itoa(int(v)),lat,lon})
	}
	w.Flush()
}