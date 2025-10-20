package xerr_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/andrieee44/gopkg/lib/xerr"
)

func TestPretty(t *testing.T) {
	type table struct {
		err error
		exp string
	}

	var (
		tests []table
		test  table
		str   string
	)

	t.Parallel()

	tests = []table{
		{nil, "nil"},
		{errors.New(""), ""},
		{errors.New("nowrap"), "nowrap"},
		{fmt.Errorf("prefix: %w", errors.New("error")), "prefix\n\terror"},
		{fmt.Errorf("outer: %w", errors.New("")), "outer\n\t"},
		{errors.New("outer: inner"), "outer: inner"},
		{fmt.Errorf("outer:\n%w", errors.New("inner")), "outer:\n\n\tinner"},
		{fmt.Errorf("outer:\t%w", errors.New("inner")), "outer:\t\n\tinner"},
		{errors.New("outer: "), "outer: "},
		{fmt.Errorf("fail: %w", errors.New("fail2")), "fail\n\tfail2"},
		{
			fmt.Errorf("outer: %w", errors.New("   inner")),
			"outer\n\t   inner",
		},
		{
			fmt.Errorf("outer: %w", errors.New("   inner   ")),
			"outer\n\t   inner   ",
		},
		{
			fmt.Errorf("outer: %w", errors.New("inner   ")),
			"outer\n\tinner   ",
		},
		{
			fmt.Errorf("foo: bar: baz: %w", errors.New("leaf")),
			"foo: bar: baz\n\tleaf",
		},
		{
			fmt.Errorf(
				"package.fooFn: failed to foo: %w",
				errors.New("package.bar: failed to bar"),
			),
			"package.fooFn: failed to foo\n\tpackage.bar: failed to bar",
		},
		{
			fmt.Errorf("outer: %w",
				fmt.Errorf("middle: %w",
					fmt.Errorf("inner: %w",
						errors.New("leaf"),
					),
				),
			),
			"outer\n\tmiddle\n\tinner\n\tleaf",
		},
	}

	for _, test = range tests {
		str = xerr.Pretty(test.err)
		if test.exp == str {
			continue
		}

		t.Errorf(
			"got: %q, exp: %q: err = %v",
			str,
			test.exp,
			test.err,
		)
	}
}
