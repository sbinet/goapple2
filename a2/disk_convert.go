package main

import (
	"log"

	"github.com/gonuts/commander"
	"github.com/zellyn/goapple2/disk"
)

var cmdDiskConvert = &commander.Command{
	Run:       runDiskConvert,
	UsageLine: "diskconvert infile outfile",
	Short:     "convert apple II disk images",
	Long: `
DiskConvert is a simple disk conversion utility.
`,
}

var diskVolume uint

func init() {
	cmdDiskConvert.Flag.UintVar(&diskVolume, "v", 0, "The volume of the disk, or 0 for default.")
}

func runDiskConvert(cmd *commander.Command, args []string) {
	if len(args) != 2 {
		cmd.Usage()
		return
	}
	if diskVolume > 254 {
		log.Fatalf("disk volume must be 0-254, got %d", diskVolume)
	}
	nyb, err := disk.DiskFromFile(args[0], byte(diskVolume))
	if err != nil {
		log.Fatal(err)
	}
	if err = disk.DiskToFile(args[1], nyb); err != nil {
		log.Fatal(err)
	}
}
