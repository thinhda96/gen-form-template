# gen-form-template
Generate live template for form of Decision Engine
## Install

```bash
go install github.com/thinhda96/gen-form-template@latest
```

## Usage

- Generate settings file

```bash
cd decision_engine
gen-form-template gen --o .
```

- Import setting

    1. Open `Goland`
    2. Select `Files` -> `Manage IDE Settings` -> `Import Settings...`
    3. Select `settings.zip` file in `decision_engine` folder
    4. If your settings is ok, `Select components to Import` popup will be shown
    5. Select `OK`
    6. Select `Import and Restart`
    7. After IDE is restarted, enjoy your life.

Note: GoLand don't support import settings if sync settings is enabled. `Files` -> `Manage IDE Settings` -> `IDE Settings Sync` -> `Disable Sync...`. After importing live templates, you can enable sync normally.