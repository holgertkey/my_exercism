package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birdCounter := 0
	for i := 0; i < len(birdsPerDay); i++ {
		birdCounter += birdsPerDay[i]
	}
	return birdCounter
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdCounter := 0
	start := 7 * (week - 1)
	end := start + 7
	for i := start; i < end; i++ {
		birdCounter += birdsPerDay[i]
	}
	return birdCounter
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i]++
    }
    return birdsPerDay
}
