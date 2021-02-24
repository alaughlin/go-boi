package gameboy

import (
	"fmt"
	"math/bits"
)

type cpu struct {
	a  *byte
	b  *byte
	c  *byte
	d  *byte
	e  *byte
	f  *byte
	h  *byte
	l  *byte
	sp uint16
	pc uint16
}

var (
	flagZ int = 7
	flagN int = 6
	flagH int = 5
	flagC int = 4
)

func initializeCPU() *cpu {
	c := byte(0x13)
	e := byte(0xD8)
	f := byte(0xB0)
	h := byte(0x01)
	l := byte(0x4D)

	return &cpu{
		c:  &c,
		e:  &e,
		f:  &f,
		h:  &h,
		l:  &l,
		sp: 0xFFFE,
		pc: 0x100,
	}
}

func (cpu *cpu) af() uint16 {
	return double(*cpu.a, *cpu.f)
}

func (cpu *cpu) bc() uint16 {
	return double(*cpu.b, *cpu.c)
}

func (cpu *cpu) de() uint16 {
	return double(*cpu.d, *cpu.e)
}

func (cpu *cpu) hl() uint16 {
	return double(*cpu.h, *cpu.l)
}

func double(n1 byte, n2 byte) uint16 {
	return uint16(n1)<<8 | uint16(n2)
}

func (cpu *cpu) getBit(n byte, bitPos int) byte {
	return (n >> bitPos) & 1
}

func (cpu *cpu) setBit(n byte, bitPos int) byte {
	return n | (1 << bitPos)
}

func (cpu *cpu) clearBit(n byte, bitPos int) byte {
	return n &^ (1 << bitPos)
}

func (cpu *cpu) setFlag(flagPos int) {
	*cpu.f = cpu.setBit(*cpu.f, flagPos)
}

func (cpu *cpu) clearFlag(flagPos int) {
	*cpu.f = cpu.clearBit(*cpu.f, flagPos)
}

