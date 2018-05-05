package main
import (
    "github.com/fatih/color"
)

func shouldBeBold(styles []color.Attribute) bool {
    badPairs := make([][]color.Attribute, 1)
    badPairs[0] = []color.Attribute{color.FgBlue, color.BgGreen}
    if len(styles) != 2 {
        return true
    } else {
        for _, pair := range badPairs {
            if styles[0] == pair[0] && styles[1] == pair[1] {
                return false
            } else if styles[1] == pair[1] && styles[0] == pair[0] {
                return false
            }
        }
        return true
    }
}

func ApplyStyles(s string, styles ...color.Attribute) string {
    c := color.New(styles...)
    if shouldBeBold(styles) {
        c.Add(color.Bold)    
    }
    c.EnableColor()
    return c.Sprint(s)
}

func StylePrompt(prompt Prompt, default_color int, symlink_color int, open_write int) {
    last_fs_root := "/"
    current_fs_root := "/"
    proper_path := "/"
    mounts := GetMounts()

    for i := 1; i < len(prompt); i += 1 {
        proper_path += prompt[i].Name
        resolved_path := proper_path
        style := prompt[i].NameStyle

        if !prompt[i].Shadowed {
            if IsLink(proper_path) {
                style = append(style, color.Attribute(symlink_color))
                resolved_path = ResolvedPath(proper_path)
            } else {
                style = append(style, color.Attribute(default_color))
            }

            if IsOpenWrite(proper_path) {
                style = append(style, color.Attribute(open_write))
            }
        }
        current_fs_root = GetPathRoot(resolved_path, mounts)
        if current_fs_root != last_fs_root {
            style = append(style, color.Underline)
            last_fs_root = current_fs_root
        }
        prompt[i].NameStyle = style
        proper_path += "/"
    }
}
