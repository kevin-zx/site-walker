package sitewalker

type DomainFilter interface {
	IsAllowed(domain string) bool
	Add(domain string)
}

type domainFilterImpl struct {
	domains map[string]bool
}

func NewDomainFilter() DomainFilter {
	return &domainFilterImpl{
		domains: make(map[string]bool),
	}
}

func (d *domainFilterImpl) IsAllowed(domain string) bool {
	_, ok := d.domains[domain]
	return ok
}

func (d *domainFilterImpl) Add(domain string) {
	d.domains[domain] = true
}
