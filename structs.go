package vkapi

import (
	"encoding/json"
)

const (
	_ = iota
	PlatformMobile
	PlatformIPhone
	PlatfromIPad
	PlatformAndroid
	PlatformWPhone
	PlatformWindows
	PlatformWeb
)

type APIResponse struct {
	Response json.RawMessage `json:"response"`
}

type Token struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	UID              int    `json:"user_id"`
	Error            string `json:"error"`
	ErorrDescription string `json:"error"`
}

type User struct {
	UID            int       `json:"uid"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Sex            int       `json:"sex"`
	Nickname       string    `json:"nickname"`
	ScreenName     string    `json:"screen_name"`
	BDate          string    `json:"bdate"`
	City           int       `json:"city"`
	Country        int       `json:"country"`
	Photo          string    `json:"photo"`
	PhotoMedium    string    `json:"photo_medium"`
	PhotoBig       string    `json:"photo_big"`
	HasMobile      int       `json:"has_mobile"`
	Online         int       `json:"online"`
	CanPost        int       `json:"can_post"`
	CanSeeAllPosts int       `json:"can_see_all_posts"`
	Status         string    `json:"activity"`
	LastOnline     *LastSeen `json:"last_seen"`
	Hidden         int       `json:"hidden"`
}

type LastSeen struct {
	Time     int64 `json:"time"`
	Platform int   `json:"platform"`
}

type Message struct {
	MID               int                `json:"mid"`
	Date              int64              `json:"date"`
	Out               int                `json:"out"`
	UID               int                `json:"uid"`
	ReadState         int                `json:"read_state"`
	Title             string             `json:"title"`
	Body              string             `json:"body"`
	ChatID            int                `json:"chat_id"`
	ChatActive        string             `json:"chat_active"`
	PushSettings      *Push              `json:"push_settings"`
	UsersCount        int                `json:"users_count"`
	AdminID           int                `json:"admin_id"`
	Photo50           string             `json:"photo_50"`
	Photo100          string             `json:"photo_100"`
	Photo200          string             `json:"photo_200"`
	ForwardedMessages []ForwardedMessage `json:"fwd_messages"`
	Attachments       []Attachment       `json:"attachments"`
}

type Push struct {
	Sound         int   `json:"sound"`
	DisabledUntil int64 `json:"disabled_until"`
}

type ForwardedMessage struct {
	UID  int    `json:"uid"`
	Date int64  `json:"date"`
	Body string `json:"body"`
}

type Attachment struct {
	Type     string           `json:"type"`
	Audio    *AudioAttachment `json:"audio"`
	Video    *VideoAttachment `json:"video"`
	Photo    *PhotoAttachment `json:"photo"`
	Document *DocAttachment   `json:"doc"`
	Link     *LinkAttachment  `json:"link"`
}

type AudioAttachment struct {
	AudioID   int    `json:"aid"`
	OwnerID   int    `json:"owner_id"`
	Artist    string `json:"artist"`
	Title     string `json:"title"`
	Duration  int    `json:"duration"`
	URL       string `json:"url"`
	Performer string `json:"performer"`
}

type VideoAttachment struct {
	VideoID     int    `json:"vid"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Duration    int    `json:"duration"`
	Description string `json:"description"`
	Date        int64  `json:"date"`
	Views       int    `json:"views"`
	Image       string `json:"image"`
	ImageBig    string `json:"image_big"`
	ImageSmall  string `json:"image_small"`
	ImageXBig   string `json:"image_xbig"`
	AccessKey   string `json:"access_key"`
	Platform    string `json:"platform"`
	CanEdit     int    `json:"can_edit"`
}

type PhotoAttachment struct {
	PhotoID     int    `json:"pid"`
	AID         int    `json:"aid"`
	OwnerID     int    `json:"owner_id"`
	Source      string `json:"src"`
	SourceBig   string `json:"src_big"`
	SourceSmall string `json:"src_small"`
	SourceXBig  string `json:"src_xbig"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Text        string `json:"text"`
	Created     int64  `json:"created"`
	AccessKey   string `json:"access_key"`
}

type DocAttachment struct {
	DocID      int    `json:"did"`
	OwnerID    int    `json:"owner_id"`
	Title      string `json:"title"`
	Size       int    `json:"size"`
	Extenstion string `json:"ext"`
	URL        string `json:"url"`
	Date       int64  `json:"date"`
	AccessKey  string `json:"access_key"`
}

type LinkAttachment struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Target      string `json:"target"`
}

type WallPost struct {
	ID           int          `json:"id"`
	FromID       int          `json:"from_id"`
	ToID         int          `json:"to_id"`
	Date         int64        `json:"date"`
	MarkedAsAd   int          `json:"marked_as_ads"`
	PostType     string       `json:"post_type"`
	CopyPostDate int64        `json:"copy_post_date"`
	CopyPostType string       `json:"copy_post_type"`
	CopyOwnerID  int          `json:"copy_owner_id"`
	CopyPostID   int          `json:"copy_post_id"`
	CreatedBy    int          `json:"created_by"`
	Text         string       `json:"text"`
	CanDelete    int          `json:"can_delete"`
	CanPin       int          `json:"can_pin"`
	Attachments  []Attachment `json:"attachments"`
	Comments     *Comment     `json:"comments"`
	Likes        *Like        `json:"likes"`
	Reposts      *Repost      `json:"reposts"`
	Online       int          `json:"online"`
	ReplyCount   int          `json:"reply_count"`
}

type Comment struct {
	Count   int `json:"count"`
	CanPost int `json:"can_post"`
}

type Like struct {
	Count      int `json:"count"`
	UserLikes  int `json:"user_likes"`
	CanLike    int `json:"can_like"`
	CanPublish int `json:"can_publish"`
}

type Repost struct {
	Count        int `json:"count"`
	UserReposted int `json:"user_reposted"`
}

type LongPollServer struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	TS     int64  `json:"ts"`
}

type LongPollUpdate struct {
	Failed  int             `json:"failed"`
	TS      int64           `json:"ts"`
	Updates [][]interface{} `json:"updates"`
}

type LongPollMessage struct {
	MessageID   int
	UserID      int
	Date        int64
	Title       string
	Body        string
	Attachments map[string]string
}

type LongPollChannel <-chan LongPollMessage
