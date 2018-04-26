# csv2qif
Go library for converting CSV files (from various banks) to QIF

The lib takes an io.Reader() (i.e. a file loaded into memory somehow) and outputs a []string in QIF format. From there a QIF file can be written.
