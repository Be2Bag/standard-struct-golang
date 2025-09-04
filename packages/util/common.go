package util

import (
	"net"
	"os"
	"strconv"
	"time"
)

// แปลง string เป็น duration
func ParseDuration(t string) time.Duration {
	d, _ := time.ParseDuration(t)
	return d
}

// ค้นหา Env ด้วย Key คืนค่า fallback เมื่อหาไม่เจอ
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// ตรวจสอบว่า item ใน slice ไหม
func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

// แปลง string เป็น int
func AtoI(s string, v int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return v
	}
	return i
}

// แปลง string เป็น float64
func AtoF(s string, v float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return v
	}
	return f
}

// ตรวจสอบว่าเป็น IPv4 ไหม
func IPv4Tester(ip string) bool {
	return net.ParseIP(ip) != nil
}
