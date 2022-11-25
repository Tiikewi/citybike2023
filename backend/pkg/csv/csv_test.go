package csv

import (
	"bufio"
	"citybike/pkg/types"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func BenchmarkProcessFile(b *testing.B) {
// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		ProcessFile()
// 	}
// }

func TestConvertInts(t *testing.T) {
	strNums := []string{"1", "10", "20"}
	expectedNums := []int{1, 10, 20}

	result, err := convertInts(strNums[:])
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, result, expectedNums)

	strNums = []string{"1", "1s", "20"}
	_, err = convertInts(strNums)
	if err == nil {
		t.Error("int conversion didn't return error")
	}
}

func TestWriteJourneyCSV(t *testing.T) {
	fileName := "data/test-valitated-journeys.csv"

	os.Setenv("ENVIROMENT", "test")
	defer os.Unsetenv("ENVIROMENT")

	journey := types.Journey{
		DepTime:        "2008-10-29 14:56:59",
		RetTime:        "2008-10-29 16:56:59",
		DepStationId:   1,
		DepStationName: "test station 1",
		RetStationId:   2,
		RetStationName: "test station 2",
		Distance:       100,
		Duration:       200,
	}

	journeys := []*types.Journey{&journey}

	writeJourneyCSV(journeys)

	f, err := os.Open(fileName)
	if err != nil {
		t.Error(err)
	}

	s := bufio.NewScanner(f)
	s.Scan()
	line := s.Text()

	expectedStr := "2008-10-29 14:56:59," +
		"2008-10-29 16:56:59,1," +
		"test station 1,2,test station 2,100,200"
	assert.Equal(t, line, expectedStr)
}

func TestValidateTime(t *testing.T) {
	time := "2006-01-02T15:04:05"

	result := validateTime(time)
	assert.Equal(t, result, true)

	time = "200601-02T15:04:05"
	result = validateTime(time)
	assert.Equal(t, result, false)

	time = "2006-01-02"
	result = validateTime(time)
	assert.Equal(t, result, false)
}
