package killswitch

import "context"

type Killswitch struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewKillswitch() *Killswitch {
	kill := &Killswitch{}
	kill.ctx, kill.cancel = context.WithCancel(context.Background())
	return kill
}

func (k *Killswitch) Err() error {
	return k.ctx.Err()
}

func (k *Killswitch) Cancel() {
	k.cancel()
}
func (k *Killswitch) Done() <-chan struct{} {
	return k.ctx.Done()
}
