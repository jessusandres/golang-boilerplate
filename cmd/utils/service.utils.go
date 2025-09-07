package utils

import (
	"log"
	"runtime"
	"time"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func PrintMemUsage() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// Show metrics in Mib
	log.Printf("Alloc = %v MiB", bToMb(memStats.Alloc))
	log.Printf("\t TotalAlloc = %v MiB", bToMb(memStats.TotalAlloc))
	log.Printf("\t Sys = %v MiB \n", bToMb(memStats.Sys))
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
