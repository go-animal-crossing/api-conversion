package main

import (
	"acnh/data/config"
	"acnh/data/source"
	"acnh/data/target"
	"acnh/pages"
	"acnh/utility"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/afero"
)

func relatedLinkIs(is string) string {
	return fmt.Sprintf("View <a href='/%s/%s'>northern</a> or <a href='/%s/%s'>southern</a> hemisphere.",
		is, "northern",
		is, "southern")
}

func relatedLinkIsMonth(is string, ms string) string {
	return fmt.Sprintf("View <a href='/%s/%s/%s'>northern</a> or <a href='/%s/%s/%s'>southern</a> hemisphere.",
		is, ms, "northern",
		is, ms, "southern")
}

func relatedLinkIsMonthHem(is string, ms string, hs config.HemisphereItem) string {
	ah := altHemisphere(hs)
	return fmt.Sprintf("View <a href='/%s/%s/%s'>%s</a> hemisphere.",
		is, ms, ah.Slug, ah.Name)
}

func relatedLinkIsHem(is string, hs config.HemisphereItem) string {
	ah := altHemisphere(hs)
	return fmt.Sprintf("View <a href='/%s/%s'>%s</a> hemisphere.",
		is, ah.Slug, ah.Name)
}

func relatedLinkType(title string, slug string) string {
	return fmt.Sprintf("View <a href='/%s/%s/%s'>northern</a> or <a href='/%s/%s/%s'>southern</a> hemisphere %s.",
		slug, "available", "northern",
		slug, "available", "southern",
		title)
}

func relatedLinkTypeIs(title string, slug string, is string) string {
	return fmt.Sprintf("View <a href='/%s/%s/%s'>northern</a> or <a href='/%s/%s/%s'>southern</a> hemisphere %s.",
		slug, is, "northern",
		slug, is, "southern",
		title)
}

func relatedLinkTypeIsMonth(title string, slug string, is string, month string) string {
	return fmt.Sprintf("View <a href='/%s/%s/%s/%s'>northern</a> or <a href='/%s/%s/%s/%s'>southern</a> hemisphere %s.",
		slug, is, month, "northern",
		slug, is, month, "southern",
		title)
}

func relatedLinkTypeIsMonthHem(title string, slug string, is string, month string, hs config.HemisphereItem) string {
	ah := altHemisphere(hs)
	return fmt.Sprintf("View <a href='/%s/%s/%s/%s'>%s</a> hemisphere %s.",
		slug, is, month, ah.Slug,
		ah.Name,
		title)
}

func altHemisphere(h config.HemisphereItem) config.HemisphereItem {
	if h.ID == config.North {
		return config.SouthHemisphere
	}
	return config.NorthHemisphere
}