func (cpu *cpu) ExecuteOpcode(memory *memory) {
	opcode := memory.read(cpu.pc)

	switch opcode {
	// 8-Bit Loads
	case 0x06:
		cpu.ld_r(cpu.b, memory.read(cpu.pc+1), 2)
	case 0x0E:
		cpu.ld_r(cpu.c, memory.read(cpu.pc+1), 2)
	case 0x16:
		cpu.ld_r(cpu.d, memory.read(cpu.pc+1), 2)
	case 0x1E:
		cpu.ld_r(cpu.e, memory.read(cpu.pc+1), 2)
	case 0x26:
		cpu.ld_r(cpu.h, memory.read(cpu.pc+1), 2)
	case 0x2E:
		cpu.ld_r(cpu.l, memory.read(cpu.pc+1), 2)
	case 0x7F:
		cpu.ld_r(cpu.a, *cpu.a, 1)
	case 0x78:
		cpu.ld_r(cpu.a, *cpu.b, 1)
	case 0x79:
		cpu.ld_r(cpu.a, *cpu.c, 1)
	case 0x7A:
		cpu.ld_r(cpu.a, *cpu.d, 1)
	case 0x7B:
		cpu.ld_r(cpu.a, *cpu.e, 1)
	case 0x7C:
		cpu.ld_r(cpu.a, *cpu.h, 1)
	case 0x7D:
		cpu.ld_r(cpu.a, *cpu.l, 1)
	case 0x7E:
		cpu.ld_r(cpu.a, memory.read(cpu.hl()), 1)
	case 0x40:
		cpu.ld_r(cpu.b, *cpu.b, 1)
	case 0x41:
		cpu.ld_r(cpu.b, *cpu.c, 1)
	case 0x42:
		cpu.ld_r(cpu.b, *cpu.d, 1)
	case 0x43:
		cpu.ld_r(cpu.b, *cpu.e, 1)
	case 0x44:
		cpu.ld_r(cpu.b, *cpu.h, 1)
	case 0x45:
		cpu.ld_r(cpu.b, *cpu.l, 1)
	case 0x46:
		cpu.ld_r(cpu.b, memory.read(cpu.hl()), 1)
	case 0x48:
		cpu.ld_r(cpu.c, *cpu.b, 1)
	case 0x49:
		cpu.ld_r(cpu.c, *cpu.c, 1)
	case 0x4A:
		cpu.ld_r(cpu.c, *cpu.d, 1)
	case 0x4B:
		cpu.ld_r(cpu.c, *cpu.e, 1)
	case 0x4C:
		cpu.ld_r(cpu.c, *cpu.h, 1)
	case 0x4D:
		cpu.ld_r(cpu.c, *cpu.l, 1)
	case 0x4E:
		cpu.ld_r(cpu.c, memory.read(cpu.hl()), 1)
	case 0x50:
		cpu.ld_r(cpu.d, *cpu.b, 1)
	case 0x51:
		cpu.ld_r(cpu.d, *cpu.c, 1)
	case 0x52:
		cpu.ld_r(cpu.d, *cpu.d, 1)
	case 0x53:
		cpu.ld_r(cpu.d, *cpu.e, 1)
	case 0x54:
		cpu.ld_r(cpu.d, *cpu.h, 1)
	case 0x55:
		cpu.ld_r(cpu.d, *cpu.l, 1)
	case 0x56:
		cpu.ld_r(cpu.d, memory.read(cpu.hl()), 1)
	case 0x58:
		cpu.ld_r(cpu.e, *cpu.b, 1)
	case 0x59:
		cpu.ld_r(cpu.e, *cpu.c, 1)
	case 0x5A:
		cpu.ld_r(cpu.e, *cpu.d, 1)
	case 0x5B:
		cpu.ld_r(cpu.e, *cpu.e, 1)
	case 0x5C:
		cpu.ld_r(cpu.e, *cpu.h, 1)
	case 0x5D:
		cpu.ld_r(cpu.e, *cpu.l, 1)
	case 0x5E:
		cpu.ld_r(cpu.e, memory.read(cpu.hl()), 1)
	case 0x60:
		cpu.ld_r(cpu.h, *cpu.b, 1)
	case 0x61:
		cpu.ld_r(cpu.h, *cpu.c, 1)
	case 0x62:
		cpu.ld_r(cpu.h, *cpu.d, 1)
	case 0x63:
		cpu.ld_r(cpu.h, *cpu.e, 1)
	case 0x64:
		cpu.ld_r(cpu.h, *cpu.h, 1)
	case 0x65:
		cpu.ld_r(cpu.h, *cpu.l, 1)
	case 0x66:
		cpu.ld_r(cpu.h, memory.read(cpu.hl()), 1)
	case 0x68:
		cpu.ld_r(cpu.l, *cpu.b, 1)
	case 0x69:
		cpu.ld_r(cpu.l, *cpu.c, 1)
	case 0x6A:
		cpu.ld_r(cpu.l, *cpu.d, 1)
	case 0x6B:
		cpu.ld_r(cpu.l, *cpu.e, 1)
	case 0x6C:
		cpu.ld_r(cpu.l, *cpu.h, 1)
	case 0x6D:
		cpu.ld_r(cpu.l, *cpu.l, 1)
	case 0x6E:
		cpu.ld_r(cpu.l, memory.read(cpu.hl()), 1)
	case 0x70:
		cpu.ld_addr(cpu.hl(), *cpu.b, memory, 1)
	case 0x71:
		cpu.ld_addr(cpu.hl(), *cpu.c, memory, 1)
	case 0x72:
		cpu.ld_addr(cpu.hl(), *cpu.d, memory, 1)
	case 0x73:
		cpu.ld_addr(cpu.hl(), *cpu.e, memory, 1)
	case 0x74:
		cpu.ld_addr(cpu.hl(), *cpu.h, memory, 1)
	case 0x75:
		cpu.ld_addr(cpu.hl(), *cpu.l, memory, 1)
	case 0x36:
		cpu.ld_addr(cpu.hl(), memory.read(cpu.pc+1), memory, 2)
	case 0x0A:
		cpu.ld_r(cpu.a, memory.read(cpu.bc()), 1)
	case 0x1A:
		cpu.ld_r(cpu.a, memory.read(cpu.de()), 1)
	case 0xFA:
		cpu.ld_r(cpu.a, memory.read(uint16(cpu.pc+2)<<8|uint16(cpu.pc+1)), 3)
	case 0x3E:
		cpu.ld_r(cpu.a, memory.read(cpu.pc+1), 2)
	case 0x47:
		cpu.ld_r(cpu.b, *cpu.a, 1)
	case 0x4F:
		cpu.ld_r(cpu.c, *cpu.a, 1)
	case 0x57:
		cpu.ld_r(cpu.d, *cpu.a, 1)
	case 0x5F:
		cpu.ld_r(cpu.e, *cpu.a, 1)
	case 0x67:
		cpu.ld_r(cpu.h, *cpu.a, 1)
	case 0x6F:
		cpu.ld_r(cpu.l, *cpu.a, 1)
	case 0x02:
		cpu.ld_addr(cpu.bc(), *cpu.a, memory, 1)
	case 0x12:
		cpu.ld_addr(cpu.de(), *cpu.a, memory, 1)
	case 0x77:
		cpu.ld_addr(cpu.hl(), *cpu.a, memory, 1)
	case 0xEA:
		cpu.ld_addr(uint16(cpu.pc+2)<<8|uint16(cpu.pc+1), *cpu.a, memory, 3)
	case 0xF2:
		cpu.ld_r(cpu.a, memory.read(0xFF00+uint16(*cpu.c)), 1)
	case 0xE2:
		cpu.ld_addr(0xFF00+uint16(*cpu.c), *cpu.a, memory, 1)
	case 0x3A:
		cpu.ld_r(cpu.a, memory.read(cpu.hl()), 1)
		memory.decrement(cpu.hl())
	case 0x32:
		cpu.ld_addr(cpu.hl(), *cpu.a, memory, 1)
		memory.decrement(cpu.hl())
	case 0x2A:
		cpu.ld_r(cpu.a, memory.read(cpu.hl()), 1)
		memory.increment(cpu.hl())
	case 0x22:
		cpu.ld_addr(cpu.hl(), *cpu.a, memory, 1)
		memory.increment(cpu.hl())
	case 0xE0:
		cpu.ld_addr(0xFF00+uint16(memory.read(cpu.pc+1)), *cpu.a, memory, 1)
	case 0xF0:
		cpu.ld_r(cpu.a, memory.read(0xFF00+uint16(memory.read(cpu.pc+1))), 1)
	// 16-Bit Loads
	case 0x01:
		cpu.ld_r_double(cpu.b, cpu.c, memory.readDouble(cpu.pc+1), 3)
	case 0x11:
		cpu.ld_r_double(cpu.d, cpu.e, memory.readDouble(cpu.pc+1), 3)
	case 0x21:
		cpu.ld_r_double(cpu.h, cpu.l, memory.readDouble(cpu.pc+1), 3)
	case 0x31:
		cpu.ld_sp(memory.readDouble(cpu.pc+1), 3)
	case 0xF9:
		cpu.ld_sp(cpu.hl(), 1)
	case 0xF8:
		// TODO: this is awful. check this. also need to set flags.
		cpu.ld_r_double(cpu.h, cpu.l, uint16(int16(cpu.sp)+int16(memory.read(cpu.pc+1))), 2)
	case 0x08:
		cpu.ld_addr_double(double(memory.read(cpu.pc+1), memory.read(cpu.pc+2)), cpu.sp, memory, 3)
	case 0xF5:
		cpu.push(memory.readDouble(cpu.af()), memory)
	case 0xC5:
		cpu.push(memory.readDouble(cpu.bc()), memory)
	case 0xD5:
		cpu.push(memory.readDouble(cpu.de()), memory)
	case 0xE5:
		cpu.push(memory.readDouble(cpu.hl()), memory)
	case 0xF1:
		cpu.pop(cpu.a, cpu.f, memory)
	case 0xC1:
		cpu.pop(cpu.b, cpu.c, memory)
	case 0xD1:
		cpu.pop(cpu.d, cpu.e, memory)
	case 0xE1:
		cpu.pop(cpu.h, cpu.l, memory)
	case 0x87:
		cpu.add_r(cpu.a, *cpu.a, 1)
	case 0x80:
		cpu.add_r(cpu.a, *cpu.b, 1)
	case 0x81:
		cpu.add_r(cpu.a, *cpu.c, 1)
	case 0x82:
		cpu.add_r(cpu.a, *cpu.d, 1)
	case 0x83:
		cpu.add_r(cpu.a, *cpu.e, 1)
	case 0x84:
		cpu.add_r(cpu.a, *cpu.h, 1)
	case 0x85:
		cpu.add_r(cpu.a, *cpu.l, 1)
	case 0x86:
		cpu.add_r(cpu.a, memory.read(cpu.hl()), 1)
	case 0xC6:
		cpu.add_r(cpu.a, memory.read(cpu.pc+1), 2)
	case 0x8F:
		cpu.adc_r(cpu.a, *cpu.a, 1)
	case 0x88:
		cpu.adc_r(cpu.a, *cpu.b, 1)
	case 0x89:
		cpu.adc_r(cpu.a, *cpu.c, 1)
	case 0x8A:
		cpu.adc_r(cpu.a, *cpu.d, 1)
	case 0x8B:
		cpu.adc_r(cpu.a, *cpu.e, 1)
	case 0x8C:
		cpu.adc_r(cpu.a, *cpu.h, 1)
	case 0x8D:
		cpu.adc_r(cpu.a, *cpu.l, 1)
	case 0x8E:
		cpu.adc_r(cpu.a, memory.read(cpu.hl()), 1)
	case 0xCE:
		cpu.adc_r(cpu.a, memory.read(cpu.pc+1), 2)
	case 0x97:
		cpu.sub_r(cpu.a, *cpu.a, 1)
	case 0x90:
		cpu.sub_r(cpu.a, *cpu.b, 1)
	case 0x91:
		cpu.sub_r(cpu.a, *cpu.c, 1)
	case 0x92:
		cpu.sub_r(cpu.a, *cpu.d, 1)
	case 0x93:
		cpu.sub_r(cpu.a, *cpu.e, 1)
	case 0x94:
		cpu.sub_r(cpu.a, *cpu.h, 1)
	case 0x95:
		cpu.sub_r(cpu.a, *cpu.l, 1)
	case 0x96:
		cpu.sub_r(cpu.a, memory.read(cpu.hl()), 1)
	case 0xD6:
		cpu.sub_r(cpu.a, memory.read(cpu.pc+1), 2)
	case 0x9F:
		cpu.sbc_r(cpu.a, *cpu.a, 1)
	case 0x98:
		cpu.sbc_r(cpu.a, *cpu.b, 1)
	case 0x99:
		cpu.sbc_r(cpu.a, *cpu.c, 1)
	case 0x9A:
		cpu.sbc_r(cpu.a, *cpu.d, 1)
	case 0x9B:
		cpu.sbc_r(cpu.a, *cpu.e, 1)
	case 0x9C:
		cpu.sbc_r(cpu.a, *cpu.h, 1)
	case 0x9D:
		cpu.sbc_r(cpu.a, *cpu.l, 1)
	case 0x9E:
		cpu.sbc_r(cpu.a, memory.read(cpu.hl()), 1)
	case 0xA7:
		cpu.and_r(cpu.a, *cpu.a, 1)
	case 0xA0:
		cpu.and_r(cpu.a, *cpu.b, 1)
	case 0xA1:
		cpu.and_r(cpu.a, *cpu.c, 1)
	case 0xA2:
		cpu.and_r(cpu.a, *cpu.d, 1)
	case 0xA3:
		cpu.and_r(cpu.a, *cpu.e, 1)
	case 0xA4:
		cpu.and_r(cpu.a, *cpu.h, 1)
	case 0xA5:
		cpu.and_r(cpu.a, *cpu.l, 1)
	case 0xA6:
		cpu.and_r(cpu.a, memory.read(cpu.hl()), 1)
	case 0xE6:
		cpu.and_r(cpu.a, memory.read(cpu.pc+1), 2)
	case 0xB7:
		cpu.or_r(cpu.a, *cpu.a, 1)
	case 0xB0:
		cpu.or_r(cpu.a, *cpu.b, 1)
	case 0xB1:
		cpu.or_r(cpu.a, *cpu.c, 1)
	case 0xB2:
		cpu.or_r(cpu.a, *cpu.d, 1)
	case 0xB3:
		cpu.or_r(cpu.a, *cpu.e, 1)
	case 0xB4:
		cpu.or_r(cpu.a, *cpu.h, 1)
	case 0xB5:
		cpu.or_r(cpu.a, *cpu.l, 1)
	case 0xB6:
		cpu.or_r(cpu.a, memory.read(cpu.hl()), 1)
	case 0xF6:
		cpu.or_r(cpu.a, memory.read(cpu.pc+1), 2)
	case 0xAF:
		cpu.xor_r(cpu.a, *cpu.a, 1)
	case 0xA8:
		cpu.xor_r(cpu.a, *cpu.b, 1)
	case 0xA9:
		cpu.xor_r(cpu.a, *cpu.c, 1)
	case 0xAA:
		cpu.xor_r(cpu.a, *cpu.d, 1)
	case 0xAB:
		cpu.xor_r(cpu.a, *cpu.e, 1)
	case 0xAC:
		cpu.xor_r(cpu.a, *cpu.h, 1)
	case 0xAD:
		cpu.xor_r(cpu.a, *cpu.l, 1)
	case 0xAE:
		cpu.xor_r(cpu.a, memory.read(cpu.hl()), 1)
	case 0xEE:
		cpu.xor_r(cpu.a, memory.read(cpu.pc+1), 2)
	case 0xBF:
		cpu.cp_r(cpu.a, *cpu.a, 1)
	case 0xB8:
		cpu.cp_r(cpu.a, *cpu.b, 1)
	case 0xB9:
		cpu.cp_r(cpu.a, *cpu.c, 1)
	case 0xBA:
		cpu.cp_r(cpu.a, *cpu.d, 1)
	case 0xBB:
		cpu.cp_r(cpu.a, *cpu.e, 1)
	case 0xBC:
		cpu.cp_r(cpu.a, *cpu.h, 1)
	case 0xBD:
		cpu.cp_r(cpu.a, *cpu.l, 1)
	case 0xBE:
		cpu.cp_r(cpu.a, memory.read(cpu.hl()), 1)
	case 0xFE:
		cpu.cp_r(cpu.a, memory.read(cpu.pc+1), 2)
	case 0x3C:
		cpu.inc_r(cpu.a)
	case 0x04:
		cpu.inc_r(cpu.b)
	case 0x0C:
		cpu.inc_r(cpu.c)
	case 0x14:
		cpu.inc_r(cpu.d)
	case 0x1C:
		cpu.inc_r(cpu.e)
	case 0x24:
		cpu.inc_r(cpu.h)
	case 0x2C:
		cpu.inc_r(cpu.l)
	case 0x34:
		cpu.inc_addr(cpu.hl(), memory)
	case 0x3D:
		cpu.dec_r(cpu.a)
	case 0x05:
		cpu.dec_r(cpu.b)
	case 0x0D:
		cpu.dec_r(cpu.c)
	case 0x15:
		cpu.dec_r(cpu.d)
	case 0x1D:
		cpu.dec_r(cpu.e)
	case 0x25:
		cpu.dec_r(cpu.h)
	case 0x2D:
		cpu.dec_r(cpu.l)
	case 0x35:
		cpu.dec_addr(cpu.hl(), memory)
	case 0x09:
		cpu.add_r_double(cpu.h, cpu.l, cpu.bc())
	case 0x19:
		cpu.add_r_double(cpu.h, cpu.l, cpu.de())
	case 0x29:
		cpu.add_r_double(cpu.h, cpu.l, cpu.hl())
	case 0x39:
		cpu.add_r_double(cpu.h, cpu.l, cpu.sp)
	case 0xE8:
		cpu.add_sp(memory.read(cpu.pc + 1))
	case 0x03:
		cpu.inc_r_double(cpu.b, cpu.c)
	case 0x13:
		cpu.inc_r_double(cpu.d, cpu.e)
	case 0x23:
		cpu.inc_r_double(cpu.h, cpu.l)
	case 0x33:
		cpu.inc_sp()
	case 0x0B:
		cpu.dec_r_double(cpu.b, cpu.c)
	case 0x1B:
		cpu.dec_r_double(cpu.d, cpu.e)
	case 0x2B:
		cpu.dec_r_double(cpu.h, cpu.l)
	case 0x3B:
		cpu.dec_sp()
	case 0xCB:
		switch memory.read(cpu.pc + 1) {
		case 0x37:
			cpu.swap_r(cpu.a)
		case 0x30:
			cpu.swap_r(cpu.b)
		case 0x31:
			cpu.swap_r(cpu.c)
		case 0x32:
			cpu.swap_r(cpu.d)
		case 0x33:
			cpu.swap_r(cpu.e)
		case 0x34:
			cpu.swap_r(cpu.h)
		case 0x35:
			cpu.swap_r(cpu.l)
		case 0x36:
			cpu.swap_addr(cpu.hl(), memory)
		case 0x07:
			cpu.rlc_r(cpu.a)
		case 0x00:
			cpu.rlc_r(cpu.b)
		case 0x01:
			cpu.rlc_r(cpu.c)
		case 0x02:
			cpu.rlc_r(cpu.d)
		case 0x03:
			cpu.rlc_r(cpu.e)
		case 0x04:
			cpu.rlc_r(cpu.h)
		case 0x05:
			cpu.rlc_r(cpu.l)
		case 0x06:
			cpu.rlc_addr(cpu.hl(), memory)
		case 0x17:
			cpu.rl_r(cpu.a)
		case 0x10:
			cpu.rl_r(cpu.b)
		case 0x11:
			cpu.rl_r(cpu.c)
		case 0x12:
			cpu.rl_r(cpu.d)
		case 0x13:
			cpu.rl_r(cpu.e)
		case 0x14:
			cpu.rl_r(cpu.h)
		case 0x15:
			cpu.rl_r(cpu.l)
		case 0x16:
			cpu.rl_addr(cpu.hl(), memory)
		case 0x0F:
			cpu.rrc_r(cpu.a)
		case 0x08:
			cpu.rrc_r(cpu.b)
		case 0x09:
			cpu.rrc_r(cpu.c)
		case 0x0A:
			cpu.rrc_r(cpu.d)
		case 0x0B:
			cpu.rrc_r(cpu.e)
		case 0x0C:
			cpu.rrc_r(cpu.h)
		case 0x0D:
			cpu.rrc_r(cpu.l)
		case 0x0E:
			cpu.rrc_addr(cpu.hl(), memory)
		case 0x1F:
			cpu.rr_r(cpu.a)
		case 0x18:
			cpu.rr_r(cpu.b)
		case 0x19:
			cpu.rr_r(cpu.c)
		case 0x1A:
			cpu.rr_r(cpu.d)
		case 0x1B:
			cpu.rr_r(cpu.e)
		case 0x1C:
			cpu.rr_r(cpu.h)
		case 0x1D:
			cpu.rr_r(cpu.l)
		case 0x1E:
			cpu.rr_addr(cpu.hl(), memory)
		case 0x27:
			cpu.sla_r(cpu.a)
		case 0x20:
			cpu.sla_r(cpu.b)
		case 0x21:
			cpu.sla_r(cpu.c)
		case 0x22:
			cpu.sla_r(cpu.d)
		case 0x23:
			cpu.sla_r(cpu.e)
		case 0x24:
			cpu.sla_r(cpu.h)
		case 0x25:
			cpu.sla_r(cpu.l)
		case 0x26:
			cpu.sla_addr(cpu.hl(), memory)
		case 0x2F:
			cpu.sra_r(cpu.a)
		case 0x28:
			cpu.sra_r(cpu.b)
		case 0x29:
			cpu.sra_r(cpu.c)
		case 0x2A:
			cpu.sra_r(cpu.d)
		case 0x2B:
			cpu.sra_r(cpu.e)
		case 0x2C:
			cpu.sra_r(cpu.h)
		case 0x2D:
			cpu.sra_r(cpu.l)
		case 0x2E:
			cpu.sra_addr(cpu.hl(), memory)
		case 0x3F:
			cpu.srl_r(cpu.a)
		case 0x38:
			cpu.srl_r(cpu.b)
		case 0x39:
			cpu.srl_r(cpu.c)
		case 0x3A:
			cpu.srl_r(cpu.d)
		case 0x3B:
			cpu.srl_r(cpu.e)
		case 0x3C:
			cpu.srl_r(cpu.h)
		case 0x3D:
			cpu.srl_r(cpu.l)
		case 0x3E:
			cpu.srl_addr(cpu.hl(), memory)
		case 0x47:
			cpu.bit_r(cpu.a, memory.read(cpu.pc+1))
		case 0x40:
			cpu.bit_r(cpu.b, memory.read(cpu.pc+1))
		case 0x41:
			cpu.bit_r(cpu.c, memory.read(cpu.pc+1))
		case 0x42:
			cpu.bit_r(cpu.d, memory.read(cpu.pc+1))
		case 0x43:
			cpu.bit_r(cpu.e, memory.read(cpu.pc+1))
		case 0x44:
			cpu.bit_r(cpu.h, memory.read(cpu.pc+1))
		case 0x45:
			cpu.bit_r(cpu.l, memory.read(cpu.pc+1))
		case 0x46:
			cpu.bit_addr(cpu.hl(), memory.read(cpu.pc+1), memory)
		case 0xC7:
			cpu.set_r(cpu.a, memory.read(cpu.pc+1))
		case 0xC0:
			cpu.set_r(cpu.b, memory.read(cpu.pc+1))
		case 0xC1:
			cpu.set_r(cpu.c, memory.read(cpu.pc+1))
		case 0xC2:
			cpu.set_r(cpu.d, memory.read(cpu.pc+1))
		case 0xC3:
			cpu.set_r(cpu.e, memory.read(cpu.pc+1))
		case 0xC4:
			cpu.set_r(cpu.h, memory.read(cpu.pc+1))
		case 0xC5:
			cpu.set_r(cpu.l, memory.read(cpu.pc+1))
		case 0xC6:
			cpu.set_addr(cpu.hl(), memory.read(cpu.pc+1), memory)
		case 0x87:
			cpu.res_r(cpu.a, memory.read(cpu.pc+1))
		case 0x80:
			cpu.res_r(cpu.b, memory.read(cpu.pc+1))
		case 0x81:
			cpu.res_r(cpu.c, memory.read(cpu.pc+1))
		case 0x82:
			cpu.res_r(cpu.d, memory.read(cpu.pc+1))
		case 0x83:
			cpu.res_r(cpu.e, memory.read(cpu.pc+1))
		case 0x84:
			cpu.res_r(cpu.h, memory.read(cpu.pc+1))
		case 0x85:
			cpu.res_r(cpu.l, memory.read(cpu.pc+1))
		case 0x86:
			cpu.res_addr(cpu.hl(), memory.read(cpu.pc+1), memory)
		default:
			panic(fmt.Sprintf("unknown instruction: CB %X", opcode))
		}
	case 0x27:
		cpu.da_r(cpu.a)
	case 0x2F:
		cpu.cpl_r(cpu.a)
	case 0x3F:
		cpu.ccf()
	case 0x37:
		cpu.scf()
	case 0x00:
		cpu.nop()
	case 0x76:
		cpu.halt()
	case 0x10:
		cpu.stop()
	case 0xF3:
		cpu.di()
	case 0xFB:
		cpu.ei()
	case 0x07:
		cpu.rlc_r(cpu.a)
	case 0x17:
		cpu.rl_r(cpu.a)
	case 0x0F:
		cpu.rrc_r(cpu.a)
	case 0x1F:
		cpu.rr_r(cpu.a)
	case 0xC3:
		cpu.jp_nn(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)))
	case 0xC2:
		cpu.jp_nn_cc(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), flagZ, 0)
	case 0xCA:
		cpu.jp_nn_cc(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), flagZ, 1)
	case 0xD2:
		cpu.jp_nn_cc(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), flagC, 0)
	case 0xDA:
		cpu.jp_nn_cc(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), flagC, 1)
	case 0xE9:
		cpu.jp_nn(memory.readDouble(cpu.hl()))
	case 0x18:
		cpu.jp_nn(cpu.pc + uint16(memory.read(cpu.pc)))
	case 0x20:
		cpu.jp_nn_cc(cpu.pc+uint16(memory.read(cpu.pc)), flagZ, 0)
	case 0x28:
		cpu.jp_nn_cc(cpu.pc+uint16(memory.read(cpu.pc)), flagZ, 1)
	case 0x30:
		cpu.jp_nn_cc(cpu.pc+uint16(memory.read(cpu.pc)), flagC, 0)
	case 0x38:
		cpu.jp_nn_cc(cpu.pc+uint16(memory.read(cpu.pc)), flagC, 1)
	case 0xCD:
		cpu.call_nn(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), memory.read(cpu.pc+3), memory)
	case 0xC4:
		cpu.call_cc_nn(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), memory.read(cpu.pc+3), flagZ, 0, memory)
	case 0xCC:
		cpu.call_cc_nn(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), memory.read(cpu.pc+3), flagZ, 1, memory)
	case 0xD4:
		cpu.call_cc_nn(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), memory.read(cpu.pc+3), flagC, 0, memory)
	case 0xDC:
		cpu.call_cc_nn(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), memory.read(cpu.pc+3), flagC, 1, memory)
	default:
		panic(fmt.Sprintf("unknown instruction: %X", opcode))
	}
}

