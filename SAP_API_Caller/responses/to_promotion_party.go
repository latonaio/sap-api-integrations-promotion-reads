package responses

type ToPromotionParty struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID       string `json:"ObjectID"`
			ParentObjectID string `json:"ParentObjectID"`
			ETag           string `json:"ETag"`
			ID             string `json:"ID"`
			RoleCode       string `json:"RoleCode"`
			RoleCodeText   string `json:"RoleCodeText"`
			Name           string `json:"Name"`
			MainIndicator  bool   `json:"MainIndicator"`
		} `json:"results"`
	} `json:"d"`
}