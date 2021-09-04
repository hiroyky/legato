package gqlmodel

const TrackName = "Track"
const AlbumName = "Album"
const AlbumArtistName = "AlbumArtist"
const GenreName = "Genre"

type Track struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Artist        string `json:"artist"`
	Composer      string `json:"composer"`
	TrackNo       int    `json:"trackNo"`
	Lyrics        string `json:"lyrics"`
	Comment       string `json:"comment"`
	Year          int    `json:"year"`
	URL           string `json:"url"`
	AlbumID       string `json:"album"`
	GenreID       string `json:"genre"`
	AlbumArtistID string `json:"albumArtist"`
}

func (Track) IsNode() {}

type Album struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	DiskNo        int    `json:"diskNo"`
	DiskTotal     int    `json:"diskTotal"`
	AlbumArtistID string `json:"albumArtist"`
}

func (Album) IsNode() {}

type AlbumArtist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (AlbumArtist) IsNode() {}

type Genre struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Genre) IsNode() {}
