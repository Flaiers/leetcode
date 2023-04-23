package main

import (
	"fmt"
)

type Interval struct {
	Start, End int
}

func canAttendMeetings(intervals []*Interval) bool {
	sort(intervals)
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i].End > intervals[i+1].Start {
			return false
		}
	}

	return true
}

func sort(s []*Interval) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i].Start < s[j].Start {
				t := s[i]
				s[i] = s[j]
				s[j] = t
			}
		}
	}
}

func main() {
	intervals := []*Interval{
		{Start: 0, End: 3},
		{Start: 5, End: 9},
		{Start: 2, End: 6},
		{Start: 6, End: 7},
	}
	fmt.Println(canAttendMeetings(intervals))
}
