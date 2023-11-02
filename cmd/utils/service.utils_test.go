package utils

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestBToMb(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected uint64
	}{
		{"Zero bytes", 0, 0},
		{"1024 bytes (1KB)", 1024, 0},
		{"1048576 bytes (1MB)", 1048576, 1},
		{"5242880 bytes (5MB)", 5242880, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := bToMb(tt.input)

			if result != tt.expected {
				t.Errorf("bToMb(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestPrintMemUsage(t *testing.T) {
	// Redirect stdout to capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintMemUsage()

	w.Close()
	// Reset stdout
	os.Stdout = oldStdout

	// Read captured output into a buffer
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Alloc =") {
		t.Errorf("PrintMemUsage() output doesn't contain 'Alloc =': %s", output)
	}

	if !strings.Contains(output, "TotalAlloc =") {
		t.Errorf("PrintMemUsage() output doesn't contain 'TotalAlloc =': %s", output)
	}

	if !strings.Contains(output, "Sys =") {
		t.Errorf("PrintMemUsage() output doesn't contain 'Sys =': %s", output)
	}
}

func TestReleaseVariableMemory(t *testing.T) {
	oldStdout := os.Stdout
	oldLog := log.Writer()
	r, w, _ := os.Pipe()
	os.Stdout = w
	log.SetOutput(w)

	testVar := "test string"
	ReleaseVariableMemory(&testVar)

	w.Close()
	os.Stdout = oldStdout
	log.SetOutput(oldLog)

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if testVar != "" {
		t.Errorf("ReleaseVariableMemory() didn't reset variable, got: %s", testVar)
	}

	if !strings.Contains(output, "Calling Garbage Collector") {
		t.Errorf("ReleaseVariableMemory() output doesn't contain garbage collector message: %s", output)
	}
}

func TestGetYearsAgoPgFormat(t *testing.T) {
	tests := []struct {
		name  string
		years int
	}{
		{"One year ago", 1},
		{"Five years ago", 5},
		{"Ten years ago", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetYearsAgoPgFormat(tt.years)

			_, err := time.Parse("2006-01-02", result)
			if err != nil {
				t.Errorf("GetYearsAgoPgFormat(%d) returned invalid date format: %s", tt.years, result)
			}

			expectedDate := time.Now().AddDate(-tt.years, 0, 0)

			parsedResult, _ := time.Parse("2006-01-02", result)

			diff := expectedDate.Sub(parsedResult)

			if diff < -24*time.Hour || diff > 24*time.Hour {
				t.Errorf("GetYearsAgoPgFormat(%d) returned unexpected date: got %s, expected around %s",
					tt.years, result, expectedDate.Format("2006-01-02"))
			}
		})
	}
}

func TestDateToUnix(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		expected int64
	}{
		{"Epoch start", "1970-01-01", 0},
		{"Sample date", "2023-05-15", 1684108800},
		{"Future date", "2030-12-31", 1924905600},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DateToUnix(tt.date)
			if result != tt.expected {
				t.Errorf("DateToUnix(%s) = %d; want %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestInitLogger(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := initLogger("test-service", "test-component")
	logger.Info("test message")

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "test-service") {
		t.Errorf("initLogger() didn't set service name correctly: %s", output)
	}

	if !strings.Contains(output, "test-component") {
		t.Errorf("initLogger() didn't set component name correctly: %s", output)
	}

	if !strings.Contains(output, "test message") {
		t.Errorf("Logger didn't output test message: %s", output)
	}
}
