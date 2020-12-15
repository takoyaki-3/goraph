package h3

import (
	h3 "github.com/uber/h3-go"
	"github.com/takoyaki-3/goraph"
	"github.com/takoyaki-3/goraph/geometry"
	"math"
)

type H3index h3.H3Index

func MakeH3Index(g goraph.Graph,resolution int)map[h3.H3Index][]int64{
	h3index := map[h3.H3Index][]int64{}

	for k,v := range g.LatLons {
		hex := h3.FromGeo(h3.GeoCoord{v.Lat,v.Lon}, resolution)

		if _,ok := h3index[hex];!ok{
			h3index[hex] = []int64{}
		}
		h3index[hex] = append(h3index[hex],int64(k))
	}
	return h3index
}

func Find(g goraph.Graph,h3indexes map[h3.H3Index][]int64,latlon goraph.LatLon,resolution int)int64{
	h3index := h3.FromGeo(h3.GeoCoord{latlon.Lat,latlon.Lon}, resolution)

	hexes,_ := h3.HexRing(h3index,1)
	hexes = append(hexes,h3index)

	min_node := int64(-1)
	min_d := math.MaxFloat64

	for _,v := range hexes{
		if _,ok := h3indexes[v];!ok{
			continue
		}
		for _,v := range h3indexes[v]{
			d := geometry.HubenyDistance(g.LatLons[v],latlon)
			if min_d > d {
				min_node = v
				min_d = d
			}
		}
	}

	return min_node
}