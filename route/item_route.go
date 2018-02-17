package route

import "bitbucket.org/hayum/hayum-service/util"

func (r *Router) ItemRoute(basePath string) {
	r.GET(util.ConstructEndpoint(basePath, "/:id"), getItem)
}