// 8-Bit Loads
func (cpu *cpu) ld_r(r *byte, n byte, incrementBy uint16) {
	*r = n
	cpu.pc += incrementBy
}

func (cpu *cpu) ld_addr(addr uint16, n byte, memory *memory, incrementBy uint16) {
	memory.write(addr, n)
	cpu.pc += incrementBy
}

// 16-Bit Loads
func (cpu *cpu) ld_r_double(r1 *byte, r2 *byte, nn uint16, incrementBy uint16) {
	*r1 = uint8(nn >> 8)
	*r2 = uint8(nn & 0x00FF)
	cpu.pc += incrementBy
}

func (cpu *cpu) ld_addr_double(addr uint16, nn uint16, memory *memory, incrementBy uint16) {
	memory.writeDouble(addr, nn)
	cpu.pc += incrementBy
}

func (cpu *cpu) ld_sp(nn uint16, incrementBy uint16) {
	cpu.sp = nn
	cpu.pc += incrementBy
}

func (cpu *cpu) push(nn uint16, memory *memory) {
	cpu.sp--
	memory.write(cpu.sp, uint8(nn>>8))
	cpu.sp--
	memory.write(cpu.sp, uint8(nn&0x00FF))
	cpu.pc++
}

func (cpu *cpu) pop(r1 *byte, r2 *byte, memory *memory) {
	*r1 = memory.read(cpu.sp)
	cpu.sp++
	*r2 = memory.read(cpu.sp)
	cpu.sp++
	cpu.pc++
}

