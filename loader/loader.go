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