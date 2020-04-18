package Ch9

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

//validates the ability for GET to download content & handles different status conditions
func TestDownload2(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.goinggo.net/feeds/posts/default?alt=rss",
			http.StatusOK,
		},
		{
			"http://rss.cnn.com/rss/ann_topstbadurl.rss",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading different content")
	{
		for _, u := range urls {
			t.Logf("\t When checking \"%s\" for status code \"%d\"", u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould not be able to Get the url", ballotX, err)
				}
				t.Log("\t\tShould be able to get the url", checkMark)

				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould have a \"%d\" status %v", u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status %v %v", u.statusCode, ballotX, resp.StatusCode)
				}
			}
		}
	}
}
