package types

import (
	"time"
)

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
}

type UserResponse struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Name         string `json:"name"`
	PicturePath  string `json:"picture_path"`
}

type UpdateUserRequest struct {
	Password string `json:"password"`
	Name     string `json:"name"`
}

type OperationResponse struct {
	Value      float64 `json:"value"`
	StatusCode int     `json:"status_code"`
	Message    string  `json:"message"`
}

type DeleteUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type EditProfileRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	PicturePath string `json:"picture_path"`
}

type ChangePasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}

type ClassResponse struct {
	ID               int    `json:"id"`
	OwnerID          int    `json:"owner_id"`
	OwnerName        string `json:"owner_name"`
	MemberAmount     int64  `json:"member_amount"`
	Code             string `json:"code"`
	Topic            string `json:"topic"`
	Description      string `json:"description"`
	Status           int    `json:"status"`
	Favorite         int    `json:"favorite"`
	BannerID         int    `json:"banner_id"`
	GoogleCourseID   string `json:"google_course_id"`
	GoogleCourseLink string `json:"google_course_link"`
	GoogleSyncedAt   string `json:"google_synced_at"`
}

type CreateClassRequest struct {
	Topic            string `json:"topic" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoogleCourseID   string `json:"google_course_id"`
	GoogleCourseLink string `json:"google_course_link"`
	BannerID         int    `json:"banner_id"`
	Code             string `json:"code"`
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
	ID          int                    `json:"id"`
	ClassID     int                    `json:"class_id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	DueDate     string                 `json:"due_date"`
	MaxAttempt  int                    `json:"max_attempt"`
	Settings    map[string]interface{} `json:"settings"`
	Condition   map[string]interface{} `json:"condition"`
	Grade       int                    `json:"grade"` // total grade of the assignment
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
	Title       string                 `json:"title" binding:"required"`
	Description string                 `json:"description"`
	DueDate     time.Time              `json:"due_date"`
	MaxAttempt  int                    `json:"max_attempt"`
	Settings    map[string]interface{} `json:"settings"`
	Condition   map[string]interface{} `json:"condition"`
	Grade       int                    `json:"grade"` // total grade of the assignment
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

// type AssignmentSettings struct {
// 	GradePolicy    GradePolicy    `json:"grade_policy"`
// 	TestCasePolicy TestCasePolicy `json:"test_case_policy"`
// 	FEBehavior     FEBehavior     `json:"fe_behavior"`
// }

type EditAssignmentRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	MaxAttempt  int    `json:"max_attempt"`
	Grade       int    `json:"grade"` // total grade of the assignment
	// Setting     AssignmentSettings  `json:"settings"`
	// Condition   AssignmentCondition `json:"condition"`
	Setting   map[string]interface{} `json:"setting"`
	Condition map[string]interface{} `json:"condition"`
}

type MemberResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PicturePath string    `json:"picture_path"`
	Role        string    `json:"role"`
	JoinAt      time.Time `json:"join_at"`
}

type ClassMeResponse struct {
	ID               int    `json:"id"`
	OwnerID          int    `json:"owner_id"`
	OwnerName        string `json:"owner_name"`
	MemberAmount     int64  `json:"member_amount"`
	Code             string `json:"code"`
	Topic            string `json:"topic"`
	Description      string `json:"description"`
	Status           int    `json:"status"`
	Favorite         int    `json:"favorite"`
	BannerID         int    `json:"banner_id"`
	GoogleCourseID   string `json:"google_course_id"`
	GoogleCourseLink string `json:"google_course_link"`
	GoogleSyncedAt   string `json:"google_synced_at"`
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
	Memory   map[string]int         `json:"memory" binding:"required"`
	Register map[string]int         `json:"register" binding:"required"`
	Flags    map[string]int         `json:"flags" binding:"required"`
	IOInput  map[string]string      `json:"io_input,omitempty"`
	Meta     map[string]interface{} `json:"_meta,omitempty"`
}

type TestCaseAssert struct {
	Memory   map[string]int    `json:"memory" binding:"required"`
	Register map[string]int    `json:"register" binding:"required"`
	Flags    map[string]int    `json:"flags" binding:"required"`
	Halted   bool              `json:"halted"`
	IOOutput map[string]string `json:"io_output,omitempty"`
}

type PlaygroundRequest struct {
	AssignmentID int                    `json:"assignment_id" binding:"required"`
	Item         map[string]interface{} `json:"item" binding:"required"`
	Status       string                 `json:"status" binding:"required"` // e.g., "in_progress", "completed", "failed"
}

type PlaygroundData struct {
	Items    map[string]interface{} `json:"items"`
	MetaData map[string]interface{} `json:"meta_data"`
	UI       map[string]interface{} `json:"ui"`
}

