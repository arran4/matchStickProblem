package main

import (
	"log"
	"reflect"
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

func TestCountthem(t *testing.T) {
	tests := []struct {
		input []bool
		wantT int
		wantF int
	}{
		{[]bool{}, 0, 0},
		{[]bool{true}, 1, 0},
		{[]bool{false}, 0, 1},
		{[]bool{true, false}, 1, 1},
		{[]bool{true, true, true}, 3, 0},
		{[]bool{false, false}, 0, 2},
		{[]bool{true, false, true, false, true}, 3, 2},
	}

	for i, tc := range tests {
		gotT, gotF := countthem(tc.input)
		if gotT != tc.wantT || gotF != tc.wantF {
			t.Errorf("Test #%d: countthem(%v) = (%d, %d); want (%d, %d)", i, tc.input, gotT, gotF, tc.wantT, tc.wantF)
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

func TestFindThem(t *testing.T) {
	tests := []struct {
		name      string
		input     []bool
		wantTrue  []int
		wantFalse []int
	}{
		{
			name:      "mixed",
			input:     []bool{true, false, true, false, true},
			wantTrue:  []int{0, 2, 4},
			wantFalse: []int{1, 3},
		},
		{
			name:      "all true",
			input:     []bool{true, true, true},
			wantTrue:  []int{0, 1, 2},
			wantFalse: nil,
		},
		{
			name:      "all false",
			input:     []bool{false, false},
			wantTrue:  nil,
			wantFalse: []int{0, 1},
		},
		{
			name:      "empty",
			input:     []bool{},
			wantTrue:  nil,
			wantFalse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTrue, gotFalse := findthem(tt.input)
			if !reflect.DeepEqual(gotTrue, tt.wantTrue) {
				t.Errorf("findthem() gotTrue = %v, want %v", gotTrue, tt.wantTrue)
			}
			if !reflect.DeepEqual(gotFalse, tt.wantFalse) {
				t.Errorf("findthem() gotFalse = %v, want %v", gotFalse, tt.wantFalse)
			}
		})
	}
}
