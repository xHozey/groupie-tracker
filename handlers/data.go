package handlers

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type Locations struct {
	Id        int            `json:"id"`
	Locations map[int]string `json:"locations"`
	Dates     string         `json:"dates"`
}
type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Dates struct {
	Index []Date
}
type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
type Relations struct {
	Index []Relation
}