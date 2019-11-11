package service

import (
	"backend_api/domain/entity"
	"backend_api/domain/repository"
)

type ProjectService struct {
	ProjectRepo         repository.ProjectRepository
	ProjectPositionRepo repository.ProjectPositionRepository
	ProjectSkillRepo    repository.ProjectSkillRepository
	ProjectMentorRepo   repository.ProjectMentorRepository
	ProjectUserRepo     repository.ProjectUserRepository
	PositionRepo        repository.PositionRepository
	SkillRepo           repository.SkillRepository
	MentorRepo          repository.MentorRepository
	UserRepo            repository.UserRepository
}

func NewProjectService(
	projectRepo repository.ProjectRepository,
	projectPositionRepo repository.ProjectPositionRepository,
	projectSkillRepo repository.ProjectSkillRepository,
	projectMentorRepo repository.ProjectMentorRepository,
	projectUserRepo repository.ProjectUserRepository,
	positionRepo repository.PositionRepository,
	skillRepo repository.SkillRepository,
	mentorRepo repository.MentorRepository,
	userRepo repository.UserRepository,
) *ProjectService {
	return &ProjectService{
		ProjectRepo:         projectRepo,
		ProjectPositionRepo: projectPositionRepo,
		ProjectSkillRepo:    projectSkillRepo,
		ProjectMentorRepo:   projectMentorRepo,
		ProjectUserRepo:     projectUserRepo,
		PositionRepo:        positionRepo,
		SkillRepo:           skillRepo,
		MentorRepo:          mentorRepo,
		UserRepo:            userRepo,
	}
}

// プロジェクトIDを指定して取得
func (svc *ProjectService) GetProjectResource(projectId int) (*entity.ProjectEntity, error) {
	var positionIds, skillIds, userIds, limitNums []int
	var skillNames []string
	var positionEntity []*entity.PositionEntity
	var skillEntity *entity.SkillEntity
	var userEntity []*entity.UserEntity

	// ProjectData
	project, err := svc.ProjectRepo.SelectProjectByProjectId(projectId)
	if err != nil {
		return nil, err
	}

	// ProjectPosition
	projectPosition, err := svc.ProjectPositionRepo.SelectProjectPositionByProjectId(projectId)
	if err != nil {
		return nil, err
	}

	// ProjectSkill
	projectSkill, err := svc.ProjectSkillRepo.SelectProjectSkillByProjectId(projectId)
	if err != nil {
		return nil, err
	}

	// ProjectMentor
	projectMentor, err := svc.ProjectMentorRepo.SelectProjectMentorByProjectId(projectId)
	if err != nil {
		return nil, err
	}

	// ProjectUsers
	projectUsers, err := svc.ProjectUserRepo.SelectProjectUserByProjectId(projectId)
	if err != nil {
		return nil, err
	}

	// Position
	for i, _ := range projectPosition {
		positionIds = append(positionIds, projectPosition[i].PositionId)
		limitNums = append(limitNums, projectPosition[i].LimitNum)
	}
	positions, err := svc.PositionRepo.SelectPositionsByPositionIDs(positionIds)
	if err != nil {
		return nil, err
	}
	for i, _ := range positions {
		positionEntity = append(positionEntity, &entity.PositionEntity{
			PositionName: positions[i].PositionName,
			LimitNum:     limitNums[i],
		})
	}

	// Skill
	for i, _ := range projectSkill {
		skillIds = append(skillIds, projectSkill[i].SkillId)
	}
	skills, err := svc.SkillRepo.SelectSkillsBySkillIDs(skillIds)
	if err != nil {
		return nil, err
	}
	for i, _ := range skills {
		skillNames = append(skillNames, skills[i].SkillName)
	}
	skillEntity = &entity.SkillEntity{
		Skill: skillNames,
	}

	// Mentor
	mentorData, err := svc.MentorRepo.SelectMentorById(projectMentor.MentorId)
	if err != nil {
		return nil, err
	}
	mentorEntity := entity.MentorEntity{Name: mentorData.Name}

	// User
	for i, _ := range projectUsers {
		userIds = append(userIds, projectUsers[i].UserId)
	}
	users, err := svc.UserRepo.SelectUsersByUserIDs(userIds)
	if err != nil {
		return nil, err
	}
	for i, _ := range users {
		userEntity = append(userEntity, &entity.UserEntity{
			UserName: users[i].Name,
		})
	}

	return &entity.ProjectEntity{
		ProjectData: project,
		Position:    positionEntity,
		Skill:       skillEntity,
		Mentor:      &mentorEntity,
		Users:       userEntity,
	}, nil
}

