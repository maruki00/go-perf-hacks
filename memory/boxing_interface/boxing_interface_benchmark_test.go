package boxinginterface

import "testing"


const MAX_ALLOC = 1_000_000

func BenchmarkBoxedLargeJobs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		workers := make([]Worker, 0, MAX_ALLOC)
		for j := 0; j < MAX_ALLOC; j++ {
			workers = append(workers, NewLargeJob())
		}
	}
}

func BenchmarkNonBoxedLargeJobs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jobs := make([]*LargeJob, 0, MAX_ALLOC)
		for j := 0; j < MAX_ALLOC; j++ {
			jobs = append(jobs, NewLargeJob())
		}
	}
}

func BenchmarkBoxedLargeJobsWithWork(b *testing.B) {
	for i := 0; i < b.N; i++ {
		workers := make([]Worker, 0, MAX_ALLOC)
		for j := 0; j < MAX_ALLOC; j++ {
			workers = append(workers, NewLargeJob())
		}
		for _, w := range workers {
			w.Work()
		}
	}
}

func BenchmarkNonBoxedLargeJobsWithWork(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jobs := make([]*LargeJob, 0, MAX_ALLOC)
		for j := 0; j < MAX_ALLOC; j++ {
			jobs = append(jobs, NewLargeJob())
		}
		for _, job := range jobs {
			job.Work()
		}
	}
}

