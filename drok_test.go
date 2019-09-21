package drok

import "testing"

func TestFormatQuery(t *testing.T) {
	tables := []struct {
		input  float32
		output string
	}{
		{5.25, "0525"},
		{61, "6100"},
		{0, "0000"},
		{.99, "0099"},
	}

	for _, table := range tables {
		got := formatQuery(table.input)
		if got != table.output {
			t.Errorf("formatQuery(%f) = %s; want %s", table.input, got, table.output)
		}
	}
}

func TestParseOutput(t *testing.T) {
	tables := []struct {
		input  []byte
		output float32
	}{
		{[]byte("#ro00000000001"), 1},
		{[]byte("#ro00000000000"), 0},
		{[]byte("#ru00000000523"), 5.23},
		{[]byte("#ru00000005320"), 53.20},
		{[]byte("#ri00000000024"), 0.24},
		{[]byte("#ri00000000124"), 1.24},
		{[]byte("#ri00000000000"), 0},
		{[]byte("gibberish"), 0},
		{[]byte("#wook"), 1},
	}

	for _, table := range tables {
		got, _ := parseOutput(table.input)
		if got != table.output {
			t.Errorf("formatQuery(%s) = %f; want %f", string(table.input), got, table.output)
		}
	}
}
