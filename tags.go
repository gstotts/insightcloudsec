package insightcloudsec

type Tags interface {
	Create(m map[string]string) []Tag
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Tags_Response struct {
	Tags []Tag `json:"resource_tags"`
}

func Create(m map[string]string) []Tag {
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
