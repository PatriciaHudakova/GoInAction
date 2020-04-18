package Ch9

import (
	"net/http"
	"testing"
)

//validates the ability for GET to download content
func TestDownload(t *testing.T) {

	const checkMark = "\u2713"
	const ballotX = "\u2717"

	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	t.Log("Given the need to test downloading content")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call", ballotX, err)
			}
			t.Log("\t\tShould be able to make the get call", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status, %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
