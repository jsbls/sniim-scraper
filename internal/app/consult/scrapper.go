package consult

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

//go:embed endpoints.yaml
var endpointsStr string

func getUrl(consult Consult) (url string, err error) {
	url = ""
	var data []CategoryEndpoint
	err = yaml.Unmarshal([]byte(endpointsStr), &data)

	if err != nil {
		return url, err
	}

	for _, endpoint := range data {
		if endpoint.Category == consult.Category {
			for _, subcat := range endpoint.Subcategories {
				if subcat.Name == consult.SubCategory {
					url = endpoint.UrlPrefix + subcat.Url + consult.ToUrl()
					break
				}
			}
		}
	}

	return url, nil
}

func Scrap(sniimAddr string, consult Consult) error {
	subUrl, err := getUrl(consult)
	if err != nil {
		return err
	}
	url := sniimAddr + subUrl

	fmt.Println(url)
	return nil
}
