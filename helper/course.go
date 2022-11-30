package helper

import "golang/models/dto"

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