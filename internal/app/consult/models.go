package consult

import (
	"fmt"
	"time"

	"github.com/everitosan/sniim-scrapper/internal/app/form"
	"github.com/sirupsen/logrus"
)

const Now = "now"

/*
* The following structs are defined for the endpoints yaml
 */

type SubCategoryEndopoint struct {
	Name     string        `json:"name"`
	Url      string        `json:"url"`
	FormType form.FormType `json:"formType"`
}
type CategoryEndpoint struct {
	Category      string                 `json:"category"`
	UrlPrefix     string                 `json:"urlprefix"`
	Subcategories []SubCategoryEndopoint `json:"subcategories"`
}

/*
* Consult is a request that should be made
 */

type Consult struct {
	Category    string            `json:"category"`    // Category for the consult
	SubCategory string            `json:"subcategory"` // Subcategory for the consult
	Params      map[string]string `json:"params"`      // Parameters for the requests
	TextParams  []RegisterConcept `json:"textParams"`  // Parameters in human language for the requests, the slice is to maintain the order
}

func NewConsult(cat, subcat string) *Consult {
	return &Consult{
		Category:    cat,
		SubCategory: subcat,
		Params:      make(map[string]string),
		TextParams:  make([]RegisterConcept, 0),
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

func (c *Consult) AddTextParameter(key, val string) {
	newRegister := RegisterConcept{Name: key, Value: val}
	c.TextParams = append(c.TextParams, newRegister)
}

// this funcionn is used when we want to generate a result Row (after a consult)
func (c *Consult) GetParamsAsConcepts() []RegisterConcept {
	concepts := make([]RegisterConcept, 0, len(c.TextParams))

	for _, registerConcept := range c.TextParams {
		realValue := registerConcept.Value

		if registerConcept.Value == Now {
			realValue = time.Now().Format("02/01/2006")
		}
		concepts = append(concepts, RegisterConcept{Name: registerConcept.Name, Value: realValue})
	}
	return concepts
}

func (c *Consult) String() string {
	str := fmt.Sprintf("%s/%s \n\t ▶️ [", c.Category, c.SubCategory)
	for _, registerConcept := range c.TextParams {
		str = str + registerConcept.Name + "=" + registerConcept.Value + ", "
	}

	str = str[0 : len(str)-2]
	str = str + "]"
	return str
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

/*
* Register concept represents the relation of a concept and it's value in a result table
 */

type RegisterConcept struct {
	Name  string `json:"concept"` // Concept of the result
	Value string `json:"value"`   // Value related to the concept
}
