package main
import (
    "github.com/fatih/color"
)

func contains(arr []color.Attribute, needle color.Attribute) bool {
    for _, a := range arr {
        if needle == a {
            return true
        }
    }
    return false
}
func isSubset(sub []color.Attribute, super []color.Attribute) bool {
    for _, sb := range sub {
        if !contains(super, sb) {
            return false
        }
    }
    return true
}

func shouldBeBold(styles []color.Attribute) bool {
    terrible_contrast := make([][]color.Attribute, 1)
    terrible_contrast[0] = []color.Attribute{color.FgBlue, color.BgGreen}

    for _, pair := range terrible_contrast {
        if isSubset(pair, styles) {
            return false
        }
    }
    return true
}

func ApplyStyles(s string, styles ...color.Attribute) string {
    if s == "" || len(styles) == 0 {
        return s
    } else {
        c := color.New(styles...)
        if shouldBeBold(styles) {
            c.Add(color.Bold)    
        }
        c.EnableColor()
        return c.Sprint(s)
    }
}

func StylePrompt(prompt Prompt, default_color int, symlink_color int, open_write int) {
    last_fs_root := "/"
    current_fs_root := "/"
    path := "/"
    mounts := GetMounts()

    for i := 1; i < len(prompt); i += 1 {
        path += prompt[i].Name
        style := prompt[i].NameStyle

        if !prompt[i].Shadowed {
            if IsLink(path) {
                style = append(style, color.Attribute(symlink_color))
                path = ResolvedPath(path)
            } else {
                style = append(style, color.Attribute(default_color))
            }

            if IsOpenWrite(path) {
                style = append(style, color.Attribute(open_write))
            }
        }
        current_fs_root = GetPathRoot(path, mounts)
        if current_fs_root != last_fs_root {
            style = append(style, color.Underline)
            last_fs_root = current_fs_root
        }
        prompt[i].NameStyle = style
        path += "/"
    }
}
