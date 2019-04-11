package main

import (
	//"github.com/alauda/bergamot/loggo"
	"log"
	"time"
)

func main() {

	start := time.Now()
	//strResult := ""
	for start.Add(time.Second * 15).After(time.Now()) {
		//if strResult, err = c.redis.Writer().LPop(c.Opts.GetTaskList()).Result(); err != nil {
		//	c.log().StError("redis lpop err", loggo.Fields{
		//		"err": err,
		//	})
		//}
		//if strResult != "" {
		//	result = []string{strResult}
		//	break
		//}
		log.Println("1111")

		time.Sleep(time.Second)
	}

}
