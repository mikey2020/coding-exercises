package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    sum := 0
    for _, value := range birdsPerDay {
       sum += value 
    }
    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    end := 7 * week
    start := end - 7
    birdsInWeek := birdsPerDay[start:end]
    return TotalBirdCount(birdsInWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i, v := range birdsPerDay {
        j := i + 1
        if j % 2 != 0 {
            birdsPerDay[i] = v + 1
        }
    }
    return birdsPerDay
}
