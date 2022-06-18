package main

import "fmt"

func main() {
	var bucket1, bucket2, bucket3 int
	fmt.Print("Enter the capacity of bucket 1: ")
	fmt.Scan(&bucket1)
	fmt.Print("Enter the capacity of bucket 2: ")
	fmt.Scan(&bucket2)
	fmt.Print("Enter the capacity of bucket 3: ")
	fmt.Scan(&bucket3)
	// round1, round2, round3 are variables that
	// show in which iteration we could use bucket to fill it
	// with one apple
	var round1, round2, round3 int = 1, 2, 3
	// iteration is a variable that shows current loop iteration
	var iteration int = 0
	for {
		iteration++
		// if current iteration is equal to iteration that noted to
		// bucket. We fill bucket with one apple in this iteration.
		// Bucket's capacity is decreased by one.
		// New noted iteration is increased by 3, that means after 2 next
		// iterations we will fill this bucket with one apple again.
		if iteration == round1 {
			// but if bucket is empty, we skip this iteration
			if bucket1 == 0 {
				continue
			}
			bucket1--
			round1 += 3
		}
		if iteration == round2 {
			if bucket2 == 0 {
				continue
			}
			bucket2--
			round2 += 3
		}
		if iteration == round3 {
			if bucket3 == 0 {
				continue
			}
			bucket3--
			round3 += 3
		}
		// if all buckets are empty, we break the loop
		if bucket1 == 0 && bucket2 == 0 && bucket3 == 0 {
			break
		}
	}
	fmt.Println("Total number of iterations that needed to spread apples into three buckets: ", iteration)
}
