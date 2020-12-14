package search


import (
	"math"
	"github.com/takoyaki-3/goraph"
	"github.com/takoyaki-3/go_minimum_set"
)

type Output struct{
	Nodes []int64
	Cost float64
}

// 
func Search(graph goraph.Graph,query goraph.Query)Output{

	q := go_minimum_set.NewMinSet()
	cost := make([]float64,len(graph.Edges))
	flag := make([]bool,len(graph.Edges))
	before := make([]int64,len(graph.Edges))

	for k,_ := range cost{
		cost[k] = math.MaxFloat64
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

		for _,e := range graph.Edges[pos]{
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
		out.Nodes = append([]int64{bef},out.Nodes...)
		bef = before[bef]
	}
	return out
}

func Voronoi(graph goraph.Graph,bases []int64)map[int64]int64{
	// initialization
	q := go_minimum_set.NewMinSet()
	cost := make([]float64,len(graph.Edges))
	flag := make([]bool,len(graph.Edges))
	start_group := map[int64]int64{}

	counter := int64(0)
	for k,_ := range cost{
		cost[k] = math.MaxFloat64
	}

	for _,v := range bases{
		cost[v] = 0.0
		q.Add_val(v,0.0)
		start_group[int64(v)] = counter % 20
		counter++
	}

	for q.Len()>0{
		pos := q.Get_min()
		if flag[pos]{
			continue
		}
		flag[pos]=true
		
		// グラフ拡張処理
		for _,e := range graph.Edges[pos]{
			eto := e.To
			if flag[eto] {
				continue
			}
			if cost[eto] <= cost[pos]+e.Cost{
				continue
			}
			cost[eto] = cost[pos]+e.Cost
			start_group[eto] = start_group[pos]
			q.Add_val(eto,cost[pos]+e.Cost)
		}
	}

	return start_group
}