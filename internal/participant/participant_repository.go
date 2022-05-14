package participant

import (
	"github.com/rewle/service-select-participants/internal/db"
	"go.uber.org/zap"
)

type ParticipantRepository struct {
	dbRepo *db.DBRepository
	log    *zap.SugaredLogger
} // link 2 calendar

type ParticipantListResponseItem struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Calendar string `json:"calendar"`
}

type ParticipantListResponse []ParticipantListResponseItem

type ParticipantCreateRequest struct {
	Name     string `json:"name"`
	Calendar string `json:"calendar"`
}

type ParticipantCreateResponse struct {
	Err string `json:"error"`
	ID  int    `json:"id"`
}

type ParticipantRemoveRequest struct {
	ID int `json:"id"`
}

type ParticipantRemoveResponse struct{}

type Participant struct {
	ID       int
	Name     string
	Calendar string
}

func (p *ParticipantRepository) List() (*ParticipantListResponse, error) {
	conn := p.dbRepo.GetConnection()
	var participants []Participant
	err := conn.Model(&participants).Select()
	if err != nil {
		p.log.Error("Participant: DB error", err)
		return nil, err
	}
	resp := make(ParticipantListResponse, len(participants))
	for i, p := range participants {
		resp[i] = ParticipantListResponseItem(p)
	}
	return &resp, nil
}

func (p *ParticipantRepository) Create(r *ParticipantCreateRequest) (*ParticipantCreateResponse, error) {
	conn := p.dbRepo.GetConnection()
	m := Participant{
		Name:     r.Name,
		Calendar: r.Calendar,
	}
	_, err := conn.Model(&m).Insert()
	if err != nil {
		p.log.Info("participant:create", err)
		return nil, err
	}
	return &ParticipantCreateResponse{"", m.ID}, nil
}

func (p *ParticipantRepository) Remove(r *ParticipantRemoveRequest) (*ParticipantRemoveResponse, error) {
	conn := p.dbRepo.GetConnection()
	_, err := conn.Model(&Participant{}).Where("id = ?", r.ID).Delete()
	if err != nil {
		p.log.Info("participant:remove", err)
		return nil, err
	}
	return &ParticipantRemoveResponse{}, nil
}

func Init(dbRepo *db.DBRepository, log *zap.SugaredLogger) *ParticipantRepository {
	return &ParticipantRepository{
		dbRepo,
		log,
	}
}
