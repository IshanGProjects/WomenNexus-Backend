package models

type Education struct {
	School       string  `json:"school"`
	DegreeType   string  `json:"degreeType"`
	FieldOfStudy string  `json:"fieldOfStudy"`
	Minor        string  `json:"minor,omitempty"`
	StartDate    string  `json:"startDate"`
	EndDate      string  `json:"endDate"`
	Classes      []Class `json:"classes"`
}

type Class struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Experience struct {
	Title       string `json:"title"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Description string `json:"description"`
}

type Accomplishment struct {
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type Project struct {
	Title        string   `json:"title"`
	StartDate    string   `json:"startDate"`
	EndDate      string   `json:"endDate"`
	Link         string   `json:"link"`
	Contributors []string `json:"contributors"`
	Description  string   `json:"description"`
}

type UserProfile struct {
	FirstName       string           `json:"firstName"`
	LastName        string           `json:"lastName"`
	DOB             string           `json:"dob"`
	Pronouns        string           `json:"pronouns"`
	Education       []Education      `json:"education"`
	Experiences     []Experience     `json:"experiences"`
	Skills          []string         `json:"skills"`
	Interests       []string         `json:"interests"`
	Hobbies         []string         `json:"hobbies"`
	Accomplishments []Accomplishment `json:"accomplishments"`
	Projects        []Project        `json:"projects"`
	Goals           string           `json:"goals"`
	ProfileImage    string           `json:"profileImage"`
	ShortBio        string           `json:"shortBio"`
}
