package helper

import (
	"golang/models/dto"
)

// get rating from course
func GetRatingCourse(course dto.Course) float64 {
	if len(course.Ratings) == 0 {
		return 0
	}
	// get rating of course
	for _, rating := range course.Ratings {
		course.Rating += float64(rating.Rating)
	}
	// average rating
	average := course.Rating / float64(len(course.Ratings))
	return average
}

func GetFavoriteCourse(course dto.Course, customerID string) bool {
	if len(course.Favorites) == 0 {
		return false
	}
	// get favorite of course
	for _, favorite := range course.Favorites {
		if favorite.CustomerID == customerID {
			return true
		}
	}
	return false
}

// function to get enrolled course
func GetEnrolledCourse(course *dto.Course, customerID string) {
	course.StatusEnroll = false
	// get enrolled of course
	for _, customerCourse := range course.CustomerCourses {
		if customerCourse.CustomerID == customerID {
			course.StatusEnroll = customerCourse.Status
			course.ProgressModule = customerCourse.NoModule
		}
	}
}

// function to get progress of course
func GetProgressCourse(course *dto.Course) float64 {
	var ProgressPercentage float64 = 0
	if course.NumberOfModules != 0 {
		ProgressPercentage = float64(course.ProgressModule - 1) * 100 / float64(course.NumberOfModules)
	}
	return ProgressPercentage
}