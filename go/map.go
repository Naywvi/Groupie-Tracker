package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strings"
)

func JscriptStr() { //handles the map on the artist page
	var loc LattitudeLongitude
	var slice []string
	str := ""
	for limite, i := range artist.Location.Locations { //Retire cas particuliers
		for _, j := range i {
			if j == '_' {
				j = ' '
			}
			str += string(j)
		}
		json.Unmarshal([]byte(request("http://api.positionstack.com/v1/forward?access_key=9a6d5681ba2b143da463543ee17cf96e&query="+str+"&limit=1")), &loc)
		if len(artist.Location.Locations)-1 == limite { //créations jsstring
			slice = append(slice, "['"+str+"',"+fmt.Sprintf("%f", loc.Data[0].Latitude)+", "+fmt.Sprintf("%f", loc.Data[0].Longitude)+","+fmt.Sprintf("%q", artist.Location.DatesLoc.Dates[limite][1:])+"],")
		} else {
			slice = append(slice, `['`+str+"',"+fmt.Sprintf("%f", loc.Data[0].Latitude)+", "+fmt.Sprintf("%f", loc.Data[0].Longitude)+","+fmt.Sprintf("%q", artist.Location.DatesLoc.Dates[limite][1:])+"],")
		}
		str = ""
	}
	k := `<script>
	var LocationsForMap = [
		` + strings.Join(slice, "\n") + `
		];

	// Initialize and add the map
	function initMap() {
	  // The map, centered 
	  const map = new google.maps.Map(document.getElementById("map"), {
		zoom: 1,
		center: new google.maps.LatLng(36.91, 1.64),
	  });

	  var infowindow = new google.maps.InfoWindow();

	  for (i = 0; i < LocationsForMap.length; i++) {
		const marker = new google.maps.Marker({
		  position: new google.maps.LatLng(LocationsForMap[i][1], LocationsForMap[i][2]),
		  map: map,
		});
  
		google.maps.event.addListener(marker, 'click', (function(marker, i) {
		  return function() {
			infowindow.setContent(LocationsForMap[i][0]+" "+LocationsForMap[i][3]);
			infowindow.open(map, marker);
		  }
		})(marker, i));
	  }
	}
  </script>`
	artist.JsString = template.HTML(k)
}
