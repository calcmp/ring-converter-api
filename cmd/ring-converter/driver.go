package main

import (
	"encoding/json"
	"fmt"
	"log"
	"ring-converter/maps"
	"strconv"
)

func convertSizes(input string) *Conversion {
	pos, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	newConversion := &Conversion{
		InsideDiameterInches: maps.InsideDiameterInches[pos],
		InsideDiameterMillis: maps.InsideDiameterMillis[pos],
		InsideCircumInches:   maps.CircumferenceInch[pos],
		InsideCircumMillis:   maps.CircumferenceMillis[pos],
		SizeBritish:          maps.BritishMap[pos],
		SizeUSA:              maps.UsaMap[pos],
		SizeFrench:           maps.FrenchMap[pos],
		SizeGerman:           maps.GermanMap[pos],
		SizeJapanese:         maps.JapaneseMap[pos],
		SizeSwiss:            maps.SwissMap[pos],
	}
	return newConversion
}

func returnMap(input string) string {
	var mapToSend map[int]string

	switch input {
	case "British":
		mapToSend = maps.BritishMap
	case "French":
		mapToSend = maps.FrenchMap
	case "German":
		mapToSend = maps.GermanMap
	case "Japanese":
		mapToSend = maps.JapaneseMap
	case "Swiss":
		mapToSend = maps.SwissMap
	case "USA":
		mapToSend = maps.UsaMap
	case "Diameter (inch)":
		mapToSend = maps.InsideDiameterInches
	case "Diameter (mm)":
		mapToSend = maps.InsideDiameterMillis
	case "Circumference (inch)":
		mapToSend = maps.CircumferenceInch
	case "Circumference (mm)":
		mapToSend = maps.CircumferenceMillis
	}

	data, err := json.Marshal(mapToSend)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonStr := string(data)

	return jsonStr
}
