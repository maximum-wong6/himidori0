package vkapi

import (
	"encoding/json"
	"net/url"
	"strconv"
)

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

func (client *VKClient) GetWallPosts(id interface{}, count int, params url.Values) ([]WallPost, error) {
	if params == nil {
		params = url.Values{}
	}

	params.Add("count", strconv.Itoa(count))

	switch id.(type) {
	case int:
		params.Add("owner_id", strconv.Itoa(id.(int)))
	case string:
		params.Add("domain", id.(string))
	}

	resp, err := client.makeRequest("wall.get", params)
	if err != nil {
		return []WallPost{}, err
	}

	var posts []WallPost
	clearedString := deleteFirstKey(string(resp.Response))
	json.Unmarshal([]byte(clearedString), &posts)

	return posts, nil
}