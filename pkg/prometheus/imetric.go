package prometheus

import "context"

//go:generate mockgen -source=imetric.go -destination=metric_mock.go -package=prometheus

//IMerchantMetric ...
type IMerchantMetric interface {
	IncreaseRequestCnt(ctx context.Context, name string, action string, code string)
	IncreasePartnerRequestCnt(ctx context.Context, name string, partnerName string, action string, code string)
	MesureRequestDuration(ctx context.Context, name string, duration float64, lvb ...string)
	RecordsCurrentValue(ctx context.Context, name string, val float64, lvb ...string)
	IncreaseWithNameCnt(ctx context.Context, name string, lvb ...string)
}
