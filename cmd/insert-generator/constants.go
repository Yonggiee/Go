package main

import (
	"regexp"
)

// Regex
var CREATETABLEREGEX = regexp.MustCompile(`^create\s+table\s+[_a-z]?[a-z0-9]*\s*\(.*\)\s*[0-9]+\s*$`)

// Datatypes
var DATATYPES = []string{
	// numeric
	"int",
	"integer",
	"float",
	"numeric",
	"decimal",

	// character/string
	"varchar",
	"char",
	"text",

	// DateTime
	"date",
	"datetime",
	"time",
	"timestamp",
}

// Insert
var INSERT = "INSERT INTO VALUES"
