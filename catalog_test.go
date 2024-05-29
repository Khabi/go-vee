package govee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidLookupBySKU(t *testing.T) {
	for i := 1; i <= 5; i++ {
		entry := catalog[i]
		sku := entry.SKU

		lentry := LookupBySKU(sku)
		assert.Equal(t, sku, lentry.SKU)
		assert.Equal(t, entry.Kind, lentry.Kind)
		assert.Equal(t, entry.Name, lentry.Name)
		assert.Equal(t, entry.ProductURL, lentry.ProductURL)
	}
}

func TestInvalidLookupBySKU(t *testing.T) {
	lentry := LookupBySKU("FAKE")
	assert.Nil(t, lentry)
}

func TestLookupByKind(t *testing.T) {
	lentry := LookupByKind(WallLights)
	assert.NotNil(t, lentry)
	count := 0
	for _, e := range catalog {
		if e.Kind == WallLights {
			count++
		}
	}
	assert.Len(t, lentry, count)
}

func TestInvalidLookupByKind(t *testing.T) {
	var invalid Kind = "INVALID"

	lentry := LookupByKind(invalid)
	assert.Nil(t, lentry)
}
