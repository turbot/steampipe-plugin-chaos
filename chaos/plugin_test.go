package chaos

import (
	"context"
	"testing"
)

func TestValidate(t *testing.T) {
	ctx := context.Background()
	validationErrors := Plugin(ctx).Validate()
	if validationErrors != "" {
		t.Errorf("plugin failed validation: \n%s", validationErrors)
	}
}
