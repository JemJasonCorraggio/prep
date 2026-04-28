package main

import (
	"fmt"
	"slices"
)

type Status string

type BuildJob struct{
	developer string
	project string
	durationMinutes int
	status Status
}

const (
	SuccessStatus Status = "success"
	FailureStatus Status = "failed"
)


type Report struct {
	TotalBuilds int
	AverageDuration int
	FailedBuilds int
	SlowestProject []string
	MostActiveDeveloper []string
}

func AnalyzeBuilds(jobs []BuildJob) Report{
	report := Report{}

	if len(jobs) == 0 {
		return report
	}

	report.TotalBuilds = len(jobs)

	buildTimeByProject := make(map[string]int)
	activityByDeveloper := make(map[string]int)
	sum := 0

	for _, job := range jobs{
		sum += job.durationMinutes
		buildTimeByProject[job.project] = buildTimeByProject[job.project] + job.durationMinutes
		activityByDeveloper[job.developer] = activityByDeveloper[job.developer] + 1
		if job.status == FailureStatus {
			report.FailedBuilds++
		}
	}

	report.AverageDuration = sum / report.TotalBuilds

	slowestProjectTime := 0

	for k, v := range buildTimeByProject{
		if v > slowestProjectTime{
			report.SlowestProject = []string{k}
			slowestProjectTime = v
		} else if v == slowestProjectTime{
			report.SlowestProject = append(report.SlowestProject, k)
		}
	}

	mostActiveDeveloperCount := 0

	for k, v := range activityByDeveloper{
		if v > mostActiveDeveloperCount{
			report.MostActiveDeveloper = []string{k}
			mostActiveDeveloperCount = v
		} else if v == mostActiveDeveloperCount{
			report.MostActiveDeveloper = append(report.MostActiveDeveloper, k)
		}
	}

	slices.Sort(report.SlowestProject)
	slices.Sort(report.MostActiveDeveloper)

	return report
}

func main() {
	jobs1 := []BuildJob{
		{"Alice", "API", 12, "success"},
		{"Bob", "Web", 25, "failed"},
		{"Alice", "API", 8, "success"},
		{"Cara", "Mobile", 40, "success"},
		{"Bob", "Web", 30, "success"},
	}

	jobs2 := []BuildJob{}

	fmt.Print(AnalyzeBuilds(jobs1))
	fmt.Print(AnalyzeBuilds(jobs2))
}