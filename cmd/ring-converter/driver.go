package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func convertSizes(input string) *Conversion {
	log.Println(input)
	pos, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	newConversion := &Conversion{
		InsideDiameterInches: InsideDiameterInches[pos],
		InsideDiameterMillis: InsideDiameterMillis[pos],
		InsideCircumInches:   CircumferenceInch[pos],
		InsideCircumMillis:   CircumferenceMillis[pos],
		SizeBritish:          BritishMap[pos],
		SizeUSA:              UsaMap[pos],
		SizeFrench:           FrenchMap[pos],
		SizeGerman:           GermanMap[pos],
		SizeJapanese:         JapaneseMap[pos],
		SizeSwiss:            SwissMap[pos],
	}
	return newConversion
}

func returnMap(input string) string {
	var mapToSend map[int]string

	switch input {
	case "British":
		mapToSend = BritishMap
	case "French":
		mapToSend = FrenchMap
	case "German":
		mapToSend = GermanMap
	case "Japanese":
		mapToSend = JapaneseMap
	case "Swiss":
		mapToSend = SwissMap
	case "USA":
		mapToSend = UsaMap
	case "DiamInch":
		mapToSend = InsideDiameterInches
	case "DiamMM":
		mapToSend = InsideDiameterMillis
	case "CircInch":
		mapToSend = CircumferenceInch
	case "CircMM":
		mapToSend = CircumferenceMillis
	}

	data, err := json.Marshal(mapToSend)
	if err != nil {
		fmt.Println(err.Error())
	}

	jsonStr := string(data)

	return jsonStr
}
