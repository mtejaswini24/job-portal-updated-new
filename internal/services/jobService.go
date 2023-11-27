package services

import (
	"context"
	"encoding/json"
	"job-portal-api/internal/models"
	"sync"

	"github.com/redis/go-redis/v9"
)

func (s *Service) CreateJob(ctx context.Context, jobData models.NewJob, cid uint64) (models.Job, error) {
	jobDetails := models.Job{
		JobRole:     jobData.JobRole,
		Description: jobData.Description,
		Min_Np:      jobData.Min_Np,
		Max_Np:      jobData.Max_Np,
		Budget:      jobData.Budget,
		MinExp:      jobData.MinExp,
		MaxExp:      jobData.MaxExp,
		Cid:         uint(cid),
	}
	for _, v := range jobData.JobLocation {
		tempLocation := models.Location{
			Id: v,
		}
		jobDetails.JobLocation = append(jobDetails.JobLocation, tempLocation)
	}
	for _, v := range jobData.Technology {
		tempTechnolgy := models.Technology{
			Id: v,
		}
		jobDetails.Technology = append(jobDetails.Technology, tempTechnolgy)
	}
	for _, v := range jobData.WorkMode {
		tempWorkMode := models.WorkMode{
			Id: v,
		}
		jobDetails.WorkMode = append(jobDetails.WorkMode, tempWorkMode)
	}
	for _, v := range jobData.Qualification {
		tempQualification := models.Qualification{
			Id: v,
		}
		jobDetails.Qualification = append(jobDetails.Qualification, tempQualification)
	}
	for _, v := range jobData.Shift {
		tempShift := models.Shift{
			Id: v,
		}
		jobDetails.Shift = append(jobDetails.Shift, tempShift)
	}
	for _, v := range jobData.JobType {
		tempJobType := models.JobType{
			Id: v,
		}
		jobDetails.JobType = append(jobDetails.JobType, tempJobType)
	}
	jobDetails, err := s.userRepo.CreateJob(jobDetails)
	if err != nil {
		return models.Job{}, err
	}
	return jobDetails, nil
}
func (s *Service) FetchJob() ([]models.Job, error) {
	jobDetails, err := s.userRepo.ViewAllJobs()
	if err != nil {
		return nil, err
	}
	return jobDetails, nil
}
func (s *Service) FetchJobById(jid uint64) (models.Job, error) {
	jobDetails, err := s.userRepo.ViewJobDetailsById(jid)
	if err != nil {
		return models.Job{}, err
	}
	return jobDetails, nil
}
func (s *Service) FetchJobByCompanyId(cid uint64) ([]models.Job, error) {
	jobDetails, err := s.userRepo.ViewJobByCompanyId(cid)
	if err != nil {
		return nil, err
	}
	return jobDetails, nil
}
func (s *Service) ProcessJob(requestData []models.RequestJob) ([]models.RequestJob, error) {
	var finalData []models.RequestJob

	ch := make(chan models.RequestJob)
	wg := new(sync.WaitGroup)

	for _, v := range requestData {
		wg.Add(1)

		go func(v models.RequestJob) {
			defer wg.Done()
			var jobData models.Job
			val, err := s.rdb.GetTheCacheData(context.Background(), uint(v.Id))
			if err == redis.Nil {
				data, err := s.userRepo.ViewJobDetailsById(uint64(v.Id))
				if err != nil {
					return
				}
				err = s.rdb.AddToTheCache(context.Background(), uint(v.Id), data)
				if err != nil {
					return
				}
				jobData = data
			} else {
				err = json.Unmarshal([]byte(val), &jobData)
				if err == redis.Nil {
					return
				}
				if err != nil {
					return
				}
			}

			check, _ := s.compare(v, jobData)
			if check {
				ch <- v
			}

		}(v)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for applications := range ch {
		finalData = append(finalData, applications)
	}
	return finalData, nil
}
func (s *Service) compare(requestData models.RequestJob, jobData models.Job) (bool, models.RequestJob) {
	if requestData.Budget > jobData.Budget {
		return false, models.RequestJob{}
	}
	if requestData.Exp < jobData.MinExp || requestData.Exp > jobData.MaxExp {
		return false, models.RequestJob{}
	}
	if requestData.NoticePeriod < uint64(jobData.Min_Np) {
		return false, models.RequestJob{}
	}
	count := 0
	for _, v := range requestData.JobLocation {
		for _, v1 := range jobData.JobLocation {
			if v == v1.Id {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.RequestJob{}
	}
	count = 0
	for _, v := range requestData.Technology {
		for _, v1 := range jobData.Technology {
			if v == v1.Id {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.RequestJob{}
	}
	count = 0
	for _, v := range requestData.WorkMode {
		for _, v1 := range jobData.WorkMode {
			if v == v1.Id {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.RequestJob{}
	}
	count = 0
	for _, v := range requestData.Qualification {
		for _, v1 := range jobData.Qualification {
			if v == v1.Id {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.RequestJob{}
	}
	count = 0
	for _, v := range requestData.Shift {
		for _, v1 := range jobData.Shift {
			if v == v1.Id {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.RequestJob{}
	}
	count = 0
	for _, v := range requestData.JobType {
		for _, v1 := range jobData.JobType {
			if v == v1.Id {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.RequestJob{}
	}
	return true, requestData
}
