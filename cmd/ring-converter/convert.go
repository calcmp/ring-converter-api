package main

func convertSizes(step int) Conversion {
	var newConversion Conversion

	newConversion.InsideDiameterInches = InsideDiameterInches[step]
	newConversion.InsideDiameterMillis = InsideDiameterMillis[step]
	newConversion.InsideCircumInches = CircumferenceInch[step]
	newConversion.InsideCircumMillis = CircumferenceMillis[step]

	newConversion.SizeBritish = britishMap[step]
	newConversion.SizeUSA = usaMap[step]
	newConversion.SizeFrench = frenchMap[step]
	newConversion.SizeGerman = germanMap[step]
	newConversion.SizeJapanese = japaneseMap[step]
	newConversion.SizeSwiss = swissMap[step]

	return newConversion
}
