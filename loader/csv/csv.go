package csv

import (
	"os"
	"log"
	"strconv"
	"encoding/csv"
	"github.com/takoyaki-3/goraph"
	"github.com/takoyaki-3/go_replace"
)

// Load CSV
func LoadEdge(filename string)goraph.Graph{
	
	replace_nodeid := *go_replace.NewReplace()
	edges := [][]goraph.Edge{}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	defer file.Close()

	counter := -1
	titles := map[string]int{}
	
	for {
		counter++
		line, err := reader.Read()
		if err != nil {
			break
		}
		if counter==0{
			for k,v:=range line{
				titles[v]=k
			}
			continue
		}
		from := replace_nodeid.AddReplace(line[titles["from"]])
		to   := replace_nodeid.AddReplace(line[titles["to"]])
		cost := 1.0
		if v,ok:=titles["cost"];ok{
			cost,_ = strconv.ParseFloat(line[v], 64)
		}

		for int64(len(edges)) <= from{
			edges = append(edges,[]goraph.Edge{})
		}
		edge := goraph.Edge{}
		edge.To = to
		edge.Cost = cost
		edges[from] = append(edges[from],edge)
	}

	g := goraph.Graph{}
	g.Edges = edges
	return g
}

func WriteEdge(filename string,graph goraph.Graph){
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	wr := csv.NewWriter(f)
	wr.Write([]string{"from","to","cost"})
	
	for k,v:=range graph.Edges{
		for _,v:=range v{
			cost := strconv.FormatFloat(v.Cost, 'f', -1, 64)
			line := []string{strconv.FormatInt(int64(k), 10),strconv.FormatInt(v.To, 10),cost}
			wr.Write(line)
		}
	}

	wr.Flush()
	if err := wr.Error(); err != nil {
			log.Fatal(err)
	}
}

func WriteLatLon(filename string,graph goraph.Graph){
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	wr := csv.NewWriter(f)
	wr.Write([]string{"latlon_id","lat","lon"})
	
	for k,v:=range graph.LatLons{
		lat := strconv.FormatFloat(v.Lat, 'f', -1, 64)
		lon := strconv.FormatFloat(v.Lon, 'f', -1, 64)
		line := []string{strconv.FormatInt(int64(k), 10),lat,lon}
		wr.Write(line)
	}

	wr.Flush()
	if err := wr.Error(); err != nil {
		log.Fatal(err)
	}
}