// 8-Bit ALU
func (cpu *cpu) add_r(r *byte, n byte, incrementBy uint16) {
	*r += n
	// TODO: set flags
	cpu.pc += incrementBy
}

func (cpu *cpu) adc_r(r *byte, n byte, incrementBy uint16) {
	*r += (n + (*cpu.f & byte(flagC)))
	// TODO: set flags
	cpu.pc += incrementBy
}

func (cpu *cpu) sub_r(r *byte, n byte, incrementBy uint16) {
	*r -= n
	// TODO: set flags
	cpu.pc += incrementBy
}

func (cpu *cpu) sbc_r(r *byte, n byte, incrementBy uint16) {
	*r -= (n + (*cpu.f & byte(flagC)))
	// TODO: set flags
	cpu.pc += incrementBy
}

func (cpu *cpu) and_r(r *byte, n byte, incrementBy uint16) {
	*r &= n
	// TODO: set flags
	cpu.pc += incrementBy
}

func (cpu *cpu) or_r(r *byte, n byte, incrementBy uint16) {
	*r |= n
	// TODO: set flags
	cpu.pc += incrementBy
}

func (cpu *cpu) xor_r(r *byte, n byte, incrementBy uint16) {
	*r ^= n
	// TODO: set flags
	cpu.pc += incrementBy
}

