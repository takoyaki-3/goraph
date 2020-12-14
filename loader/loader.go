package loader

import (
	"log"
	"io/ioutil"
	"github.com/takoyaki-3/goraph"
	"github.com/takoyaki-3/goraph/pb"
	"github.com/golang/protobuf/proto"
)

// Load Protocol Buffer
func LoadFormPtotocolBuffer(filename string)goraph.Graph{
	// Read the existing graph.
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	graph := &pb.Graph{}
	if err := proto.Unmarshal(in, graph); err != nil {
		log.Fatalln("Failed to parse graph:", err)
	}

	edges := [][]goraph.Edge{}

	for _,v:=range graph.Edge{
		for int64(len(edges)) <= v.From{
			edges = append(edges,[]goraph.Edge{})
		}
		edge := goraph.Edge{}
		edge.To = v.To
		edge.Cost = v.Cost
		edges[v.From] = append(edges[v.From],edge)
	}

	g := goraph.Graph{}
	g.Edges = edges
	return g
}

// Write to Protocol Buffer
func Write(filename string,g goraph.Graph){
	graph := &pb.Graph{}
	// ...
	id := int64(0)
	edge := []*pb.Edge{}
	for k,v:=range g.Edges{
		for _,v:=range v{
			e := &pb.Edge{}
			e.EdgeId = id
			e.From = int64(k)
			e.To = v.To
			e.Cost = v.Cost
			edge = append(edge,e)
			id++
		}
	}
	graph.Edge = edge
	latlon := []*pb.LatLon{}
	for k,v:=range g.LatLons{
		e := &pb.LatLon{}
		e.LatlonId = int64(k)
		e.Lat = v.Lat
		e.Lon = v.Lon
		latlon = append(latlon,e)
	}
	graph.Latlon = latlon
	
	// Write the new address book back to disk.
	out, err := proto.Marshal(graph)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}