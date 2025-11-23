package helper

import (
	"foodieland/model/domain"
	"foodieland/model/web"
)

func PatchToDirection(req []web.DirectionPatchRequest) []domain.Direction {
    res := make([]domain.Direction, len(req))
    for i, d := range req {
        res[i] = domain.Direction{
            Step:        derefUint(d.Step),
            Description: derefString(d.Description),
            Img:         derefString(d.Img),
        }
    }
    return res
}

func derefUint(v *uint) uint {
    if v == nil { return 0 }
    return *v
}

func derefString(v *string) string {
    if v == nil { return "" }
    return *v
}


func ToDomainDirections(req []web.DirectionRequest) []domain.Direction {
    res := make([]domain.Direction, len(req))
    for i, d := range req {
        res[i] = domain.Direction{
            Step:        d.Step,
            Description: d.Description,
            Img:         d.Img,
        }
    }
    return res
}