func (cpu *cpu) cp_r(r *byte, n byte, incrementBy uint16) {
	// TODO: set flags
	cpu.pc += incrementBy
}

func (cpu *cpu) inc_r(n *byte) {
	*n++
	// TODO: set flags
	cpu.pc++
}

func (cpu *cpu) inc_addr(addr uint16, memory *memory) {
	memory.increment(addr)
	// TODO: set flags
	cpu.pc++
}

func (cpu *cpu) dec_r(n *byte) {
	*n--
	// TODO: set flags
	cpu.pc++
}

func (cpu *cpu) dec_addr(addr uint16, memory *memory) {
	memory.decrement(addr)
	// TODO: set flags
	cpu.pc++
}

// 16-Bit ALU
func (cpu *cpu) add_r_double(r1 *byte, r2 *byte, nn uint16) {
	res := double(*r1, *r2) + nn
	*r1 = uint8(res >> 8)
	*r2 = uint8(res & 0x00FF)
	// TODO: set flags
	cpu.pc++
}

func (cpu *cpu) add_sp(n byte) {
	cpu.sp += uint16(n)
	// TODO: set flags
	cpu.pc += 2
}

func (cpu *cpu) inc_r_double(r1 *byte, r2 *byte) {
	res := double(*r1, *r2) + 1
	*r1 = uint8(res >> 8)
	*r2 = uint8(res & 0x00FF)
	cpu.pc++
}

