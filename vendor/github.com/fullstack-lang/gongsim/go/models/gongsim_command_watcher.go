package models

import (
	"log"
	"runtime/pprof"
	"time"
)

var lastSimTime = EngineSingloton.currentTime

var DisplayWatch bool
var CpuProfile bool

// the watcher thread inspects the status of the simulation
func (gongsimCommand *GongsimCommand) watcher() {

	// The period shall not too short for performance sake but not too long because the end user needs a responsive application
	//
	// checkoutSchedulerPeriod is the period of the "checkout scheduler"
	watcherPeriod := 500 * time.Millisecond
	var WatcherSchedulerPeriod = time.NewTicker(watcherPeriod)

	realtimeSimStart := time.Now()
	for {
		select {
		case t := <-WatcherSchedulerPeriod.C:

			_ = t

			// const layout = "Jan 2, 2006 at 15:04:05 (MST)"
			const layout = "15:04:05.999"
			measuredSimSpeed := float64(EngineSingloton.currentTime.Sub(lastSimTime)) / float64(watcherPeriod)
			if DisplayWatch {
				log.Printf("time %s, next %s, status %s, speed %f, speed request %f, Sim %s, Ho %s",
					time.Now().Format(layout), EngineSingloton.nextRealtimeHorizon.Format(layout),
					EngineSingloton.State, measuredSimSpeed, EngineSingloton.Speed,
					EngineSingloton.currentTime.Format(layout), EngineSingloton.nextSimulatedTimeHorizon.Format(layout))
			}
			lastSimTime = EngineSingloton.currentTime

			if CpuProfile {
				if time.Since(realtimeSimStart) > 20*time.Second {
					pprof.StopCPUProfile()
					log.Println("generated CPU profile")
				}
			}
		}
	}
}
