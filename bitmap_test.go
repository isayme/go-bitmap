package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBitmap(t *testing.T) {
	t.Run("Set 0-7", func(t *testing.T) {
		require := require.New(t)

		bitmap := New()

		var i uint64
		var data byte
		for i = 0; i < 8; i++ {
			bitmap.Set(i)
			require.Equal(bitmap.length, i+1)
			require.Equal(bitmap.Lenght(), i+1)
			require.Equal(len(bitmap.data), 1)

			data |= (0x01 << i)
			require.Equal(bitmap.data[0], data)
		}
	})

	t.Run("Set 8", func(t *testing.T) {
		require := require.New(t)

		bitmap := New()

		bitmap.Set(8)
		require.Equal(bitmap.length, uint64(1))
		require.Equal(len(bitmap.data), 2)
		require.Equal(bitmap.data[0], byte(0x00))
		require.Equal(bitmap.data[1], byte(0x01))
	})

	t.Run("Has", func(t *testing.T) {
		require := require.New(t)

		bitmap := New()
		bitmap.Set(0)

		require.Equal(bitmap.Has(0), true)
		require.Equal(bitmap.Has(1), false)
		require.Equal(bitmap.Has(100), false)
	})

	t.Run("duplicate set", func(t *testing.T) {
		require := require.New(t)

		bitmap := New()
		bitmap.Set(0)
		bitmap.Set(0)
		require.Equal(bitmap.length, uint64(1))
		require.Equal(len(bitmap.data), 1)
		require.Equal(bitmap.data[0], byte(0x01))
	})
}

func BenchmarkGet(b *testing.B) {
	bitmap := New()
	for i := 0; i < b.N; i++ {
		bitmap.Set(uint64(i))
	}
}
