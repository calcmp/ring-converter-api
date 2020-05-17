package main

// InputPos obj ...
type InputPos struct {
	Position int `json:"input"`
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
