// https://github.com/zekroTJA/ken/tree/master/middlewares/ratelimit
package andrusharatelimit

import "time"

type Limiter struct {
	burst       int
	restoration time.Duration

	tokens         int
	lastActivation time.Time
}

func NewLimiter(burst int, restoration time.Duration) *Limiter {
	return new(Limiter).setParams(burst, restoration)
}

func (l *Limiter) setParams(burst int, restoration time.Duration) *Limiter {
	l.burst = burst
	l.restoration = restoration
	l.tokens = burst
	l.lastActivation = time.Time{}
	return l
}

func (l *Limiter) Take() (ok bool, next time.Duration) {
	tokens := l.getVirtualTokens()
	if tokens == 0 {
		next = l.restoration - time.Since(l.lastActivation)
		return
	}

	l.tokens = tokens - 1
	l.lastActivation = time.Now()
	ok = true

	return
}

func (l *Limiter) getVirtualTokens() int {
	tokens := int(time.Since(l.lastActivation)/l.restoration) + l.tokens
	if tokens > l.burst {
		return l.burst
	}
	return tokens
}
