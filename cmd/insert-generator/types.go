package main

type column struct {
	name        string
	typ         string
	constraints []string
	specialTag  string
}

type table struct {
	name                  string
	cols                  []column
	dependsOn             []table
	insertCount           int
	additionalConstraints []string
}
