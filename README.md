# Siri Shortcuts

This is a system to expose an API for custom Siri shortcuts.


## File Structure

- `shortcuts` Actual shortcuts
- `server` The API, exposed over HTTP and WebSockets, for receiving shortcut action requests
- `client` The application that processes targetted action requests
