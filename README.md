# csv2qif
[![Go Report Card](https://goreportcard.com/badge/github.com/dgjnpr/csv2qif)](https://goreportcard.com/report/github.com/dgjnpr/csv2qif)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/dgjnpr/csv2qif)
[![Release](https://img.shields.io/github/release/dgjnpr/csv2qif.svg)](https://github.com/dgjnpr/csv2qif/releases/latest)
<!-- [![license](https://img.shields.io/github.com/license/dgjnpr/csv2qif.svg)](https://github.com/dgjnpr/csv2qif/LICENSE) -->

Go library for converting CSV files (from various banks) to QIF

The lib takes an io.Reader() (i.e. a file loaded into memory somehow) and outputs a []string in QIF format. From there a QIF file can be written.
