package example

// PID Controller, see https://en.wikipedia.org/wiki/PID_controller
type PID struct {
	Target         float64
	OutMax, OutMin float64 // min and max output value

	// P accounts for present values of the error.
	// For example, if the error is large and positive, the control output will also be large and positive.
	kp float64

	// I accounts for past values of the error.
	// For example, if the current output is not sufficiently strong,
	// the integral of the error will accumulate over time,
	// and the controller will respond by applying a stronger action.
	ki float64

	// D accounts for possible future trends of the error, based on its current rate of change.
	kd float64

	accErr    float64 // accumulated error value
	lastInput float64
}

func NewPID(kp, ki, kd, min, max float64) *PID {
	return &PID{
		kp:     kp,
		ki:     ki,
		kd:     kd,
		OutMin: min,
		OutMax: max,
	}
}

func clamp(input float64, min float64, max float64) float64 {
	if input > max {
		return max
	}
	if input < min {
		return min
	}
	return input
}

func (p *PID) Compute(input float64) (output float64) {
	errVal := p.Target - input
	p.accErr = clamp(p.accErr+errVal, p.OutMin, p.OutMax)
	deltaInput := input - p.lastInput
	// Kp * error value + Ki * accumulated error + Kd * delta input
	output = clamp(p.kp*errVal+p.ki*p.accErr+p.kd*deltaInput, p.OutMin, p.OutMax)
	p.lastInput = input
	return
}
