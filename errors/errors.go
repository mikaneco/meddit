package errors

import "errors"

var (
	ErrInvalidAppointmentTime = errors.New("invalid appointment time")
	ErrInvalidCounselor       = errors.New("invalid counselor")
	ErrInvalidUser            = errors.New("invalid user")
	ErrInvalidCounselingMenu  = errors.New("invalid counseling menu")
	ErrAppointmentOverlapping = errors.New("appointment overlapping")
	ErrCounselorNotAvailable  = errors.New("counselor not available")
)
