package driver

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/config"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/driver/controller"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/driver/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/manager"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type serverRequirment struct {
	engine  *gin.Engine
	cfg     *config.Config
	loger   common.MyLogger
	jwtAuth common.JwtToken
	host    string
	uc      manager.UseCaseManager
}

func (s *serverRequirment) setUpController() {
	v1 := s.engine.Group(config.ApiGroup)

	authMiddleware := middleware.NewMiddlewareAuth(s.jwtAuth)

	controller.NewAuthController(s.uc.AuthUsecase(), v1, authMiddleware).Router()
	controller.NewUserController(s.uc.UserUsecase(), v1, authMiddleware).RouterUser()
	controller.NewStackController(s.uc.StackUseCase(), v1, authMiddleware).Route()
	controller.NewNoteController(s.uc.NoteUseCase(), v1, authMiddleware).Route()
	controller.NewScheduleController(s.uc.ScheduleUseCase(), v1, authMiddleware).Route()
	controller.NewQuestionController(s.uc.QuestionUsecase(), v1, authMiddleware).QuestionRouter()
	controller.NewAbsencesController(s.uc.AbsenceUseCase(), v1, authMiddleware).Route()
	controller.NewScheduleApproveController(s.uc.ScheduleApproveUseCase(), v1, authMiddleware).Route()
}

func (s *serverRequirment) Run() {
	s.engine.Use(middleware.NewLogMiddleware(s.loger).LogRequest())

	s.setUpController()

	if err := s.engine.Run(":" + s.host); err != nil {
		panic(err)
	}
}

func NewServer() *serverRequirment {
	eng := gin.Default()
	cfg := config.NewConfig()
	token := common.NewJwtToken(cfg)
	myLog := common.NewMyLogger(cfg.LogConfig)
	infra := manager.NewInfraManager(cfg)
	repo := manager.NewRepoManager(infra)
	uc := manager.NewUseCaseManager(repo, token)

	return &serverRequirment{
		engine:  eng,
		cfg:     cfg,
		loger:   myLog,
		host:    cfg.ApiPort,
		uc:      uc,
		jwtAuth: token,
	}
}
