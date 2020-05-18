package main

func convertSizes(step int) *Conversion {
	newConversion := &Conversion{
		InsideDiameterInches: InsideDiameterInches[step],
		InsideDiameterMillis: InsideDiameterMillis[step],
		InsideCircumInches:   CircumferenceInch[step],
		InsideCircumMillis:   CircumferenceMillis[step],
		SizeBritish:          BritishMap[step],
		SizeUSA:              UsaMap[step],
		SizeFrench:           FrenchMap[step],
		SizeGerman:           GermanMap[step],
		SizeJapanese:         JapaneseMap[step],
		SizeSwiss:            SwissMap[step],
	}
	return newConversion
}
