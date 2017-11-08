package snowflakework

import (
	"os"
	"time"
)

//0 41 10 12

var shiftTime uint64 = 1483200000000 //2017-01-01
var count uint64 = 0
var workdId uint64 = 0 //

func getNowTimeStamp() uint64 {
	timestamp := uint64(time.Now().UnixNano() / 1000000)
	return timestamp - shiftTime
}

func getWorkId() uint64 {
	var mask uint64 = 0x3ff
	if workdId > 0 {
		return workdId
	}
	//
	pid := uint64(os.Getpid())
	pid = mask & pid
	if pid == 0 {
		return 1
	}
	return pid
}

func getCount(timestamp uint64) uint64 {
	var mask uint64 = 0xfff
	count++
	count = count & mask
	if count == 0 {
		//wait next
		for {
			nowTimeStamp := getNowTimeStamp()
			if nowTimeStamp > timestamp {
				break
			}
		}
		return 0
	}
	return count
}

func SetWorkId(id uint64) {
	workdId = id
}

func MakeId() uint64 {
	time := getNowTimeStamp()
	workid := getWorkId()
	count := getCount(time)

	id := (time << 22) | (workid << 12) | count

	return id
}