func (cpu *cpu) inc_sp() {
	cpu.sp++
	cpu.pc++
}

func (cpu *cpu) dec_r_double(r1 *byte, r2 *byte) {
	res := double(*r1, *r2) - 1
	*r1 = uint8(res >> 8)
	*r2 = uint8(res & 0x00FF)
	cpu.pc++
}

func (cpu *cpu) dec_sp() {
	cpu.sp--
	cpu.pc++
}

// Misc
func (cpu *cpu) swap_r(r *byte) {
	*r = *r&0x0F<<4 | *r>>4
	// TODO: set flags
	cpu.pc += 2
}

func (cpu *cpu) swap_addr(addr uint16, memory *memory) {
	n := memory.read(addr)
	memory.write(addr, n&0x0F<<4|n>>4)
	// TODO: set flags
	cpu.pc += 2
}

func (cpu *cpu) da_r(r *byte) {
	// TODO
	// TODO: set flags
	cpu.pc++
}

func (cpu *cpu) cpl_r(r *byte) {
	*r ^= *r
	// TODO: set flags
	cpu.pc++
}

func (cpu *cpu) ccf() {
	// TODO
	// TODO: set flags
	cpu.pc++
}

func (cpu *cpu) scf() {
	// TODO
	// TODO: set flags
	cpu.pc++
}

