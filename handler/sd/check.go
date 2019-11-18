package sd

import (
	"github.com/coffee377/voc-admin/model/account"
	"github.com/coffee377/voc-admin/result"
	"github.com/coffee377/voc-admin/service"
	"github.com/gin-gonic/gin"
	"log"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// HealthCheck shows `OK` as the ping-pong result.
func HealthCheck(c *gin.Context) {
	id := c.Param("id")
	//c.Query()
	log.Printf(id)
	r := result.Of(service.Login(account.Credentials{}))
	c.JSON(r.Response())
	//handler.RestResponse(c, r, result.OK)
}

//// DiskCheck checks the disk usage.
//func DiskCheck(c *gin.Context) {
//	u, _ := disk.Usage("/")
//
//	usedMB := int(u.Used) / MB
//	usedGB := int(u.Used) / GB
//	totalMB := int(u.Total) / MB
//	totalGB := int(u.Total) / GB
//	usedPercent := int(u.UsedPercent)
//
//	status := http.StatusOK
//	text := "OK"
//
//	if usedPercent >= 95 {
//		status = http.StatusOK
//		text = "CRITICAL"
//	} else if usedPercent >= 90 {
//		status = http.StatusTooManyRequests
//		text = "WARNING"
//	}
//
//	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
//	c.String(status, "\n"+message)
//}
//
//// CPUCheck checks the cpu usage.
//func CPUCheck(c *gin.Context) {
//	cores, _ := cpu.Counts(false)
//
//	a, _ := load.Avg()
//	l1 := a.Load1
//	l5 := a.Load5
//	l15 := a.Load15
//
//	status := http.StatusOK
//	text := "OK"
//
//	if l5 >= float64(cores-1) {
//		status = http.StatusInternalServerError
//		text = "CRITICAL"
//	} else if l5 >= float64(cores-2) {
//		status = http.StatusTooManyRequests
//		text = "WARNING"
//	}
//
//	message := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | Cores: %d", text, l1, l5, l15, cores)
//	c.String(status, "\n"+message)
//}
//
//// RAMCheck checks the disk usage.
//func RAMCheck(c *gin.Context) {
//	u, _ := mem.VirtualMemory()
//
//	usedMB := int(u.Used) / MB
//	usedGB := int(u.Used) / GB
//	totalMB := int(u.Total) / MB
//	totalGB := int(u.Total) / GB
//	usedPercent := int(u.UsedPercent)
//
//	status := http.StatusOK
//	text := "OK"
//
//	if usedPercent >= 95 {
//		status = http.StatusInternalServerError
//		text = "CRITICAL"
//	} else if usedPercent >= 90 {
//		status = http.StatusTooManyRequests
//		text = "WARNING"
//	}
//
//	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
//	c.String(status, "\n"+message)
//}