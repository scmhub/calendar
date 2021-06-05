package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
	https://en.wikipedia.org/wiki/Equinox
	event	equinox		solstice	equinox		solstice
	month	March		June		September	December
	year	day	time	day	time	day	time	day	time
	2020	20	03:50	20	21:43	22	13:31	21	10:03
	2021	20	09:37	21	03:32	22	19:21	21	15:59
	2022	20	15:33	21	09:14	23	01:04	21	21:48
	2023	20	21:25	21	14:58	23	06:50	22	03:28
	2024	20	03:07	20	20:51	22	12:44	21	09:20
	2025	20	09:02	21	02:42	22	18:20	21	15:03
*/
func TestSolsticesAndEquinoxes(t *testing.T) {
	assert := assert.New(t)
	// 2020
	assert.Equal(time.Date(2020, 3, 20, 3, 50, 0, 0, time.UTC), northwardEquinox(2020))
	assert.Equal(time.Date(2020, 6, 20, 21, 43, 0, 0, time.UTC), northernSolstice(2020))
	assert.Equal(time.Date(2020, 9, 22, 13, 31, 0, 0, time.UTC), southwardEquinox(2020))
	assert.Equal(time.Date(2020, 12, 21, 10, 03, 0, 0, time.UTC), southernSolstice(2020))
	// 2021
	assert.Equal(time.Date(2021, 3, 20, 9, 37, 0, 0, time.UTC), northwardEquinox(2021))
	assert.Equal(time.Date(2021, 6, 21, 3, 32, 0, 0, time.UTC), northernSolstice(2021))
	assert.Equal(time.Date(2021, 9, 22, 19, 21, 0, 0, time.UTC), southwardEquinox(2021))
	assert.Equal(time.Date(2021, 12, 21, 15, 59, 0, 0, time.UTC), southernSolstice(2021))
	// 2022
	assert.Equal(time.Date(2022, 3, 20, 15, 33, 0, 0, time.UTC), northwardEquinox(2022))
	assert.Equal(time.Date(2022, 6, 21, 9, 13, 0, 0, time.UTC), northernSolstice(2022)) // vs 9:14
	assert.Equal(time.Date(2022, 9, 23, 1, 4, 0, 0, time.UTC), southwardEquinox(2022))
	assert.Equal(time.Date(2022, 12, 21, 21, 48, 0, 0, time.UTC), southernSolstice(2022))
	// 2023
	assert.Equal(time.Date(2023, 3, 20, 21, 24, 0, 0, time.UTC), northwardEquinox(2023)) // vs 21:25
	assert.Equal(time.Date(2023, 6, 21, 14, 57, 0, 0, time.UTC), northernSolstice(2023)) // vs 14:58
	assert.Equal(time.Date(2023, 9, 23, 6, 50, 0, 0, time.UTC), southwardEquinox(2023))
	assert.Equal(time.Date(2023, 12, 22, 3, 27, 0, 0, time.UTC), southernSolstice(2023)) // vs 3:28
	// 2024
	assert.Equal(time.Date(2024, 3, 20, 3, 6, 0, 0, time.UTC), northwardEquinox(2024)) // vs 3:07
	assert.Equal(time.Date(2024, 6, 20, 20, 51, 0, 0, time.UTC), northernSolstice(2024))
	assert.Equal(time.Date(2024, 9, 22, 12, 43, 0, 0, time.UTC), southwardEquinox(2024)) // vs12:44
	assert.Equal(time.Date(2024, 12, 21, 9, 20, 0, 0, time.UTC), southernSolstice(2024))
	// 2025
	assert.Equal(time.Date(2025, 3, 20, 9, 1, 0, 0, time.UTC), northwardEquinox(2025)) // vs 9:02
	assert.Equal(time.Date(2025, 6, 21, 2, 42, 0, 0, time.UTC), northernSolstice(2025))
	assert.Equal(time.Date(2025, 9, 22, 18, 19, 0, 0, time.UTC), southwardEquinox(2025)) //vs 18:20
	assert.Equal(time.Date(2025, 12, 21, 15, 3, 0, 0, time.UTC), southernSolstice(2025))
}

func TestAbs(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1.0, abs(1.0))
	assert.Equal(0.0, abs(0.0))
	assert.Equal(1.0, abs(-1.0))
}
