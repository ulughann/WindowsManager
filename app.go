package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jaypipes/ghw"
	. "github.com/klauspost/cpuid/v2"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/sys/windows/registry"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Install() string {
	return InstallChoco()
}

func (a *App) OsInfo() ([]string, error) {
	return getWindowsVersion(), nil
}

func (a *App) Memory() string {
	return getMemory()
}

func (a *App) CpuInfo() []string {
	return getCpuInfo()
}

func (a *App) CpuFreq() []float64 {
	i, _ := cpu.Percent(1*time.Second, true)
	return i
}

func (a *App) GpuInfo() []string {
	m := []string{}
	gpu, err := ghw.GPU()
	if err != nil {
		fmt.Printf("Error getting GPU info: %v", err)
	}

	fmt.Printf("%v\n", gpu)

	for _, card := range gpu.GraphicsCards {
		m = append(m, card.DeviceInfo.Product.Name)
	}

	return m
}

func getMemory() string {
	memory, err := ghw.Memory()
	if err != nil {
		fmt.Printf("Error getting memory info: %v", err)
	}
	return memory.String()
}

func getWindowsVersion() []string {
	a := []string{}
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	cv, _, err := k.GetStringValue("CurrentVersion")
	if err != nil {
		log.Fatal(err)
	}
	a = append(a, cv)

	pn, _, err := k.GetStringValue("ProductName")
	if err != nil {
		log.Fatal(err)
	}
	a = append(a, pn)

	maj, _, err := k.GetIntegerValue("CurrentMajorVersionNumber")
	if err != nil {
		log.Fatal(err)
	}
	a = append(a, fmt.Sprintf("%d", maj))

	min, _, err := k.GetIntegerValue("CurrentMinorVersionNumber")
	if err != nil {
		log.Fatal(err)
	}
	a = append(a, fmt.Sprintf("%d", min))

	cb, _, err := k.GetStringValue("CurrentBuild")
	if err != nil {
		log.Fatal(err)
	}
	a = append(a, cb)

	return a
}

func getCpuInfo() []string {
	a := []string{}
	a = append(a, CPU.BrandName)
	a = append(a, fmt.Sprintf("%d", CPU.PhysicalCores))
	a = append(a, fmt.Sprintf("%d", CPU.ThreadsPerCore))
	a = append(a, fmt.Sprintf("%d", CPU.LogicalCores))
	a = append(a, fmt.Sprintf("%d", CPU.PhysicalCores))
	a = append(a, fmt.Sprintf("%d", CPU.Family))
	a = append(a, strings.Join(CPU.FeatureSet(), ","))
	a = append(a, fmt.Sprintf("%d", CPU.CacheLine))
	a = append(a, fmt.Sprintf("%d", CPU.Cache.L1D))
	a = append(a, fmt.Sprintf("%d", CPU.Cache.L1I))
	a = append(a, fmt.Sprintf("%d", CPU.Cache.L2))
	a = append(a, fmt.Sprintf("%d", CPU.Cache.L3))
	a = append(a, fmt.Sprintf("%d", CPU.Hz))
	a = append(a, fmt.Sprintf("%d", CPU.BoostFreq))

	if CPU.Supports(SSE, SSE2) {
		a = append(a, "SSE")
	}
	return a
}
func (a *App) GetCPUFrequencies() []float64 {
	info, err := cpu.Info()
	if err != nil {
		fmt.Println("Error:", err)
	}

	var freqs []float64
	for _, i := range info {
		if i.Mhz > 0 {
			freqs = append(freqs, i.Mhz)
		}
	}

	if len(freqs) == 0 {
		fmt.Println("No CPU frequencies found")
	} else {
		return freqs
	}
	return freqs
}

func (a *App) GetMemUsage() []any {
	v, _ := mem.VirtualMemory()

	return []any{v, v, v, v.UsedPercent}
}

func (a *App) GetStorage() []any {
	c := []any{}
	block, err := ghw.Block()
	if err != nil {
		return c
	}

	for _, disk := range block.Disks {
		for _, part := range disk.Partitions {
			c = append(c, part)
		}
	}
	return c
}
