package model

import (
	"encoding/json"
	"testing"
)

func TestAccountableResponseEntityUnmarshalOverviewPayload(t *testing.T) {
	payload := []byte(`{"total_invested":440421331772,"total_invested_in_usd":440421.331772,"unrealized_pnl_usd":803.509864}`)

	var entity AccountableResponseEntity
	if err := json.Unmarshal(payload, &entity); err != nil {
		t.Fatalf("expected payload to unmarshal, got error: %v", err)
	}

	if entity.TotalInvested != 440421331772 {
		t.Fatalf("expected TotalInvested to be 440421331772, got %v", entity.TotalInvested)
	}

	if entity.TotalInvestedInUsd != 440421.331772 {
		t.Fatalf("expected TotalInvestedInUsd to be 440421.331772, got %v", entity.TotalInvestedInUsd)
	}

	if entity.UnrealizedPnlUsd != 803.509864 {
		t.Fatalf("expected UnrealizedPnlUsd to be 803.509864, got %v", entity.UnrealizedPnlUsd)
	}
}
