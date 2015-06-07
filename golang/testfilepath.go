package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Base(`/usr/local/ccc.json`))
	fmt.Println(filepath.Dir(`/backends/geoip/42269b048910:condescending_fermat:50000`))
}