func (cpu *cpu) nop() {
	cpu.pc++
}

func (cpu *cpu) halt() {
	// TODO: disable opcode execution until interrupt
	cpu.pc++
}

func (cpu *cpu) stop() {
	/// TODO: disable opcode execution until button pressed
	cpu.pc++
}

func (cpu *cpu) di() {
	// TODO: disable interrupts AFTER following instruction
	cpu.pc++
}

func (cpu *cpu) ei() {
	// TODO: enable interrupts AFTER following instruction
	cpu.pc++
}

// Rotates and Shifts
func (cpu *cpu) rlc_r(r *byte) {
	bit7 := 128 & *r
	if bit7 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	*r = byte(bits.RotateLeft8(uint8(*r), 1))
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) rlc_addr(addr uint16, memory *memory) {
	n := memory.read(addr)
	bit7 := 128 & n
	if bit7 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	memory.write(addr, byte(bits.RotateLeft8(uint8(n), 1)))
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) rl_r(r *byte) {
	bit7 := 128 & *r
	if bit7 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	*r = byte(bits.RotateLeft8(uint8(*r), 1))
	if *cpu.f&16 == 0 {
		*r = cpu.clearBit(*r, 0)
	} else {
		*r = cpu.setBit(*r, 0)
	}
	cpu.clearFlag(flagZ)
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) rl_addr(addr uint16, memory *memory) {
	n := memory.read(addr)
	bit7 := 128 & n
	if bit7 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	n = byte(bits.RotateLeft8(uint8(n), 1))
	if *cpu.f&16 == 0 {
		n = cpu.clearBit(n, 0)
	} else {
		n = cpu.setBit(n, 0)
	}
	memory.write(addr, n)
	cpu.clearFlag(flagZ)
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) rrc_r(r *byte) {
	bit0 := 1 & *r
	if bit0 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	*r = byte(bits.RotateLeft8(uint8(*r), -1))
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) rrc_addr(addr uint16, memory *memory) {
	n := memory.read(addr)
	bit0 := 1 & n
	if bit0 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	n >>= 1
	memory.write(addr, n)
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) rr_r(r *byte) {
	bit0 := 1 & *r
	if bit0 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	*r >>= 1
	if *cpu.f&16 == 0 {
		*r = cpu.clearBit(*r, 7)
	} else {
		*r = cpu.setBit(*r, 7)
	}
	cpu.clearFlag(flagZ)
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) rr_addr(addr uint16, memory *memory) {
	n := memory.read(addr)
	bit0 := 1 & n
	if bit0 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	n >>= 1
	if *cpu.f&16 == 0 {
		n = cpu.clearBit(n, 7)
	} else {
		n = cpu.setBit(n, 7)
	}
	cpu.clearFlag(flagZ)
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) sla_r(r *byte) {
	bit7 := 128 & *r
	if bit7 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	*r <<= 1
	*r = cpu.clearBit(*r, 0)
	cpu.clearFlag(flagZ)
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) sla_addr(addr uint16, memory *memory) {
	n := memory.read(addr)
	bit7 := 128 & n
	if bit7 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	n <<= 1
	n = cpu.clearBit(n, 0)
	memory.write(addr, n)
	cpu.clearFlag(flagZ)
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) sra_r(r *byte) {
	bit0 := 1 & *r
	if bit0 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	*r >>= 1
	cpu.clearFlag(flagZ)
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) sra_addr(addr uint16, memory *memory) {
	n := memory.read(addr)
	bit0 := 1 & n
	if bit0 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	n >>= 1
	memory.write(addr, n)
	cpu.clearFlag(flagZ)
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) srl_r(r *byte) {
	bit0 := 1 & *r
	if bit0 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	*r >>= 1
	*r = cpu.clearBit(*r, 7)
	cpu.clearFlag(flagZ)
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) srl_addr(addr uint16, memory *memory) {
	n := memory.read(addr)
	bit0 := 1 & n
	if bit0 == 0 {
		cpu.clearFlag(flagC)
	} else {
		cpu.setFlag(flagC)
	}
	n >>= 1
	n = cpu.clearBit(n, 7)
	memory.write(addr, n)
	cpu.clearFlag(flagZ)
	cpu.clearFlag(flagN)
	cpu.clearFlag(flagH)
	cpu.pc += 2
}

