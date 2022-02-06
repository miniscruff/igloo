package mathf

import "math"

type EaseFunc func(float64) float64

func EaseLinear(t float64) float64 {
	return t
}

func EaseInQuad(t float64) float64 {
	return t * t
}

func EaseOutQuad(t float64) float64 {
	return -t * (t - 2)
}

func EaseInOutQuad(t float64) float64 {
	if t < 0.5 {
		return 2 * t * t
	} else {
		t = 2*t - 1
		return -0.5 * (t*(t-2) - 1)
	}
}

func EaseInCubic(t float64) float64 {
	return t * t * t
}

func EaseOutCubic(t float64) float64 {
	t -= 1
	return t*t*t + 1
}

func EaseInOutCubic(t float64) float64 {
	t *= 2
	if t < 1 {
		return 0.5 * t * t * t
	} else {
		t -= 2
		return 0.5 * (t*t*t + 2)
	}
}

func EaseInQuart(t float64) float64 {
	return t * t * t * t
}

func EaseOutQuart(t float64) float64 {
	t -= 1
	return -(t*t*t*t - 1)
}

func EaseInOutQuart(t float64) float64 {
	t *= 2
	if t < 1 {
		return 0.5 * t * t * t * t
	} else {
		t -= 2
		return -0.5 * (t*t*t*t - 2)
	}
}

func EaseInQuint(t float64) float64 {
	return t * t * t * t * t
}

func EaseOutQuint(t float64) float64 {
	t -= 1
	return t*t*t*t*t + 1
}

func EaseInOutQuint(t float64) float64 {
	t *= 2
	if t < 1 {
		return 0.5 * t * t * t * t * t
	} else {
		t -= 2
		return 0.5 * (t*t*t*t*t + 2)
	}
}

func EaseInSine(t float64) float64 {
	return -1*math.Cos(t*math.Pi/2) + 1
}

func EaseOutSine(t float64) float64 {
	return math.Sin(t * math.Pi / 2)
}

func EaseInOutSine(t float64) float64 {
	return -0.5 * (math.Cos(math.Pi*t) - 1)
}

func EaseInExpo(t float64) float64 {
	if t == 0 {
		return 0
	} else {
		return math.Pow(2, 10*(t-1))
	}
}

func EaseOutExpo(t float64) float64 {
	if t == 1 {
		return 1
	} else {
		return 1 - math.Pow(2, -10*t)
	}
}

func EaseInOutExpo(t float64) float64 {
	if t == 0 {
		return 0
	} else if t == 1 {
		return 1
	} else {
		if t < 0.5 {
			return 0.5 * math.Pow(2, (20*t)-10)
		} else {
			return 1 - 0.5*math.Pow(2, (-20*t)+10)
		}
	}
}

func EaseInCirc(t float64) float64 {
	return -1 * (math.Sqrt(1-t*t) - 1)
}

func EaseOutCirc(t float64) float64 {
	t -= 1
	return math.Sqrt(1 - (t * t))
}

func EaseInOutCirc(t float64) float64 {
	t *= 2
	if t < 1 {
		return -0.5 * (math.Sqrt(1-t*t) - 1)
	} else {
		t = t - 2
		return 0.5 * (math.Sqrt(1-t*t) + 1)
	}
}

func EaseInElastic(t float64) float64 {
	return EaseInElasticFunc(0.5)(t)
}

func EaseOutElastic(t float64) float64 {
	return EaseOutElasticFunc(0.5)(t)
}

func EaseInOutElastic(t float64) float64 {
	return EaseInOutElasticFunc(0.5)(t)
}

func EaseInElasticFunc(period float64) EaseFunc {
	p := period
	return func(t float64) float64 {
		t -= 1
		return -1 * (math.Pow(2, 10*t) * math.Sin((t-p/4)*(2*math.Pi)/p))
	}
}

func EaseOutElasticFunc(period float64) EaseFunc {
	p := period
	return func(t float64) float64 {
		return math.Pow(2, -10*t)*math.Sin((t-p/4)*(2*math.Pi/p)) + 1
	}
}

func EaseInOutElasticFunc(period float64) EaseFunc {
	p := period
	return func(t float64) float64 {
		t *= 2
		if t < 1 {
			t -= 1
			return -0.5 * (math.Pow(2, 10*t) * math.Sin((t-p/4)*2*math.Pi/p))
		} else {
			t -= 1
			return math.Pow(2, -10*t)*math.Sin((t-p/4)*2*math.Pi/p)*0.5 + 1
		}
	}
}

func EaseInBack(t float64) float64 {
	s := 1.70158
	return t * t * ((s+1)*t - s)
}

func EaseOutBack(t float64) float64 {
	s := 1.70158
	t -= 1
	return t*t*((s+1)*t+s) + 1
}

func EaseInOutBack(t float64) float64 {
	s := 1.70158
	t *= 2
	if t < 1 {
		s *= 1.525
		return 0.5 * (t * t * ((s+1)*t - s))
	} else {
		t -= 2
		s *= 1.525
		return 0.5 * (t*t*((s+1)*t+s) + 2)
	}
}

func EaseInBounce(t float64) float64 {
	return 1 - EaseOutBounce(1-t)
}

func EaseOutBounce(t float64) float64 {
	if t < 4/11.0 {
		return (121 * t * t) / 16.0
	} else if t < 8/11.0 {
		return (363 / 40.0 * t * t) - (99 / 10.0 * t) + 17/5.0
	} else if t < 9/10.0 {
		return (4356 / 361.0 * t * t) - (35442 / 1805.0 * t) + 16061/1805.0
	} else {
		return (54 / 5.0 * t * t) - (513 / 25.0 * t) + 268/25.0
	}
}

func EaseInOutBounce(t float64) float64 {
	if t < 0.5 {
		return EaseInBounce(2*t) * 0.5
	} else {
		return EaseOutBounce(2*t-1)*0.5 + 0.5
	}
}
