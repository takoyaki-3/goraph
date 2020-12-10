package goraph

import (
	"math"
	"github.com/takoyaki-3/go_minimum_set"
)

type Edge struct{
	Cost float64
	To int64
}

type Graph struct{
	Edges [][]Edge
}

type Query struct{
	From int64
	To int64
}

type Output struct{
	Nodes []int64
	Cost float64
}

// 
func Search(graph Graph,query Query)Output{
	edges := graph.Edges

	q := go_minimum_set.NewMinSet()
	cost := make([]float64,len(edges))
	flag := make([]bool,len(edges))
	before := make([]int64,len(edges))

	maxcost := math.MaxFloat64

	for k,_ := range cost{
		cost[k] = maxcost
	}

	cost[query.From] = 0.0
	before[query.From] = -2

	q.Add_val(query.From,0.0)

	var pos int64
	for q.Len()>0{
		pos = q.Get_min()
		if flag[pos]{
			continue
		}
		flag[pos]=true

		if pos == query.To{
			break
		}

		for _,e := range edges[pos]{
			eto := e.To
			if flag[eto]{
				continue
			}
			if cost[eto] <= cost[pos]+e.Cost{
				continue
			}
			cost[eto] = cost[pos]+e.Cost
			before[eto] = pos
			q.Add_val(eto,cost[eto])
		}
	}

	// 出力
	out := Output{}
	out.Cost = cost[pos]
	out.Nodes = append(out.Nodes,pos)

	bef := before[pos]
	for bef!=-2 {
		out.Nodes = append(out.Nodes,bef)
		bef = before[bef]
	}
	return out
}
