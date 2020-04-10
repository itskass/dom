//+build !js

package js

import (
	"testing"

	"github.com/itskass/dom/js/jstest"
)

func TestJS(t *testing.T) {
	jstest.RunTestNodeJS(t)
}
