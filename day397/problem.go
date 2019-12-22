package day397

import "sort"

// Job has a start and end time.
type Job struct {
	Start, End int
}

// LargestSubsetCompatibleJobs returns the largest
// subset of compatible jobs. Compatible jobs are non-overlapping.
func LargestSubsetCompatibleJobs(jobs []Job) []Job {
	copied := append([]Job{}, jobs...)
	sort.Slice(copied, func(i, j int) bool {
		if copied[i].End < copied[j].End {
			return true
		} else if copied[i].End == copied[j].End {
			return copied[i].Start < copied[j].Start
		}

		return false
	})

	possible := make([][]Job, 0)

	for _, job := range copied {
		added := false

		for i, p := range possible {
			left := p[len(p)-1]

			if job.Start >= left.End {
				possible[i] = append(p, job)
			}
		}

		if !added {
			possible = append(possible, []Job{job})
		}
	}

	largestSize := 0

	var largest []Job

	for _, p := range possible {
		if len(p) > largestSize {
			largestSize = len(p)
			largest = p
		}
	}

	return largest
}
