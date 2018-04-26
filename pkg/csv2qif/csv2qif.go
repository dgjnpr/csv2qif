package csv2qif

// This package defines the interface for the bank specific implementations
// of converting a CSV file into the QIF format.

// Converter is the interface all implementations must adhere to.
type Converter interface {
	Csv2qif() string
}