// プロジェクト全件取得
func (svc *ProjectService) SelectAllProjects() ([]*entity.ProjectEntity, error) {
	var results []*entity.ProjectEntity

	// ProjectDataの配列
	projects, err := svc.ProjectRepo.SelectAllProjects()
	if err != nil {
		return nil, err
	}

	for i, _ := range projects {

		var positionIds, skillIds, userIds, limitNums []int
		var skillNames []string
		var positionEntity []*entity.PositionEntity
		var skillEntity *entity.SkillEntity
		var userEntity []*entity.UserEntity

		// ProjectPosition
		projectPosition, err := svc.ProjectPositionRepo.SelectProjectPositionByProjectId(projects[i].Id)
		if err != nil {
			return nil, err
		}

		// ProjectSkill
		projectSkill, err := svc.ProjectSkillRepo.SelectProjectSkillByProjectId(projects[i].Id)
		if err != nil {
			return nil, err
		}

		// ProjectMentor
		projectMentor, err := svc.ProjectMentorRepo.SelectProjectMentorByProjectId(projects[i].Id)
		if err != nil {
			return nil, err
		}

		// ProjectUsers
		projectUsers, err := svc.ProjectUserRepo.SelectProjectUserByProjectId(projects[i].Id)
		if err != nil {
			return nil, err
		}

		// Position
		for i, _ := range projectPosition {
			positionIds = append(positionIds, projectPosition[i].PositionId)
			limitNums = append(limitNums, projectPosition[i].LimitNum)
		}
		positions, err := svc.PositionRepo.SelectPositionsByPositionIDs(positionIds)
		if err != nil {
			return nil, err
		}
		for i, _ := range positions {
			positionEntity = append(positionEntity, &entity.PositionEntity{
				PositionName: positions[i].PositionName,
				LimitNum:     limitNums[i],
			})
		}

		// Skill
		for i, _ := range projectSkill {
			skillIds = append(skillIds, projectSkill[i].SkillId)
		}
		skills, err := svc.SkillRepo.SelectSkillsBySkillIDs(skillIds)
		if err != nil {
			return nil, err
		}
		for i, _ := range skills {
			skillNames = append(skillNames, skills[i].SkillName)
		}
		skillEntity = &entity.SkillEntity{
			Skill: skillNames,
		}

		// Mentor
		mentorData, err := svc.MentorRepo.SelectMentorById(projectMentor.MentorId)
		if err != nil {
			return nil, err
		}
		mentorEntity := entity.MentorEntity{Name: mentorData.Name}

		// User
		for i, _ := range projectUsers {
			userIds = append(userIds, projectUsers[i].UserId)
		}
		users, err := svc.UserRepo.SelectUsersByUserIDs(userIds)
		if err != nil {
			return nil, err
		}
		for i, _ := range users {
			userEntity = append(userEntity, &entity.UserEntity{
				UserName: users[i].Name,
			})
		}

		results = append(results, &entity.ProjectEntity{
			ProjectData: projects[i],
			Position:    positionEntity,
			Skill:       skillEntity,
			Mentor:      &mentorEntity,
			Users:       userEntity,
		})
	}

	return results, nil
}
