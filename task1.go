package main

import (
	"fmt"
)

// Function to calculate average grade
func calculateAverage(grades []float64) float64 {
	var total float64
	for _, grade := range grades {
		total += grade
	}
	return total / float64(len(grades))
}

func main() {
	var studentName string
	var numSubjects int

	fmt.Println(" Student Grade Calculator ")

	// Input: student name
	fmt.Print("Enter your name: ")
	fmt.Scanln(&studentName)

	// Input: number of subjects
	fmt.Print("Enter the number of subjects: ")
	fmt.Scanln(&numSubjects)

	// Validate number of subjects
	if numSubjects <= 0 {
		fmt.Println("Invalid number of subjects.")
		return
	}

	// Declare a map to store subject names and grades
	subjectGrades := make(map[string]float64)
	var grades []float64

	// Loop through each subject
	for i := 1; i <= numSubjects; i++ {
		var subjectName string
		var grade float64

		fmt.Printf("\nEnter name of subject #%d: ", i)
		fmt.Scanln(&subjectName)

		// Validate grade input
		for {
			fmt.Printf("Enter grade for %s (0 - 100): ", subjectName)
			fmt.Scanln(&grade)
			if grade >= 0 && grade <= 100 {
				break
			} else {
				fmt.Println("Invalid grade. Please enter a value between 0 and 100.")
			}
		}

		subjectGrades[subjectName] = grade
		grades = append(grades, grade)
	}

	// Calculate average
	average := calculateAverage(grades)

	// Display results
	fmt.Println("\n====== Student Grade Summary ======")
	fmt.Printf("Student Name: %s\n", studentName)
	fmt.Println("-----------------------------------")
	for subject, grade := range subjectGrades {
		fmt.Printf("%-15s : %.2f\n", subject, grade)
	}
	fmt.Println("-----------------------------------")
	fmt.Printf("Average Grade  : %.2f\n", average)
	fmt.Println("===================================")

	// Optional: Grade evaluation
	if average >= 90 {
		fmt.Println("Excellent work! ")
	} else if average >= 75 {
		fmt.Println("Good job! ")
	} else if average >= 50 {
		fmt.Println("You passed, but there's room for improvement.")
	} else {
		fmt.Println("You need to work harder next time.")
	}
}
