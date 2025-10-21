package types

type CatResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type OAuthRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
}
type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
}

type UserResponse struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	Tel          string `json:"tel"`
	Picture_path string `json:"picture_path"`
}

type UpdateUserRequest struct {
	Password string `json:"password"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
}

type OperationResponse struct {
	Value      float64 `json:"value"`
	StatusCode int     `json:"status_code"`
	Message    string  `json:"message"`
}

// DeleteUserRequest represents the request body for deleting a user
type DeleteUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type EditProfileRequest struct {
	Username     string `json:"username"`
	Name         string `json:"name"`
	Tel          string `json:"tel"`
	Picture_path string `json:"picture_path"`
}

type ChangePasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}

type ClassResponse struct {
	ID               int    `json:"id"`
	Topic            string `json:"topic"`
	Description      string `json:"description"`
	GoogleCourseID   string `json:"google_course_id"`
	GoogleCourseLink string `json:"google_course_link"`
	GoogleSyncedAt   string `json:"google_synced_at"`
	FavScore         int64  `json:"fav_score"`
	Owner            int    `json:"owner"`
	Status           int    `json:"status"`
}

type CreateClassRequest struct {
	Topic            string `json:"topic" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoogleCourseID   string `json:"google_course_id"`
	GoogleCourseLink string `json:"google_course_link"`
	Status           int    `json:"status"` // public or private
}

type UpdateClassRequest struct {
	Topic            string `json:"topic"`
	Description      string `json:"description"`
	GoogleCourseID   string `json:"google_course_id"`
	GoogleCourseLink string `json:"google_course_link"`
	Status           int    `json:"status"` // public or private
}

