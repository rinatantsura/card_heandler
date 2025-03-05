package pkg

import "errors"

func checkSum(num []int64) error {
	var sum int64
	j := 1
	for i := len(num) - 1; i >= 0; i-- {
		if j%2 == 0 {
			num[i] = num[i] * 2
			if num[i] >= 10 {
				num[i] = num[i]%10 + num[i]/10
			}
		}
		sum += num[i]
		j++
	}
	if sum%10 == 0 {
		return nil
	} else {
		return errors.New("invalid card number")
	}
}
