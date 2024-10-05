package docker

import (
	"fmt"
	"strings"
	"time"
)

func humanReadableSize(bytes int64) string {
	const (
		KB = 1000
		MB = 1000 * KB
		GB = 1000 * MB
	)

	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.1fGB", float64(bytes)/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.1fMB", float64(bytes)/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.1fKB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%dB", bytes)
	}
}

func timeAgo(unixTime int64) string {
	now := time.Now()
	t := time.Unix(unixTime, 0)
	diff := now.Sub(t)

	switch {
	case diff < time.Minute:
		return fmt.Sprintf("%d seconds ago", int(diff.Seconds()))
	case diff < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(diff.Minutes()))
	case diff < time.Hour*24:
		return fmt.Sprintf("%d hours ago", int(diff.Hours()))
	case diff < time.Hour*24*7:
		return fmt.Sprintf("%d days ago", int(diff.Hours()/24))
	case diff < time.Hour*24*30*2:
		return fmt.Sprintf("%d weeks ago", int(diff.Hours()/24/7))
	case diff < time.Hour*24*365*2:
		return fmt.Sprintf("%d months ago", int(diff.Hours()/24/30))
	default:
		return fmt.Sprintf("%d years ago", int(diff.Hours()/24/365))
	}
}

func shortID(id string) string {
	return strings.Split(id, ":")[1][:12]
}
