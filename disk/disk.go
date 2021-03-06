package disk

const (
	NUM_TRACKS           = 35
	DOS_DISK_BYTES       = 143360 // 35 tracks * 16 sectors * 256 bytes
	DOS_TRACK_BYTES      = DOS_DISK_BYTES / 35
	NYBBLE_DISK_BYTES    = 232960
	NYBBLE_TRACK_BYTES   = NYBBLE_DISK_BYTES / 35
	NYBBLE_ADDRESS_BYTES = 14
	NYBBLE_DATA_BYTES    = 349
	DEFAULT_VOLUME       = 254
	DEFAULT_PRESYNC      = 20
	DEFAULT_INTRASYNC    = 8
	MAX_PRE_INTRA_SYNC   = (NYBBLE_TRACK_BYTES / 16) - NYBBLE_ADDRESS_BYTES - NYBBLE_DATA_BYTES
)
