package snowflake

import (
	"errors"
	"example/fundemo01/web-app/settings"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sync"
	"time"
)

const (
	centerIdBits  int64 = 4                      //中心码位数
	workerIdBits  int64 = 6                      //机器码位数
	sequenceBits  int64 = 12                      //序列号位数
	timestampBits int64 = 41
	maxWorkerId int64 = -1 ^ (-1 << workerIdBits)   //机器码最大值（即1023）
	maxCenterId int64 = -1 ^ (-1 << centerIdBits)   //序列号最大值（即4095）
	workerIdShift   = sequenceBits// 机器id左移偏移量
	centerIdShift   = sequenceBits+workerIdBits // 数据中心机房id左移偏移量
    sequenceMask int64 = -1 ^ (-1 << sequenceBits)                        // 计算毫秒内，最大的序列号
	timestampShift = sequenceBits + workerIdBits + centerIdBits // 时间截向左移22位
	maxTimeStamp int64 = -1 ^ (-1 << timestampBits)
)

var	SF *snowFlake


type snowFlake struct {
	Epoch     int64 // 起始时间戳
	Nowtimestamp int64 // 当前时间戳，毫秒
	CenterId  int64 // 数据中心机房ID
	WorkerId  int64 // 机器ID
	Sequence  int64 // 毫秒内序列号
	LastTimestamp  int64 // 上一次生成ID的时间戳
	Err             error
	Lock sync.Mutex // 锁
}

func  getStartTimestamp() (int64){
	var location = viper.GetString("snowflake.location")
	var starttime = viper.GetString("snowflake.starttime")
	var loc,_ = time.LoadLocation(location)
	st,err :=  time.ParseInLocation("2006-01-02",starttime,loc)
	if err != nil {
		return 0
		zap.L().Error("func getStartTimestamp error!")
	}
	return st.UnixNano()/1e6
}

func getNowTimestamp() (int64){
	var location = viper.GetString("snowflake.location")
	loc,_ := time.LoadLocation(location)
	currentTime := time.Now()
	nt,err := time.ParseInLocation("2006-01-02 15:04:05",currentTime.Format("2006-01-02 15:04:05"),loc)
	if err != nil {
		return 0
		zap.L().Error("func getNowTimestamp error!")
	}
	return nt.UnixNano()/1e6
}

func  Init(cfg *settings.SnowFlake) (*snowFlake,error) {
	var err error
	if int(cfg.CenterId ) > int(maxCenterId) || cfg.CenterId  < 0 {
		err = errors.New(fmt.Sprintf("Center ID can't be greater than %d or less than 0", maxCenterId))
		return &snowFlake{},err
	}
	if int(cfg.WorkerId) > int(maxWorkerId) || cfg.WorkerId < 0 {
		err =  errors.New(fmt.Sprintf("Worker ID can't be greater than %d or less than 0", maxWorkerId))
		return &snowFlake{},err
	}

	SF = &snowFlake{
		Epoch: getStartTimestamp(),
		Nowtimestamp: getNowTimestamp(),
		CenterId: cfg.CenterId,
		WorkerId: cfg.WorkerId,
		Sequence: -1,
		LastTimestamp: getNowTimestamp()-1,
		Err: nil,
	}
	return SF,nil
}

func  GenID(s *snowFlake) (int64,error) {
	s.Lock.Lock() //设置锁，保证线程安全
	defer s.Lock.Unlock()
	if s.Nowtimestamp < s.LastTimestamp {             // 如果当前时间小于上一次 ID 生成的时间戳，说明发生时钟回拨
		return 0, errors.New(fmt.Sprintf("Clock moved backwards. Refusing to generate id for %d milliseconds", s.LastTimestamp-s.Nowtimestamp))
	}

	t := s.Nowtimestamp - s.Epoch
	if t > maxTimeStamp {
		return 0, errors.New(fmt.Sprintf("epoch must be between 0 and %d", maxTimeStamp-1))
	}

	// 同一时间生成的，则序号+1
	if s.LastTimestamp == s.Nowtimestamp {
		s.Sequence = (s.Sequence + 1) & sequenceMask
		// 毫秒内序列溢出：超过最大值; 阻塞到下一个毫秒，获得新的时间戳
		if s.Sequence == 0 {
			for s.Nowtimestamp <= s.LastTimestamp {
				s.Nowtimestamp = getNowTimestamp()
			}
		}
	} else {
		s.Sequence = 0 // 时间戳改变，序列重置
	}
	// 保存本次的时间戳
	s.LastTimestamp = s.Nowtimestamp

	// 根据偏移量，向左位移达到
	return (t << timestampShift) | (s.CenterId << centerIdShift) | (s.WorkerId << workerIdShift) | s.Sequence,nil
}