type PlaygroundItem struct {
	ID          int                  `json:"id"`
	Instruction string               `json:"instruction"`
	Label       string               `json:"label"`
	Operands    []PlaygroundOperands `json:"operands"`
	Next        *int                 `json:"next"`
	NextTrue    *int                 `json:"next_true"`
	NextFalse   *int                 `json:"next_false"`
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
	ID           int                    `json:"id"`
	AssignmentID int                    `json:"assignment_id"`
	UserID       int                    `json:"user_id"`
	Item         map[string]interface{} `json:"item"`
	Status       string                 `json:"status"` // e.g., "in_progress", "completed", "failed"
}

type PlaygroundMeRequest struct {
	AssignmentID int `json:"assignment_id" binding:"required"`
}

type ErrorStateDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	NodeID  int    `json:"node_id"`
}

type ExecutionStepLog struct {
	StepIndex   int            `json:"step_index"`
	NodeID      int            `json:"node_id"`
	Operation   string         `json:"operation"`
	Registers   map[string]int `json:"registers"`
	Flags       map[string]int `json:"flags"`
	MemoryDelta map[string]int `json:"memory_delta"`
	Stdout      []string       `json:"stdout"`
	Timestamp   time.Time      `json:"timestamp"`
}

type CourseData struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	AlternateLink  string `json:"alternateLink"`
	EnrollmentCode string `json:"enrollmentCode"`
}

type CourseWorkData struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"dueDate"`
	AlternateLink string `json:"alternateLink"`
	MaxPoints     int    `json:"maxPoints"`
}

type CourseWorkListResponse struct {
	CourseWork []CourseWorkData `json:"courseWork"`
}

type NewRoleRequest struct {
	ClassID int    `json:"class_id" binding:"required"`
	UserID  int    `json:"user_id" binding:"required"`
	NewRole string `json:"new_role" binding:"required"`
}

type RemoveMemberRequest struct {
	ClassID int `json:"class_id" binding:"required"`
	UserID  int `json:"user_id" binding:"required"`
}

type CreateSubmissionRequest struct {
	AssignmentID  int                    `json:"assignment_id"`
	PlaygroundID  int                    `json:"playground_id"`
	AttemptNumber int                    `json:"attempt_no"`
	ItemSnapshot  map[string]interface{} `json:"item_snapshot"`
	ClientResult  map[string]interface{} `json:"client_result"`
	ServerResult  map[string]interface{} `json:"server_result"`
	Score         float64                `json:"score"`
	Status        string                 `json:"status"`
	IsVerified    bool                   `json:"is_verified"`
	DurationMS    int                    `json:"duration_ms"`
}

type UpdateSubmissionRequest struct {
	AssignmentID  int                    `json:"assignment_id"`
	PlaygroundID  int                    `json:"playground_id"`
	AttemptNumber int                    `json:"attempt_no"`
	ItemSnapshot  map[string]interface{} `json:"item_snapshot"`
	ClientResult  map[string]interface{} `json:"client_result"`
	ServerResult  map[string]interface{} `json:"server_result"`
	Score         float64                `json:"score"`
	Status        string                 `json:"status"`
	IsVerified    bool                   `json:"is_verified"`
	DurationMS    int                    `json:"duration_ms"`
}

type SubmissionResponse struct {
	ID            int                    `json:"id"`
	AssignmentID  int                    `json:"assignment_id"`
	UserID        int                    `json:"user_id"`
	PlaygroundID  int                    `json:"playground_id"`
	AttemptNumber int                    `json:"attempt_no"`
	ItemSnapshot  map[string]interface{} `json:"item_snapshot"`
	ClientResult  map[string]interface{} `json:"client_result"`
	ServerResult  map[string]interface{} `json:"server_result"`
	Score         float64                `json:"score"`
	Status        string                 `json:"status"`
	IsVerified    bool                   `json:"is_verified"`
	DurationMS    int                    `json:"duration_ms"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
}

type NotificationRequest struct {
	UserID  int                    `json:"user_id" binding:"required"`
	Type    string                 `json:"type" binding:"required"`
	Title   string                 `json:"title" binding:"required"`
	Message string                 `json:"message" binding:"required"`
	Data    map[string]interface{} `json:"data"`
}

type NotificationResponse struct {
	ID        int                    `json:"id"`
	UserID    int                    `json:"user_id"`
	Type      string                 `json:"type"`
	Title     string                 `json:"title"`
	Message   string                 `json:"message"`
	Data      map[string]interface{} `json:"data"`
	IsRead    bool                   `json:"is_read"`
}