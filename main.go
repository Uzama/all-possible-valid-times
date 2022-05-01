package main

import (
	"fmt"
	"strconv"
)

/*
Task: Given a list of 4 integers, find all possible valid 24 hour times (eg: 12:34) that the given integers can be assembled into and return the total number of valid times.
You can not use the same number twice.
Times such as 34:12 and 12:60 are not valid.
Provided integers can not be negative.
Notes: Input integers can not be negative.
Input integers can yeald 0 possible valid combinations.
Example:
	Input: [1, 2, 3, 4]
	Valid times: ["12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "23:14", "23:41", "21:34", "21:43"]
	Return: 10
*/

// assumtion: all integers must be within 0-9
func possibleTimes(digits []int) int {

	channel := make(chan []int)

	go func() {
		permutations(digits, len(digits), channel)
		close(channel)
	}()

	// creating a set data structure (will maintain unique values)
	validTimes := make(map[string]struct{})

	// recive permutations from the channel and check whether each permutation is a valid time
	for permutation := range channel {

		hh := strconv.Itoa(permutation[0]) + strconv.Itoa(permutation[1])
		mm := strconv.Itoa(permutation[2]) + strconv.Itoa(permutation[3])

		hours, _ := strconv.Atoi(hh)
		minutes, _ := strconv.Atoi(mm)

		// check whether hours are valid
		if hours > 23 || hours < 0 {
			continue
		}

		// check whether minutes are valid
		if minutes > 59 || minutes < 0 {
			continue
		}

		key := hh + ":" + mm

		validTimes[key] = struct{}{}
	}

	return len(validTimes)
}

// Uses Heap's algorithm to generate permutations
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func permutations(arr []int, n int, channel chan []int) {

	if n == 1 {

		out := make([]int, len(arr))
		copy(out, arr)

		// send the permutation via channel
		channel <- out

		return
	}

	for i := 0; i < n; i++ {

		permutations(arr, n-1, channel)

		if i < (n - 1) {

			if (n % 2) == 0 {
				swap(arr, i, n-1)
				continue
			}

			swap(arr, 0, n-1)
		}
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	// Example test cases.
	fmt.Println(possibleTimes([]int{1, 2, 3, 4}))
	fmt.Println(possibleTimes([]int{9, 1, 2, 0}))
	fmt.Println(possibleTimes([]int{2, 2, 1, 9}))
}