// main func
func main() {

	// download and convert
	source := source.Source{}
	source.Load()
	converted, e := source.Convert()

	if e != nil {
		log.Fatalf("ERROR: %v\n", e)
		os.Exit(1)
	}

	fs := afero.NewOsFs()
	dir := "./src/data"
	now := time.Now()
	// now we save each converted item as a page
	var pg pages.Page

	for _, item := range converted.All {
		//fmt.Printf("Saving [%d/%d] [%v of type %v] as page to file\n", i+1, l, item.Title, item.Type.Title)
		pg.IsList = false
		pg.Items = []target.Item{item}
		pg.Meta = pages.Meta{Type: item.Type.Slug}

		fmt.Printf("Generating %s\n", pg.GetURL())

		err := pg.Save(fs, dir)
		if err != nil {
			log.Fatalf("ERROR: %v\n", err)
			os.Exit(1)
		}
	}

	// work out Links
	monthlinks := make([]pages.MonthLink, 0)
	for i := 1; i <= 12; i++ {
		m := time.Month(i)
		t := time.Date(2000, m, 1, 1, 1, 1, 1, time.UTC)
		monthlinks = append(monthlinks, pages.MonthLink{
			MonthLong:  t.Format("January"),
			MonthShort: t.Format("Jan")})
	}

	// now these are the list pages
	// => /
	// => /mine
	// => /{type}/
	// => /{type}/{new|leaving|available}/
	// => /{type}/{new|leaving|available}/{north|south}
	// => /{type}/{new|leaving|available}/{month}/
	// => /{type}/{new|leaving|available}/{month}/{north|south}
	//
	// => /{new|leaving|available}/
	// => /{new|leaving|available}/{month}/
	// => /{new|leaving|available}/{month}/{north|south}
	// => /{new|leaving|available}/{north|south}

	// => /
	available := converted.Filter(target.Filter{Is: config.Available})
	fmt.Printf("Generating / -> found (%d) items\n", len(available))
	home := pages.Page{
		Title:    "Homepage",
		URI:      "/",
		Template: "home.html",
		IsList:   true,
		Items:    available,
		Type:     target.Type{Title: "home"},
		Meta: pages.Meta{
			Type: "home",
			Links: pages.Links{
				Months: monthlinks}},
		Grid: pages.Grid{
			Is:         true,
			Type:       true,
			Hemisphere: true}}
	home.Save(fs, dir)

	// => /mine
	all := converted.Filter(target.Filter{})
	fmt.Printf("Generating /mine -> found (%d) items\n", len(all))
	mine := pages.Page{
		ID:       "shareable",
		Title:    "My Island",
		URI:      "/mine",
		Template: "shareable.html",
		IsList:   true,
		Items:    all,
		Type:     target.Type{Title: "home"},
		Meta: pages.Meta{
			Type: "mixed",
			Links: pages.Links{
				Months: monthlinks}},
		Grid: pages.Grid{
			Is:         false,
			Type:       true,
			Hemisphere: false}}
	mine.Save(fs, dir)

	fmt.Printf("Generating /shared/ -> found (%d) items\n", len(all))
	shared := pages.Page{
		ID:       "shared",
		Title:    "Shared Island",
		URI:      "/shared",
		Template: "shared.html",
		IsList:   true,
		Items:    all,
		Type:     target.Type{Title: "home"},
		Meta: pages.Meta{
			Type: "mixed",
			Links: pages.Links{
				Months: monthlinks}},
		Grid: pages.Grid{
			Is:         false,
			Type:       true,
			Hemisphere: false}}
	shared.Save(fs, dir)

	// => /{type}/
	for _, m := range config.Config.ModelConfigs {

		byType := converted.Filter(target.Filter{
			Month: now.Month(),
			Type:  m.IsA})

		fmt.Printf("Generating /{%s}/ -> found (%d) items\n", m.Slug, len(byType))
		title := fmt.Sprintf("%s", m.Title)
		pg = pages.Page{
			Title:    title,
			H1:       title,
			URI:      utility.URL(m.Slug),
			Template: "list.html",
			IsList:   true,
			Items:    byType,
			Has:      m.Has,
			Type: target.Type{
				Title: m.Title,
				Slug:  m.Slug},
			Meta: pages.Meta{
				Type: m.Slug,
				Links: pages.Links{
					Months:  monthlinks,
					Related: relatedLinkType(m.Title, m.Slug)}}}
		pg.Save(fs, dir)

		// => /{type}/{new|leaving|available}/
		for _, is := range config.Config.IsOptions {
			// reset month to current
			byTypeAndIs := converted.Filter(target.Filter{
				Type:  m.IsA,
				Month: now.Month(),
				Is:    is.ID})
			fmt.Printf("Generating /{%s}/%s -> found (%d) items\n", m.Slug, is.Slug, len(byTypeAndIs))

			title = fmt.Sprintf("%s %s", is.Name, m.Title)
			pg.Title = title
			pg.H1 = title
			pg.Items = byTypeAndIs
			pg.URI = utility.URL(m.Slug, is.Slug)
			pg.Meta.Is = is.Slug
			// rebuild related link
			pg.Meta.Links.Related = relatedLinkTypeIs(m.Title, m.Slug, is.Slug)

			pg.Save(fs, dir)

			// => /{type}/{new|leaving|available}/{north|south}
			for _, h := range config.Config.HemisphereOptions {
				byTypeIsAndH := converted.Filter(target.Filter{
					Month:      now.Month(),
					Type:       m.IsA,
					Hemisphere: h.ID,
					Is:         is.ID})
				fmt.Printf("Generating /{%s}/%s/%s -> found (%d) items\n", m.Slug, is.Slug, h.Slug, len(byTypeIsAndH))
				title = fmt.Sprintf("%s %s in the %s hemisphere", is.Name, m.Title, h.Name)
				p := pg
				p.Title = title
				p.H1 = title
				p.Items = byTypeIsAndH
				p.Meta.Hemisphere = h.Slug
				//
				ah := altHemisphere(h)

				p.Meta.Links.Related = fmt.Sprintf("View <a href='%s/%s/%s'>%s</a> hemisphere %s here.",
					utility.URL(m.Slug), is.Slug, ah.Slug,
					ah.Name,
					m.Title)
				p.URI = utility.URL(m.Slug, is.Slug, h.Slug)
				p.Save(fs, dir)
			}

			// => /{type}/{new|leaving|available}/{month}/
			for monthN := 1; monthN <= 12; monthN++ {
				mth := time.Month(monthN)
				ms := utility.Slugify(mth.String())
				byTypeIsAndM := converted.Filter(target.Filter{
					Type:  m.IsA,
					Month: mth,
					Is:    is.ID})

				fmt.Printf("Generating /{%s}/%s/%s -> found (%d) items\n", m.Slug, is.Slug, ms, len(byTypeIsAndM))

				title = fmt.Sprintf("%s %s in %s", is.Name, m.Title, mth.String())
				p := pg
				p.Meta.Month = ms
				p.Title = title
				p.H1 = title
				p.Items = byTypeIsAndM
				p.URI = utility.URL(m.Slug, is.Slug, ms)
				pg.Meta.Links.Related = relatedLinkTypeIsMonth(m.Title, m.Slug, is.Slug, ms)

				p.Save(fs, dir)

				// => /{type}/{new|leaving|available}/{month}/{north|south}
				for _, hs := range config.Config.HemisphereOptions {
					byTypeIsHAndM := converted.Filter(target.Filter{
						Type:       m.IsA,
						Month:      mth,
						Hemisphere: hs.ID,
						Is:         is.ID})

					fmt.Printf("Generating /{%s}/%s/%s/%s -> found (%d) items\n", m.Slug, is.Slug, ms, hs.Slug, len(byTypeIsHAndM))

					title = fmt.Sprintf("%s %s in %s for the %s hemisphere", is.Name, m.Title, mth.String(), hs.Name)
					p.Meta.Hemisphere = hs.Slug
					p.Title = title
					p.H1 = title
					p.Items = byTypeIsAndM
					p.URI = utility.URL(m.Slug, is.Slug, ms, hs.Slug)

					pg.Meta.Links.Related = relatedLinkTypeIsMonthHem(m.Title, m.Slug, is.Slug, ms, hs)

					p.Save(fs, dir)

				}

			}

		}
	}

	// => /{new|leaving|available}/
	for _, is := range config.Config.IsOptions {
		byIs := converted.Filter(target.Filter{
			Month: now.Month(),
			Is:    is.ID})
		fmt.Printf("Generating /{%s}/ -> found (%d) items\n", is.Slug, len(byIs))
		title := fmt.Sprintf("%s", is.Name)
		pg = pages.Page{
			Title:    title,
			H1:       title,
			URI:      utility.URL(is.Slug),
			Template: "list-mixed.html",
			IsList:   true,
			Items:    byIs,
			Has:      config.Has{},
			Type:     target.Type{},
			Meta: pages.Meta{
				Is: is.Slug,
				Links: pages.Links{
					Months:  monthlinks,
					Related: relatedLinkIs(is.Slug)}}}

		pg.Save(fs, dir)

		// => /{new|leaving|available}/{north|south}

		for _, h := range config.Config.HemisphereOptions {
			byIsAndH := converted.Filter(target.Filter{
				Hemisphere: h.ID,
				Month:      now.Month(),
				Is:         is.ID})

			fmt.Printf("Generating /{%s}/%s -> found (%d) items\n", is.Slug, h.Slug, len(byIsAndH))

			title = fmt.Sprintf("%s in the %s hemisphere", is.Name, h.Name)
			p := pg
			p.Meta.Hemisphere = h.Slug
			p.Title = title
			p.H1 = title
			p.Items = byIsAndH
			p.URI = utility.URL(is.Slug, h.Slug)
			p.Meta.Links.Related = relatedLinkIsHem(is.Slug, h)
			p.Save(fs, dir)
		}

		// => /{new|leaving|available}/{month}/
		for monthN := 1; monthN <= 12; monthN++ {
			mth := time.Month(monthN)
			ms := utility.Slugify(mth.String())

			byIsAndM := converted.Filter(target.Filter{
				Month: mth,
				Is:    is.ID})

			fmt.Printf("Generating /{%s}/%s -> found (%d) items\n", is.Slug, ms, len(byIsAndM))
			title = fmt.Sprintf("%s in %s", is.Name, mth.String())
			p := pg
			p.Meta.Month = ms
			p.Title = title
			p.H1 = title
			p.URI = utility.URL(is.Slug, ms)
			p.Items = byIsAndM
			p.Meta.Links.Related = relatedLinkIsMonth(is.Slug, ms)
			p.Save(fs, dir)

			// => /{new|leaving|available}/{month}/{north|south}
			for _, h := range config.Config.HemisphereOptions {
				byIsMAndH := converted.Filter(target.Filter{
					Month:      mth,
					Hemisphere: h.ID,
					Is:         is.ID})

				fmt.Printf("Generating /{%s}/%s/%s/ -> found (%d) items\n", is.Slug, ms, h.Slug, len(byIsMAndH))

				title = fmt.Sprintf("%s in %s for the %s hemisphere", is.Name, mth.String(), h.Name)
				p.Meta.Hemisphere = h.Slug
				p.Title = title
				p.H1 = title
				p.Items = byIsMAndH
				p.URI = utility.URL(is.Slug, ms, h.Slug)
				p.Meta.Links.Related = relatedLinkIsMonthHem(is.Slug, ms, h)
				p.Save(fs, dir)
			}

		}

	}

}
