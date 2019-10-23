package rnames

import (
	"errors"
	"github.com/reconditematter/mym"
	"math"
	"math/rand"
	"time"
)

// HumanName -- the name and gender of a person.
type HumanName struct {
	Family string `json:"family"`
	Given  string `json:"given"`
	Gender string `json:"gender"`
}

const (
	GenBoth = iota // specifies both genders
	GenF           // specifies the female gender
	GenM           // specifies the male gender
)

// Gen -- generates `count` random names.
// `gengen` specifies the names gender.
// This function returns an error when count∉{1,...,1000}.
func Gen(count int, gengen int) ([]HumanName, error) {
	if !(1 <= count && count <= 1000) {
		return nil, errors.New("bad count")
	}
	//
	names := make(map[[2]string]string)
	src := mym.MT19937()
	src.Seed(time.Now().UnixNano())
	rng := rand.New(src)
	//
	for len(names) < count {
		i := rng.Intn(1000)
		j := rng.Intn(1000)
		switch gengen {
		case GenF:
			name := [2]string{family[i], givenf[j]}
			names[name] = "female"
		case GenM:
			name := [2]string{family[i], givenm[j]}
			names[name] = "male"
		default:
			k := rng.Uint64() < math.MaxUint64/2
			if k {
				name := [2]string{family[i], givenf[j]}
				names[name] = "female"
			} else {
				name := [2]string{family[i], givenm[j]}
				names[name] = "male"
			}
		}
	}
	//
	hnames := make([]HumanName, 0, count)
	for name, gender := range names {
		hn := HumanName{Family: name[0], Given: name[1], Gender: gender}
		hnames = append(hnames, hn)
	}
	//
	return hnames, nil
}
