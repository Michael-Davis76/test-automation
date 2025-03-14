package go_unit_test_bootcamp

func FindMissingDrone(droneIds []int) int {
	if len(droneIds) == 0 {
		return -1
	}
	// slice to store how many times drone has come back
	checker := make([]int, len(droneIds))
	// loop to iterate through the droneIds
	for id := 0; id < len(droneIds); id++ {
		// increment the value of the id in the checker slice
		if checker[id] == 0 {
			for u := id + 1; u < len(droneIds); u++ {
				if droneIds[id] == droneIds[u] {
					checker[id]++
					checker[u]++
				}
			}
		}
	}
	for i := 0; i < len(checker); i++ {
		// if the value of the id is 0, return the id
		if checker[i] == 0 {
			return droneIds[i]
		}
	}
	// if no missing drone is found,
	return -1
}
