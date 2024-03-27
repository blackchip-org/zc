//go:build !proj

package ext

import "github.com/blackchip-org/zc/v6/pkg/zc"

func ProjTransform(p0, p1 float64, source, target string) (float64, float64, error) {
	return 0, 0, zc.ErrFeatureNotSupported(FeatureProj)
}
