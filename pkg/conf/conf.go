package conf

import "github.com/manatriton/skf/pkg/api"

type Conf struct {
	API   *api.API
	Token string
	URL   string
}
