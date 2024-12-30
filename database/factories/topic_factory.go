package factories

import (
	"github.com/bxcodec/faker/v3"
	"gohub/app/models/topic"
	"gohub/pkg/helpers"
)

func MakeTopics(count int) []topic.Topic {

	var objs []topic.Topic

	for i := 0; i < count; i++ {
		topicModel := topic.Topic{
			Title:      faker.Sentence(),
			Body:       faker.Paragraph(),
			CategoryID: string(helpers.RandomNumber(1)),
			UserID:     string(helpers.RandomNumber(1)),
		}
		objs = append(objs, topicModel)
	}

	return objs
}
