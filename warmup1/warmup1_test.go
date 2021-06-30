package warmup1

import "testing"

func TestSleepIn(t *testing.T) {
	var tests = []struct {
		w    bool
		v    bool
		want bool
	}{
		{false, false, true},
		{false, true, true},
		{true, false, false},
		{true, true, true},
	}
	for _, test := range tests {
		if got := SleepIn(test.w, test.v); got != test.want {
			t.Errorf("sleepIn(%v, %v) = %v [wanted: %v]", test.w, test.v, got, test.want)
		}
	}
}

func TestMonkeyTrouble(t *testing.T) {
	var tests = []struct {
		aSmile bool
		bSmile bool
		want bool
	}{
		{true, true, true},
		{false, false, true},
		{true, false, false},
		{false, true, false},
	}
	for _, test := range tests {
		if got := MonkeyTrouble(test.aSmile, test.bSmile); got != test.want {
			t.Errorf("MonkeyTrouble(%v, %v) = %v [wanted: %v]", test.aSmile, test.bSmile, got, test.want)
		}
	}
}


func TestSumDouble(t *testing.T) {
	var tests = []struct{
		a int
		b int
		want int
	}{
		{1,2,3},
		{3,2,5},
		{2,2,8},
		{-1,0,-1},
		{3,3,12},
		{0,0,0},
		{0,1,1},
		{3,4,7},
	}
	for _, test := range tests {
		if got := Sumdouble(test.a, test.b); got != test.want {
			t.Errorf("MonkeyTrouble(%v, %v) = %v [wanted: %v]", test.a, test.b, got, test.want)
		}
	}
}


func TestDiff21(t *testing.T) {
	var tests = []struct{
		n int
		want int
	}{
		{19,2},
		{10,11},
		{21,0},
		{22,2},
		{25,8},
		{30,18},
		{0,21},
		{1,20},
		{2,19},
		{-1,22},
		{-2,23},
		{50,58},
	}
	for _, test := range tests {
		if got := Diff21(test.n); got != test.want {
			t.Errorf("Diff21(%v) = %v [wanted: %v]", test.n, got, test.want)
		}
	}
}


func TestParrotTrouble(t *testing.T) {
	var tests = []struct {
		talking bool
		hour int
		want   bool
	}{
		{true, 6, true},
		{true, 7, false},
		{false, 6, false},
		{true, 21, true},
		{false, 21, false},
		{false, 20, false},
		{true, 23, true},
		{false, 23, false},
		{true, 20, false},
		{false, 12, false},
	}
	for _, test := range tests {
		if got := ParrotTrouble(test.talking, test.hour); got != test.want {
			t.Errorf("ParrotTrouble(%v, %d) = %v [wanted %v]", test.talking, test.hour, got, test.want)
		}
	}
}


func TestMakes10(t *testing.T) {
	var tests = []struct {
		a int
		b int
		want bool
	}{
		{9, 10, true},
		{9, 9, false},
		{1, 9, true},
		{10, 1, true},
		{10, 10, true},
		{8, 2, true},
		{8, 3, false},
		{10, 42, true},
		{12, -2, true},
	}
	for _, test := range tests {
		if got := Makes10(test.a, test.b); got != test.want {
			t.Errorf("Makes10(%d, %d) = %v [wanted %v]", test.a, test.b, got, test.want)
		}
	}
}


func TestNearHundred(t *testing.T) {
	var tests = []struct {
		n int
		want  bool
	}{
		{93, true},
		{90, true},
		{89, false},
		{0, false},
		{-101, false},
		{-209, false},
		{209, true},
		{211, false},
		{191, true},
		{200, true},
		{100, true},
	}
	for _, test := range tests {
		if got := NearHundred(test.n); got != test.want {
			t.Errorf("NearHundred(%d) = %v [wanted: %v]", test.n, got, test.want)
		}
	}
}



func TestPosNeg(t *testing.T) {
	var tests = []struct {
		a int
		b int
		negative bool
		want   bool
	}{
		{1, -1, false, true},
		{1, 1, false, false},
		{-1, 1, false, true},
		{-4, -5, true, true},
		{-4, 5, true, false},
		{-4, 5, false, true},
	}
	for _, test := range tests {
		if got := PosNeg(test.a, test.b, test.negative); got != test.want {
			t.Errorf("PosNeg(%d, %d, %v) = %v [wanted %v]", test.a, test.b, test.negative, got, test.want)
		}
	}
}


