package model

// Book object [API MODEL]
type Book struct {
	ID       string   `json:"_id" bson:"_id,omitempty"`
	Title    string   `json:"title"`
	Author   *Author  `json:"author"`
	Category string   `json:"category"`
	Tags     []string `json:"tags,omitempty"`
}

// Author object [API MODEL]
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Config object represtons configuration of application
type Config struct {
	Database struct {
		Connection  string   `json:"connection"`
		Name        string   `json:"name"`
		Collcetions []string `json:"collcetions"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}
