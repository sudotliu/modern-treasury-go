package moderntreasury_test

import (
	"context"
	"errors"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestLedgerAccountNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccounts.New(context.TODO(), moderntreasury.LedgerAccountNewParams{Name: moderntreasury.F("string"), Description: moderntreasury.F("string"), NormalBalance: moderntreasury.F(moderntreasury.LedgerAccountNewParamsNormalBalanceCredit), LedgerID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Currency: moderntreasury.F("string"), CurrencyExponent: moderntreasury.F(int64(0)), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"})})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountGetWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccounts.Get(
		context.TODO(),
		"string",
		moderntreasury.LedgerAccountGetParams{Balances: moderntreasury.F(moderntreasury.LedgerAccountGetParamsBalances{AsOfDate: moderntreasury.F(time.Now()), EffectiveAt: moderntreasury.F(time.Now()), EffectiveAtLowerBound: moderntreasury.F(time.Now()), EffectiveAtUpperBound: moderntreasury.F(time.Now())})},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccounts.Update(
		context.TODO(),
		"string",
		moderntreasury.LedgerAccountUpdateParams{Name: moderntreasury.F("string"), Description: moderntreasury.F("string"), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"})},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccounts.List(context.TODO(), moderntreasury.LedgerAccountListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), Metadata: moderntreasury.F(map[string]string{"foo": "string"}), ID: moderntreasury.F("string"), Name: moderntreasury.F("string"), LedgerID: moderntreasury.F("string"), Balances: moderntreasury.F(moderntreasury.LedgerAccountListParamsBalances{AsOfDate: moderntreasury.F(time.Now()), EffectiveAt: moderntreasury.F(time.Now()), EffectiveAtLowerBound: moderntreasury.F(time.Now()), EffectiveAtUpperBound: moderntreasury.F(time.Now())}), UpdatedAt: moderntreasury.F(map[string]time.Time{"foo": time.Now()}), LedgerAccountCategoryID: moderntreasury.F("string")})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountDelete(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccounts.Delete(
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