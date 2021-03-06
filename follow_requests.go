package mastodon

import (
	"net/url"
)

// FollowRequests implements methods under /follow_requests.
type FollowRequests struct {
	api *API
}

// Get returns an slice of accounts which have requested to follow the
// authenticated user.
func (followRequests FollowRequests) Get() ([]Account, error) {
	a := []Account{}
	return a, followRequests.api.Get("follow_requests", nil, &a)
}

// Authorize authorizes a follow request.
func (followRequests FollowRequests) Authorize(id string) error {
	v := url.Values{"id": {id}}
	return followRequests.api.Post("follow_requests/authorize", v, nil)
}

// Reject rejects a follow request.
func (followRequests FollowRequests) Reject(id string) error {
	v := url.Values{"id": {id}}
	return followRequests.api.Post("follow_requests/reject", v, nil)
}

// RejectFalseIcons rejects a follow request.
func (followRequests FollowRequests) RejectFalseIcons(id string) error {
	return followRequests.Reject(id)
}
