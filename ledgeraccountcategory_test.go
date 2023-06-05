package moderntreasury_test

import (
	"context"
	"errors"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestLedgerAccountCategoryNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccountCategories.New(context.TODO(), moderntreasury.LedgerAccountCategoryNewParams{
		Currency:         moderntreasury.F("string"),
		LedgerID:         moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Name:             moderntreasury.F("string"),
		NormalBalance:    moderntreasury.F(moderntreasury.LedgerAccountCategoryNewParamsNormalBalanceCredit),
		CurrencyExponent: moderntreasury.F(int64(0)),
		Description:      moderntreasury.F("string"),
		Metadata:         moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}),
		IdempotencyKey:   moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountCategoryGetWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccountCategories.Get(
		context.TODO(),
		"string",
		moderntreasury.LedgerAccountCategoryGetParams{
			Balances: moderntreasury.F(moderntreasury.LedgerAccountCategoryGetParamsBalances{AsOfDate: moderntreasury.F(time.Now()), EffectiveAt: moderntreasury.F(time.Now())}),
		},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountCategoryUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccountCategories.Update(
		context.TODO(),
		"string",
		moderntreasury.LedgerAccountCategoryUpdateParams{
			Balances:    moderntreasury.F(moderntreasury.LedgerAccountCategoryUpdateParamsBalances{AsOfDate: moderntreasury.F(time.Now()), EffectiveAt: moderntreasury.F(time.Now())}),
			Description: moderntreasury.F("string"),
			Metadata:    moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}),
			Name:        moderntreasury.F("string"),
		},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountCategoryListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccountCategories.List(context.TODO(), moderntreasury.LedgerAccountCategoryListParams{
		AfterCursor:                   moderntreasury.F("string"),
		LedgerID:                      moderntreasury.F("string"),
		Metadata:                      moderntreasury.F(map[string]string{"foo": "string"}),
		Name:                          moderntreasury.F("string"),
		ParentLedgerAccountCategoryID: moderntreasury.F("string"),
		PerPage:                       moderntreasury.F(int64(0)),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountCategoryDeleteWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccountCategories.Delete(
		context.TODO(),
		"string",
		moderntreasury.LedgerAccountCategoryDeleteParams{
			Balances: moderntreasury.F(moderntreasury.LedgerAccountCategoryDeleteParamsBalances{AsOfDate: moderntreasury.F(time.Now()), EffectiveAt: moderntreasury.F(time.Now())}),
		},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountCategoryAddLedgerAccount(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	err := c.LedgerAccountCategories.AddLedgerAccount(
		context.TODO(),
		"string",
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

func TestLedgerAccountCategoryAddNestedCategory(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	err := c.LedgerAccountCategories.AddNestedCategory(
		context.TODO(),
		"string",
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

func TestLedgerAccountCategoryRemoveLedgerAccount(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	err := c.LedgerAccountCategories.RemoveLedgerAccount(
		context.TODO(),
		"string",
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

func TestLedgerAccountCategoryRemoveNestedCategory(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	err := c.LedgerAccountCategories.RemoveNestedCategory(
		context.TODO(),
		"string",
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
