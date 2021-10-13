package conf

import "skf/pkg/api"

type Conf struct {
	API   *api.API
	Token string
	URL   string
}
