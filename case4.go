package main

func D(n uint) uint {
	sum := n

	for n > 0 {
		sum += (n % 10)
		n /= 10
	}

	return sum
}

func position(firstElement, positionElement uint) uint {
	return positionElement - firstElement + 1
}

func arithmeticSequenceSum(firstElement, positionElement uint) uint {
	return (position(firstElement, positionElement) / 2) * (firstElement + positionElement)
}

func SumSelfNumbers(min, max uint) uint {
	var (
		existingSum   uint
		hasGenerators = map[uint]bool{}
	)

	for i := min; i <= max; i++ {
		digitSum := D(i)
		if _, ok := hasGenerators[digitSum]; !ok {
			hasGenerators[digitSum] = true
			existingSum += digitSum
		}
	}

	return arithmeticSequenceSum(min, max) - existingSum
}
