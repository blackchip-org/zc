//go:build proj

package ext

import "github.com/twpayne/go-proj/v10"

func ProjTransform(p0, p1 float64, source, target string) (float64, float64, error) {
	sCoord := proj.NewCoord(p0, p1, 0, 0)
	pj, err := proj.NewCRSToCRS(source, target, nil)
	if err != nil {
		return 0, 0, err
	}
	tCoord, err := pj.Forward(sCoord)
	if err != nil {
		return 0, 0, err
	}
	return tCoord.X(), tCoord.Y(), nil
}
