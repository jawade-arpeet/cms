package countries

import "sort"

var (
	byISO2  = make(map[string]Country, len(allCountries))
	byISO3  = make(map[string]Country, len(allCountries))
	byGroup = make(map[string]Group, len(groupMembers))
)

func init() {
	for _, country := range allCountries {
		if country.ISO2 == "" || country.ISO3 == "" {
			panic("countries: country with empty ISO code")
		}
		if _, exists := byISO2[country.ISO2]; exists {
			panic("countries: duplicate ISO2 code " + country.ISO2)
		}
		if _, exists := byISO3[country.ISO3]; exists {
			panic("countries: duplicate ISO3 code " + country.ISO3)
		}

		byISO2[country.ISO2] = country
		byISO3[country.ISO3] = country
	}

	for code, memberCodes := range groupMembers {
		if _, exists := byISO2[code]; exists {
			panic("countries: group code collides with ISO2 code " + code)
		}
		if _, exists := byISO3[code]; exists {
			panic("countries: group code collides with ISO3 code " + code)
		}

		seen := make(map[string]struct{}, len(memberCodes))
		members := make([]Country, 0, len(memberCodes))
		for _, memberCode := range memberCodes {
			if _, exists := seen[memberCode]; exists {
				panic("countries: duplicate member " + memberCode + " in group " + code)
			}
			seen[memberCode] = struct{}{}

			country, exists := byISO2[memberCode]
			if !exists {
				panic("countries: unknown member " + memberCode + " in group " + code)
			}
			members = append(members, country)
		}

		sort.Slice(members, func(i, j int) bool {
			return members[i].ISO2 < members[j].ISO2
		})

		byGroup[code] = Group{
			Code:    code,
			Name:    groupNames[code],
			Members: members,
		}
	}
}

func IsValidISO2(code string) bool {
	_, ok := byISO2[code]
	return ok
}

func IsValidISO3(code string) bool {
	_, ok := byISO3[code]
	return ok
}

func FromISO2(code string) (Country, bool) {
	country, ok := byISO2[code]
	return country, ok
}

func GroupMembers(groupCode string) ([]Country, bool) {
	group, ok := byGroup[groupCode]
	if !ok {
		return nil, false
	}

	members := make([]Country, len(group.Members))
	copy(members, group.Members)
	return members, true
}
