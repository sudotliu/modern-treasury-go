package moderntreasury_test

import (
	"context"
	"errors"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestLedgerTransactionNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerTransactions.New(context.TODO(), moderntreasury.LedgerTransactionNewParams{Description: moderntreasury.F("string"), Status: moderntreasury.F(moderntreasury.LedgerTransactionNewParamsStatusArchived), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), EffectiveDate: moderntreasury.F(time.Now()), LedgerEntries: moderntreasury.F([]moderntreasury.LedgerTransactionNewParamsLedgerEntries{{Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(moderntreasury.LedgerTransactionNewParamsLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}, {Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(moderntreasury.LedgerTransactionNewParamsLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}, {Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(moderntreasury.LedgerTransactionNewParamsLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}}), ExternalID: moderntreasury.F("string"), LedgerableType: moderntreasury.F(moderntreasury.LedgerTransactionNewParamsLedgerableTypeCounterparty), LedgerableID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerTransactionGet(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerTransactions.Get(
		context.TODO(),
		"string",
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerTransactionUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerTransactions.Update(
		context.TODO(),
		"string",
		moderntreasury.LedgerTransactionUpdateParams{Description: moderntreasury.F("string"), Status: moderntreasury.F(moderntreasury.LedgerTransactionUpdateParamsStatusArchived), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), LedgerEntries: moderntreasury.F([]moderntreasury.LedgerTransactionUpdateParamsLedgerEntries{{Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(moderntreasury.LedgerTransactionUpdateParamsLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}, {Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(moderntreasury.LedgerTransactionUpdateParamsLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}, {Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(moderntreasury.LedgerTransactionUpdateParamsLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}})},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerTransactionListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerTransactions.List(context.TODO(), moderntreasury.LedgerTransactionListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), ID: moderntreasury.F(map[string]string{"foo": "string"}), Metadata: moderntreasury.F(map[string]string{"foo": "string"}), LedgerID: moderntreasury.F("string"), LedgerAccountID: moderntreasury.F("string"), EffectiveAt: moderntreasury.F(map[string]string{"foo": "string"}), EffectiveDate: moderntreasury.F(map[string]time.Time{"foo": time.Now()}), PostedAt: moderntreasury.F(map[string]time.Time{"foo": time.Now()}), UpdatedAt: moderntreasury.F(map[string]time.Time{"foo": time.Now()}), OrderBy: moderntreasury.F(moderntreasury.LedgerTransactionListParamsOrderBy{CreatedAt: moderntreasury.F(moderntreasury.LedgerTransactionListParamsOrderByCreatedAtAsc), EffectiveAt: moderntreasury.F(moderntreasury.LedgerTransactionListParamsOrderByEffectiveAtAsc)}), Status: moderntreasury.F(moderntreasury.LedgerTransactionListParamsStatusPending), ExternalID: moderntreasury.F("string"), LedgerAccountCategoryID: moderntreasury.F("string")})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}