// progress.go
package main

import (
	"fmt"
)

// Progress represents a simple progress bar
type Progress struct {
	total   int
	current int
}

// NewProgress initializes a new progress bar
func NewProgress(total int) *Progress {
	return &Progress{total: total, current: 0}
}

// Increment increases the current progress by one step
func (p *Progress) Increment() {
	if p.current < p.total {
		p.current++
		p.display()
	}
}

// display shows the current progress in the console
func (p *Progress) display() {
	percentage := float64(p.current) / float64(p.total) * 100
	barLength := 50
	progress := int(percentage / 100 * float64(barLength))

	fmt.Printf("\r[")
	for i := 0; i < barLength; i++ {
		if i < progress {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %.2f%%", percentage)
}

// Complete finalizes the progress display
func (p *Progress) Complete() {
	p.current = p.total
	p.display()
	fmt.Println("\nBuild process completed!")
}