func TestNotString(t *testing.T) {
	var tests = []struct {
		str string
		want  string
	}{
		{"candy", "not candy"},
		{"x", "not x"},
		{"not bad", "not bad"},
		{"bad", "not bad"},
		{"is not", "not is not"},
		{"no", "not no"},
	}
	for _, test := range tests {
		if got := NotString(test.str); got != test.want {
			t.Errorf("NotString(%q) = %q [wanted: %q]", test.str, got, test.want)
		}
	}
}


func TestMissingChar(t *testing.T) {
	var tests = []struct {
		str string
		n int
		want   string
	}{
		{"kitten", 1, "ktten"},
		{"kitten", 0, "itten"},
		{"kitten", 4, "kittn"},
		{"Hi", 0, "i"},
		{"Hi", 1, "H"},
		{"code", 0, "ode"},
		{"code", 1, "cde"},
		{"code", 2, "coe"},
		{"code", 3, "cod"},
		{"chocolate", 8, "chocolat"},
	}
	for _, test := range tests {
		if got := MissingChar(test.str, test.n); got != test.want {
			t.Errorf("MissingChar(%q, %d) = %q [wanted %q]", test.str, test.n, got, test.want)
		}
	}
}


func TestFrontBack(t *testing.T) {
	var tests = []struct {
		str string
		want  string
	}{
		{"code", "eodc"},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"", ""},
		{"Chocolate", "ehocolatC"},
		{"aavJ", "Java"},
		{"hello", "oellh"},
	}
	for _, test := range tests {
		if got := FrontBack(test.str); got != test.want {
			t.Errorf("FrontBack(%q) = %q [wanted %q]", test.str, got, test.want)
		}
	}
}


func TestFront3(t *testing.T) {
	var tests = []struct {
		str string
		want  string
	}{
		{"Java", "JavJavJav"},
		{"Chocolate", "ChoChoCho"},
		{"abc", "abcabcabc"},
		{"abcXYZ", "abcabcabc"},
		{"ab", "ababab"},
		{"a", "aaa"},
		{"", ""},
	}
	for _, test := range tests {
		if got := Front3(test.str); got != test.want {
			t.Errorf("Front3(%q) = %q [wanted %q]", test.str, got, test.want)
		}
	}
}


func TestBackAround(t *testing.T) {
	var tests = []struct{
		str string
		want string
	}{
		{"cat","tcatt"},
		{"Hello","oHelloo"},
		{"a","aaa"},
		{"abc","cabcc"},
		{"read","dreadd"},
		{"boo","obooo"},
	}

	for _, test := range tests {
		if got := BackAround(test.str); got != test.want {
			t.Errorf("BackAround(%q) = %q [wanted %q]", test.str, got, test.want)
		}
	}
}


func TestOr35(t *testing.T) {
	var tests = []struct{
		n int
		want bool
	}{
		{3, true},
		{10, true},
		{8, false},
		{15, true},
		{5, true},
		{9, true},
		{4, false},
		{7, false},
		{6, true},
		{17, false},
		{18,true},
		{29, false},
		{20, true},
		{21, true},
		{22, false},
		{45, true},
		{99, true},
		{100, true},
		{101, false},
		{121, false},
		{122, false},
		{123, true},
	}

	for _, test := range tests {
		if got := Or35(test.n); got != test.want {
			t.Errorf("Or35(%v) = %v [wanted %v]", test.n, got, test.want)
		}
	}
}


func TestFront22(t *testing.T) {
	var tests = []struct{
		str string
		want string
	}{
		{"kitten", "kikittenki"},
		{"Ha", "HaHaHa"},
		{"abc","ababcab"},
		{"ab","ababab"},
		{"a","aaa"},
		{"",""},
		{"Logic","LoLogicLo"},
	}

	for _, test := range tests {
		if got := Front22(test.str); got != test.want {
			t.Errorf("Front22(%q) = %v [wanted %v]", test.str, got, test.want)
		}
	}
}


func TestStartHi(t *testing.T) {
	var tests = []struct{
		str string
		want bool
	}{
		{"hi there",true},
		{"hi", true},
		{"hello hi", false},
		{"he", false},
		{"h", false},
		{"", false},
		{"ho hi", false},
		{"hi ho", true},
	}

	for _, test := range tests {
		if got := StartHi(test.str); got != test.want {
			t.Errorf("StartHi(%q) = %v [wanted %v]", test.str, got, test.want)
		}
	}
}


