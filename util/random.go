package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomPath() string {
	rand.Seed(time.Now().UnixNano())

	dir1 := RandomString(5)
	dir2 := RandomString(5)
	dir3 := RandomString(5)

	file := fmt.Sprintf("%d", rand.Intn(100000))

	return fmt.Sprintf("/%s/%s/%s/%s", dir1, dir2, dir3, file)
}

func RandomVersion() string {
	rand.Seed(time.Now().UnixNano())

	major := rand.Intn(10) + 1
	minor := rand.Intn(10)
	patch := rand.Intn(10)

	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
