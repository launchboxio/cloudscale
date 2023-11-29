package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

const (
	TargetGroupsBucket = "TargetGroups"
	CertificateBucket  = "Certificates"
	ListenersBucket    = "Listeners"
)

type Service struct {
	Db      *bbolt.DB
	channel chan struct{}
}

// SetChannel configures the event channel to emit create / update /delete
// events onto
func (s *Service) SetChannel(channel chan struct{}) {
	s.channel = channel
}

func (s *Service) Init() error {
	tx, err := s.Db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.CreateBucketIfNotExists([]byte(TargetGroupsBucket)); err != nil {
		return err
	}
	if _, err := tx.CreateBucketIfNotExists([]byte(CertificateBucket)); err != nil {
		return err
	}
	if _, err := tx.CreateBucketIfNotExists([]byte(ListenersBucket)); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *Service) ListCertificates() ([]*Certificate, error) {
	var certificates []*Certificate
	err := s.readAll(CertificateBucket, func(k, v []byte) error {
		var certificate *Certificate
		if err := json.Unmarshal(v, &certificate); err != nil {
			return err
		}
		certificate.Id = string(k)
		certificates = append(certificates, certificate)
		return nil
	})
	return certificates, err
}

func (s *Service) GetCertificate(certificateId string) (*Certificate, error) {
	var certificate *Certificate
	data, err := s.getRecord(CertificateBucket, certificateId)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, certificate)
	return certificate, err
}

func (s *Service) CreateCertificate(certificate *Certificate) (*Certificate, error) {
	certificate.Id = uuid.New().String()
	err := s.Db.Update(func(tx *bbolt.Tx) error {
		data, err := json.Marshal(certificate)
		if err != nil {
			return err
		}
		return tx.Bucket([]byte(CertificateBucket)).Put([]byte(certificate.Id), data)
	})
	if err != nil {
		return nil, err
	}
	s.emitUpdate()
	return certificate, nil
}

func (s *Service) UpdateCertificate(certificate *Certificate) (*Certificate, error) {
	err := s.Db.Update(func(tx *bbolt.Tx) error {
		data, err := json.Marshal(certificate)
		if err != nil {
			return err
		}

		return tx.Bucket([]byte(CertificateBucket)).Put([]byte(certificate.Id), data)
	})
	if err != nil {
		return nil, err
	}
	s.emitUpdate()
	return certificate, nil
}

func (s *Service) DeleteCertificate(certificateId string) error {
	err := s.Db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(CertificateBucket)).Delete([]byte(certificateId))
	})
	if err == nil {
		s.emitUpdate()
	}
	return err
}

func (s *Service) ListTargetGroups() ([]*TargetGroup, error) {
	var targetGroups []*TargetGroup
	err := s.readAll(TargetGroupsBucket, func(k, v []byte) error {
		var targetGroup *TargetGroup
		if err := json.Unmarshal(v, &targetGroup); err != nil {
			return err
		}
		targetGroup.Id = string(k)
		targetGroups = append(targetGroups, targetGroup)
		return nil
	})
	return targetGroups, err
}

func (s *Service) GetTargetGroup(targetGroupId string) (*TargetGroup, error) {
	var targetGroup *TargetGroup
	data, err := s.getRecord(TargetGroupsBucket, targetGroupId)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, targetGroup)
	return targetGroup, err
}

func (s *Service) CreateTargetGroup(group *TargetGroup) (*TargetGroup, error) {
	group.Id = uuid.New().String()
	err := s.Db.Update(func(tx *bbolt.Tx) error {
		data, err := json.Marshal(group)
		if err != nil {
			return err
		}
		return tx.Bucket([]byte(TargetGroupsBucket)).Put([]byte(group.Id), data)
	})
	if err != nil {
		return nil, err
	}
	s.emitUpdate()
	return group, nil
}

func (s *Service) UpdateTargetGroup(group *TargetGroup) (*TargetGroup, error) {
	err := s.Db.Update(func(tx *bbolt.Tx) error {
		data, err := json.Marshal(group)
		if err != nil {
			return err
		}

		return tx.Bucket([]byte(TargetGroupsBucket)).Put([]byte(group.Id), data)
	})
	if err != nil {
		return nil, err
	}
	s.emitUpdate()
	return group, nil
}

func (s *Service) DestroyTargetGroup(targetGroupId string) error {
	err := s.Db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(TargetGroupsBucket)).Delete([]byte(targetGroupId))
	})
	if err == nil {
		s.emitUpdate()
	}
	return err
}

func (s *Service) ListListeners() ([]*Listener, error) {
	var listeners []*Listener
	err := s.readAll(ListenersBucket, func(k, v []byte) error {
		var listener *Listener
		if err := json.Unmarshal(v, &listener); err != nil {
			return err
		}
		listener.Id = string(k)
		listeners = append(listeners, listener)
		return nil
	})
	return listeners, err
}

func (s *Service) GetListener(listenerId string) (*Listener, error) {
	var listener *Listener
	data, err := s.getRecord(ListenersBucket, listenerId)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, listener)
	return listener, err
}

func (s *Service) CreateListener(listener *Listener) (*Listener, error) {
	listener.Id = uuid.New().String()
	err := s.Db.Update(func(tx *bbolt.Tx) error {
		data, err := json.Marshal(listener)
		if err != nil {
			return err
		}
		return tx.Bucket([]byte(ListenersBucket)).Put([]byte(listener.Id), data)
	})
	if err != nil {
		return nil, err
	}
	s.emitUpdate()
	return listener, nil
}

func (s *Service) UpdateListener(listener *Listener) (*Listener, error) {
	err := s.Db.Update(func(tx *bbolt.Tx) error {
		data, err := json.Marshal(listener)
		if err != nil {
			return err
		}

		return tx.Bucket([]byte(ListenersBucket)).Put([]byte(listener.Id), data)
	})
	if err != nil {
		return nil, err
	}
	s.emitUpdate()
	return listener, nil
}

func (s *Service) DestroyListener(listenerId string) error {
	err := s.Db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(ListenersBucket)).Delete([]byte(listenerId))
	})
	if err == nil {
		s.emitUpdate()
	}
	return err
}

func (s *Service) readAll(bucket string, f func(k, v []byte) error) error {
	return s.Db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).ForEach(f)
	})
}

func (s *Service) getRecord(bucket string, identifier string) ([]byte, error) {
	var data []byte
	err := s.Db.View(func(tx *bbolt.Tx) error {
		data = tx.Bucket([]byte(bucket)).Get([]byte(identifier))
		return nil
	})
	return data, err
}

func (s *Service) emitUpdate() {
	if s.channel != nil {
		s.channel <- struct{}{}
	}
}
