package main

// Size obj ...
type Size struct {
	Millis int `json:"mm,string"`
	Cms    int `json:"cm,string"`
}

// Conversion obj ...
type Conversion struct {
	InsideDiameterInches int
	InsideCircumInches   int
	InsideDiameterMillis int
	InsideCircumMillis   int
	SizeUSA              int
	SizeBritish          int
	SizeFrench           int
	SizeGerman           int
	SizeJapanese         int
	SizeSwiss            int
}
