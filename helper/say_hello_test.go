package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Start Test")

	m.Run()

	fmt.Println("End Test")
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Kresna")
	}
}

func BenchmarkSubSayHello(b *testing.B) {
	b.Run("Kresna", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Kresna")
		}
	})

	b.Run("Cahyono", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Cahyono")
		}
	})
}

func BenchmarkTableSayHello(b *testing.B) {
	bencmark := []struct {
		Name string
	}{
		{
			Name: "Kresna",
		},
		{
			Name: "Cahyono",
		},
		{
			Name: "Joko",
		},
		{
			Name: "Hadi",
		},
		{
			Name: "Kusumo",
		},
	}

	for _, v := range bencmark {
		b.Run(v.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(v.Name)
			}
		})
	}
}

func TestSayHelloTable(t *testing.T) {
	tests := []struct {
		Name, Request, Expected string
	}{
		{
			Name:     "Kresna",
			Request:  "Kresna",
			Expected: "Hello Kresna",
		},
		{
			Name:     "Cahyono",
			Request:  "Cahyono",
			Expected: "Hello Cahyono",
		},
		{
			Name:     "Joko",
			Request:  "Joko",
			Expected: "Hello Joko",
		},
		{
			Name:     "Budi",
			Request:  "Budi",
			Expected: "Hello Budi",
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			result := SayHello(v.Request)
			assert.Equal(t, v.Expected, result, "Result Must Be "+v.Expected)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Kresna", func(t *testing.T) {
		result := SayHello("Kresna")
		assert.Equal(t, "Hello Kresna", result, "Result Must Be 'Hello Kresna'")
	})
	t.Run("Cahyono", func(t *testing.T) {
		result := SayHello("Cahyono")
		assert.Equal(t, "Hello Cahyono", result, "Result Must Be 'Hello Cahyono'")
	})
}

func TestSayHello(t *testing.T) {
	string := SayHello("Kresna")
	if string != "Hello Kresna" {
		t.Error("Harusnya 'Hello Kresna'") // tidak langsung dihentikan
		t.Fatal("Harusnya 'Hello Kresna'") // langsung dihentikan
	}
}

func TestWithAssert(t *testing.T) {
	result := SayHello("Kresna")
	assert.Equal(t, "Hello Kresna", result, "Result Must Be 'Hello Kresna'") // tidak langsung exit
}

func TestWithRequire(t *testing.T) {
	result := SayHello("Kresna")
	require.Equal(t, "Hello Kresna", result, "Result Must Be 'Hello Kresna'") // langsung exit
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Application cannot running on windows")
	}

	result := SayHello("Kresna")
	require.Equal(t, "Hello Kresna", result, "Result Must Be 'Hello Kresna'") // langsung exit
}
