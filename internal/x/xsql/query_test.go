package xsql

import "testing"

func TestListQuery(t *testing.T) {
	tests := []struct {
		query string
		ns    []int
		want  string
	}{
		{"", []int{}, ""},
		{"x", []int{1, 2, 3}, "x"},
		{"{{?}}", []int{1}, "?"},
		{"{{$$}}", []int{3}, "$$,$$,$$"},
		{"({{?}}) and ({{(?, ?)}})", []int{2, 1}, "(?,?) and ((?, ?))"},
		{"select * from users where id in ({{?}}) or name in ({{?}})", []int{2, 3},
			"select * from users where id in (?,?) or name in (?,?,?)"},
	}

	for _, test := range tests {
		actual := ListQuery(test.query, test.ns...)
		if actual != test.want {
			t.Errorf("ListQuery(%q, %v) = %q (want %q)", test.query, test.ns, actual, test.want)
		}
	}
}
