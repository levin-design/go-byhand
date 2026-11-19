package main

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

var reportRows = []Record{
	{"apple", 2, 3.5},
	{"banana", 10, 12},
	{"a very long product name", 7, 0.99},
}

const wantReport = `NAME         QTY      PRICE
---------- ----- ----------
apple          2       3.50
banana        10      12.00
a very lon     7       0.99
`

func TestFormatReport(t *testing.T) {
	if got := FormatReport(reportRows); got != wantReport {
		t.Fatalf("FormatReport:\n got %q\nwant %q", got, wantReport)
	}
}

func TestAppendReport(t *testing.T) {
	got := string(AppendReport([]byte("HEAD\n"), reportRows))
	if got != "HEAD\n"+wantReport {
		t.Fatalf("AppendReport:\n got %q\nwant %q", got, "HEAD\n"+wantReport)
	}
}

func TestDump(t *testing.T) {
	var buf bytes.Buffer
	n, err := Dump(&buf, reportRows)
	if err != nil {
		t.Fatal(err)
	}
	if n != buf.Len() {
		t.Fatalf("n=%d, wrote %d bytes", n, buf.Len())
	}
	if buf.String() != wantReport {
		t.Fatalf("Dump:\n got %q\nwant %q", buf.String(), wantReport)
	}
}

func TestMoney(t *testing.T) {
	m := Money(1234)
	cases := []struct{ format, want string }{
		{"%v", "$12.34"},
		{"%s", "$12.34"},
		{"%+v", "12.34 CNY"},
		{"%d", "1234"},
		{"%.1f", "12.3"},
		{"%10.2f", "     12.34"},
	}
	for _, c := range cases {
		if got := fmt.Sprintf(c.format, m); got != c.want {
			t.Errorf("Sprintf(%q, Money(1234)) = %q, want %q", c.format, got, c.want)
		}
	}
}

func TestErrorChain(t *testing.T) {
	err := Load("config.yaml")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("errors.Is(err, ErrNotFound) = false, want true")
	}
	var le *LoadError
	if !errors.As(err, &le) || le.Path != "config.yaml" {
		t.Fatalf("errors.As did not recover *LoadError with Path config.yaml, got %#v", le)
	}
	if got, want := err.Error(), `load "config.yaml": open "config.yaml": not found`; got != want {
		t.Fatalf("Error() = %q, want %q", got, want)
	}
}

func TestBrokenChain(t *testing.T) {
	if errors.Is(BrokenLoad("x"), ErrNotFound) {
		t.Fatalf("BrokenLoad used %%v; errors.Is must be false")
	}
}

func TestBoth(t *testing.T) {
	err := Both()
	if !errors.Is(err, ErrA) || !errors.Is(err, ErrB) {
		t.Fatalf("multiple %%w should match both sentinels: %v", err)
	}
}
