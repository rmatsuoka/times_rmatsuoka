package diff

import (
	"fmt"
	"strings"
	"testing"
)

func TestOrderedMaps(t *testing.T) {
	tests := []struct {
		m, n string
		want string
	}{
		{m: "", n: "", want: ""},
		{m: "0=zero", n: "", want: "- 0: zero\n"},
		{m: "", n: "1=one", want: "+ 1: one\n"},
		{
			m: "0=zero 1=one 2=two 3=three 4=four",
			n: "0=zero 1=one 2=two 4=four",
			want: `  0: zero
  1: one
  2: two
- 3: three
  4: four
`,
		},
		{
			m: "0=zero 1=one 2=two 4=four",
			n: "0=zero 1=one 2=two 3=three 4=four",
			want: `  0: zero
  1: one
  2: two
+ 3: three
  4: four
`,
		},
		{
			m: "0=zero 1=one 2=two",
			n: "0=zero 1=1 2=two",
			want: `  0: zero
- 1: one
+ 1: 1
  2: two
`,
		},
	}

	for _, test := range tests {
		m := mustParseMapString(test.m)
		n := mustParseMapString(test.n)
		got := OrderedMaps(m, n)

		if got != test.want {
			t.Errorf("OrderedMap(%#v, %#v) returns %s, want %s", test.m, test.n, got, test.want)
		}
	}
}

func mustParseMapString(s string) map[string]string {
	m, err := parseMapString(s)
	if err != nil {
		panic(err)
	}
	return m
}

func parseMapString(s string) (map[string]string, error) {
	m := make(map[string]string)
	for _, kv := range strings.Fields(s) {
		k, v, ok := strings.Cut(kv, "=")
		if !ok {
			return nil, fmt.Errorf("failed to parse key value: %s", kv)
		}
		m[k] = v
	}
	return m, nil
}
