package tool

import (
	"github.com/takoyaki-3/goraph"
)

func CutGoraph(g *goraph.Graph,leftUp goraph.LatLon,rightDown goraph.LatLon)error{
	old2new := map[int]int{}
	newG := goraph.Graph{}
	for old, v := range g.LatLons {
		if leftUp.Lon <= v.Lon && v.Lon <= rightDown.Lon {
			if rightDown.Lat <= v.Lat && v.Lat <= leftUp.Lat {
				old2new[old] = len(newG.LatLons)
				newG.LatLons = append(newG.LatLons, v)
			}
		}
	}
	for old, es := range g.Edges {
		if _, ok := old2new[old]; ok {
			edges := []goraph.Edge{}
			for _, v := range es {
				if n, ok := old2new[int(v.To)]; ok {
					v.To = int64(n)
					edges = append(edges, v)
				}
			}
			newG.Edges = append(newG.Edges, edges)
		}
	}

	*g = newG

	return nil
}
