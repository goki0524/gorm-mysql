package utils

import (
	"fmt"
	"strconv"
)

func ConvertIntToString(i *int) string {
	return fmt.Sprint(*i)
}

func ConvertStringToInt(s *string) (int, error) {
	num, err := strconv.Atoi(*s)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func ConvertInt64ToString(i *int64) string {
	return fmt.Sprint(*i)
}

func ConvertStringToInt64(s *string) (int64, error) {
	num, err := strconv.Atoi(*s)
	if err != nil {
		return 0, err
	}
	return int64(num), nil
}
