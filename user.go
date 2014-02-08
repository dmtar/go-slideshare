package slideshare

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UserFavorites struct {
	Values []UserFavorite `xml:"favorite"`
}
type UserFavorite struct {
	SlideshowID uint64 `xml:"slideshow_id"`
	TagText     string `xml:"tag_text"`
}
type UserContacts struct {
	Values []UserContact `xml:"Contact"`
}
type UserContact struct {
	Username      string `xml:"Username"`
	NumSlideshows uint32 `xml:"NumSlideshows"`
	NumComments   uint32 `xml:"NumComments"`
}

type Groups struct {
	Groups []Group `xml:"group"`
}
type Group struct {
	Name          string `xml:"name"`
	NumPosts      uint32 `xml:"numposts"`
	NumSlideshows uint32 `xml:"numslideshows"`
	NumMembers    uint32 `xml:"nummembers"`
	Created       string `xml:"created"`
	QueryName     string `xml:"queryname"`
	Url           string `xml:"url"`
}

// Returns user favorites
// username_for required, username of user whose favorites are being requested.
func (s *Service) GetUserFavorites(username_for string) (UserFavorites, error) {
	args := make(map[string]string)
	args["username_for"] = username_for
	url := s.generateUrl("get_user_favorites", args)
	resp, err := http.Get(url)
	if err != nil {
		return UserFavorites{}, err
	}
	favorites := UserFavorites{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &favorites)
	}
	return favorites, err
}

// Returns user contacts
// username_for required, username of user whose contacts are being requested
func (s *Service) GetUserContacts(username_for string, limitOffset ...int) (UserContacts, error) {
	args := make(map[string]string)
	if limitOffset != nil {
		switch len(limitOffset) {
		case 1:
			args["limit"] = strconv.Itoa(limitOffset[0])
			break
		case 2:
			args["limit"] = strconv.Itoa(limitOffset[0])
			args["offset"] = strconv.Itoa(limitOffset[1])
			break
		default:
		}
	}
	args["username_for"] = username_for
	url := s.generateUrl("get_user_contacts", args)
	resp, err := http.Get(url)
	if err != nil {
		return UserContacts{}, err
	}
	contacts := UserContacts{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &contacts)
	}
	return contacts, err
}

/*
// Returns user groups
// username_for required, username of user whose groups are being requested
func (s *Service) GetUserGroups(username_for string) (Groups, error)

// Returns user tags
// username required, username of user whose tags are being requested
// password required, password of user whose tags are being requested
func (s *Service) GetUserTags(username string, password string) (Tags, error)
*/
