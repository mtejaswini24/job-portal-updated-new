package services

// import (
// 	"context"
// 	"errors"
// 	"job-portal-api/internal/models"
// 	"job-portal-api/internal/repository"
// 	"reflect"
// 	"testing"

// 	"go.uber.org/mock/gomock"
// 	"gorm.io/gorm"
// )

// // func TestService_FetchJobById(t *testing.T) {
// // 	type args struct {
// // 		jid uint64
// // 	}
// // 	tests := []struct {
// // 		name             string
// // 		args             args
// // 		want             models.Job
// // 		wantErr          bool
// // 		mockRepoResponse func() (models.Job, error)
// // 	}{
// // 		{
// // 			name: "success",
// // 			want: models.Job{
// // 				Company: models.Company{
// // 					CompanyName: "hp",
// // 					Location:    "musuru",
// // 				},
// // 				Cid:            1,
// // 				JobRole:        "front-end",
// // 				JobDescription: "something",
// // 			},
// // 			args: args{
// // 				jid: 15,
// // 			},
// // 			wantErr: false,
// // 			mockRepoResponse: func() (models.Job, error) {
// // 				return models.Job{
// // 					Company: models.Company{
// // 						CompanyName: "hp",
// // 						Location:    "musuru",
// // 					},
// // 					Cid:            1,
// // 					JobRole:        "front-end",
// // 					JobDescription: "something",
// // 				}, nil
// // 			},
// // 		},
// // 		{
// // 			name: "invalid job id",
// // 			want: models.Job{},
// // 			args: args{
// // 				jid: 5,
// // 			},
// // 			mockRepoResponse: func() (models.Job, error) {
// // 				return models.Job{}, errors.New("error test")
// // 			},
// // 			wantErr: true,
// // 		},
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			mc := gomock.NewController(t)
// // 			mockRepo := repository.NewMockUserRepo(mc)
// // 			if tt.mockRepoResponse != nil {
// // 				mockRepo.EXPECT().ViewJobDetailsById(tt.args.jid).Return(tt.mockRepoResponse()).AnyTimes()
// // 			}
// // 			s, err := NewService(mockRepo)
// // 			if err != nil {
// // 				t.Errorf("error in initializing the repo layer")
// // 				return
// // 			}
// // 			got, err := s.FetchJobById(tt.args.jid)
// // 			if (err != nil) != tt.wantErr {
// // 				t.Errorf("Service.FetchJobById() error = %v, wantErr %v", err, tt.wantErr)
// // 				return
// // 			}
// // 			if !reflect.DeepEqual(got, tt.want) {
// // 				t.Errorf("Service.FetchJobById() = %v, want %v", got, tt.want)
// // 			}
// // 		})
// // 	}
// // }

// // func TestService_FetchJob(t *testing.T) {
// // 	tests := []struct {
// // 		name             string
// // 		want             []models.Job
// // 		wantErr          bool
// // 		mockRepoResponse func() ([]models.Job, error)
// // 	}{
// // 		{
// // 			name: "database success",
// // 			want: []models.Job{
// // 				{
// // 					Cid:            1,
// // 					JobRole:        "test",
// // 					JobDescription: "something",
// // 				},
// // 				{
// // 					Cid:            2,
// // 					JobRole:        "putvi",
// // 					JobDescription: "having",
// // 				},
// // 			},
// // 			wantErr: false,
// // 			mockRepoResponse: func() ([]models.Job, error) {
// // 				return []models.Job{
// // 					{
// // 						Cid:            1,
// // 						JobRole:        "test",
// // 						JobDescription: "something",
// // 					},
// // 					{
// // 						Cid:            2,
// // 						JobRole:        "putvi",
// // 						JobDescription: "having",
// // 					},
// // 				}, nil
// // 			},
// // 		},
// // 		{
// // 			name:    "database failure",
// // 			want:    nil,
// // 			wantErr: true,
// // 			mockRepoResponse: func() ([]models.Job, error) {
// // 				return nil, errors.New("error test")
// // 			},
// // 		},
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			mc := gomock.NewController(t)
// // 			mockRepo := repository.NewMockUserRepo(mc)
// // 			if tt.mockRepoResponse != nil {
// // 				mockRepo.EXPECT().ViewAllJobs().Return(tt.mockRepoResponse()).AnyTimes()
// // 			}
// // 			s, err := NewService(mockRepo)
// // 			if err != nil {
// // 				t.Errorf("error in initializing the repo layer")
// // 				return
// // 			}
// // 			got, err := s.FetchJob()
// // 			if (err != nil) != tt.wantErr {
// // 				t.Errorf("Service.FetchJob() error = %v, wantErr %v", err, tt.wantErr)
// // 				return
// // 			}
// // 			if !reflect.DeepEqual(got, tt.want) {
// // 				t.Errorf("Service.FetchJob() = %v, want %v", got, tt.want)
// // 			}
// // 		})
// // 	}
// // }

