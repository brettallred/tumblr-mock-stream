package main

type ActivityEnvelope struct {
	Id              string   `json:"id"`
	Timestamp       int64    `json:"timestamp"`
	Version         string   `json:"version"`
	ActivityPrivacy string   `json:"activity_privacy"`
	ActivityType    string   `json:"activity_type"`
	Activity        Activity `json:"activity"`
}

type Activity struct {
	Id     string `json:"id"`
	Action string `json:"action"`
	Blog   Blog   `json:"blog"`
}
