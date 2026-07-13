package matchstickproblem

import (
	"log"
	"reflect"
	"testing"
)

func makeDigit(on ...int) []bool {
	arr := make([]bool, 16)
	for _, i := range on {
		arr[i] = true
	}
	return arr
}

func concat(slices ...[]bool) []bool {
	var res []bool
	for _, s := range slices {
		res = append(res, s...)
	}
	return res
}

func TestIsANumber(t *testing.T) {
	d0 := makeDigit(0, 1, 2, 3, 4, 5, 6, 7)
	d1 := makeDigit(2, 3)
	d1Alt := makeDigit(6, 7)
	d0Slash := makeDigit(0, 1, 2, 3, 4, 5, 6, 7, 12, 13)
	d1Center := makeDigit(11, 14)
	d1Flag := makeDigit(10, 11, 14)
	d2Diag := makeDigit(0, 1, 4, 5, 12, 13)
	d7Cross := makeDigit(0, 1, 2, 3, 8, 9)
	d7Diag := makeDigit(0, 1, 12, 13)
	d2 := makeDigit(0, 1, 2, 8, 9, 6, 4, 5)
	d3 := makeDigit(0, 1, 2, 8, 9, 3, 4, 5)
	d4 := makeDigit(7, 8, 9, 2, 3)
	d5 := makeDigit(0, 1, 7, 8, 9, 3, 4, 5)
	d6 := makeDigit(0, 1, 7, 6, 4, 5, 3, 8, 9)
	d7 := makeDigit(0, 1, 2, 3)
	d8 := makeDigit(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	d9 := makeDigit(0, 1, 2, 3, 4, 5, 7, 8, 9)
	d11 := makeDigit(2, 3, 6, 7)
	empty := makeDigit()
	invalid := makeDigit(15) // K is not used alone in any standard digit here

	expected := []struct {
		b     int
		ok    bool
		input []bool
	}{
		{0, false, invalid},
		{0, false, empty},
		{1, true, d1},
		{1, true, d1Alt},
		{0, true, d0Slash},
		{1, true, d1Center},
		{1, true, d1Flag},
		{2, true, d2Diag},
		{7, true, d7Cross},
		{7, true, d7Diag},
		{2, true, d2},
		{3, true, d3},
		{4, true, d4},
		{5, true, d5},
		{6, true, d6},
		{7, true, d7},
		{8, true, d8},
		{9, true, d9},
		{0, true, d0},
		{11, true, d11},
		{11, true, concat(d1, d1)},
		{1111, true, concat(d1, d1, d1, d1)},
	}
	for i, each := range expected {
		if b, ok := isANumber(each.input); b != each.b || ok != each.ok {
			log.Printf("Failed on #%d (expected %d %v) got (%d %v)", i, each.b, each.ok, b, ok)
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
	d0 := makeDigit(0, 1, 2, 3, 4, 5, 6, 7)
	d1 := makeDigit(2, 3)
	d1Alt := makeDigit(6, 7)
	d0Slash := makeDigit(0, 1, 2, 3, 4, 5, 6, 7, 12, 13)
	d1Center := makeDigit(11, 14)
	d1Flag := makeDigit(10, 11, 14)
	d2Diag := makeDigit(0, 1, 4, 5, 12, 13)
	d7Cross := makeDigit(0, 1, 2, 3, 8, 9)
	d7Diag := makeDigit(0, 1, 12, 13)
	d2 := makeDigit(0, 1, 2, 8, 9, 6, 4, 5)
	d3 := makeDigit(0, 1, 2, 8, 9, 3, 4, 5)
	d4 := makeDigit(7, 8, 9, 2, 3)
	d5 := makeDigit(0, 1, 7, 8, 9, 3, 4, 5)
	d6 := makeDigit(0, 1, 7, 6, 4, 5, 3, 8, 9)
	d7 := makeDigit(0, 1, 2, 3)
	d8 := makeDigit(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	d9 := makeDigit(0, 1, 2, 3, 4, 5, 7, 8, 9)
	d11 := makeDigit(2, 3, 6, 7)
	empty := makeDigit()
	invalid := makeDigit(15)

	expected := []struct {
		b     string
		ok    bool
		input []bool
	}{
		{"", false, invalid},
		{"", true, empty},
		{"1", true, d1},
		{"1", true, d1Alt},
		{"0", true, d0Slash},
		{"1", true, d1Center},
		{"1", true, d1Flag},
		{"2", true, d2Diag},
		{"7", true, d7Cross},
		{"7", true, d7Diag},
		{"2", true, d2},
		{"3", true, d3},
		{"4", true, d4},
		{"5", true, d5},
		{"6", true, d6},
		{"7", true, d7},
		{"8", true, d8},
		{"9", true, d9},
		{"0", true, d0},
		{"11", true, d11},
	}
	for i, each := range expected {
		if b, ok := isADigit(each.input); string(b) != each.b || ok != each.ok {
			log.Printf("Failed on #%d (expected %s %v) got (%s %v)", i, each.b, each.ok, b, ok)
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
