package timejkt

import "time"

// Now ...
func Now() time.Time {
	return time.Now().In(Location())
}

// Location ...
func Location() *time.Location {
	return time.FixedZone("Asia/Jakarta", 7*60*60)
}
