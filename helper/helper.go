package helper

import (
	"strconv"
	"strings"
)

func ParseHospitalIDs(hospitalIDs string) []int {
	ids := strings.Split(hospitalIDs, ",")
	var result []int
	for _, id := range ids {
		parsedID, err := strconv.Atoi(id)
		if err == nil {
			result = append(result, parsedID)
		}
	}

	return result
}
