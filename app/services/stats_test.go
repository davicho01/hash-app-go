package services

import (
	"testing"
)

func TestStats(t *testing.T) {

	statsService := NewStats()

	statsService.AddStats(10)
	statsService.AddStats(7)
	statsService.AddStats(3)

	stats := statsService.GetStats()

	if stats.Total != 3 {
		t.Error("Stats Total not 3, Found:", stats.Total)
		return
	}

	if stats.Average != 6.67 {
		t.Error("Stats Average not 6.67, Found:", stats.Average)
		return
	}
}