// // func TestService_FetchJobByCompanyId(t *testing.T) {
// // 	type args struct {
// // 		cid uint64
// // 	}
// // 	tests := []struct {
// // 		name             string
// // 		args             args
// // 		want             []models.Job
// // 		wantErr          bool
// // 		mockRepoResponse func() ([]models.Job, error)
// // 	}{
// // 		{
// // 			name: "success",
// // 			args: args{
// // 				cid: 2,
// // 			},
// // 			want: []models.Job{
// // 				{
// // 					Cid:            2,
// // 					JobRole:        "xyz",
// // 					JobDescription: "sleeping",
// // 				},
// // 				{
// // 					Cid:            2,
// // 					JobRole:        "abc",
// // 					JobDescription: "walking",
// // 				},
// // 			},
// // 			wantErr: false,
// // 			mockRepoResponse: func() ([]models.Job, error) {
// // 				return []models.Job{
// // 					{
// // 						Cid:            2,
// // 						JobRole:        "xyz",
// // 						JobDescription: "sleeping",
// // 					},
// // 					{
// // 						Cid:            2,
// // 						JobRole:        "abc",
// // 						JobDescription: "walking",
// // 					},
// // 				}, nil
// // 			},
// // 		},
// // 		{
// // 			name: "failure",
// // 			args: args{
// // 				cid: 10,
// // 			},
// // 			want:    nil,
// // 			wantErr: true,
// // 			mockRepoResponse: func() ([]models.Job, error) {
// // 				return nil, errors.New("data is not there")
// // 			},
// // 		},
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			mc := gomock.NewController(t)
// // 			mockRepo := repository.NewMockUserRepo(mc)
// // 			if tt.mockRepoResponse != nil {
// // 				mockRepo.EXPECT().ViewJobByCompanyId(tt.args.cid).Return(tt.mockRepoResponse()).AnyTimes()
// // 			}
// // 			s, err := NewService(mockRepo)
// // 			if err != nil {
// // 				t.Errorf("error in initializing the repo layer")
// // 				return
// // 			}
// // 			got, err := s.FetchJobByCompanyId(tt.args.cid)
// // 			if (err != nil) != tt.wantErr {
// // 				t.Errorf("Service.FetchJobByCompanyId() error = %v, wantErr %v", err, tt.wantErr)
// // 				return
// // 			}
// // 			if !reflect.DeepEqual(got, tt.want) {
// // 				t.Errorf("Service.FetchJobByCompanyId() = %v, want %v", got, tt.want)
// // 			}
// // 		})
// // 	}
// // }

