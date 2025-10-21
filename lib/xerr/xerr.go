// Package xerr provides minimal error handling helpers for loud, readable
// diagnostics and fail-fast workflows.
//
// These utilities reduce boilerplate while preserving error semantics. They
// wrap, print, panic, or exit based on error presence, enabling concise,
// expressive control flow without sacrificing clarity.
//
// xerr favors explicit contracts and predictable behavior. It does not
// validate error quality or structure, it formats and reacts to whatever is
// given.
package xerr

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Pretty formats an error chain into a readable, newline-separated string.
//
// This helper unwraps the error using [errors.Unwrap] and prints each link
// on its own line. The first line is unindented; subsequent lines are
// indented with a single tab. Each link is derived by trimming the next
// error from the current error string, then stripping trailing colons and
// spaces.
//
// If the error is nil, it returns "nil". Empty error messages still produce
// a line. This function does not validate error quality or structure, it
// formats whatever it's given.
//
// Use Pretty to visualize wrapped errors without losing chain semantics or
// flooding logs with redundant output.
func Pretty(err error) string {
	var (
		builder strings.Builder
		nxtErr  error
		str     string
		first   bool
	)

	if err == nil {
		return "nil"
	}

	first = true

	for {
		if !first {
			builder.WriteString("\n\t")
		}

		nxtErr = errors.Unwrap(err)
		if nxtErr == nil {
			builder.WriteString(err.Error())

			return builder.String()
		}

		str = strings.TrimSuffix(err.Error(), nxtErr.Error())
		str = strings.TrimRight(str, " ")
		str = strings.TrimSuffix(str, ":")
		builder.WriteString(str)

		err = nxtErr
		first = false
	}
}

// ExitCodeIf prints the error to stderr and exits with the given code if
// the error is non-nil.
//
// This helper enforces fail-fast behavior with explicit exit codes. If err
// is nil, it does nothing. Otherwise, it prints the error to [os.Stderr] and
// terminates the process using [os.Exit](code).
//
// Note: [os.Exit] bypasses Go's normal shutdown sequence. Deferred
// functions will not run.
//
// Use ExitCodeIf when failure should abort execution with a specific code,
// especially in CLI tools, init routines, or scripts where error propagation
// is irrelevant.
func ExitCodeIf(err error, code int) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(code)
	}
}

// ExitIf prints the error to [os.Stderr] and exits with code 1 if the error
// is non-nil.
//
// This helper is a shorthand for [ExitCodeIf] with a default exit code of 1.
// If err is nil, it does nothing. Otherwise, it prints the error to
// [os.Stderr] and terminates the process.
//
// Note: [os.Exit] bypasses Go's normal shutdown sequence. Deferred
// functions will not run.
//
// Use ExitIf when failure should abort execution with a standard error code
// and no further cleanup is needed.
func ExitIf(err error) {
	ExitCodeIf(err, 1)
}

// PanicIf panics if the given error is non-nil.
//
// This helper enforces fail-fast behavior in workflows where errors must
// not propagate silently. If err is nil, it does nothing. Otherwise, it
// panics with the error value.
//
// Use PanicIf when failure is unrecoverable or diagnostic visibility is
// critical. It’s ideal for init-time checks, assertions, or scaffolding
// where error handling would be noise.
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

// WrapIf conditionally wraps a non-nil error with the given prefix.
//
// This helper reduces boilerplate when annotating failures for diagnostics.
// It preserves compatibility with [errors.Is] and [errors.As] by using
// [fmt.Errorf] with %w wrapping.
func WrapIf(prefix string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", prefix, err)
	}

	return nil
}

// WrapIf0 executes the provided function and conditionally wraps its failure.
//
// If the function signals failure, the result is annotated with the given
// prefix using [fmt.Errorf], preserving compatibility with [errors.Is] and
// [errors.As]. This helper reduces boilerplate in workflows that wrap
// failures for diagnostics without cluttering control flow.
func WrapIf0(prefix string, fn func() error) error {
	return WrapIf(prefix, fn())
}

// WrapIf1 executes the provided function and returns its result.
//
// If the function signals failure, the zero value of T is returned
// alongside a formatted diagnostic that includes the given prefix.
// This helper reduces repetitive conditional handling in workflows
// that produce a single result.
func WrapIf1[T any](prefix string, fn func() (T, error)) (T, error) {
	var (
		result, zero T
		err          error
	)

	result, err = fn()
	if err != nil {
		return zero, fmt.Errorf("%s: %w", prefix, err)
	}

	return result, nil
}

// WrapIf2 executes the provided function and returns its two results.
//
// If the function signals failure, the zero values of T1 and T2 are
// returned alongside a formatted diagnostic that includes the given
// prefix. This helper reduces repetitive conditional handling in
// workflows that produce two results.
func WrapIf2[T1, T2 any](prefix string, fn func() (T1, T2, error)) (T1, T2, error) {
	var (
		result1, zero1 T1
		result2, zero2 T2
		err            error
	)

	result1, result2, err = fn()
	if err != nil {
		return zero1, zero2, fmt.Errorf("%s: %w", prefix, err)
	}

	return result1, result2, nil
}
