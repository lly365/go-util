package util

import "testing"

func TestStr2Time(t *testing.T) {
	tt := Str2Time("2017-12-13 15:30:41")
	if tt != 1513150241 {
		t.Fatal("invalid timestamp")
	}

	if Str2TimeWithFormat("2017-12-13", DateOnlyFormat, "CST") != 1513094400 {
		t.Fatal("invalid timestamp")
	}

}

func TestTime2Str(t *testing.T) {
	s := Time2Str(1513150241)
	if s != "2017-12-13 15:30:41" {
		t.Fatal("invalid time string")
	}

	s = Time2StrWithFormat(1513150241, DateOnlyFormat)
	if s != "2017-12-13" {
		t.Fatal("invalid time string")
	}

	s = Time2StrWithFormat(1513150241, TimeOnlyFormat)
	if s!= "15:30:41" {
		t.Fatal("invalid time string")
	}
}
