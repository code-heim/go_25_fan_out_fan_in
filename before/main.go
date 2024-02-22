package main

import (
	imageprocessing "fan_out_fan_in/image_processing"

	"image"
	"strings"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) []Job {
	var jobs []Job

	// For each input path create job struct and add it to the list
	for _, p := range paths {
		job := Job{InputPath: p,
			OutPath: strings.Replace(p, "images/", "images/output/", 1)}
		job.Image = imageprocessing.ReadImage(p)
		jobs = append(jobs, job)
	}
	return jobs
}

func resize(jobs *[]Job) {
	// For each input job, change the image in the job struct
	for index := range *jobs {
		(*jobs)[index].Image = imageprocessing.Resize((*jobs)[index].Image)
	}
}

func saveImages(jobs *[]Job) {
	for _, job := range *jobs {
		imageprocessing.WriteImage(job.OutPath, job.Image)
	}
}

func main() {

	imagePaths := []string{"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	jobs := loadImage(imagePaths)
	resize(&jobs)
	saveImages(&jobs)

}

