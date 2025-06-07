package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
    parsedTime, err := time.Parse(layout, date)
    if err != nil {
        panic("Invalid date format")
    }
    return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
    appointmentTime, err := time.Parse(layout, date)
    if err != nil {
        panic("Invalid date format")
    }
    return appointmentTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    appointmentTime, err := time.Parse(layout, date)
    if err != nil {
        panic("Invalid date format")
    }
    hour := appointmentTime.Hour()
    return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    appointmentTime, err := time.Parse(layout, date)
    if err != nil {
        panic("Invalid date format")
    }
    return "You have an appointment on " + appointmentTime.Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
    dateStr := time.Date(currentYear, 9, 15, 0, 0, 0, 0, time.UTC)
    return dateStr
}
