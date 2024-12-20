package frickelcloud

// AvailableRessources computes the Host's available ressources.
func (h Host) AvailableRessources() Server {
	// TODO: sum up allocated ressources of h.Guest
	var cpu uint8 = 0
	var mem uint32 = 0
	var ssd uint16 = 0

	for _, vm := range h.Guests {
		cpu = cpu + vm.CPU
		mem = mem + vm.RAM
		ssd = ssd + vm.SSD
	}

	// TODO: subtract allocted ressources from h.Hardware
	cpu = h.Hardware.CPU - cpu
	mem = h.Hardware.RAM - mem
	ssd = h.Hardware.SSD - ssd

	// TODO: return Server struct with the ressources
	return Server{cpu, mem, ssd}
}

// CanFitIn returns true if the host has enough available ressources for the
// given vm, and false otherwise.
func (h Host) CanFitIn(vm Server) bool {
	// TODO: call AvailableRessources on the host
	f := h.AvailableRessources()
	// TODO: check that all the host's available ressources are larger than the given vm
	if (f.CPU >= vm.CPU) && (f.RAM >= vm.RAM) && (f.SSD >= vm.SSD) {
		return true
	}
	return false
}
