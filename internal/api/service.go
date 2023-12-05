package api

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	Db      *gorm.DB
	channel chan struct{}
}

type Base struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	base.ID = id
	return nil
}

// SetChannel configures the event channel to emit create / update /delete
// events onto
func (s *Service) SetChannel(channel chan struct{}) {
	s.channel = channel
}

func (s *Service) Init() error {
	return s.Db.AutoMigrate(
		&Certificate{},
		&TargetGroup{},
		&TargetGroupAttachment{},
		&Listener{},
		&Rule{},
	)
}

func (s *Service) ListCertificates() ([]*Certificate, error) {
	var certificates []*Certificate
	if err := s.Db.Find(&certificates).Error; err != nil {
		return nil, err
	}
	return certificates, nil
}

func (s *Service) GetCertificate(certificateId string) (*Certificate, error) {
	var certificate *Certificate
	if err := s.Db.First(&certificate, "id = ?", certificateId).Error; err != nil {
		return nil, err
	}
	return certificate, nil
}

func (s *Service) CreateCertificate(certificate *Certificate) (*Certificate, error) {
	result := s.Db.Create(&certificate)
	s.emitUpdate()
	return certificate, result.Error
}

func (s *Service) UpdateCertificate(certificate *Certificate) (*Certificate, error) {
	result := s.Db.Save(certificate)
	s.emitUpdate()
	return certificate, result.Error
}

func (s *Service) DeleteCertificate(certificateId string) error {
	return s.Db.Delete(&Certificate{}, certificateId).Error
}

func (s *Service) ListTargetGroups() ([]*TargetGroup, error) {
	var targetGroups []*TargetGroup
	if err := s.Db.Preload("Attachments").Find(&targetGroups).Error; err != nil {
		return nil, err
	}
	return targetGroups, nil
}

func (s *Service) GetTargetGroup(targetGroupId string) (*TargetGroup, error) {
	var targetGroup *TargetGroup
	if err := s.Db.Preload("Attachments").First(&targetGroup, "id = ?", targetGroupId).Error; err != nil {
		return nil, err
	}
	return targetGroup, nil
}

func (s *Service) CreateTargetGroup(group *TargetGroup) (*TargetGroup, error) {
	result := s.Db.Create(&group)
	s.emitUpdate()
	return group, result.Error
}

func (s *Service) UpdateTargetGroup(group *TargetGroup) (*TargetGroup, error) {
	result := s.Db.Save(group)
	s.emitUpdate()
	return group, result.Error
}

func (s *Service) DestroyTargetGroup(targetGroupId string) error {
	return s.Db.Delete(&TargetGroup{}, targetGroupId).Error
}

func (s *Service) ListListeners() ([]*Listener, error) {
	var listeners []*Listener
	if err := s.Db.Preload("Rules").Find(&listeners).Error; err != nil {
		return nil, err
	}
	return listeners, nil
}

func (s *Service) GetListener(listenerId string) (*Listener, error) {
	var listener *Listener
	if err := s.Db.Preload("Rules").First(&listener, "id = ?", listenerId).Error; err != nil {
		return nil, err
	}
	return listener, nil
}

func (s *Service) CreateListener(listener *Listener) (*Listener, error) {
	result := s.Db.Create(&listener)
	s.emitUpdate()
	return listener, result.Error
}

func (s *Service) UpdateListener(listener *Listener) (*Listener, error) {
	result := s.Db.Save(listener)
	s.emitUpdate()
	return listener, result.Error
}

func (s *Service) DestroyListener(listenerId string) error {
	err := s.Db.Delete(&Listener{}, "id = ?", listenerId).Error
	s.emitUpdate()
	return err
}

func (s *Service) emitUpdate() {
	if s.channel != nil {
		s.channel <- struct{}{}
	}
}
