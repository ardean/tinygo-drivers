package hts221

const (

	// I2C address
	HTS221_ADDRESS = 0x5F

	// control/status registers
	HTS221_WHO_AM_I_REG  = 0x0F
	HTS221_AV_CONF_REG   = 0x10
	HTS221_CTRL1_REG     = 0x20
	HTS221_CTRL2_REG     = 0x21
	HTS221_STATUS_REG    = 0x27
	HTS221_HUMID_OUT_REG = 0x28
	HTS221_TEMP_OUT_REG  = 0x2A

	// calibration registers
	HTS221_H0_rH_x2_REG   = 0x30
	HTS221_H1_rH_x2_REG   = 0x31
	HTS221_T0_degC_x8_REG = 0x32
	HTS221_T1_degC_x8_REG = 0x33
	HTS221_T1_T0_MSB_REG  = 0x35
	HTS221_H0_T0_OUT_REG  = 0x36
	HTS221_H1_T0_OUT_REG  = 0x3A
	HTS221_T0_OUT_REG     = 0x3C
	HTS221_T1_OUT_REG     = 0x3E
)
