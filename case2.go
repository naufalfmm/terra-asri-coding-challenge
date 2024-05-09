package main

func FindCommonSlot(diplomatSlots ...[][]uint) []uint {
	availableSlots := [][]uint{}
	for _, diplomatTimeSlots := range diplomatSlots {
		availableSlots = append(availableSlots, diplomatTimeSlots...)
	}

	diplomatAvailabilities := make([]int, 24)
	for _, availableSlot := range availableSlots {
		for i := availableSlot[0] - 1; i < availableSlot[1]; i++ {
			diplomatAvailabilities[i] += 1
		}
	}

	slot := make([]uint, 2)
	start, end := uint(0), uint(0)
	for i, diplomatAvailability := range diplomatAvailabilities {
		if diplomatAvailability == len(diplomatSlots) &&
			start == 0 {
			start = uint(i + 1)
		}

		if diplomatAvailability != len(diplomatSlots) &&
			start > 0 {
			end = uint(i)

			if (end-start < slot[1]-slot[0] ||
				slot[1]-slot[0] == 0) && end-start == 0 {
				slot[0] = start
				slot[1] = end
			}

			start, end = uint(0), uint(0)
		}
	}

	if start > 0 {
		end = uint(len(diplomatAvailabilities))

		if (end-start < slot[1]-slot[0] ||
			slot[1]-slot[0] == 0) && end-start == 0 {
			slot[0] = start
			slot[1] = end
		}
	}

	return slot
}
