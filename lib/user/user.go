package user

// User details required to create user
type User struct {
	Profile     Profile   `yaml:"profile"`
	Gender      Gender    `yaml:"gender"`
	Ethnicity   Ethnicity `yaml:"ethnicity"`
	Country     string    `yaml:"country"`
	Religon     Religon   `yaml:"religon"`
	Personality []string  `yaml:"personality"` // tags of words
	Age         Age       `yaml:"age"`
	Height      int       `yaml:"height"` // Height in CM
	Sexuality   Sexuality `yaml:"sexuality"`
}

// Profile a profile of the user
type Profile struct {
	AvatarName   string       `yaml:"avatar_name"`
	Initals      string       `yaml:"initals"`
	Verification Verification `yaml:"verification"`
}

// Gender the users gender
type Gender struct {
	Male   bool `yaml:"male"`
	Female bool `yaml:"female"`
}

// Ethnicity the users ethnicity
type Ethnicity struct {
	ethnicGroup map[string]string `yaml:"ethnic_group"`
	/*
		White
			English / Welsh / Scottish / Northern Irish / British
			Irish
			Gypsy or Irish Traveller
			Any other White background
		Mixed / Multiple ethnic groups
			White and Black Caribbean
			White and Black African
			White and Asian
			Any other Mixed / Multiple ethnic background
		Asian
			Indian
			Pakistani
			Bangladeshi
			Chinese
			Any other Asian background
		Black
			African
			Caribbean
			Other
		Other
			Arab
			Any other ethnic group
	*/
}

// Religon the users religon
type Religon struct {
	Any          bool      `yaml:"any"`
	Christian    Christian `yaml:"christian"`
	Islam        Islam     `yaml:"islam"`
	Agnostic     bool      `yaml:"agnostic"`
	Nonreligious bool      `yaml:"nonreligious"`
	Hinduism     Hinduism  `yaml:"hinduism"`
	Confucianism bool      `yaml:"confucianism"`
	Taoism       bool      `yaml:"taoism"`
	Buddhism     Buddhism  `yaml:"buddhism"`
	Sikhism      bool      `yaml:"sikhism"`
	Juche        bool      `yaml:"juche"`
	Spiritism    bool      `yaml:"spiritism"`
	Judaism      bool      `yaml:"judaism"`
	Bahai        bool      `yaml:"bahai"`
	Jainism      bool      `yaml:"jainism"`
	Shinto       bool      `yaml:"shinto"`
	CaoDai       bool      `yaml:"caoDai"`
	Other        bool      `yaml:"other"`
}

// Buddhism type of buddist
type Buddhism struct {
	Community string `yaml:"community"`
}

// Christian type of christan
type Christian struct {
	Community string `yaml:"community"`
}

// Hinduism type of Hindu
type Hinduism struct {
	Community string `yaml:"community"`
}

// Islam type of islam
type Islam struct {
	Community string `yaml:"community"`
}

// Age type of age
type Age struct {
	Day   int32 `yaml:"day"`
	Month int32 `yaml:"month"`
	Year  int32 `yaml:"year"`
}

// Verification confirms users identity
type Verification struct {
	EncryptedEmail    string `yaml:"encrypted_email"`
	EncryptedMobileNo string `yaml:"encrypted_mobile_number"`
	EncryptedPass     string `yaml:"encrypted_pass"`
	IncorrectCounts   uint8  `yaml:"incorrect_counts"`
}

// Sexuality the users sexuality
type Sexuality struct {
	Heterosexual bool `yaml:"heterosexual"`
	Homosexual   bool `yaml:"homosexual"`
	Bisexual     bool `yaml:"bisexual"`
}

func (v *Verification) wrongPassword() {
	if v.IncorrectCounts < 4 {
		v.IncorrectCounts++
	}
}
