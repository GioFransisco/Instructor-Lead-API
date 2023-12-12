package utilsmodel

const (
	//User
	UserLogin           = "Select id,password,email,role From users Where username=$1 or email=$1 or phone_number=$1"
	UserRegist          = "Insert Into users (name,email,phone_number,username,password,age,address,gender,role,updated_at) Values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id,created_at,updated_at"
	UserUpdate          = "Update users Set "
	UserUpdateReturning = " RETURNING id,name,email,phone_number,username,age,gender,address,created_at,updated_at"
	UserGetByEmail      = "Select id,name,email,phone_number,username,age,address,gender,role From users Where email=$1"
	UserChangePassword  = "Update users Set password=$1,updated_at=$2 Where id=$3 Returning id,name,email,phone_number,username,age,address,gender,role,created_at,updated_at"

	//Question
	QuestionCreate                = "INSERT INTO student_questions (schedule_details_id,student_id,question,status,updated_at) VALUES($1,$2,$3,$4,$5) RETURNING id,created_at,updated_at"
	DeleteQuestion                = "Delete From student_questions where id=$1"
	GetQuestionByIDScheduleDetail = "Select sq.id,sq.schedule_details_id,sq.student_id,sq.question,sq.status,sq.created_at,sq.updated_at,u.name,u.email,u.phone_number,u.username,u.age,u.address,u.gender,u.role,u.created_at,u.updated_at from student_questions sq join users u ON sq.student_id = u.id where sq.schedule_details_id=$1"
	UpdateQuestion                = "Update student_questions SET question=$1,updated_at=$2 WHERE id=$3 RETURNING id,schedule_details_id,student_id,question,status,created_at,updated_at"
	UpdateStatusQuestion          = "Update student_questions SET status=$1,updated_at=$2 WHERE id=$3 RETURNING id,schedule_details_id,student_id,question,status,created_at,updated_at"
	SelectIdQuestionById          = "Select id From student_questions Where id=$1"
	SelectQuestionById            = "Select * From student_questions Where id=$1"

	//Absences
	CreateAbsences               = "INSERT INTO absences (schedule_details_id, student_id, description, updated_at) VALUES ($1,$2,$3,$4) RETURNING id,created_at,updated_at"
	GetAbsences                  = "SELECT a.id, u.id, u.name, u.email, u.phone_number, u.username, u.password, u.age, u.address, u.gender, u.role, u.created_at, u.updated_at, a.description, a.created_at, a.updated_at FROM absences AS a JOIN users AS u ON u.id = a.student_id WHERE a.schedule_details_id = $1"
	GetScheduleDetailAbsenceById = "SELECT sd.id, s.id, s.name, s.date_activity, s.created_at, s.updated_at, u.id, u.name, u.email, u.phone_number, u.username, u.password, u.age, u.address, u.gender, u.role, u.created_at, u.updated_at, st.id, st.name, st.status, st.created_at, st.updated_at, sd.start_time, sd.end_time, sd.created_at, sd.updated_at FROM schedule_details AS sd JOIN schedules AS s ON s.id = sd.schedule_id JOIN users AS u ON u.id = sd.trainer_id JOIN stacks AS st ON st.id = sd.stack_id WHERE sd.id = $1"

	//Note
	NoteCreate     = "INSERT INTO notes (id, schedule_details_id, email, note, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, schedule_details_id, email, note, created_at, updated_at"
	NoteUpdate     = "UPDATE notes SET note = $1, updated_at = $2 WHERE id = $3 RETURNING id, email, note, schedule_details_id, created_at, updated_at"
	NoteList       = "SELECT id, schedule_details_id, note, email, created_at, updated_at FROM notes"
	NoteFindById   = "SELECT * FROM notes WHERE id = $1"
	DeleteNoteById = "DELETE FROM notes WHERE id = $1"
	
	//Stack
	StackCreate			 = "INSERT INTO stacks (name, status, updated_at) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
	StackList			 = "SELECT * FROM stacks ORDER BY created_at DESC"
	StackFindById		 = "SELECT * FROM stacks WHERE id = $1"
	StackUpdate          = "UPDATE stacks SET "
	StackUpdateReturning = " RETURNING id, name, status, created_at, updated_at"
	DeleteStackById		 = "DELETE FROM stacks WHERE id = $1"
	DefaultStatus		 = "Active"

	//Schedule
	ScheduleCreate          = "INSERT INTO schedules (name, date_activity, updated_at) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
	ScheduleGet             = "SELECT * FROM schedules ORDER BY created_at DESC"
	ScheduleGetById         = "SELECT * FROM schedules WHERE id = $1"
	ScheduleUpdate          = "UPDATE schedules SET "
	ScheduleUpdateReturning = " RETURNING id, name, date_activity, created_at, updated_at"

	//Schedule Detail
	ScheduleDetailCreate          = "INSERT INTO schedule_details (schedule_id , trainer_id, stack_id, start_time, end_time, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at"
	ScheduleDetailGet             = "SELECT sd.id, sd.start_time, sd.end_time, sd.created_at, sd.updated_at, t.id, t.name, t.email, t.phone_number, t.username, t.age, t.address, t.gender, t.role, t.created_at, t.updated_at, s.id, s.name, s.status, s.created_at, s.updated_at FROM schedule_details sd JOIN users t ON sd.trainer_id = t.id JOIN stacks s on sd.stack_id = s.id WHERE sd.schedule_id = $1 "
	ScheduleDetailGetById         = "SELECT * FROM schedule_details WHERE id = $1"
	ScheduleDetailUpdate          = "UPDATE schedule_details SET "
	ScheduleDetailUpdateReturning = " RETURNING id, trainer_id, stack_id, start_time, end_time, created_at, updated_at"

	//Schedule Approve
	GetScheduleApproveById = "SELECT * FROM schedule_approve WHERE schedule_details_id = $1"
	CreateScheduleApprove  = "INSERT INTO schedule_approve (schedule_details_id, schedule_approve, updated_at) VALUES ($1,$2,$3) RETURNING id, created_at, updated_at"
)
