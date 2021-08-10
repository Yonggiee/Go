package main

type column struct {
	name        string
	typ         string
	constraints []string
}

type table struct {
	cols      []column
	dependsOn []table
}
