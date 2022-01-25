package schema

import (
	"database/sql/driver"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"math"
	"street/pkg/utils"
	"strings"
)

type IDMixin struct {
	mixin.Schema
}

type ID string

func (i ID) Value() (driver.Value, error) {

	return i.ToInt64()
}

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (i *ID) ToInt64() (int64, error) {
	str := string(*i)
	result := int64(0)
	for i := 0; i < len(str); i++ {
		result *= int64(len(alphabet))
		value := int64(strings.IndexByte(alphabet, str[i]))
		if value == -1 {
			return 0, errors.New("not a valid base62")
		}
		result += value
	}
	return result, nil
}
func (i *ID) Scan(src interface{}) error {
	switch v := src.(type) {
	case nil:
		return nil
	case int64:
		*i = ID(Int64ToString(v))
		return nil
	}
	return errors.New("not a valid base62")
}

func Int64ToString(value int64) string {
	digit := int(math.Log(float64(value))/math.Log(float64(len(alphabet)))) + 1
	data := make([]uint8, digit)
	r := int64(len(alphabet))
	for i := 0; i < len(data); i++ {
		data[len(data)-i-1] = alphabet[value%r]
		value = value / r
	}
	if value != 0 {
		fmt.Println("something wrong at id")
	}
	return string(data)
}

func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("SID").GoType(ID("")).DefaultFunc(func() ID {
			return ID(utils.RandomString(8))
		}).Unique(),
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
	}
}
