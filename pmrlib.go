package pmrlib

import(
	"encoding/json"
)

type Report struct {
	Changed bool		`json:"changed"`
	Hostname string		`json:"hostname"`
	MAC []string		`json:"mac"`
	IP []string		`json:"ip"`
	Distribution string	`json:"dist"`
	DistName string		`json:"distName"`
	DistVersion string	`json:"distVersion"`
	Architecture string	`json:"arch"`
	Packages []Package	`json:"packages"`
	Updates []Update	`json:"updates"`
	Repositories []string	`json:"repositories"`
	Errors []Error		`json:"errors"`
}

type Package struct {
	Name string		`json:"name"`
	Version string		`json:"version"`
	Architecture string	`json:"arch"`
	Status string		`json:"status"`
}

type Update struct {
	Name string		`json:"name"`
	Version string		`json:"version"`
	Available string	`json:"available"`
	Architecture string	`json:"arch"`
}

type Error struct {
	Code int		`json:"code"`
	Message string		`json:"message"`
}

func ToJSON(report Report) (string, error) {
	jsonReport, err := json.Marshal(report)
	if err != nil {
		return "", err
	}
	return string(jsonReport), nil
}

func (r *Report) FromJSON(data []byte) error {
	var rep Report
	if err := json.Unmarshal(data, &rep); err != nil {
		return err
	}

	r.Changed = rep.Changed
	r.Hostname = rep.Hostname
	r.Distribution = rep.Distribution
	r.DistName = rep.DistName
	r.DistVersion = rep.DistVersion
	r.Architecture = rep.Architecture
	r.MAC = make([]string, 0)
	r.IP = make([]string, 0)
	r.Packages = make([]Package, 0)
	r.Updates = make([]Update, 0)
	r.Repositories = make([]string, 0)
	r.Errors = make([]Error, 0)

	for i := range rep.MAC {
		if len(rep.MAC) <= 0 {
			continue
		}
		r.MAC = append(r.MAC, rep.MAC[i])
	}

	for i := range rep.IP {
		if len(rep.IP) <= 0 {
			continue
		}
		r.IP = append(r.IP, rep.IP[i])
	}

	for i := range rep.Packages {
		if len(rep.Packages) <= 0 {
			continue
		}
		var pkg Package
		pkg.Name = rep.Packages[i].Name
		pkg.Version = rep.Packages[i].Version
		pkg.Architecture = rep.Packages[i].Architecture
		pkg.Status = rep.Packages[i].Status
		r.Packages = append(r.Packages, pkg)
	}

	for i := range rep.Updates {
		if len(rep.Updates) <= 0 {
			continue
		}
		var upd Update
		upd.Name = rep.Updates[i].Name
		upd.Version = rep.Updates[i].Version
		upd.Architecture = rep.Updates[i].Architecture
		upd.Available = rep.Updates[i].Available
		r.Updates = append(r.Updates, upd)
	}

	for i := range rep.Errors {
		if len(rep.Errors) <= 0 {
			continue
		}
		var rerr Error
		rerr.Code = rep.Errors[i].Code
		rerr.Message = rep.Errors[i].Message
	}

	for i := range rep.Repositories {
		if len(rep.Repositories) <= 0 {
			continue
		}
		r.Repositories = append(r.Repositories, rep.Repositories[i])
	}
	return nil
}
