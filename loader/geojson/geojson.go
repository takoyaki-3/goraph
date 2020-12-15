package geojson

import (
	"log"
	"bytes"
	"strconv"
  "io/ioutil"
  "crypto/sha256"
  "github.com/paulmach/orb"
  "github.com/paulmach/orb/geojson"
	"github.com/takoyaki-3/goraph"
	"github.com/takoyaki-3/goraph/geometry"
)

type Sha256 [32]byte

func Node2Hash(node goraph.LatLon)Sha256{
  flat := strconv.FormatFloat(node.Lat, 'f', -1, 64)
  flon := strconv.FormatFloat(node.Lon, 'f', -1, 64)
  b := bytes.Join([][]byte{[]byte(flat),[]byte(flon)}, []byte{})
  return sha256.Sum256(b)
}

func Load(filename string)goraph.Graph{
	g := goraph.Graph{}

  latlon2id := map[Sha256]int{}

  // JSONファイル読み込み
  rawJSON, err := ioutil.ReadFile(filename)
  if err != nil {
    log.Fatal(err)
	}

	fc, _ := geojson.UnmarshalFeatureCollection(rawJSON)

  func(){
    for _,f := range fc.Features{
      line := []int{}
      for _,v := range f.Geometry.(orb.LineString){
        node := goraph.LatLon{v[1],v[0]}
        h := Node2Hash(node)
        if id,ok:=latlon2id[h];!ok{
          id = len(g.LatLons)
          latlon2id[h]=id
          g.LatLons = append(g.LatLons,node)
          line = append(line,id)
        } else {
          line = append(line,id)
        }
      }
  
      for k,_ := range line{
        if k==0{
          continue
				}
				e := goraph.Edge{}

				node1,node2 := int64(line[k-1]),int64(line[k])
				for len(g.Edges) <= int(node1) || len(g.Edges) <= int(node2){
					g.Edges = append(g.Edges,[]goraph.Edge{})
				}

				e.Cost = geometry.HubenyDistance(g.LatLons[node1],g.LatLons[node2])
				e.To = node2
				g.Edges[node1] = append(g.Edges[node1],e)
				e.To = node1
				g.Edges[node2] = append(g.Edges[node2],e)
      }
    }
	}()
	
	return g
}
