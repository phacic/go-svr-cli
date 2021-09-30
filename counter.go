package main

import (
	"errors"
)

// CountUp counts from zero to stop
func CountUp(stop int, verbose bool) ([]int, error) {
	if stop <= 0 {
		return nil, errors.New("stop cannot be negative")
	}

	var s []int
	Logger.Info("counting up to ... ", stop)
	for i := 1; i <= stop; i++ {
		if verbose {
			Logger.Info(i)
		}
		s = append(s, i)
	}

	return s, nil
}

// CountDown counts from start to zero
func CountDown(start int, verbose bool) ([]int, error) {
	if start <= 0 {
		return nil, errors.New("start cannot be negative")
	}

	var s []int
	Logger.Info("counting up from ... ", start)
	for i := start; i > 0; i-- {
		if verbose {
			Logger.Info(i)
		}
		s = append(s, i)
	}
	return s, nil
}
