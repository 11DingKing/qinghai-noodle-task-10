package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask10(t *testing.T) {
	now := time.Now()
	s := NewService(NewRegistry(), func() time.Time { return now })
	lot := IngredientLot{SupplierID: "sup", CertificateID: "cert", IngredientCode: "yak"}
	cert := SupplierCertificate{ID: "cert", SupplierID: "sup", Scope: []string{"yak"}, EffectiveAt: now.Add(-time.Hour), ExpiresAt: now.Add(time.Hour)}
	require.NoError(t, s.CheckSupplierCertificate(context.Background(), lot, cert))
}
