package util

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	randomStringCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	random      = rand.New(rand.NewSource(time.Now().UnixNano()))
	randomMutex = sync.Mutex{}
)

// FileExists checks if a file exists, and returns true if it does
func FileExists(filename string) bool {
	stat, _ := os.Stat(filename)
	return stat != nil
}

// InStringList returns true if needle is contained in haystack
func InStringList(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

// RandomString returns a random string with a given length
func RandomString(length int) string {
	randomMutex.Lock() // Who would have thought that random.Intn() is not thread-safe?!
	defer randomMutex.Unlock()
	b := make([]byte, length)
	for i := range b {
		b[i] = randomStringCharset[random.Intn(len(randomStringCharset))]
	}
	return string(b)
}

// DurationToHuman converts a duration to a human readable format
func DurationToHuman(d time.Duration) (str string) {
	if d == 0 {
		return "0"
	}

	d = d.Round(time.Second)
	days := d / time.Hour / 24
	if days > 0 {
		str += fmt.Sprintf("%dd", days)
	}
	d -= days * time.Hour * 24

	hours := d / time.Hour
	if hours > 0 {
		str += fmt.Sprintf("%dh", hours)
	}
	d -= hours * time.Hour

	minutes := d / time.Minute
	if minutes > 0 {
		str += fmt.Sprintf("%dm", minutes)
	}
	d -= minutes * time.Minute

	seconds := d / time.Second
	if seconds > 0 {
		str += fmt.Sprintf("%ds", seconds)
	}
	return
}
