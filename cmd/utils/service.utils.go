package utils

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"runtime"
	"time"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func PrintMemUsage() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// In Mib
	fmt.Printf("Alloc = %v MiB", bToMb(memStats.Alloc))
	fmt.Printf("\t TotalAlloc = %v MiB", bToMb(memStats.TotalAlloc))
	fmt.Printf("\t Sys = %v MiB \n", bToMb(memStats.Sys))
}

func ReleaseVariableMemory[T any](variable *T) {
	PrintMemUsage()

	var zero T
	// Clean variable for release memory
	*variable = zero

	log.Println("Calling Garbage Collector 🚜")
	runtime.GC()
	PrintMemUsage()
}

func GetYearsAgoPgFormat(years int) string {
	return time.Now().AddDate(-years, 0, 0).Format("2006-01-02")
}

func DateToUnix(date string) int64 {
	t, _ := time.Parse("2006-01-02", date)
	return t.Unix()
}

func initLogger(service string, component string) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})).With(
		"service", service,
		"component", component,
	)
}
