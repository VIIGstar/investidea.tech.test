package system_handlers

import build_info "investidea.tech.test/pkg/build-info"

type systemHandler struct {
	buildInfo build_info.BuildInfo
}

func New(buildInfo build_info.BuildInfo) systemHandler {
	return systemHandler{
		buildInfo: buildInfo,
	}
}
