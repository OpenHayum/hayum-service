package route

import "bitbucket.org/hayum/hayum-service/util"

func (r *Router) UserRoute(basePath string) {
	r.GET(util.ConstructEndpoint(basePath, "/:name"), getUser)
}
