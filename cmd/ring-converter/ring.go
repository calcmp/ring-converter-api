package main

// Input obj ...
type Input struct {
	Input string `json:"input"`
}

// Conversion obj ...
type Conversion struct {
	InsideDiameterInches string
	InsideCircumInches   string
	InsideDiameterMillis string
	InsideCircumMillis   string
	SizeUSA              string
	SizeBritish          string
	SizeFrench           string
	SizeGerman           string
	SizeJapanese         string
	SizeSwiss            string
}

// SizeMap obj ...
type SizeMap struct {
	size string
}
