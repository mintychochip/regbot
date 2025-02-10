package data

type Day int
type DayComparator int
type Term int
type Comparator int
type CourseComponent int
type ModeOfInstruction int

// Day
const (
	MONDAY Day = iota + 1
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

// DayComparator
const (
	EXCLUDE_ANY_OF_THESE DayComparator = iota + 1
	EXCLUDE_ONLY_THESE
	INCLUDE_ANY_OF_THESE
	INCLUDE_ONLY_THESE
)

// Term
const (
	WINTER = iota
	SPRING
	SUMMER
	FALL
)

// Mode of Instruction
const (
	FACE_TO_FACE ModeOfInstruction = iota
	HYBRID
	ITV
	MINDEPENDENT_STUDY
	ONLINE
)

const (
	POSTBAC = iota
	UNDERGRAD
	EXTENDED
)

const (
	ACTIVITY CourseComponent = iota
	CLINICAL
	CONTINUANCE
	DISCUSSION
	FIELD_STUDIES
	INDEPENDENT_STUDY
	LABORATORY
	LECTURE
	PRACTICUM
	RESEARCH
	SEMINAR
	SUPERVISION
	THESIS_RESEARCH
	TUTORIAL
)

type comparator int

const ()

type DayComparable struct {
	DayComparator DayComparator `json:"day_comparator"`
	Days          []Day         `json:"days"`
}
type Comparable[T any] struct {
	instance T
}
type SearchQuery struct {
	Subject      string `json:"subject"`
	CourseNumber uint16 `json:"course_number"`
	CourseCareer Term   `json:"course_career"`
	//MeetingStartTime
	//MeetingEndTime
	DaysOfWeek DayComparable `json:"days_of_week"`
	//InstructorLastName
	ClassNbr      uint32 `json:"class_nbr"`
	CourseKeyword string `json:"course_keyword"`
	//MinimumUnits
	//MaximumUnits
	CourseComponent CourseComponent `json:"course_component"`
	ModeOfInstruction
}
