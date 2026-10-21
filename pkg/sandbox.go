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

func (m *Manager) WaitReady(ctx context.Context, id string) error {
	for {
		sb, err := m.store.Get(ctx, id)
		if err != nil {
			m.log.Errorf("sandbox %s not found while waiting for readiness: %v", id, err)
			return err
		}
		m.log.Infof("sandbox readiness poll: %+v", sb)
		if sb.Ready {
			return nil
		}
		time.Sleep(3 * time.Second)
	}
}
