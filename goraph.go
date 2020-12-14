package goraph

type Edge struct{
	Cost float64
	To int64
}

type LatLon struct{
	Lat float64
	Lon float64
}

type Graph struct{
	Edges [][]Edge
	LatLons []LatLon
}

type Query struct{
	From int64
	To int64
}