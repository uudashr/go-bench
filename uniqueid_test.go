package bench_test

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"

	hashicorpuuid "github.com/hashicorp/go-uuid"
	"github.com/rs/xid"

	satoriuuid "github.com/satori/go.uuid"

	googleuuid "github.com/google/uuid"
)

func BenchmarkUniqueID_UUID_Google(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := googleuuid.New().String()
		_ = s
	}
}

func BenchmarkUniqueID_UUID_Hashicorp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, err := hashicorpuuid.GenerateUUID()
		if err != nil {
			b.Fatal("Fail to generate UUID:", err)
		}
		_ = s
	}
}

func BenchmarkUniqueID_UUID_Satori(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := satoriuuid.NewV4().String()
		_ = s
	}
}

func BenchmarkUniqueID_UUID_Custom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, err := newUUID()
		if err != nil {
			b.Fatal("Fail to generate UUID:", err)
		}
		_ = s
	}
}

func BenchmarkUniqueID_XID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := xid.New().String()
		_ = s
	}
}

// newUUID generates a random UUID according to RFC 4122
// see: https://play.golang.org/p/4FkNSiUDMg
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