type AssignmentResponse struct {
	ID          int                 `json:"id"`
	ClassID     int                 `json:"class_id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	DueDate     string              `json:"due_date"`
	MaxAttempt  int                 `json:"max_attempt"`
	Grade       int                 `json:"grade"` // total grade of the assignment
	Settings    AssignmentSettings  `json:"settings"`
	Condition   AssignmentCondition `json:"condition"`
}

type AllowedInstructions struct {
	System                   map[string]int `json:"system"`                     // LABEL, NOP, HLT
	DataMovement             map[string]int `json:"data_movement"`              // LOAD, STORE, MOV
	Arithmetic               map[string]int `json:"arithmetic"`                 // ADD, SUB, INC, DEC, MUL, DIV
	ComparisonAndConditional map[string]int `json:"comparison_and_conditional"` // CMP, JMP, JC, JNZ, JZ, JNC
}

type Memory struct {
	Address int `json:"address"`
	Value   int `json:"value"`
}

type InitialState struct {
	Memory []Memory `json:"memory"`
}

type ExecutionConstraints struct {
	RegisterCount int          `json:"register_count"`
	MemoryNode    int          `json:"memory_node"`
	InitialState  InitialState `json:"initial_state"`
}

type AssignmentCondition struct {
	AllowedInstructions  AllowedInstructions  `json:"allowed_instructions"`
	ExecutionConstraints ExecutionConstraints `json:"execution_constraints"`
}

type CreateAssignmentRequest struct {
	Title       string              `json:"title" binding:"required"`
	Description string              `json:"description" binding:"required"`
	MaxAttempt  int                 `json:"max_attempt"`
	Grade       int                 `json:"grade"` // total grade of the assignment
	Settings    AssignmentSettings  `json:"settings"`
	Condition   AssignmentCondition `json:"condition"`
}

type GradePolicy struct {
	Mode   string       `json:"mode"`
	Weight WeightPolicy `json:"weight"`
}

type WeightPolicy struct {
	TestCaseWeight         float64 `json:"test_case"`
	NumberOfNodeUsedWeight float64 `json:"number_of_node_used"`
}

type TestCasePolicy struct {
	VisibleToStudent bool `json:"visible_to_student"`
}

type FEBehavior struct {
	LockAfterSubmit       bool `json:"lock_after_submit"`
	AllowResubmitAfterDue bool `json:"allow_resubmit_after_due"`
}

type AssignmentSettings struct {
	GradePolicy    GradePolicy    `json:"grade_policy"`
	TestCasePolicy TestCasePolicy `json:"test_case_policy"`
	FEBehavior     FEBehavior     `json:"fe_behavior"`
}

type EditAssignmentRequest struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	MaxAttempt  int                 `json:"max_attempt"`
	Grade       int                 `json:"grade"` // total grade of the assignment
	Setting     AssignmentSettings  `json:"settings"`
	Condition   AssignmentCondition `json:"condition"`
}

type MemberResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Picture_path string `json:"picture_path"`
}

type ClassMeResponse struct {
	ID          int    `json:"id"`
	Topic       string `json:"topic"`
	Description string `json:"description"`
	FavScore    int64  `json:"fav_score"`
	Owner       int    `json:"owner"`
}

type InvitationResponse struct {
	ID               int    `json:"id"`
	ClassID          int    `json:"class_id"`
	UserID           int    `json:"user_id"`
	InvitationEmail  string `json:"invitation_email"`
	GoogleInviteCode string `json:"google_invite_code"`
	Status           string `json:"status"`
}

type UploadAvatarRequest struct {
	AvatarURL string `json:"avatar_url" binding:"required,url"`
}

type TestSuiteRequest struct {
	Name string `json:"name" binding:"required"`
}

type TestSuiteResponse struct {
	ID           int    `json:"id"`
	AssignmentID int    `json:"assignment_id"`
	Name         string `json:"name"`
}

type TestCaseRequest struct {
	Name   string         `json:"name" binding:"required"`
	Init   TestCaseInit   `json:"init" binding:"required"`
	Assert TestCaseAssert `json:"assert" binding:"required"`
}

type TestCaseResponse struct {
	ID          int            `json:"id"`
	TestSuiteID int            `json:"test_suite_id"`
	Name        string         `json:"name" binding:"required"`
	Init        TestCaseInit   `json:"init"`
	Assert      TestCaseAssert `json:"assert"`
}

type TestCaseInit struct {
	Memory   map[string]int `json:"memory" binding:"required"`
	Register map[string]int `json:"register" binding:"required"`
	Flags    map[string]int `json:"flags" binding:"required"`
}

type TestCaseAssert struct {
	Memory   map[string]int `json:"memory" binding:"required"`
	Register map[string]int `json:"register" binding:"required"`
	Flags    map[string]int `json:"flags" binding:"required"`
	Halted   bool           `json:"halted" binding:"required"`
}

type PlaygroundRequest struct {
	AssignmentID int            `json:"assignment_id" binding:"required"`
	AttemptNO    int            `json:"attempt_no" binding:"required"`
	Item         PlaygroundData `json:"item" binding:"required"`
	Status       string         `json:"status" binding:"required"` // e.g., "in_progress", "completed", "failed"
}

type PlaygroundData struct {
	Items    []PlaygroundItem   `json:"items"`
	MetaData PlaygroundMetaData `json:"meta_data"`
	UI       PlaygroundUI       `json:"ui"`
}

type PlaygroundItem struct {
	ID          int                  `json:"id"`
	Instruction string               `json:"instruction"`
	Label       string               `json:"label"`
	Operands    []PlaygroundOperands `json:"operands"`
	Next        int                  `json:"next"`
	NextTrue    int                  `json:"next_true"`
	NextFalse   int                  `json:"next_false"`
}

type PlaygroundOperands struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type PlaygroundMetaData struct {
	ProgramName string `json:"program_name"`
	AuthorID    int    `json:"author_id"`
	Timestamp   string `json:"timestamp"`
}

type PlaygroundPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PlaygroundUI struct {
	Position map[string]PlaygroundPosition `json:"position"`
	Zoom     float64                       `json:"zoom"`
	Pan      PlaygroundPosition            `json:"pan"`
}

type PlaygroundResponse struct {
	ID           int            `json:"id"`
	AssignmentID int            `json:"assignment_id"`
	UserID       int            `json:"user_id"`
	AttemptNO    int            `json:"attempt_no"`
	Item         PlaygroundData `json:"item"`
	Status       string         `json:"status"` // e.g., "in_progress", "completed", "failed"
}

type PlaygroundMeRequest struct {
	AssignmentID int `json:"assignment_id" binding:"required"`
}
