package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	api := rest.NewApi()
	api.Use(&rest.AccessLogApacheMiddleware{})
	api.Use(rest.DefaultCommonStack...)
	router, err := rest.MakeRouter(
		rest.Get("/stream", StreamActivities),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9999", api.MakeHandler()))
}

func StreamActivities(w rest.ResponseWriter, r *rest.Request) {
	for {
		now := time.Now()
		seconds := now.Unix()
		activityTypes := []string{"blog", "post", "like"}

		w.WriteJson(
			&ActivityEnvelope{
				Id:              strconv.FormatInt(rand.Int63(), 10),
				Timestamp:       seconds,
				Version:         "2.2",
				ActivityPrivacy: "public",
				ActivityType:    activityTypes[0],
				Activity: Activity{
					Id:     strconv.FormatInt(rand.Int63(), 10),
					Action: "create",
					Blog: Blog{
						Created:     4787198327,
						Updated:     4787198327,
						Name:        "Blog name",
						Url:         "http://blog-name.tumblr.com",
						Title:       "Title of blog",
						Description: "Blog description",
						IsGroupBlog: false,
						IsPrimary:   false,
						Timezone:    "US/EST",
						PostCount:   744,
						LastPost:    4787198327,
						Language:    "English",
					},
				},
			},
		)

		w.(http.ResponseWriter).Write([]byte("\n"))
		w.(http.Flusher).Flush()
		time.Sleep(time.Duration(3) * time.Second)
	}
}
