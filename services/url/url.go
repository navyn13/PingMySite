package url

import (
	"fmt"
	"net/http"
	"time"
)

// Job represents a single URL to process
type Job struct {
	URL string
}

// Result holds the output of a processed Job
type Result struct {
	Output string
}

// GetMetrics fetches a URL and returns its response time and status
func (h *Handler) GetMetrics(url string) string {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("URL: %s | Error: %s", url, err.Error())
	}
	defer resp.Body.Close()

	responseTime := time.Since(start).Milliseconds()
	return fmt.Sprintf("URL: %s | ResponseTime: %dms | Status: %s",
		url, responseTime, http.StatusText(resp.StatusCode))
}

// worker function that processes jobs from the jobs channel
func (h *Handler) worker(jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		output := h.GetMetrics(job.URL)
		results <- Result{Output: output}
	}
}

// GetBatchMetrics runs GetMetrics concurrently on multiple URLs
func (h *Handler) GetBatchMetrics(urls []string) []string {
	numJobs := len(urls)
	numWorkers := 8

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// Start workers
	for w := 0; w < numWorkers; w++ {
		go h.worker(jobs, results)
	}

	// Send jobs
	for _, u := range urls {
		jobs <- Job{URL: u}
	}
	close(jobs)

	// Collect results
	batchResults := make([]string, 0, numJobs)
	for i := 0; i < numJobs; i++ {
		res := <-results
		batchResults = append(batchResults, res.Output)
	}

	return batchResults
}