// func TestService_CreateJob(t *testing.T) {
// 	type args struct {
// 		ctx     context.Context
// 		jobData models.NewJob
// 		cid     uint64
// 	}
// 	tests := []struct {
// 		name             string
// 		args             args
// 		want             models.Job
// 		wantErr          bool
// 		mockRepoResponse func() (models.Job, error)
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				ctx: context.Background(),
// 				jobData: models.NewJob{
// 					JobRole:       "java developer",
// 					Description:   "app development",
// 					Min_Np:        2,
// 					Max_Np:        5,
// 					Budget:        250000,
// 					JobLocation:   []uint{1, 2},
// 					Technology:    []uint{1, 2},
// 					WorkMode:      []uint{1, 2},
// 					MinExp:        2,
// 					MaxExp:        5,
// 					Qualification: []uint{1, 2},
// 					Shift:         []uint{1, 2},
// 					JobType:       []uint{1, 2},
// 				},
// 				cid: 1,
// 			},
// 			want: models.Job{
// 				JobRole:       "java developer",
// 				Description:   "app development",
// 				Min_Np:        2,
// 				Max_Np:        5,
// 				Budget:        250000,
// 				MinExp:        uint64(2),
// 				MaxExp:        uint64(5),
// 				JobLocation:   []models.Location{{Id: 1}, {Id: 2}},
// 				Technology:    []models.Technology{{Id: 1}, {Id: 2}},
// 				WorkMode:      []models.WorkMode{{Id: 1}, {Id: 2}},
// 				Qualification: []models.Qualification{{Id: 1}, {Id: 2}},
// 				Shift:         []models.Shift{{Id: 1}, {Id: 2}},
// 				JobType:       []models.JobType{{Id: 1}, {Id: 2}},
// 				Cid:           1,
// 			},
// 			wantErr: false,
// 			mockRepoResponse: func() (models.Job, error) {
// 				return models.Job{
// 					JobRole:       "java developer",
// 					Description:   "app development",
// 					Min_Np:        2,
// 					Max_Np:        5,
// 					Budget:        250000,
// 					MinExp:        uint64(2),
// 					MaxExp:        uint64(5),
// 					JobLocation:   []models.Location{{Id: 1}, {Id: 2}},
// 					Technology:    []models.Technology{{Id: 1}, {Id: 2}},
// 					WorkMode:      []models.WorkMode{{Id: 1}, {Id: 2}},
// 					Qualification: []models.Qualification{{Id: 1}, {Id: 2}},
// 					Shift:         []models.Shift{{Id: 1}, {Id: 2}},
// 					JobType:       []models.JobType{{Id: 1}, {Id: 2}},
// 					Cid:           1,
// 				}, nil
// 			},
// 		},
// 		{
// 			name: "failure",
// 			args: args{
// 				ctx: context.Background(),
// 				jobData: models.NewJob{
// 					JobRole:       "java developer",
// 					Description:   "app development",
// 					Min_Np:        2,
// 					Max_Np:        5,
// 					Budget:        250000,
// 					JobLocation:   []uint{1, 2},
// 					Technology:    []uint{1, 2},
// 					WorkMode:      []uint{1, 2},
// 					MinExp:        2,
// 					MaxExp:        5,
// 					Qualification: []uint{1, 2},
// 					Shift:         []uint{1, 2},
// 					JobType:       []uint{1, 2},
// 				},
// 				cid: 1,
// 			},
// 			want:    models.Job{},
// 			wantErr: true,
// 			mockRepoResponse: func() (models.Job, error) {
// 				return models.Job{}, errors.New("job is not created")
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			if tt.mockRepoResponse != nil {
// 				mockRepo.EXPECT().CreateJob(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
// 			}
// 			s, err := NewService(mockRepo)
// 			if err != nil {
// 				t.Errorf("error in initializing the repo layer")
// 				return
// 			}
// 			got, err := s.CreateJob(tt.args.ctx, tt.args.jobData, tt.args.cid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.CreateJob() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.CreateJob() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_ProcessJob(t *testing.T) {
// 	type args struct {
// 		requestData []models.RequestJob
// 	}
// 	tests := []struct {
// 		name             string
// 		s                *Service
// 		args             args
// 		want             []models.RequestJob
// 		wantErr          bool
// 		mockRepoResponse func(mockRepo *repository.MockUserRepo)
// 	}{
// 		{
// 			name: "failure case",
// 			args: args{
// 				requestData: []models.RequestJob{
// 					{
// 						Id:            1,
// 						JobRole:       "software engineer",
// 						Description:   "app development",
// 						NoticePeriod:  12,
// 						Budget:        23000,
// 						JobLocation:   []uint{1, 2},
// 						Technology:    []uint{1, 2},
// 						WorkMode:      []uint{1, 2},
// 						Exp:           2,
// 						Qualification: []uint{1, 2},
// 						Shift:         []uint{1, 2, 3},
// 						JobType:       []uint{1, 2},
// 					},
// 					// {

// 					// 	Id:            2,
// 					// 	JobRole:       "software engineer",
// 					// 	Description:   "app development",
// 					// 	NoticePeriod:  1,
// 					// 	Budget:        23000333,
// 					// 	JobLocation:   []uint{1, 2},
// 					// 	Technology:    []uint{1, 2},
// 					// 	WorkMode:      []uint{1, 2},
// 					// 	Exp:           2,
// 					// 	Qualification: []uint{1, 2},
// 					// 	Shift:         []uint{1, 2, 3},
// 					// 	JobType:       []uint{1, 2},
// 					// },
// 					// {

// 					// 	Id:            3,
// 					// 	JobRole:       "software engineer",
// 					// 	Description:   "app development",
// 					// 	NoticePeriod:  1,
// 					// 	Budget:        230,
// 					// 	JobLocation:   []uint{1, 2},
// 					// 	Technology:    []uint{1, 2},
// 					// 	WorkMode:      []uint{1, 2},
// 					// 	Exp:           2,
// 					// 	Qualification: []uint{1, 2},
// 					// 	Shift:         []uint{1, 2, 3},
// 					// 	JobType:       []uint{1, 2},
// 					// },
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockRepoResponse: func(mockRepo *repository.MockUserRepo) {
// 				mockRepo.EXPECT().ViewJobDetailsById(uint64(1)).Return(models.Job{
// 					JobRole:       "software engineer",
// 					Model:         gorm.Model{ID: uint(1)},
// 					Description:   "app development",
// 					Min_Np:        0,
// 					Max_Np:        5,
// 					Budget:        25000,
// 					JobLocation:   []models.Location{{Id: 1}},
// 					Technology:    []models.Technology{{Id: 1}},
// 					WorkMode:      []models.WorkMode{{Id: 1}},
// 					Qualification: []models.Qualification{{Id: 1}},
// 					Shift:         []models.Shift{{Id: 1}},
// 					JobType:       []models.JobType{{Id: 1}},
// 				}, nil).Times(1)

// 				// mockRepo.EXPECT().ViewJobDetailsById(uint64(2)).Return(models.Job{
// 				// 	JobRole:       "software engineer",
// 				// 	Model:         gorm.Model{ID: uint(1)},
// 				// 	Description:   "app development",
// 				// 	Min_Np:        2,
// 				// 	Max_Np:        5,
// 				// 	Budget:        25000,
// 				// 	JobLocation:   []models.Location{{Id: 2}},
// 				// 	Technology:    []models.Technology{{Id: 2}},
// 				// 	WorkMode:      []models.WorkMode{{Id: 2}},
// 				// 	Qualification: []models.Qualification{{Id: 2}},
// 				// 	Shift:         []models.Shift{{Id: 2}},
// 				// 	JobType:       []models.JobType{{Id: 2}},
// 				// }, nil).Times(1)

// 				// mockRepo.EXPECT().ViewJobDetailsById(uint64(3)).Return(models.Job{
// 				// 	JobRole:       "software engineer",
// 				// 	Model:         gorm.Model{ID: uint(3)},
// 				// 	Description:   "app development",
// 				// 	Min_Np:        2,
// 				// 	Max_Np:        5,
// 				// 	Budget:        25000,
// 				// 	JobLocation:   []models.Location{{Id: 3}},
// 				// 	Technology:    []models.Technology{{Id: 3}},
// 				// 	WorkMode:      []models.WorkMode{{Id: 3}},
// 				// 	Qualification: []models.Qualification{{Id: 3}},
// 				// 	Shift:         []models.Shift{{Id: 3}},
// 				// 	JobType:       []models.JobType{{Id: 3}},
// 				// }, nil).Times(1)
// 			},
// 		},
// 		// {
// 		// 	name: "success",
// 		// 	args: args{
// 		// 		requestData: []models.RequestJob{
// 		// 			{
// 		// 				Id:            1,
// 		// 				JobRole:       "software engineer",
// 		// 				Description:   "app development",
// 		// 				NoticePeriod:  2,
// 		// 				Budget:        23000,
// 		// 				JobLocation:   []uint{1, 2},
// 		// 				Technology:    []uint{1, 2},
// 		// 				WorkMode:      []uint{1, 2},
// 		// 				Exp:           2,
// 		// 				Qualification: []uint{1, 2},
// 		// 				Shift:         []uint{1, 2, 3},
// 		// 				JobType:       []uint{1, 2},
// 		// 			},
// 		// 		},
// 		// 	},
// 		// 	want: []models.RequestJob{
// 		// 		{
// 		// 			Id:            1,
// 		// 			JobRole:       "software engineer",
// 		// 			Description:   "app development",
// 		// 			NoticePeriod:  2,
// 		// 			Budget:        23000,
// 		// 			JobLocation:   []uint{1, 2},
// 		// 			Technology:    []uint{1, 2},
// 		// 			WorkMode:      []uint{1, 2},
// 		// 			Exp:           2,
// 		// 			Qualification: []uint{1, 2},
// 		// 			Shift:         []uint{1, 2, 3},
// 		// 			JobType:       []uint{1, 2},
// 		// 		},
// 		// 	},
// 		// 	wantErr: false,
// 		// 	mockRepoResponse: func(mockRepo *repository.MockUserRepo) {
// 		// 		mockRepo.EXPECT().ViewJobDetailsById(uint64(1)).Return(models.Job{
// 		// 			JobRole:       "software engineer",
// 		// 			Model:         gorm.Model{ID: uint(1)},
// 		// 			Description:   "app development",
// 		// 			Min_Np:        2,
// 		// 			Max_Np:        5,
// 		// 			Budget:        23000,
// 		// 			JobLocation:   []models.Location{{Id: 1}},
// 		// 			Technology:    []models.Technology{{Id: 1}},
// 		// 			WorkMode:      []models.WorkMode{{Id: 1}},
// 		// 			Qualification: []models.Qualification{{Id: 1}},
// 		// 			Shift:         []models.Shift{{Id: 1}},
// 		// 			JobType:       []models.JobType{{Id: 1}},
// 		// 		}, nil).Times(1)
// 		// 	},
// 		// },
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			tt.mockRepoResponse(mockRepo)

// 			s, err := NewService(mockRepo)
// 			if err != nil {
// 				t.Errorf("error in initializing the repo layer")
// 				return
// 			}
// 			got, err := s.ProcessJob(tt.args.requestData)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.ProcessJob() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.ProcessJob() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
