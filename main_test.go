package main

import (
	"log"
	"testing"
)

func TestIsANumber(t *testing.T) {
	expected := []struct {
		b     int
		ok    bool
		input []bool
	}{
		{0, false, []bool{
			false,
			false, false,
			false,
			false, false,
			true,
		}},
		{0, false, []bool{
			false,
			false, false,
			false,
			false, false,
			false,
		}},
		{1, true, []bool{
			false,
			true, false,
			false,
			true, false,
			false,
		}},
		{1, true, []bool{
			false,
			false, true,
			false,
			false, true,
			false,
		}},
		{2, true, []bool{
			true,
			false, true,
			true,
			true, false,
			true,
		}},
		{3, true, []bool{
			true,
			false, true,
			true,
			false, true,
			true,
		}},
		{4, true, []bool{
			false,
			true, true,
			true,
			false, true,
			false,
		}},
		{5, true, []bool{
			true,
			true, false,
			true,
			false, true,
			true,
		}},
		{6, true, []bool{
			true,
			true, false,
			true,
			true, true,
			true,
		}},
		{7, true, []bool{
			true,
			false, true,
			false,
			false, true,
			false,
		}},
		{8, true, []bool{
			true,
			true, true,
			true,
			true, true,
			true,
		}},
		{9, true, []bool{
			true,
			true, true,
			true,
			false, true,
			true,
		}},
		{9, true, []bool{
			true,
			true, true,
			true,
			false, true,
			false,
		}},
		{0, true, []bool{
			true,
			true, true,
			false,
			true, true,
			true,
		}},
		{11, true, []bool{
			false,
			true, true,
			false,
			true, true,
			false,
		}},
		{1111, true, []bool{
			false,
			true, true,
			false,
			true, true,
			false,
			false,
			true, true,
			false,
			true, true,
			false,
		}},
	}
	for i, each := range expected {
		if b, ok := isANumber(each.input); b != each.b || ok != each.ok {
			log.Printf("Failed on #%d (expected %d) got (%d %v)", i, each.b, b, ok)
			t.Fail()
		}
	}
}

func TestIsADigit(t *testing.T) {
	expected := []struct {
		b     string
		ok    bool
		input []bool
	}{
		{"", false, []bool{
			false,
			false, false,
			false,
			false, false,
			true,
		}},
		{"", true, []bool{
			false,
			false, false,
			false,
			false, false,
			false,
		}},
		{"1", true, []bool{
			false,
			true, false,
			false,
			true, false,
			false,
		}},
		{"1", true, []bool{
			false,
			false, true,
			false,
			false, true,
			false,
		}},
		{"2", true, []bool{
			true,
			false, true,
			true,
			true, false,
			true,
		}},
		{"3", true, []bool{
			true,
			false, true,
			true,
			false, true,
			true,
		}},
		{"4", true, []bool{
			false,
			true, true,
			true,
			false, true,
			false,
		}},
		{"5", true, []bool{
			true,
			true, false,
			true,
			false, true,
			true,
		}},
		{"6", true, []bool{
			true,
			true, false,
			true,
			true, true,
			true,
		}},
		{"7", true, []bool{
			true,
			false, true,
			false,
			false, true,
			false,
		}},
		{"8", true, []bool{
			true,
			true, true,
			true,
			true, true,
			true,
		}},
		{"9", true, []bool{
			true,
			true, true,
			true,
			false, true,
			true,
		}},
		{"9", true, []bool{
			true,
			true, true,
			true,
			false, true,
			false,
		}},
		{"0", true, []bool{
			true,
			true, true,
			false,
			true, true,
			true,
		}},
		{"11", true, []bool{
			false,
			true, true,
			false,
			true, true,
			false,
		}},
	}
	for i, each := range expected {
		if b, ok := isADigit(each.input); string(b) != each.b || ok != each.ok {
			log.Printf("Failed on #%d (expected %s) got (%s %v)", i, each.b, b, ok)
			t.Fail()
		}
	}
}

func TestCountThem(t *testing.T) {
	expected := []struct {
		t     int
		f     int
		input []bool
	}{
		{0, 0, []bool{}},
		{1, 0, []bool{true}},
		{0, 1, []bool{false}},
		{1, 1, []bool{true, false}},
		{2, 1, []bool{true, false, true}},
		{3, 0, []bool{true, true, true}},
		{0, 3, []bool{false, false, false}},
	}
	for i, each := range expected {
		if tr, fr := countthem(each.input); tr != each.t || fr != each.f {
			log.Printf("Failed on #%d (expected %d, %d) got (%d, %d)", i, each.t, each.f, tr, fr)
			t.Fail()
		}
	}
}
