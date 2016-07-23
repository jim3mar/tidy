package checkin

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//"gopkg.in/mgo.v2"

// Comment data
// the ID is the same as CheckinID
type Comment struct {
	ID       bson.ObjectId   `bson:"_id" json:"id"`
	Comments []SingleComment `bson:"comments" json:"comments"`
	//CheckinID bson.ObjectId   `bson:"checkin_id" json:"checkin_id"`
}

// SingleComment for Comment data
type SingleComment struct {
	UserID      bson.ObjectId `bson:"uid" json:"uid"`
	Comment     string        `bson:"comment" json:"comment"`
	CreateAt    time.Time     `bson:"create_at" json:"create_at"`
	CreateDay   int           `bson:"create_day" json:"create_day"`
	CreateMonth int           `bson:"create_month" json:"create_month"`
	CreateYear  int           `bson:"create_year" json:"create_year"`
	CreateHour  int           `bson:"create_hour" json:"create_hour"`
	CreateMin   int           `bson:"create_min" json:"create_min"`
	CreateSec   int           `bson:"create_sec" json:"create_sec"`
	Timestamp   int64         `bson:"timestamp" json:"timestamp"`
}
