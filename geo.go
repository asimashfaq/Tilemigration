package main

import (
	"github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"
	"fmt"
	"encoding/json"
	"log"
)

type LatLng struct {
	lat1 float64
	lon1 float64
	lat2 float64
	lon2 float64
}
func main() {
	str :=`[[{
	"lat": 33.6020866,
	"lng": 73.11957699999999
	},
	{
	"lat": 33.601932,
	"lng": 73.12015149999999
	}],
	[{
	"lat": 33.601932,
	"lng": 73.12015149999999
	},
	{
	"lat": 33.6024262,
	"lng": 73.120469
	}],
	[{
	"lat": 33.6024262,
	"lng": 73.120469
	},
	{
	"lat": 33.6015895,
	"lng": 73.12341669999999
	}],
	[{
	"lat": 33.6015895,
	"lng": 73.12341669999999
	},
	{
	"lat": 33.6023443,
	"lng": 73.1235623
	}],
	[ {
	"lat": 33.6023443,
	"lng": 73.1235623
	},
	{
	"lat": 33.6026672,
	"lng": 73.12297590000001
	}],
	[{
	"lat": 33.6026672,
	"lng": 73.12297590000001
	},
	{
	"lat": 33.6048633,
	"lng": 73.13047259999999
	}],
	[{
	"lat": 33.6048633,
	"lng": 73.13047259999999
	},
	{
	"lat": 33.606583,
	"lng": 73.1302683
	}],
	[{
	"lat": 33.606583,
	"lng": 73.1302683
	},
	{
	"lat": 33.7081121,
	"lng": 73.0542299
	}],
	[{
	"lat": 33.7081121,
	"lng": 73.0542299
	},
	{
	"lat": 33.7082661,
	"lng": 73.0534802
	}],
	[{
	"lat": 33.7082661,
	"lng": 73.0534802
	},
	{
	"lat": 33.6994689,
	"lng": 73.0362457
	}],
	[{
	"lat": 33.6994689,
	"lng": 73.0362457
	},
	{
	"lat": 33.7124245,
	"lng": 73.02659849999999
	}],
	[{
	"lat": 33.7124245,
	"lng": 73.02659849999999
	},
	{
	"lat": 33.6895928,
	"lng": 72.9818745
	}],
	[{
	"lat": 33.6895928,
	"lng": 72.9818745
	},
	{
	"lat": 33.6856911,
	"lng": 72.9848837
	}]]`
	var objmap interface{}
	lists :=[]float64{}
	if err := json.Unmarshal([]byte(str), &objmap); err != nil {
		log.Fatal(err)
	}
	m := objmap.([]interface{})
	geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)

	for _,v:= range  m {
		d := v.([]interface{})
		d1  := d[0].(map[string]interface{})
		d2  := d[1].(map[string]interface{})
		distance, _ := geo1.To(d1["lat"].(float64), d1["lng"].(float64),
			d2["lat"].(float64), d2["lng"].(float64))
		inc := (int(distance)%50)
		if(inc >0){
			inc = 1
		}
		_,_,array :=geo1.Intermediate(d1["lat"].(float64), d1["lng"].(float64),
			d2["lat"].(float64), d2["lng"].(float64),(int(distance)/50)+inc)
		for _,vv := range array{
			lists = append(lists,vv)
		}



	}
	counter :=0
	for _,vv := range lists{
		if(counter == 2){
			counter = 0
			fmt.Println("")
		}

		if(counter == 1){
			fmt.Print(",")

		}

		fmt.Print(vv)
		counter = counter+1
	}



}
