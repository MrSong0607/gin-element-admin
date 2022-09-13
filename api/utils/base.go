package utils

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	BeijingTimeZone = time.FixedZone("Beijing Time", int((8 * time.Hour).Seconds()))
	rd              = rand.New(rand.NewSource(time.Now().UnixNano()))
	TimeFormat      = `2006-01-02 15:04:05`
)

func GetUUID() []byte {
	b, _ := uuid.New().MarshalBinary()
	return b
}

func GetUUIDString() string {
	return uuid.New().String()
}

func Hex(b []byte) string {
	return hex.EncodeToString(b)
}

func HexDecode(b string) ([]byte, error) {
	return hex.DecodeString(b)
}

func Base64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// Round 四舍五入
func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

// Floor 下取整
func Floor(x, unit float64) float64 {
	return math.Floor(x/unit) * unit
}

// Ceil 上取整
func Ceil(x, unit float64) float64 {
	return math.Ceil(x/unit) * unit
}

func RandNumber(min, max int) int {
	return min + rd.Intn(max-min+1)
}

func StartOfDay(t time.Time) time.Time {
	t = t.In(BeijingTimeZone)
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func ContainString(src string, target ...string) bool {
	for _, val := range target {
		if val == src {
			return true
		}
	}

	return false
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func BitOn(src, des uint64) bool {
	return (src & des) > 0
}

func RemoveAt[T any](s []T, i int) []T {
	if len(s)-1 < i || i < 0 {
		return s
	}

	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func JoinURL(base string, p ...string) (s string, e error) {
	u, e := url.Parse(base)
	if e != nil {
		return
	}

	for _, val := range p {
		u.Path = path.Join(u.Path, val)
	}

	s = u.String()
	return
}

func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func ParseIPList(s string) string {
	s = strings.TrimSpace(s)

	dealIP := func(s []string) []string {
		var r []string
		for _, str := range s {
			if net.ParseIP(str) != nil {
				r = append(r, str)
			}
		}
		return r
	}

	return strings.Join(dealIP(DeleteEmpty(strings.Split(s, " "))), " ")
}

func ValidIP(ip, whiteList string) bool {
	if len(ip) == 0 {
		return false
	}

	if len(whiteList) > 0 && !strings.Contains(whiteList, ip) {
		return false
	}

	return true
}

func Distinct(src []string) []string {
	result := make([]string, 0, len(src))
	temp := map[string]struct{}{}
	for _, item := range src {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func Contains[T comparable](src T, dst ...T) bool {
	for _, v := range dst {
		if src == v {
			return true
		}
	}
	return false
}
