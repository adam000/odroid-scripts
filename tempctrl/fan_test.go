package main

import "testing"

func TestGetHourDifference(t *testing.T) {
	tests := []struct {
		Hour1 int
		Hour2 int
		Diff  int
	}{
		{14, 2, 12},
		{2, 14, 12},
		{14, 14, 0},
		{4, 22, 6},
		{22, 4, 6},
		{2, 22, 4},
		{22, 2, 4},
		{2, 3, 1},
		{3, 2, 1},
		{0, 5, 5},
		{5, 0, 5},
		{0, 23, 1},
		{23, 0, 1},
	}

	for _, test := range tests {
		result := getHourDistance(test.Hour1, test.Hour2)
		if result != test.Diff {
			t.Errorf("Expected distance from %d to %d to be %d, got %d", test.Hour1, test.Hour2, test.Diff, result)
		}
	}
}

func TestGetLowFanTriggerTemp(t *testing.T) {
	config := Config{
		FanTrigger: NightDayVariance{
			Night: 0,
			Day:   100,
		},
	}

	tests := []struct {
		ZenithDist int
		FanLevel   int
	}{
		{0, 100},
		{12, 0},
		{6, 50},
	}

	for _, test := range tests {
		result := getLowFanTriggerTemp(config, test.ZenithDist)
		if result != test.FanLevel {
			t.Errorf("At ZenithDist %d, expected fan to be at %d but it was at %d", test.ZenithDist, test.FanLevel, result)
		}
	}
}
