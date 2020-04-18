package Ch9

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//const checkMark = "\u2713"
//const ballotX = "\u2717"

// feed is mocking the XML document we expect to receive
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
	<title>Going Go Programming<\title>
	<description>Golang : https://github.com/goinggo<\description>
	<link>http://goinggo.net/<\link>
	<item>
		<pubDate>Sun, 15 Mar 2015 15:04:00 +0000<\pubDate>
		<title>Object Oriented Programming Mechanics<\title>
		<description>Go is an object oriented language.<\description>
		<link>http://www.goingo.net/2015/03/object-oriented<\link>
	<\item>
<\channel>
<\rss>`

//mockserver returns a pointer to a server to handle the get call
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownload3(t *testing.T) {

	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(server.URL)
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
