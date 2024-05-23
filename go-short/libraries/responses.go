package libraries

type LinkResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`

	ID        string `bson:"id" json:"id,omitempty"`
	ShortLink string `bson:"shortlink" json:"shortlink,omitempty"`
}
