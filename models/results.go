package models

type Result struct {
	Id          int `json:"user_id" db:"user_id"`
	Words       int `json:"words" db:"words"`
	Digits      int `json:"digits" db:"digits"`
	SpecialChar int `json:"special_char" db:"special_char"`
	Lines       int `json:"lines" db:"lines"`
	Spaces      int `json:"spaces" db:"spaces"`
	Sentences   int `json:"sentences" db:"sentences"`
	Punctuation int `json:"punctuation" db:"punctuation"`
	Consonants  int `json:"consonants" db:"consonants"`
	Vowels      int `json:"vowels" db:"vowels"`
}
