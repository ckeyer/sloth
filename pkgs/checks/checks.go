package checks

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	// "github.com/gojp/goreportcard/check"
)

// Check describes what methods various checks (gofmt, go lint, etc.)
// should implement
type Check interface {
	Name() string
	Description() string
	Weight() float64
	// Percentage returns the passing percentage of the check,
	// as well as a map of filename to output
	Percentage() (float64, []FileSummary, error)
}

type Grade string

type score struct {
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	FileSummaries []FileSummary `json:"file_summaries"`
	Weight        float64       `json:"weight"`
	Percentage    float64       `json:"percentage"`
	Error         string        `json:"error"`
}

type checksResp struct {
	Checks               []score   `json:"checks"`
	Average              float64   `json:"average"`
	Grade                Grade     `json:"grade"`
	Files                int       `json:"files"`
	Issues               int       `json:"issues"`
	Repo                 string    `json:"repo"`
	LastRefresh          time.Time `json:"last_refresh"`
	HumanizedLastRefresh string    `json:"humanized_last_refresh"`
}

func RunChecks(dir string, filenames []string) (checksResp, error) {
	if len(filenames) == 0 {
		return checksResp{}, fmt.Errorf("no .go files found")
	}

	checks := []Check{
		GoFmt{Dir: dir, Filenames: filenames},
		GoVet{Dir: dir, Filenames: filenames},
		GoLint{Dir: dir, Filenames: filenames},
		GoCyclo{Dir: dir, Filenames: filenames},
		License{Dir: dir, Filenames: []string{}},
		Misspell{Dir: dir, Filenames: filenames},
		IneffAssign{Dir: dir, Filenames: filenames},
		ErrCheck{Dir: dir, Filenames: filenames}, // disable errcheck for now, too slow and not finalized
	}

	ch := make(chan score)
	for _, c := range checks {
		go func(c Check) {
			p, summaries, err := c.Percentage()
			errMsg := ""
			if err != nil {
				log.Errorf("RunChecks ERROR: (%s) %v", c.Name(), err)
				errMsg = err.Error()
			}
			s := score{
				Name:          c.Name(),
				Description:   c.Description(),
				FileSummaries: summaries,
				Weight:        c.Weight(),
				Percentage:    p,
				Error:         errMsg,
			}
			ch <- s
		}(c)
	}

	resp := checksResp{
		Files:                len(filenames),
		LastRefresh:          time.Now().UTC(),
		HumanizedLastRefresh: time.Now().String(),
	}

	var total float64
	var totalWeight float64
	var issues = make(map[string]bool)
	for i := 0; i < len(checks); i++ {
		s := <-ch
		log.WithFields(log.Fields{
			"name":       s.Name,
			"percentage": s.Percentage,
		}).Debugf("RunChecks.")

		resp.Checks = append(resp.Checks, s)
		total += s.Percentage * s.Weight
		totalWeight += s.Weight
		for _, fs := range s.FileSummaries {
			issues[fs.Filename] = true

			for _, ferr := range fs.Errors {
				log.WithFields(log.Fields{
					"filename": fs.Filename,
					"file URL": fs.FileURL,
					"line":     ferr.LineNumber,
				}).Debug("RunChecks FileSummaries. ", ferr.ErrorString)
			}
		}
	}
	total /= totalWeight

	// sort.Sort(ByWeight(resp.Checks))
	resp.Average = total
	resp.Issues = len(issues)

	log.WithFields(log.Fields{
		"issues": len(issues),
		"total":  total,
	}).Debug("RunChecks.")

	return resp, nil
}
