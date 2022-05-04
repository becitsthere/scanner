package registry

import (
	log "github.com/sirupsen/logrus"
)

type tagsResponse struct {
	Tags []string `json:"tags"`
}

func (r *Registry) Tags(repository string) ([]string, error) {
	url := r.url("/v2/%s/tags/list", repository)
	tags := make([]string, 0)

	var response tagsResponse
	var err error

	r.Client.SetTimeout(longTimeout)
	for {
		log.WithFields(log.Fields{"url": url, "repository": repository}).Debug()
		url, err = r.getPaginatedJson(url, &response)
		switch err {
		case ErrNoMorePages:
			tags = append(tags, response.Tags...)
			return tags, nil
		case nil:
			tags = append(tags, response.Tags...)
			continue
		default:
			return tags, nil
		}
	}
}
