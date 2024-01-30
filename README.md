# Windows Manager

Chocolatey front-end, Performence Monitor and Tweaker/Optimizer for Windows 10/11

Built with Vue, DaisyUI, Wails and Go.

![image](/marketing/Ultimate.png)

## Roadmap
- [x] Chocolatey front-end + installer
- - [ ] Updated search
- [x] Performence Monitor
- - [x] Primary cpu/ram/disk monitoring
- - [ ] Gpu monitoring
- - [ ] Network monitoring
- [ ] Tweaks/Optimizations 
- [ ] More customizations

## Installation
Just run the latest standalone executable from the [releases](https://github.com/onrirr/WindowsManager/releases) page.

## Development
### Requirements
- [Go 1.18 or higher](https://golang.org/)
- [Node.js 16 or higher](https://nodejs.org/en/)
- [Wails](https://wails.app/)
  
### Build
```bash
# Clone the repository
$ git clone <repo>

# Go into the repository
$ cd WindowsManager

# Run Wails
$ wails dev
```

To build an executable, run `wails build` and the executable will be in the `build` folder.
