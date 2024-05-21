package main

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	MembersCount int
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"`
	Concerts     []Concert
	AuthorMemb   string
	AllNames     string
}

type Concert struct {
	Country string
	City    string
	Dates   []string
}

// type Relation struct {
// 	Id            uint                   `json:"id"`
// 	DateLocations map[string]interface{} `json:"datesLocations"`
// }

// type Relation struct {
// 	Id            uint                `json:"id"`
// 	DateLocations map[string][]string `json:"datesLocations"`
// }
