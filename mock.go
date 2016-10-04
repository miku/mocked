package mocked

import (
	"fmt"
	"math/rand"

	"github.com/miku/structs"
	"github.com/satori/go.uuid"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Mold is just some struct, will be filled with random data.
type Mold struct {
	PID       string `mocked:"uuid"`
	Generator string
	S1        string
	S2        string
	S3        string
	S4        string
	AR01      []int `length:"5-1000"`
	AR02      []int `length:"5-1000"`
	AR03      []int `length:"5-1000"`
	AR04      []int `length:"5-1000"`
	AT01      int
	AT02      int
	AT03      int
	AT04      int
	AT05      int
	AT06      int
	AT07      int
	AT08      int
	AT09      int
	AT10      int
	AT11      int
	AT12      int
	AT13      int
	AT14      int
	AT15      int
	AT16      int
	AT17      int
	AT18      int
	AT19      int
	AT20      int
	AT21      int
	AT22      int
	AT23      int
	AT24      int
	AT25      int
	AT26      int
	AT27      int
	AT28      int
	AT29      int
	AT30      int
	AT31      int
	AT32      int
	AT33      int
	AT34      int
	AT35      int
	AT36      int
	AT37      int
	AT38      int
	AT39      int
	AT40      int
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Random takes an interface (a struct) and tries to fill it with random data.
func Random(s interface{}) {
	st := structs.New(s)
	for k, v := range st.Map() {
		field := st.Field(k)
		switch v.(type) {
		case int:
			field.Set(rand.Intn(1000000))
		case string:
			switch field.Tag("mocked") {
			case "uuid":
				u := uuid.NewV4()
				field.Set(fmt.Sprintf("%s", u))
			default:
				field.Set(RandStringRunes(10))
			}
		case []int:
			var vals []int
			for i := 0; i < 100; i++ {
				vals = append(vals, rand.Intn(1000000))
			}
			field.Set(vals)
		}
	}
}