// Bit Opcodes
func (cpu *cpu) bit_r(r *byte, bitPos byte) {
	if (*r>>bitPos)&1 == 0 {
		cpu.setFlag(flagZ)
	}
	cpu.clearFlag(flagN)
	cpu.setFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) bit_addr(addr uint16, bitPos byte, memory *memory) {
	n := memory.read(addr)
	if (n>>bitPos)&1 == 0 {
		cpu.setFlag(flagZ)
	}
	cpu.clearFlag(flagN)
	cpu.setFlag(flagH)
	cpu.pc += 2
}

func (cpu *cpu) set_r(r *byte, bitPos byte) {
	*r = cpu.setBit(*r, int(bitPos))
	cpu.pc += 2
}

func (cpu *cpu) set_addr(addr uint16, bitPos byte, memory *memory) {
	n := memory.read(addr)
	n = cpu.setBit(n, int(bitPos))
	memory.write(addr, n)
	cpu.pc += 2
}

func (cpu *cpu) res_r(r *byte, bitPos byte) {
	*r = cpu.clearBit(*r, int(bitPos))
	cpu.pc += 2
}

func (cpu *cpu) res_addr(addr uint16, bitPos byte, memory *memory) {
	n := memory.read(addr)
	n = cpu.clearBit(n, int(bitPos))
	memory.write(addr, n)
	cpu.pc += 2
}

// Jumps
func (cpu *cpu) jp_nn(addr uint16) {
	cpu.pc = addr
}

func (cpu *cpu) jp_nn_cc(addr uint16, flagPos int, expected byte) {
	if cpu.getBit(*cpu.f, flagPos) == expected {
		cpu.pc = addr
	} else {
		cpu.pc += 3
	}
}

// Calls
func (cpu *cpu) call_nn(addr uint16, nextInstruction byte, memory *memory) {
	cpu.sp--
	memory.write(cpu.sp, nextInstruction)
	cpu.pc = addr
}

func (cpu *cpu) call_cc_nn(addr uint16, nextInstruction byte, flagPos int, expected byte, memory *memory) {
	if cpu.getBit(*cpu.f, flagPos) == expected {
		cpu.call_nn(addr, nextInstruction, memory)
	} else {
		cpu.pc += 3
	}
}
