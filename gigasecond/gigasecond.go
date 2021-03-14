package gigasecond

import "time"

const gigaSecond = time.Second * 1000000000

func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
