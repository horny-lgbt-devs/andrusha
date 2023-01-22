// https://github.com/zekroTJA/ken/tree/master/middlewares/ratelimit
package andrusharatelimit

import "time"

type LimitedCommand interface {
	LimiterBurst() int

	LimiterRestoration() time.Duration

	IsLimiterGlobal() bool
}
