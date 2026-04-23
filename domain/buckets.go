package domain

type BucketKind string

const (
	BucketKindDebt           BucketKind = "debt"
	BucketKindLiquidAssets   BucketKind = "liquid_assets"
	BucketKindIlliquidAssets BucketKind = "illiquid_assets"
)
