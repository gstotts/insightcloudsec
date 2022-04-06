package insightcloudsec

var _ Tags = (*tags)(nil)

type Tags interface {
	Create(m map[string]string) []Tag
}

type tags struct {
	client *Client
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Tags_Response struct {
	Tags []Tag `json:"resource_tags"`
}

func (t *tags) Create(m map[string]string) []Tag {
	tags := []Tag{}
	for tag, value := range m {
		item := Tag{
			Key:   tag,
			Value: value,
		}
		tags = append(tags, item)
	}

	return tags
}
