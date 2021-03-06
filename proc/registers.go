package proc

import "fmt"
import "errors"

// Registers is an interface for a generic register type. The
// interface encapsulates the generic values / actions
// we need independant of arch. The concrete register types
// will be different depending on OS/Arch.
type Registers interface {
	PC() uint64
	SP() uint64
	CX() uint64
	TLS() uint64
	Get(int) (uint64, error)
	SetPC(*Thread, uint64) error
	String() string
}

var UnknownRegisterError = errors.New("unknown register")

// Registers obtains register values from the debugged process.
func (t *Thread) Registers() (Registers, error) {
	regs, err := registers(t)
	if err != nil {
		return nil, fmt.Errorf("could not get registers: %s", err)
	}
	return regs, nil
}

// PC returns the current PC for this thread.
func (t *Thread) PC() (uint64, error) {
	regs, err := t.Registers()
	if err != nil {
		return 0, err
	}
	return regs.PC(), nil
}
