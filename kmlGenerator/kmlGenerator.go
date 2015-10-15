package main

import (
	"encoding/csv"
	"fmt"
	//	"strings"
	//	"io/ioutil"
	"os"
	"strconv"
)

const (
	XMETER = 0.0000111411
	YMETER = 0.0000089908
)

/*type KML struct{
	XML xml.Name `xml:"kml"`

}*/
type Boundary struct {
	X float64
	Y float64
}

func main() {
	var myData string

	//Open input file
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	//Create reader for csv file
	r := csv.NewReader(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Read boundary points from input file
	inputData, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(inputData[0]) != 6 {
		fmt.Println(inputData[0], "=", len(inputData[0]))
		for _, value := range inputData[0] {
			fmt.Println(value)
		}
		fmt.Println("Input file needs to have only 2 GPS points")
		return
	}

	//Parse boundary points into Boundary's
	x1, err := strconv.ParseFloat(inputData[0][0], 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	y1, err := strconv.ParseFloat(inputData[0][1], 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	x2, err := strconv.ParseFloat(inputData[0][3], 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	y2, err := strconv.ParseFloat(inputData[0][4], 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	nwBound := Boundary{X: x1, Y: y1}
	seBound := Boundary{X: x2, Y: y2}

	currGPS := nwBound

	for i := 0; currGPS.Y > seBound.Y; {
		for ; currGPS.X < seBound.X; i++ {

			//calculate gps of corners for square
			myCoords := strconv.FormatFloat(currGPS.X, 'f', -1, 64) +","+
				strconv.FormatFloat(currGPS.Y, 'f', -1, 64) +","+
				"0 " +
				strconv.FormatFloat(currGPS.X, 'f', -1, 64) +","+
				strconv.FormatFloat(currGPS.Y-(YMETER*150), 'f', -1, 64) +","+
				"0 " +
				strconv.FormatFloat(currGPS.X+(XMETER*150), 'f', -1, 64) +","+
				strconv.FormatFloat(currGPS.Y-(YMETER*150), 'f', -1, 64) +","+
				"0 " +
				strconv.FormatFloat(currGPS.X+(XMETER*150), 'f', -1, 64) +","+
				strconv.FormatFloat(currGPS.Y, 'f', -1, 64) +","+
				"0 " +
				strconv.FormatFloat(currGPS.X, 'f', -1, 64) +","+
				strconv.FormatFloat(currGPS.Y, 'f', -1, 64) +","+
				"0 "
			myData = `<?xml version="1.0" encoding="UTF-8"?>
<kml xmlns="http://www.opengis.net/kml/2.2" xmlns:gx="http://www.google.com/kml/ext/2.2" xmlns:kml="http://www.opengis.net/kml/2.2" xmlns:atom="http://www.w3.org/2005/Atom">
<Document>
<name>` + `output` + strconv.Itoa(i) + `.kml` + `</name>
<StyleMap id="m_ylw-pushpin">
<Pair>
<key>normal</key>
<styleUrl>#s_ylw-pushpin</styleUrl>
</Pair>
<Pair>
<key>highlight</key>
<styleUrl>#s_ylw-pushpin_hl</styleUrl>
</Pair>
</StyleMap>
<Style id="s_ylw-pushpin_hl">
<IconStyle>
<scale>1.3</scale>
<Icon>
<href>http://maps.google.com/mapfiles/kml/pushpin/ylw-pushpin.png</href>
</Icon>
<hotSpot x="20" y="2" xunits="pixels" yunits="pixels"/>
</IconStyle>
<LineStyle>
<color>ff00ffff</color>
<width>1.5</width>
</LineStyle>
<PolyStyle>
<color>7f00ffff</color>
</PolyStyle>
</Style>
<Style id="s_ylw-pushpin">
<IconStyle>
<scale>1.1</scale>
<Icon>
<href>http://maps.google.com/mapfiles/kml/pushpin/ylw-pushpin.png</href>
</Icon>
<hotSpot x="20" y="2" xunits="pixels" yunits="pixels"/>
</IconStyle>
<LineStyle>
<color>ff00ffff</color>
<width>1.5</width>
</LineStyle>
<PolyStyle>
<color>7f00ffff</color>
</PolyStyle>
</Style>
<Placemark>
<name>150x150m sq approx</name>
<styleUrl>#m_ylw-pushpin</styleUrl>
<Polygon>
<tessellate>1</tessellate>
<outerBoundaryIs>
<LinearRing>
<coordinates>
` + myCoords + `
</coordinates>
</LinearRing>
</outerBoundaryIs>
</Polygon>
</Placemark>
</Document>
</kml>`

			g, err := os.Create("output" + strconv.Itoa(i) + ".kml")
			if err != nil {
				panic(err)
			}
			defer g.Close()

			g.WriteString(myData)

			//shift the currGPS.X
			currGPS.X = currGPS.X + (XMETER * 120)
		}
		//reset the currGPS.X
		currGPS.X = nwBound.X

		//shift the currGPS.Y
		currGPS.Y = currGPS.Y - (YMETER * 120)
	}
}
