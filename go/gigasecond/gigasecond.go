
package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
    const gigasecondDuration = 1e9 * time.Second
    return t.Add(gigasecondDuration)
}