func TestIcyHot(t *testing.T) {
	var tests = []struct{
		temp1 int
		temp2 int
		want bool
		}{
		{120,-1,true},
		{-1,120, true},
		{2,120, false},
		{-1,100, false},
		{-2,-2, false},
		{120,100, false},
	}

	for _, test := range tests {
		if got := IcyHot(test.temp1, test.temp2); got != test.want {
			t.Errorf("IcyHot(%v, %v) = %v [wanted %v]", test.temp1, test.temp2, got, test.want)
		}
	}
}


func TestIn1020(t *testing.T) {
	var tests = []struct{
		a,b int
		want bool
	}{
		{12,99,true},
		{21,12,true},
		{8,99,false},
		{99,10,true},
		{20,20,true},
		{21,21,false},
		{9,9,false},
	}

	for _, test := range tests {
		if got := In1020(test.a, test.b); got != test.want {
			t.Errorf("In1020(%v, %v) = %v [wanted %v]", test.a, test.b, got, test.want)
		}
	}
}


func TestHasTeen(t *testing.T) {
	var tests = []struct{
		a,b,c int
		want bool
	}{
		{13,20,10,true},
		{20,19,10,true},
		{20,10,13,true},
		{1,20,12,false},
		{19,20,12,true},
		{12,20,19,true},
		{12,9,20,false},
		{12,18,20,true},
		{14,2,20,true},
		{4,2,20,false},
		{11,22,22,false},
	}

	for _, test := range tests {
		if got := HasTeen(test.a, test.b, test.c); got != test.want {
			t.Errorf("HasTeen(%v, %v, %v) = %v [wanted %v]", test.a, test.b, test.c, got, test.want)
		}
	}
}



func TestLoneTeen(t *testing.T) {
	var tests = []struct{
		a,b int
		want bool
	}{
		{13,99,true},
		{21,19,true},
		{13,13,false},
		{14,20,true},
		{20,15,true},
		{16,17,false},
		{16,9,true},
		{16,18,false},
		{13,19,false},
		{13,20,true},
		{6,18,true},
		{99,13,true},
		{99,99,false},
	}

	for _, test := range tests {
		if got := LoneTeen(test.a, test.b); got != test.want {
			t.Errorf("LoneTeen(%v, %v) = %v [wanted %v]", test.a, test.b, got, test.want)
		}
	}
}


func TestDelDel(t *testing.T) {
	var tests = []struct{
		str string
		want string
	}{
		{"adelbc","abc"},
		{"adelHello","aHello"},
		{"adedbc","adedbc"},
		{"abcdel","abcdel"},
		{"add","add"},
		{"ad","ad"},
		{"a","a"},
		{"",""},
		{"del","del"},
		{"adel","a"},
		{"aadelbb","aadelbb"},
	}

	for _, test := range tests {
		if got := DelDel(test.str); got != test.want {
			t.Errorf("DelDel(%v) = %v [wanted %v]", test.str, got, test.want)
		}
	}
}


func TestMixStart(t *testing.T) {
	var tests = []struct{
		str string
		want bool
	}{
		{"mix snacks",true},
		{"pix snacks",true},
		{"piz snacks",false},
		{"nix",true},
		{"ni",false},
		{"n",false},
		{"",false},
	}

	for _, test := range tests {
		if got := MixStart(test.str); got != test.want {
			t.Errorf("MixStart(%v) = %v [wanted %v]", test.str, got, test.want)
		}
	}
}



func TestStartOz(t *testing.T) {
	var tests = []struct{
		str string
		want string
	}{
		{"ozymandias","oz"},
		{"bzoo","z"},
		{"oxx","o"},
		{"oz","oz"},
		{"ounce","o"},
		{"o","o"},
		{"abc",""},
		{"",""},
		{"zoo",""},
		{"aztec","z"},
		{"zzzz","z"},
		{"oznic","oz"},
	}

	for _, test := range tests {
		if got := StartOz(test.str); got != test.want {
			t.Errorf("StartOz(%v) = %v [wanted %v]", test.str, got, test.want)
		}
	}
}


