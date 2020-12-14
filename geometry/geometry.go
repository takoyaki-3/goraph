package geometry

import (
  "github.com/paulmach/orb"
  "github.com/paulmach/orb/geojson"
	"github.com/takoyaki-3/goraph"
)

func MakeLineString(graph goraph.Graph,latlons []int64)string{
	line := orb.LineString{}
	for _,v:=range latlons{
		line = append(line,orb.Point{graph.LatLons[v].Lon, graph.LatLons[v].Lat})
	}

	fc := geojson.NewFeatureCollection()
	fc.Append(geojson.NewFeature(line))
	rawJSON, _ := fc.MarshalJSON()
	return string(rawJSON)
}