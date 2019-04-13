package domain

import "fmt"

type Movie struct {
    Name        string `json:"name"`
    Duration    string `json:"duration"`
    Genre       string `json:"genre"`
    Country     string `json:"country"`
    Slogan      string `json:"slogan"`
    Director    string `json:"director"`
    Scenario    string `json:"scenario"`
    Producers   string `json:"producers"`
    Operator    string `json:"operator"`
    Artists     string `json:"artists"`
    Composer    string `json:"composer"`
    Budget      string `json:"budget"`
    Premiere    string `json:"premier"`
    Quality     string `json:"quality"`
    Sound       string `json:"sound"`
    Image       string `json:"image"`
    Description string `json:"description"` 
}

func (m Movie) String() string {
    return fmt.Sprintf("Movie(\n"+
        " Name: %s,\n"+
        " Duration: %s,\n"+
        " Genre: %s,\n"+
        " Country: %s,\n"+
        " Slogan: %s,\n"+
        " Director: %s,\n"+
        " Scenario: %s,\n"+
        " Producers: %s,\n"+
        " Operator: %s,\n"+
        " Composer: %s,\n"+
        " Budget: %s,\n"+
        " Premiere: %s,\n"+
        " Quality: %s,\n"+
        " Sound: %s,\n"+
        " Image: %s,\n"+
        " Description: %s\n"+
        ")",
        m.Name,
        m.Duration,
        m.Genre,
        m.Country,
        m.Slogan,
        m.Director,
        m.Scenario,
        m.Producers,
        m.Operator,
        m.Composer,
        m.Budget,
        m.Premiere,
        m.Quality,
        m.Sound,
        m.Image,
        m.Description)
}