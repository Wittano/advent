package three

import (
	"errors"
	"strconv"
)

type batteryPower []byte

func (b batteryPower) MaxPower() (int, error) {
	var (
		first, last byte
	)

	for i := 0; i < len(b)-1; i++ {
		prevFirst := first
		first = max(first, b[i])
		if prevFirst != first {
			last = 0
		}

		last = max(last, b[i+1])
	}

	return strconv.Atoi(strconv.Itoa(int(first)) + strconv.Itoa(int(last)))
}

func newBatteryPower(in string) (battery batteryPower, err error) {
	if len(in) == 0 {
		return nil, errors.New("input is empty")
	}

	battery = make(batteryPower, len(in))

	for i, c := range in {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			return nil, err
		}

		battery[i] = byte(num)
	}

	return
}

func SumBatteryPower(in []string) (sum uint, err error) {
	for _, raw := range in {
		battery, err := newBatteryPower(raw)
		if err != nil {
			return 0, err
		}

		power, err := battery.MaxPower()
		if err != nil {
			return 0, err
		}

		sum += uint(power)
	}

	return
}
