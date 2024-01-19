# tisea-backend

This is the main repository of the backend of Tisea.

- Built on **gin** with `Go 1.12.1`
- RESTful API design
- CLI control

## Repository Map

Project Package Name: `tisea-backend` (Used in `import`s and other package-related scenario)

- **cli** - *Manages the commandline function of the backend applicaton. Used to control running, file, database of the backend.*
  - `cli.go` - Definitions of available CLI commands.
  - others - Implementations of each command, presented as exported functions. These functions are called in `cli.go` and correspond to different commands.
- **handlers** - *Handlers of the endpoints. These handlers are eventually applied to `gin` and be called when some request is received.*
  - (handler-tag) - Directories under **handlers** are named after the tag of the set of requests handled by the code inside them.
    - Inside each directory, there is a file that is of the same name of the directory, exporting a function of which signature is like `Bind(*gin.Engine)`. This file should and should only contain this function, used to be called in the `cli/backend.go` in the format of `tag.Bind(gin)`. 
    - Other files in the directory are named after the functionality they are responsible for. For example, `handlers/auth/login.go` exposes a function `login(*gin.Context)` inside the package, and the function is expected to be called in `Bind` located at `auth/auth.go`.
- **structs** - *Definitions of structures used in passing groups of named values from front-end to backend, and to the database.*
  - (category-of-struct) - Structures are put in the file named after their category. E.g. `dynamic.go` contains all the structures used in the Dynamic module.
- **utils** - *Utilities.*
- **middlewares** - *Middleware functions.*

## RESTful API Map

Here is the map of all the API endpoints implemented for now. Modules that the endpoint belongs to is written in 3rd title, followed with the detailed endpoints explanations.

### `auth` - Authentication

- *POST* `/auth/login` JSON - Get a JSON Web Token. Returns OK with token if login credentials are valid.
- *POST* `/auth/register` JSON - Insert a unique user into the server database. Returns OK if a user is created.

### `dynamic` - Dynamic Posts

- *POST* `/dynamic` JSON - Create a new dynamic using the identity of the provided token.
- *DELETE* `/dynamic?=` Query - (Hard or Soft) Delete the dynamic with the provided `id`. "Soft Delete" actually stands for "Hide" operation.
- *GET* `/dynamic?=` Query - Get a list of dynamics. The list is built according to the arguments provided.

## Improvement & Suggestions Wanted!

This project is trying to make a concrete, inspiring and easy-to-use backend of a potential Minecraft community. Any suggestion and direct improvement (through PRs) is welcome, and we'd also like to listen to your unique opinions on this project (through Issues or Email).

## License

MIT
