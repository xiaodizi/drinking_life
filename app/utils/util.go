package utils

import "crypto/md5"
import "encoding/hex"
import "time"
import "math/rand"
import "strings"
import "github.com/revel/revel"

func HashMd5(password string) string {
	encrypt := "drinking_life_fulei"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password + "-" + encrypt))
	cipherStr := md5Ctx.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}

func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

func GetSession(key string, session revel.Session) string {
	if value, ok := session[key]; ok {
		return value
	}
	return ""
}
