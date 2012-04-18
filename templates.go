package main

import (
    "log"
    "os"
    "html/template"
    "time"
)

func parseTemplates(dir string) (t *template.Template, err error) {
    t, err = template.ParseGlob(dir + "/*")

    if err != nil {
        return
    }

    log.Print("Parsed Templates:")
    for _,templ := range t.Templates() {
        log.Print("  " + templ.Name())
    }

    return
}

func watchTemplates(dir string, tmpl **template.Template) {
    info, err := os.Stat(dir)
    if err != nil {
        log.Fatal("watchtemplates: couldn't stat template dir")
    }

    var lastModified time.Time

    for {
        if info.ModTime().After(lastModified) {
            t, err := parseTemplates("templates")

            if err != nil {
                log.Print("Template parse error: ", err)
            } else {
                *tmpl = t
            }

            lastModified = info.ModTime()
        }

        time.Sleep(1 * time.Second)
        info, err = os.Stat(dir)

        if err != nil {
            log.Print("watchtemplates: can no longer stat template dir")
        }

    }
}

