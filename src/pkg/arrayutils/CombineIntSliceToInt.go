package arrayutils

import (
	"bytes"
	"fmt"
	"strconv"
)

func ConvertIntSliceToInt(nums []int) (int, error) {
	var buf bytes.Buffer

	for i := range nums {
		buf.WriteString(fmt.Sprintf("%d", nums[i]))
	}

	result, err := strconv.Atoi(buf.String())
	if err != nil {
		return -1, err
	}

	return result, nil
}
