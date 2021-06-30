package main

type GENDER int8

const (
	MALE GENDER = iota
	FEMALE
	UNKNOWN
)

var genderText = map[GENDER]string{
	MALE:    "MALE",
	FEMALE:  "FEMALE",
	UNKNOWN: "UNKNOWN",
}

func (gender GENDER) String() string {
	return genderText[gender]
}
