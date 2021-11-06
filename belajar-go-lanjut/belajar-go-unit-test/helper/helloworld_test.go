package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Alan",
			request:  "Alan",
			expected: "Hello Alan",
		},
		{
			name:     "Nur",
			request:  "Nur",
			expected: "Hello Nur",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := helloworld(test.request)
			assert.Equal(t, test.expected, result)
		})

	}

}

func TestSubTest(t *testing.T) {
	t.Run("Alan", func(t *testing.T) {
		result := helloworld("Alan")

		require.Equal(t, "Hello Alan", result, "Result must be Hello Alan")

	})
	t.Run("Nur", func(t *testing.T) {
		result := helloworld("Nur")

		require.Equal(t, "Hello Nur", result, "Result must be Hello Alan")

	})
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := helloworld("Alan")
	require.Equal(t, "Hello Alan", result, "Result must be Hello Alan")
}

func TestHelloWorldRequire(t *testing.T) {
	result := helloworld("Alan")
	require.Equal(t, "Hello Alan", result, "Result must be Hello Alan")
	fmt.Println("TestHellowWorldRequire Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := helloworld("Alan")
	assert.Equal(t, "Hello Alan", result, "Result must be Hello Alan")
	fmt.Println("TestHellowWorldAssert Done")

}

func TestHelloWorldAlan(t *testing.T) {
	result := helloworld("Alan")
	if result != "Hello Alan" {
		t.Error("Result must be Hello Alan")
	}
	fmt.Println("TestHelloWorldAlan Done")
}

func TestHelloWorldNur(t *testing.T) {
	result := helloworld("Nur")
	if result != "Hello Nur" {
		t.Fatal("Result must be Hello Nur")
	}
	fmt.Println("TestHelloWorldNur Done")

}
