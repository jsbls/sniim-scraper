package consult

import (
	"time"

	"github.com/sirupsen/logrus"
)

const Now = "now"

type Consult struct {
	Category    string            `json:"category"`    // Category for the consult
	SubCategory string            `json:"subcategory"` // Subcategory for the consult
	Params      map[string]string `json:"params"`
}

func NewConsult(cat, subcat string) *Consult {
	return &Consult{
		Category:    cat,
		SubCategory: subcat,
		Params:      make(map[string]string),
	}
}

func (c *Consult) AddParameter(key, val string) {
	_, exists := c.Params[key]

	if !exists {
		c.Params[key] = val
	} else {
		logrus.Warnf("Duplicated key %s", key)
	}
}

func (c *Consult) ToUrl() string {

	query := "?"
	for key, val := range c.Params {
		paramVal := val

		if val == Now {
			nowDate := time.Now()
			paramVal = nowDate.Format("02/01/2006")
		}

		query = query + key + "=" + paramVal + "&"
	}
	query = query[0 : len(query)-1]

	return query
}
