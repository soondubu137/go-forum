package redis

const (
    KEY_POST_TIME_ZSET = "forum::post::time"
    KEY_POST_SCORE_ZSET = "forum::post::score"
    KEY_POST_VOTED_ZSET_PREFIX = "forum::post::voted::"
    KEY_COMMUNITY_SET_PREFIX = "forum::community::"
)