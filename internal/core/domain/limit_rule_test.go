package domain_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gastonbordet/notification_service/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetTimeUnit(t *testing.T) {
	type GetTimeUnitTemplateTest struct {
		Title            string
		ExpectedDuration time.Duration
		Given            *domain.NotificationType
	}

	tests := []*GetTimeUnitTemplateTest{{
		Title:            "Should return Status type unit",
		ExpectedDuration: time.Second,
		Given: &domain.NotificationType{
			Type: "Status",
			Limit: &domain.LimitRule{
				Rate:       1,
				Unit:       "second",
				UnitAmount: 1,
				Enabled:    true,
			},
		},
	}, {
		Title:            "Should return News type unit",
		ExpectedDuration: time.Hour,
		Given: &domain.NotificationType{
			Type: "Status",
			Limit: &domain.LimitRule{
				Rate:       1,
				Unit:       "hour",
				UnitAmount: 1,
				Enabled:    true,
			},
		},
	}, {
		Title:            "Should return Marketing type unit",
		ExpectedDuration: time.Minute,
		Given: &domain.NotificationType{
			Type: "Status",
			Limit: &domain.LimitRule{
				Rate:       1,
				Unit:       "minute",
				UnitAmount: 1,
				Enabled:    true,
			},
		},
	}, {
		Title:            "Should return Alert type unit",
		ExpectedDuration: time.Hour * 24,
		Given: &domain.NotificationType{
			Type: "Status",
			Limit: &domain.LimitRule{
				Rate:       1,
				Unit:       "day",
				UnitAmount: 1,
				Enabled:    true,
			},
		},
	}, {
		Title:            "Should return Default type unit",
		ExpectedDuration: time.Minute,
		Given: &domain.NotificationType{
			Type: "Status",
			Limit: &domain.LimitRule{
				Rate:       1,
				Unit:       "",
				UnitAmount: 1,
				Enabled:    true,
			},
		},
	}}

	for _, test := range tests {
		unit := test.Given.Limit.GetTimeUnit()
		assert.Equal(t, test.ExpectedDuration, unit, fmt.Sprintf("Test assertion failed: %s", test.Title))
	}
}

func TestTimeExcedeed(t *testing.T) {
	type TimeExcedeedTemplateTest struct {
		Title        string
		Expected     bool
		GivenDate    string
		GivenNowDate string
		GivenRule    *domain.LimitRule
	}

	tests := []*TimeExcedeedTemplateTest{{
		Title:        "Should return true when time limit is exceeded",
		Expected:     true,
		GivenDate:    "2023-11-16T21:01:06-03:00",
		GivenNowDate: "2023-11-16T21:01:06-03:00",
		GivenRule: &domain.LimitRule{
			Rate:       1,
			Unit:       "minute",
			UnitAmount: 2,
			Enabled:    true,
		},
	}, {
		Title:        "Should return false when time limit is not exceeded",
		Expected:     false,
		GivenDate:    "2023-11-16T18:01:06-03:00",
		GivenNowDate: "2023-11-16T21:01:06-03:00",
		GivenRule: &domain.LimitRule{
			Rate:       1,
			Unit:       "hour",
			UnitAmount: 1,
			Enabled:    true,
		},
	}}

	for _, test := range tests {
		now, _ := time.Parse(time.RFC3339, test.GivenNowDate)
		result := test.GivenRule.TimeExceeded(test.GivenDate, now)
		assert.Equal(t, test.Expected, result, fmt.Sprintf("Test assertion failed: %s", test.Title))
	}
}
