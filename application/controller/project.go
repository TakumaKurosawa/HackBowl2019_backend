package controller

import (
	"backend_api/application/controller/requestBody"
	"backend_api/application/controller/responseBody"
	"backend_api/domain/repository"
	"backend_api/domain/service"
	"backend_api/pkg/server/response"
	"encoding/json"
	"log"
	"net/http"
)

type ProjectCtrl struct {
	ProjectSrv *service.ProjectService
}

func NewProjectCtl(
	projectRepo repository.ProjectRepository,
	projectPositionRepo repository.ProjectPositionRepository,
	projectSkillRepo repository.ProjectSkillRepository,
	projectMentorRepo repository.ProjectMentorRepository,
	projectUserRepo repository.ProjectUserRepository,
	positionRepository repository.PositionRepository,
	skillRepository repository.SkillRepository,
	mentorRepository repository.MentorRepository,
	userRepository repository.UserRepository,
) *ProjectCtrl {
	return &ProjectCtrl{
		ProjectSrv: service.NewProjectService(
			projectRepo,
			projectPositionRepo,
			projectSkillRepo,
			projectMentorRepo,
			projectUserRepo,
			positionRepository,
			skillRepository,
			mentorRepository,
			userRepository,
		),
	}
}

// HandleProjectGet プロジェクト情報取得処理
func (ctrl *ProjectCtrl) HandleProjectGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// RequestBodyのパース
		var reqBody requestBody.Project
		err := json.NewDecoder(request.Body).Decode(&reqBody)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, "Invalid Request Body")
			return
		}

		selectedProject, err := ctrl.ProjectSrv.GetProjectResource(reqBody.ProjectId)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// レスポンスに必要な情報を詰めて返却
		response.Success(writer, responseBody.Project{
			ProjectId:   selectedProject.ProjectData.Id,
			ProjectName: selectedProject.ProjectData.Name,
			StartDate:   selectedProject.ProjectData.StartDate,
			Level:       selectedProject.ProjectData.Level,
			Position:    selectedProject.Position,
			Skill:       selectedProject.Skill,
			Mentor:      selectedProject.Mentor,
			Users:       selectedProject.Users,
		})
	}
}

// HandleAllProjectGet プロジェクト情報全件取得処理
func (ctrl *ProjectCtrl) HandleAllProjectGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var responseData []responseBody.Project

		selectedProjects, err := ctrl.ProjectSrv.SelectAllProjects()
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		for i, _ := range selectedProjects {
			responseData = append(responseData, responseBody.Project{
				ProjectId:   selectedProjects[i].ProjectData.Id,
				ProjectName: selectedProjects[i].ProjectData.Name,
				StartDate:   selectedProjects[i].ProjectData.StartDate,
				Level:       selectedProjects[i].ProjectData.Level,
				Position:    selectedProjects[i].Position,
				Skill:       selectedProjects[i].Skill,
				Mentor:      selectedProjects[i].Mentor,
				Users:       selectedProjects[i].Users,
			})
		}

		// レスポンスに必要な情報を詰めて返却
		response.Success(writer, responseBody.Projects{Projects: responseData})
	}
}
