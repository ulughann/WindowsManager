package main

import "fmt"

type pkg struct {
	displayname string
	name        string
	logo        string
}

func gen(appName, appLogo, displayName string) string {
	out := fmt.Sprintf(`
	<div class="item bg-base-200 w-[200px] h-[200px] rounded-box">
		<center>
			<div class="h-[112px]"><img style="width:108px;height:auto;padding-top:16px;" src="%v"></div>
			<br><a id="install_%v" class="btn btn-neutral">Install %v</a>
		</center>
	</div>
	`, appLogo, appName, displayName)

	return out
}

func (a *App) RenderAppsList() string {
	apps := []pkg{
		{"vscode", "vscode", "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/vscode/vscode-original.svg"},
		{"chrome", "googlechrome", "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/chrome/chrome-original.svg"},
		{"firefox", "firefox", "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/firefox/firefox-original.svg"},
		{"spyware", "opera", "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/opera/opera-original.svg"},
		{"brave", "brave", "https://brave.com/static-assets/images/brave-logo-sans-text.svg"},
		{"cpu-z", "cpu-z", "https://upload.wikimedia.org/wikipedia/commons/2/2a/CPU-Z_Icon.svg"},
		{"spotify", "spotify", "https://upload.wikimedia.org/wikipedia/commons/8/84/Spotify_icon.svg"},
		{"acrobat", "adobereader", "https://upload.wikimedia.org/wikipedia/commons/6/60/Adobe_Acrobat_Reader_icon_%282020%29.svg"},
		{"java SE 8", "jre8", "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/java/java-original.svg"},
		{"Notepad++", "notepadplusplus.install", "https://upload.wikimedia.org/wikipedia/commons/0/0f/Notepad%2B%2B_Logo.png"},
		{"vlc", "vlc", "https://upload.wikimedia.org/wikipedia/commons/e/e6/VLC_Icon.svg"},
		{"git", "git.install", "https://git-scm.com/images/logos/downloads/Git-Icon-1788C.svg"},
		{"paint.net", "paint.net", "https://softcroco.com/wp-content/uploads/2021/06/Paint.net-download.png"},
		{"notion", "notion", "https://upload.wikimedia.org/wikipedia/commons/e/e9/Notion-logo.svg"},
		{"obsidian", "obsidian", "https://upload.wikimedia.org/wikipedia/commons/1/10/2023_Obsidian_logo.svg"},
		{"cherrytree", "cherrytree", "https://www.kali.org/tools/cherrytree/images/cherrytree-logo.svg"},
	}

	out := ""
	for _, app := range apps {
		out += gen(app.name, app.logo, app.displayname)
	}

	return out
}
