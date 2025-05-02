package pkg

import "go.uber.org/zap"

func (s *Sandbox) Create(id string) error {
	s.logger.Info("Entering Create")
	if err := s.provision(id); err != nil {
		s.logger.Error("sandbox provision failed", zap.String("sandboxId", id), zap.Error(err))
		return err
	}
	s.logger.Info("sandbox created", zap.String("sandboxId", id))
	return nil
}
