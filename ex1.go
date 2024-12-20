package frickelcloud

// AvailableRessources computes the Host's available ressources.
func (h Host) AvailableRessources() Server {
	// TODO: sum up allocated ressources of h.Guest
	for _, vm := range h.Guests {
		cpu += vm.CPU
		mem += vm.Memory
		disk += vm.Disk
	}

	// TODO: subtract allocted ressources from h.Hardware
	cpu = h.Hardware.CPU - cpu
	mem = h.Hardware.Memory - mem
	disk = h.Hardware.Disk - disk

	// TODO: return Server struct with the ressources
	return Server{cpu, mem, disk}
}

// CanFitIn returns true if the host has enough available ressources for the
// given vm, and false otherwise.
func (h Host) CanFitIn(vm Server) bool {
	// TODO: call AvailableRessources on the host
	f := h.AvailableRessources()
	// TODO: check that all the host's available ressources are larger than the given vm
	if (f.CPU >= vm.CPU) && (f.Memory >= vm.Memory) && (f.Disk >= vm.Disk) {
		return true
	}
	return false
}
