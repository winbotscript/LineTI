package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
)

const (
	HOST        = "https://legy-jp-addr.line.naver.jp"
	// Post
	TIMELINE1   = HOST + "/mh/api/v50/post/update.json" // homeId + sourceType
	TIMELINE2   = HOST + "/mh/api/v50/post/create.json"
	TIMELINE3   = HOST + "/mh/api/v50/post/get.json"
	TIMELINE4   = HOST + "/mh/api/v50/post/getShareLink.json"
	TIMELINE5   = HOST + "/mh/api/v50/post/delete.json"
	TIMELINE6   = HOST + "/mh/api/v50/post/contents/hide"// contentId + serviceCode
	TIMELINE7   = HOST + "/mh/api/v50/post/shared/list.json"
	TIMELINE8   = HOST + "/mh/api/v50/post/report.json"
	TIMELINE9   = HOST + "/mh/api/v50/post/sendPostToTalk.json"
	TIMELINE10  = HOST + "/mh/api/v50/post/list.json"

	// Like
	TIMELINE11  = HOST + "/mh/api/v50/like/get.json"
	TIMELINE12  = HOST + "/mh/api/v50/like/create.json"
	TIMELINE13  = HOST + "/mh/api/v50/like/cancel.json"
	TIMELINE14  = HOST + "/mh/api/v50/like/getList.json"
	TIMELINE15  = HOST + "/mh/api/v50/like/cancel/sharedPost.json"

	// Hashtag
	TIMELINE16  = HOST + "/mh/api/v50/hashtag/search.json"
	TIMELINE17  = HOST + "/mh/api/v50/hashtag/posts.json"
	TIMELINE18  = HOST + "/mh/api/v50/hashtag/popular.json"
	TIMELINE19  = HOST + "/mh/api/v50/hashtag/suggest.json"

	// Relay
	TIMELINE20  = HOST + "/mh/api/v50/relay/update.json"
	TIMELINE21  = HOST + "/mh/api/v50/relay/create.json"
	TIMELINE22  = HOST + "/mh/api/v50/relay/update.json"
	TIMELINE23  = HOST + "/mh/api/v50/relay/members.json"
	TIMELINE24  = HOST + "/mh/api/v50/relay/delete.json"
	TIMELINE25  = HOST + "/mh/api/v50/relay/sharelist/members.json"
	TIMELINE26  = HOST + "/mh/api/v50/relay/end.json"

	// Comment
	TIMELINE27  = HOST + "/mh/api/v50/comment/create.json"
	TIMELINE28  = HOST + "/mh/api/v50/comment/getList.json"
	TIMELINE29  = HOST + "/mh/api/v50/comment/delete.json"
	TIMELINE30  = HOST + "/mh/api/v50/comment/report.json"

	// Joinedrelay
	TIMELINE31  = HOST + "/mh/api/v50/joinedrelay/update.json"
	TIMELINE32  = HOST + "/mh/api/v50/joinedrelay/create.json"
	TIMELINE33  = HOST + "/mh/api/v50/joinedrelay/list.json"
	TIMELINE34  = HOST + "/mh/api/v50/joinedrelay/delete.json"
	TIMELINE35  = HOST + "/mh/api/v50/joinedrelay/masterdelete.json"

	// Home
	TIMELINE36  = HOST + "/mh/api/v50/home/updateCover.json"
	TIMELINE37  = HOST + "/mh/api/v50/home/post/list.json"
	TIMELINE38  = HOST + "/mh/api/v50/home/post/mediatab.json"
	TIMELINE39  = HOST + "/mh/api/v1/home/post/cover.json"
	TIMELINE40  = HOST + "/mh/api/v1/home/groupprofile.json"
	TIMELINE41  = HOST + "/mh/api/v1/home/story/recentcontent.json"
	TIMELINE42  = HOST + "/mh/api/v1/home/getBirthdayBoardId.json"
	TIMELINE43  = HOST + "/mh/api/v1/home/profile.json"

	// Birthday
	TIMELINE44  = HOST + "/mh/api/v50/birthday/card/create.json"
	TIMELINE45  = HOST + "/mh/api/v50/birthday/card/delete.json"
	TIMELINE46  = HOST + "/mh/api/v50/birthday/card/list.json"
	TIMELINE47  = HOST + "/mh/api/v50/birthday/template/list.json"
	TIMELINE48  = HOST + "/mh/api/v50/birthday/delete.json"
	TIMELINE49  = HOST + "/mh/api/v50/birthday/template/list.json"

	// Announce
	TIMELINE50  = HOST + "/mh/api/v50/announce/create.json" // homeId + postId
	TIMELINE51  = HOST + "/mh/api/v50/announce/delete.json" // homeId + postId
	TIMELINE52  = HOST + "/mh/api/v50/announce/postlist.json" // homeId + scrollId
	TIMELINE53  = HOST + "/mh/api/v50/announce/list.json" // homeId

	// Feed
	TIMELINE54  = HOST + "/mh/api/v50/feed/list.get" // postLimit + likeLimit + commentLimit + requestTime
	TIMELINE55  = HOST + "/mh/api/v50/feed/get.json" // likeLimit + commentLimit
	TIMELINE56  = HOST + "/mh/api/v50/feed/newfeed.json" // requestTime
	TIMELINE57  = HOST + "/mh/api/v50/feed/carousel/oa.json" // Unknown
	TIMELINE58  = HOST + "/mh/api/v50/feed/carousel/postmerge.json" // Unknown
	TIMELINE59  = HOST + "/mh/api/v50/feed/carousel/postmergeprivate.json" // Unknown

	// Other
	TIMELINE60  = HOST + "/mh/api/v50/prefetch/post/get.json"
	TIMELINE61  = HOST + "/mh/api/v50/share/toTalk.json"
	TIMELINE62  = HOST + "/mh/api/v50/reactions/get.json"
	TIMELINE63  = HOST + "/mh/api/v50/user/report.json"
	TIMELINE64  = HOST + "/mh/api/v50/search/note.json"
)


func Newtimeline(uri, bodydata string) {
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(bodydata)))
	if err != nil {
		fmt.Println(err)
		return
	}
	headers := map[string]string{
		"X-Line-Carrier": "44010, 1",
		"X-Line-ChannelToken": "Your ChannelToken",
		"X-Line-Application": "IOS\t9.16.6\tiOS\t13.1.3",
		"X-Line-Mid": "Your Mid",
		"User-Agent": "Line/9.16.6",
		"Accept-Language": "ja-jp",
		"Connection": "Keep-Alive",
	}
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := new(http.Client)
	resp, _ := client.Do(req)

	dump, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(dump))
	return
}

func main()  {
	body := `{"postInfo": {"readPermission": {"homeId": "Your Mid"}}, "sourceType": "TIMELINE", "contents": {"text": "TEST"}}`
	Newtimeline(TIMELINE2, body)
}
