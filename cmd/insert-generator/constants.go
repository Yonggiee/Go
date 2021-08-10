package main

import (
	"regexp"
)

var createTableRegex = regexp.MustCompile(`^create\s+table\s+[_a-z]?[a-z0-9]*\s*\(.*\)$`)
var innerTableRegex = regexp.MustCompile(`\((.*)\)$`)

var dataTypes = []string{
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
