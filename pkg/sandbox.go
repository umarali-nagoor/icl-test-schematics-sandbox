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

func (m *Manager) BulkProvision(ctx context.Context, sandboxes []*Sandbox) error {
	for _, sb := range sandboxes {
		m.log.Infof("provisioning sandbox: %+v", sb)
		if err := m.provision(ctx, sb); err != nil {
			m.log.Errorf("provision failed for sandbox %+v: %v", sb, err)
			return err
		}
		m.log.Infof("sandbox ready: %+v", sb)
	}
	return nil
}
