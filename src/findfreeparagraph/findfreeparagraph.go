package findfreeparagraph

import (
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
)


func FindFromQ(golog syslog.Writer, locale string, themes string, bot string,startparameters []string) domains.Paragraph {

	redisprotocol := startparameters[0]
	redishost := startparameters[1]

	c, err := redis.Dial(redisprotocol, redishost)
	if err != nil {

		golog.Crit(err.Error())

	}

	queuename := locale + ":" + themes

	var unmarPar domains.Paragraph

	if quan_prs, err := redis.Int(c.Do("LLEN", queuename)); err != nil {

		golog.Crit(err.Error())

	} else {

		if quan_prs > 1 {

			bparagraph, _ := redis.Bytes(c.Do("LPOP", queuename))

			err := json.Unmarshal(bparagraph, &unmarPar)
			if err != nil {

				golog.Crit(err.Error())

			}
			


		} else {

			golog.Crit("need more free paragraphs!!!!")

		}

	}
	return unmarPar
}