func TestIntMax(t *testing.T) {
	var tests = []struct{
		a, b, c, want int
	}{
		{1, 2, 3, 3},
		{1, 3, 2, 3},
		{3, 2, 1, 3},
		{9, 3, 3, 9},
		{3, 3, 9, 9},
		{8, 2, 3, 8},
		{-3, -1, -2, -1},
		{6, 2, 5, 6},
		{5, 6, 2, 6},
		{5, 2, 6, 6},
	}

	for _, test := range tests {
		if got := IntMax(test.a, test.b, test.c); got != test.want {
			t.Errorf("IntMax(%v, %v, %v) = %v [wanted %v]", test.a, test.b, test.c, got, test.want)
		}
	}
}

func TestClose10(t *testing.T) {
	var tests = []struct {
		a, b, want int
	}{
		{8, 13, 8},
		{13, 8, 8},
		{13, 7, 0},
		{7, 13, 0},
		{9, 13, 9},
		{13, 8, 8},
		{10, 12, 10},
		{11, 10, 10},
		{5, 21, 5},
		{0, 20, 0},
		{10, 10, 0},
	}

	for _, test := range tests {
		if got := Close10(test.a, test.b); got != test.want {
			t.Errorf("Close10(%v, %v) = %v [wanted %v]", test.a, test.b, got, test.want)
		}
	}
}


func TestIn3050(t *testing.T) {
	var tests = []struct{
		a, b int
		want bool
	}{
		{30, 31, true},
		{30, 41, false},
		{40, 50, true},
		{40, 51, false},
		{39, 50, false},
		{50, 39, false},
		{40, 39, true},
		{49, 48, true},
		{50, 40, true},
		{50, 51, false},
		{35, 36, true},
		{35, 45, false},
	}

	for _, test := range tests {
		if got := In3050(test.a, test.b); got != test.want {
			t.Errorf("In3050(%v, %v) = %v [wanted %v]", test.a, test.b, got, test.want)
		}
	}
}


func TestMax1020(t *testing.T) {
	var tests = []struct{
		a, b, want int
	}{
		{11, 19, 19},
		{19, 11, 19},
		{11, 9, 11},
		{9, 21, 0},
		{10, 21, 10},
		{21, 10, 10},
		{9, 11, 11},
		{23, 10, 10},
		{20, 10, 20},
		{7, 20, 20},
		{17, 16, 17},
	}

	for _, test := range tests {
		if got := Max1020(test.a, test.b); got != test.want {
			t.Errorf("Max1020(%v, %v) = %v [wanted %v]", test.a, test.b, got, test.want)
		}
	}
}


func TestStringE(t *testing.T) {
	var tests = []struct{
		str string
		want bool
	}{
		{"Hello", true},
		{"Heelle", true},
		{"Heelele", false},
		{"Hll", false},
		{"e", true},
		{"", false},
	}

	for _, test := range tests {
		if got := StringE(test.str); got != test.want {
			t.Errorf("StringE(%v) = %v [wanted %v]", test.str, got, test.want)
		}
	}
}


func TestLastDigit(t *testing.T) {
	var tests = []struct{
		a, b int
		want bool
	}{
		{7, 17, true},
		{6, 17, false},
		{3, 113, true},
		{114, 113, false},
		{114, 4, true},
		{10, 0, true},
		{11, 0, false},
	}

	for _, test := range tests {
		if got := LastDigit(test.a, test.b); got != test.want {
			t.Errorf("LastDigit(%v, %v) = %v [wanted %v]", test.a, test.b, got, test.want)
		}
	}
}


func TetEndUp(t *testing.T) {
	var tests = []struct{
		str, want string
	}{
		{"Hello", "HeLLO"},
		{"hi there", "hi thERE"},
		{"hi", "HI"},
		{"woo hoo", "woo HOO"},
		{"xyz12", "xyZ12"},
		{"x", "X"},
		{"", ""},
	}

	for _, test := range tests {
		if got := EndUp(test.str); got != test.want {
			t.Errorf("Endup(%v) = %v [wanted %v]", test.str, got, test.want)
		}
	}
}


func TestEveryNth(t *testing.T) {
	var tests = []struct{
		str string
		n int
    want string
	}{
		{"Miracle", 2, "Mrce"},
		{"abcdefg", 2, "aceg"},
		{"abcdefg", 3, "adg"},
		{"Chocolate", 3, "Cca"},
		{"Chocolates", 3, "Ccas"},
		{"Chocolates", 4, "Coe"},
		{"Chocolates", 100, "C"},
	}

	for _, test := range tests {
		if got := EveryNth(test.str, test.n); got != test.want {
			t.Errorf("EveryNth(%v, %v) = %v [wanted %v]", test.str, test.n, got, test.want)
		}
	}
}
