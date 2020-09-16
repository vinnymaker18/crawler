package fetchers

import (
	"errors"

	"github.com/vinnymaker18/crawler/core"
	"github.com/vinnymaker18/crawler/utils"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Key names used in crawler config file for twitter api auth parameters.
const (
	APP_KEY_KWD    = "twitter-app-key"
	APP_SECRET_KWD = "twitter-app-secret"
	UAT_KWD        = "twitter-user-access-token"
	UAT_SECRET_KWD = "twitter-user-access-token-secret"
)

// Limit the no. of links/tweets to fetch in a single request.
const SINGLE_FETCH_LIMIT = 20

// TweetFetcher is a twitter api client for fetching urls in tweets.
type TweetFetcher struct {
	client      *twitter.Client
	lastSinceID int64
}

// TweetFetcher constructor.
func NewTweetFetcher(cfg *core.Config) (*TweetFetcher, error) {
	configKV := cfg.CfgKeyValues
	appKey := configKV[APP_KEY_KWD]
	appSecret := configKV[APP_SECRET_KWD]
	appConfig := oauth1.NewConfig(appKey, appSecret)

	uat := configKV[UAT_KWD]
	uatSecret := configKV[UAT_SECRET_KWD]
	token := oauth1.NewToken(uat, uatSecret)

	httpClient := appConfig.Client(oauth1.NoContext, token)
	twClient := twitter.NewClient(httpClient)
	return &TweetFetcher{client: twClient, lastSinceID: 0}, nil
}

// Fetcher interface for TweetFetcher
func (tweetFetcher *TweetFetcher) FetchLinks() ([]core.LinkItem, error) {
	tweets, _, err := tweetFetcher.client.Timelines.HomeTimeline(
		&twitter.HomeTimelineParams{
			Count:          SINGLE_FETCH_LIMIT,
			SinceID:        tweetFetcher.lastSinceID,
			ExcludeReplies: utils.NewBool(true)})

	if err != nil {
		return nil, errors.New("Failed fetching new tweets " +
			err.Error())
	}

	urls := make([]core.LinkItem, 0, 20)
	for _, tweet := range tweets {
		tweetFetcher.lastSinceID = utils.Max(tweetFetcher.lastSinceID,
			tweet.ID)
		links, err := utils.ExtractURLs(tweet.Text)
		if err == nil {
			for _, link := range links {
				urls = append(urls, core.LinkItem{PageURL: link})
			}
		}
	}
	return urls, nil
}
