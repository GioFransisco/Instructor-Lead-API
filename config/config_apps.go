package config

const (
	RoleAdmin      = "Admin"
	RolePartcipant = "Participant"
	RoleTrainer    = "Trainer"

	//user path
	ApiGroup           = "api/v1"
	LoginPath          = "auth/login"
	RegisterPath       = "auth/register"
	UserUpdatePath     = "users"
	UserDeleteId       = "/:id"
	UserGetById        = "/id/:id"
	UserGEtByEmail     = "/:email"
	UserChangePassword = "password"

	//schedule path
	ScheduleCreatePath       = "schedules"
	ScheduleGetPath          = "schedules"
	ScheduleUpdatePath       = "schedules/:id"
	ScheduleDetailUpdatePath = "schedule-details/:id"

	//question path
	QuestionCreatePath       = "questions"
	QuestionGetPath          = "/:scheduleDetailId"
	QuestionUpdatePath       = "/:id/question"
	QuestionUpdateStatusPath = "/:id/status"
	QuestionDeletePath       = "/:id"

	//absences path
	AbsencesCreatePath = "absences"
	AbsencesGetPath    = "absences"

	//note path
	NoteGroupPath   = "notes"
	NoteGetPath     = ""
	NoteCreatePath  = ""
	NoteUpdatePath  = ":id"
	NoteGetByIDPath = ":id"
	NoteDeletePath = ":id"

	//schedule aprove path
	ScheduleAprovePath = "schedule-aprove"

	//stack path
	StackGroupPath   = "stacks"
	StackCreatePath  = ""
	StackGetPath     = ""
	StackGetByIdPath = ":id"
	StackUpdatePath  = ":id"
	StackDeletePath  = ":id"

	//error description
	ErrorDescriptionForInvalidData = "make sure all data is filled in correctly"
)
