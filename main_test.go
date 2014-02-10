package main

import "testing"

type testpair struct {
	arabic int
	roman  string
}

var tests = []testpair{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{999, "CMXCIX"},
	{1882, "MDCCCLXXXII"},
	{1883, "MDCCCLXXXIII"},
	{1884, "MDCCCLXXXIV"},
	{1885, "MDCCCLXXXV"},
	{1886, "MDCCCLXXXVI"},
	{1887, "MDCCCLXXXVII"},
	{1888, "MDCCCLXXXVIII"},
	{1889, "MDCCCLXXXIX"},
	{1890, "MDCCCXC"},
	{1999, "MCMXCIX"},
}

func TestA2r(t *testing.T) {
	for _, pair := range tests {
		v := a2r(pair.arabic)
		if v != pair.roman {
			t.Error(
				"For", pair.arabic,
				"expected", pair.roman,
				"got", v,
			)
		}
	}
}

func TestR2a(t *testing.T) {
	for _, pair := range tests {
		v := r2a(pair.roman)
		if v != pair.arabic {
			t.Error(
				"For", pair.roman,
				"expected", pair.arabic,
				"got", v,
			)
		}
	}
}
