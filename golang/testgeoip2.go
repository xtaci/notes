package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"time"
)

func main() {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	//	ip := net.ParseIP("218.104.200.146")
	ip := net.ParseIP("62.216.125.241")

	const N = 100000
	start := time.Now()
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Now().Sub(start))
	fmt.Printf("%v\n", record.Country.GeoNameID)
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	// Output:
	// Portuguese (BR) city name: Londres
	// English subdivision name: England
	// Russian country name: Великобритания
	// ISO country code: GB
	// Time zone: Europe/London
	// Coordinates: 51.5142, -0.0931
}
