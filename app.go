package main

import "github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery"

func main() {
	// Menu Vehicle
	// VH1 -> Insert, VH2 -> List, VH3 -> Get, VH3 -> Update, VH4 -> Delete
	delivery.NewServer().Run("VH1")
}
