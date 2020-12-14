package osm

import (
	"os"
	"io"
	"fmt"
	"log"
	"runtime"
	"github.com/cheggaaa/pb"
	"github.com/dustin/go-humanize"
	"github.com/kkdd/osmpbf"
	"math"
	"github.com/takoyaki-3/goraph"
)

func Load(filename string)goraph.Graph{
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	stat, _ := f.Stat()
	filesiz := int(stat.Size()/1024)

	d := osmpbf.NewDecoder(f)
	err = d.Start(runtime.GOMAXPROCS(-1))
	if err != nil {
		log.Fatal(err)
	}

	// 一時記憶用変数
	latlons := map[int64]goraph.LatLon{}
	ways := map[int64][]int64{}
	usednode := map[int64]bool{}

	nodeid := NewReplace()

	nc,wc,rc:=int64(0),int64(0),int64(0)
	pb.New(filesiz).SetUnits(pb.U_NO)
	for i := 0; ; i++ {
		if v, err := d.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			switch v := v.(type) {
				case *osmpbf.Node:
					latlons[v.ID] = goraph.LatLon{v.Lat,v.Lon}
					nc++
				case *osmpbf.Way:
					if _,ok:=v.Tags["highway"];!ok{
						continue
					}
					nodes := []int64{}
					for _,v := range v.NodeIDs{
						nodes = append(nodes,v)
						usednode[v] = true
					}
					ways[v.ID] = nodes
					wc++
				case *osmpbf.Relation:
					rc++
				default:
					log.Fatalf("unknown type %T\n", v)
			}
		}
	}
	fmt.Printf("Nodes: %s, Ways: %s, Relations: %s\n", humanize.Comma(nc), humanize.Comma(wc), humanize.Comma(rc))

	// 結果を出力
	

	graph := goraph.Graph{}

	for _,v := range ways{
		for i:=1;i<len(v);i++{
			e := goraph.Edge{}
			e.Cost = HubenyDistance(latlons[v[i-1]],latlons[v[i]])
			node1 := nodeid.AddReplace(v[i-1])
			node2 := nodeid.AddReplace(v[i])
			for len(graph.Edges) <= int(node1) || len(graph.Edges) <= int(node2){
				graph.Edges = append(graph.Edges,[]goraph.Edge{})
			}
			e.To = node2
			graph.Edges[node1] = append(graph.Edges[node1],e)
			e.To = node1
			graph.Edges[node2] = append(graph.Edges[node2],e)
		}
	}

	return graph
}


// 置き換え関数
type Replace struct{
	Id2Str map[int64]int64
	Str2Id map[int64]int64
}

func (s *Replace)AddReplace(str int64)int64{
	if val, ok := s.Str2Id[str]; ok {
		return val
	}
	id := int64(len(s.Str2Id) + 1)
	s.Str2Id[str] = id
	s.Id2Str[id] = str
	return id
}

func (s *Replace)AddReplaceIndex(str int64,index int64)int64{
	if val, ok := s.Str2Id[str]; ok {
		return val
	}
	s.Str2Id[str] = index
	s.Id2Str[index] = str
	return index
}

func NewReplace() *Replace{
	s := new(Replace)
	s.Str2Id = map[int64]int64{}
	s.Id2Str = map[int64]int64{}
	return s
}

// 緯度経度から距離を計算する
func degree2radian(x float64) float64 {
	return x * math.Pi / 180
}

func Power2(x float64) float64 {
	return math.Pow(x, 2)
}
const (
	EQUATORIAL_RADIUS    = 6378137.0            // 赤道半径 GRS80
	POLAR_RADIUS         = 6356752.314          // 極半径 GRS80
	ECCENTRICITY         = 0.081819191042815790 // 第一離心率 GRS80
)
type Point struct {
	Lat  float64
	Lon float64
}
func HubenyDistance(src goraph.LatLon, dst goraph.LatLon) float64 {
	dx := degree2radian(dst.Lon - src.Lon)
	dy := degree2radian(dst.Lat - src.Lat)
	my := degree2radian((src.Lat + dst.Lat) / 2)

	W := math.Sqrt(1 - (Power2(ECCENTRICITY) * Power2(math.Sin(my)))) // 卯酉線曲率半径の分母
	m_numer := EQUATORIAL_RADIUS * (1 - Power2(ECCENTRICITY))         // 子午線曲率半径の分子

	M := m_numer / math.Pow(W, 3) // 子午線曲率半径
	N := EQUATORIAL_RADIUS / W    // 卯酉線曲率半径

	d := math.Sqrt(Power2(dy*M) + Power2(dx*N*math.Cos(my)))

	return d
}