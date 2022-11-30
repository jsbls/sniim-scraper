package consult

import "github.com/sirupsen/logrus"

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
		query = query + key + "=" + val + "&"
	}
	query = query[0 : len(query)-1]

	return query
}
