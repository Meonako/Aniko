# Aniko

Stable Diffusion Discord Bot written in GO. Using [discordgo](https://github.com/bwmarrin/discordgo).

This rely on [AUTOMATIC1111's stable diffusion webui](https://github.com/AUTOMATIC1111/stable-diffusion-webui).

### Currently support

- [x] Text to Image
- [x] View Progress
- [ ] Image to Image
- [ ] Extras
- [ ] PNG Info

# Installation & Usage

- Need your own discord bot. If you don't have one or don't know how to create, please search Google/YouTube for that. Too many tutorial are there already.
- If you don't know how to get Stable Diffusion Web UI up and running, I suggest you should read [THIS](https://github.com/AUTOMATIC1111/stable-diffusion-webui#installation-and-running) first

----

Assuming you already have GO working environment
1. Clone this repo
1. Create `config.env` in root folder
1. Create variable named `TOKEN` and set it value to your discord bot token.
It may look like this
```
TOKEN=Section1.Section2.Section3
```
4. ***[ OPTIONAL ]*** Change `OWNER_ID` in `config.json` to your discord ID.
5. Install dependencies
```
go mod tidy
```
6. Run using
```
go run .
```
OR
```
go build
.\Aniko
```

----

Alternatively, on Windows, you can download already build version in [releases](https://github.com/Meonako/Aniko/releases) section and follow these steps
1. Create `config.env` where `Aniko.exe` is located
1. Create variable named `TOKEN` and set it value to your discord bot token.

It may look like this
```
TOKEN=Section1.Section2.Section3
```
3. ***[ OPTIONAL ]*** Change `OWNER_ID` in `config.json` to your discord ID.
4. Run by clicking .exe
